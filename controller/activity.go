package controller

import (
	"fmt"
	"log"
	"net/http"

	"strings"

	"strconv"

	"time"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

func UploadRawActivityData(c *gin.Context) {
	var request model.ActivityRawDataRequest
	if err := c.BindJSON(&request); err != nil {
		log.Printf("Error on activity upload data. Bind with json. %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	db := database.New()
	defer db.Close()
	device, err := GetDeviceByMacID(db, request.MacID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("can't find the device by the MAC ID: %s", request.MacID),
		})
		return
	}

	indoor := strings.Split(request.Indoor, ",")

	indoorActivityLong, err := strconv.ParseInt(indoor[0], 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error on parse indoor time to Long: %s", request.Indoor),
			"error":   err,
		})
		return
	}

	var exist bool
	if err := db.Get(&exist, "SELECT EXISTS(SELECT id FROM activity_raw WHERE time = ? AND device_id = ? LIMIT 1)",
		indoorActivityLong, device.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when finding eixst activity",
			"err":     err,
		})
		return
	}

	if exist {
		c.JSON(http.StatusConflict, gin.H{
			"message": "This is a duplicate data",
		})
		return
	}

	var indoorActivity model.ActivityInsight
	indoorActivity.Steps, err = strconv.ParseInt(indoor[2], 10, 64)
	indoorActivity.Time = time.Unix(indoorActivityLong, 0)
	log.Printf("Received Indoor Activity Time: %s", indoorActivity.Time)

	outdoor := strings.Split(request.Outdoor, ",")
	outdoorActivityLong, err := strconv.ParseInt(outdoor[0], 0, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error on parse outdoor time to Long: %s", request.Indoor),
			"error":   err,
		})
		return
	}

	var outdoorActivity model.ActivityInsight
	outdoorActivity.Steps, err = strconv.ParseInt(outdoor[2], 10, 64)
	outdoorActivity.Time = time.Unix(outdoorActivityLong, 0)

	//indoorActivityLong = indoorActivityLong * 1000

	user := GetSignedInUser(c)
	result := db.MustExec("INSERT INTO activity_raw (device_id, indoor_activity, outdoor_activity, time, device_time, uploaded_user_id, date_created, last_updated) "+
		"VALUES (?, ?, ?, ?, ?, ?, Now(), Now())", device.ID, request.Indoor, request.Outdoor, indoorActivityLong, request.Time, user.ID)

	if !checkInsertResult(result) {
		log.Printf("Error on inserting raw activity. Data: %#v\n", request)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on inserting raw activity",
		})
		return
	}

	var todayActivity []model.Activity

	log.Printf("%d, %d, %d", indoorActivity.Time.Year(), indoorActivity.Time.Month(), indoorActivity.Time.Day())

	retrieveError := db.Select(&todayActivity, "SELECT a.id, steps, distance, received_date, type, d.mac_id FROM activity a "+
		"JOIN device d ON a.device_id = d.id WHERE d.id = ? AND YEAR(received_date) = ? AND MONTH(received_date) = ? AND DAY(received_date) = ?",
		device.ID, indoorActivity.Time.Year(), indoorActivity.Time.Month(), indoorActivity.Time.Day())

	if retrieveError != nil {
		log.Printf("Error on retreing today's activity. %#v", retrieveError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on retrieve today's activity",
			"error":   retrieveError,
		})
		return

	}

	if len(todayActivity) == 0 {
		_ = db.MustExec("INSERT INTO activity (steps, received_date, received_time, device_id, type, date_created, last_updated) VALUES "+
			"(?, ?, ?, ?, 'INDOOR', Now(), Now())",
			indoorActivity.Steps, indoorActivity.Time, indoorActivityLong, device.ID)

		_ = db.MustExec("INSERT INTO activity (steps, received_date, received_time, device_id, type, date_created, last_updated) VALUES "+
			"(?, ?, ?, ?, 'OUTDOOR', Now(), Now())",
			outdoorActivity.Steps, outdoorActivity.Time, outdoorActivityLong, device.ID)
	} else {
		for _, a := range todayActivity {
			if a.Type == "INDOOR" {
				a.Steps += indoorActivity.Steps
				a.ReceivedDate = indoorActivity.Time
			} else {
				a.Steps += outdoorActivity.Steps
				a.ReceivedDate = outdoorActivity.Time
			}

			result = db.MustExec("UPDATE activity SET steps = ?, received_date = ? WHERE id = ?",
				a.Steps, a.ReceivedDate, a.ID)

			if !checkInsertResult(result) {
				log.Println("Error on inserting activity.")
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error on inserting activity",
				})
				return
			}

		}
	}

	c.JSON(http.StatusOK, gin.H{})
}

