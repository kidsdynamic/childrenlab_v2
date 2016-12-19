package controller

import (
	"fmt"
	"log"
	"net/http"

	"strings"

	"strconv"

	"time"

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

	err = db.Select(&todayActivity, "SELECT a.id, steps, distance, received_date, type, a.date_created, d.mac_id FROM activity a "+
		"JOIN device d ON a.device_id = d.id WHERE d.id = ? AND YEAR(received_date) = ? AND MONTH(received_date) = ? AND DAY(received_date) = ?",
		device.ID, indoorActivity.Time.Year(), indoorActivity.Time.Month(), indoorActivity.Time.Day())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on retrieve today's activity",
			"error":   err,
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