func GetActivity(c *gin.Context) {
	user := GetSignedInUser(c)

	fmt.Printf("Daily Actiivty Request: %s, %s", c.Query("kidId"), c.Query("period"))

	kidIdString := c.Query("kidId")
	period := c.Query("period")
	kidId, err := strconv.ParseInt(kidIdString, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "kidId should be int type.",
			"error":   err,
		})
		return
	}

	var activityRequest model.ActivityRequest
	activityRequest.KidID = kidId

	if activityRequest.KidID == 0 || period == "" {
		log.Printf("Error on parsing activity request. %#v\n", activityRequest)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("One of parameter is missing: %#v\n", activityRequest),
		})
		return
	}

	db := database.New()
	defer db.Close()

	var activity []model.Activity
	switch period {
	case "DAILY":
		err = db.Select(&activity, "SELECT a.id, d.mac_id, d.kid_id, distance, a.received_date, steps, a.type FROM activity a JOIN device d ON a.device_id = d.id JOIN kids k ON "+
			"k.id = d.kid_id WHERE k.id = ? AND parent_id = ? AND a.received_Date > ?", activityRequest.KidID, user.ID, getTodayDate())
		break
	case "WEEKLY":
		err = db.Select(&activity, "SELECT a.id, d.mac_id, d.kid_id, distance, a.received_date, steps, a.type FROM activity a JOIN device d ON a.device_id = d.id JOIN kids k ON "+
			"k.id = d.kid_id WHERE k.id = ? AND parent_id = ? AND received_date > ?", activityRequest.KidID, user.ID, getBeginningOfWeek())
		break
	case "MONTHLY":
		err = db.Select(&activity, "SELECT a.id, d.mac_id, d.kid_id, distance, a.received_date, steps, a.type FROM activity a JOIN device d ON a.device_id = d.id JOIN kids k ON "+
			"k.id = d.kid_id WHERE k.id = ? AND parent_id = ? AND received_date > ?", activityRequest.KidID, user.ID, getBeginningOfMonth())
	case "YEARLY":
		err = db.Select(&activity, "SELECT a.id, d.mac_id, d.kid_id, distance, a.received_date, steps, a.type FROM activity a JOIN device d ON a.device_id = d.id JOIN kids k ON "+
			"k.id = d.kid_id WHERE k.id = ? AND parent_id = ? AND received_date > ?", activityRequest.KidID, user.ID, getBeginningOfYear())
		break
	default:
		err = errors.New(fmt.Sprintf("Can't recognize the period: %s", period))
	}

	if err != nil {
		log.Printf("Error on retrieve Activity: %#v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error on retriving activity: %#v\n", activityRequest),
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"activities": activity,
	})
}

func GetActivityByTime(c *gin.Context) {
	startTimeString := c.Query("start")
	endTimeString := c.Query("end")
	kidIdString := c.Query("kidId")

	if startTimeString == "" || endTimeString == "" || kidIdString == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Start time and end time and kid ID are required",
		})
		return
	}
	db := database.New()
	defer db.Close()

	startTimeLong, err := strconv.ParseInt(startTimeString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on parse string to int",
			"error":   err,
		})
		return
	}

	endTimeLong, err := strconv.ParseInt(endTimeString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on parse string to int",
			"error":   err,
		})
		return
	}

	kidID, err := strconv.ParseInt(kidIdString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on parse string to int",
			"error":   err,
		})
		return
	}

	start := time.Unix(startTimeLong, 0)
	end := time.Unix(endTimeLong, 0)

	user := GetSignedInUser(c)
	var activity []model.Activity
	err = db.Select(&activity, "SELECT a.id, d.mac_id, d.kid_id, distance, a.received_date, steps, a.type FROM activity a JOIN device d ON a.device_id = d.id JOIN kids k ON "+
		"k.id = d.kid_id WHERE k.id = ? AND parent_id = ? AND (a.received_Date between ? and ?)", kidID, user.ID, start, end)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on getting activities",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"activities": activity,
	})
}

func getTodayDate() *time.Time {
	now := time.Now()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, now.Location())

	fmt.Println("get beginning of day: ", today)
	return &today
}

func getBeginningOfWeek() *time.Time {
	now := time.Now()

	//now = now.Add(24 * time.Hour)

	fmt.Printf("Week of day: %d\n", int(now.Weekday()))

	days := int(now.Weekday())
	if days == 0 {
		days = 7
	}

	fmt.Printf("Days passed from fist day of week: %d\n", days)

	now = now.AddDate(0, 0, -days+1)

	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, now.Location())

	fmt.Println("get beginning of week: ", today)
	return &today
}

func getBeginningOfMonth() *time.Time {
	now := time.Now()
	year, month, _ := now.Date()
	today := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())

	fmt.Println("get beginning of Month: ", today)
	return &today
}

func getBeginningOfYear() *time.Time {
	now := time.Now()
	year, _, _ := now.Date()
	today := time.Date(year, 1, 1, 0, 0, 0, 0, now.Location())

	fmt.Println("get beginning of Year: ", today)
	return &today
}
