package controller

import (
	"fmt"
	"log"
	"net/http"

	"strings"

	"strconv"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kidsdynamic/childrenlab_v2/app/constants"
	"github.com/kidsdynamic/childrenlab_v2/app/database"
	"github.com/kidsdynamic/childrenlab_v2/app/model"
	"github.com/pkg/errors"
)

func UploadRawActivityData(c *gin.Context) {
	var request model.ActivityRawData
	if err := c.BindJSON(&request); err != nil {
		log.Printf("Error on activity upload data. Bind with json. %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	var kid model.Kid

	if err := db.Where("mac_id = ?", request.MacID).First(&kid).Error; err != nil {
		logError(errors.Wrapf(err, "can't find the device by the MAC ID: %s", request.Indoor))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("can't find the device by the MAC ID: %s", request.MacID),
		})
		return
	}

	indoor := strings.Split(request.Indoor, ",")

	indoorActivityLong, err := strconv.ParseInt(indoor[0], 10, 64)

	if err != nil {
		logError(errors.Wrapf(err, "Error on parse indoor time to Long: %s", request.Indoor))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error on parse indoor time to Long: %s", request.Indoor),
			"error":   err,
		})
		return
	}

	var exist bool
	row := db.Raw("SELECT EXISTS(SELECT id FROM activity_raw WHERE time = ? AND mac_id = ? LIMIT 1)", indoorActivityLong, kid.MacID).Row()
	if err := row.Scan(&exist); err != nil {
		logError(errors.Wrapf(err, "Something wrong when finding eixst activity: %s", indoorActivityLong))
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
	indoorActivity.Date = time.Unix(indoorActivityLong, 0)
	indoorActivity.TimeLong = indoorActivityLong
	indoorActivity.TimeZone = request.TimeZoneOffset

	outdoor := strings.Split(request.Outdoor, ",")
	outdoorActivityLong, err := strconv.ParseInt(outdoor[0], 0, 64)

	if err != nil {
		logError(errors.Wrapf(err, "Error on parse outdoor time to Long: %s", request.Indoor))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error on parse outdoor time to Long: %s", request.Indoor),
			"error":   err,
		})
		return
	}

	var outdoorActivity model.ActivityInsight
	outdoorActivity.Steps, err = strconv.ParseInt(outdoor[2], 10, 64)
	outdoorActivity.Date = time.Unix(outdoorActivityLong, 0)
	outdoorActivity.TimeZone = request.TimeZoneOffset
	outdoorActivity.TimeLong = outdoorActivityLong

	user := GetSignedInUser(c)

	request.DateCreated = GetNowTime()
	request.LastUpdated = GetNowTime()
	request.IndoorSteps = indoorActivity.Steps
	request.OutdoorSteps = outdoorActivity.Steps
	request.UserID = user.ID
	if err := db.Create(&request).Error; err != nil {
		logError(errors.Wrap(err, "Error on inserting raw activity."))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on inserting raw activity",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

	if err := calculateActivity(db, indoorActivity, outdoorActivity, kid); err != nil {
		logError(errors.Wrap(err, "Error on genreate activity."))
	}

	LogUserActivity(db, &user, fmt.Sprintf("Uplaod Raw Activity (%d)", request.ID), &kid.MacID)
}

func calculateActivity(db *gorm.DB, indoorActivity, outdoorActivity model.ActivityInsight, kid model.Kid) error {
	var todayActivity []model.Activity
	timeWithZone := indoorActivity.Date.Add(time.Duration(indoorActivity.TimeZone) * time.Minute)
	if err := db.Where("(mac_id = ? OR mac_id = REVERSE(?)) AND (YEAR(received_date) = ? AND MONTH(received_date) = ? AND DAY(received_date) = ?)", kid.MacID, kid.MacID, timeWithZone.Year(), timeWithZone.Month(), timeWithZone.Day()).
		Find(&todayActivity).Error; err != nil {
		return err
	}

	if len(todayActivity) == 0 {
		if err := db.Create(&model.Activity{
			Steps:        indoorActivity.Steps,
			ReceivedDate: timeWithZone,
			ReceivedTime: indoorActivity.TimeLong,
			KidID:        kid.ID,
			MacID:        kid.MacID,
			Type:         constants.ActivityIndoorType,
			DateCreated:  GetNowTime(),
			LastUpdated:  GetNowTime(),
		}).Error; err != nil {
			logError(errors.Wrap(err, "Error on create indoor activity record"))
			return err
		}

		if err := db.Create(&model.Activity{
			Steps:        outdoorActivity.Steps,
			ReceivedDate: timeWithZone,
			ReceivedTime: outdoorActivity.TimeLong,
			KidID:        kid.ID,
			MacID:        kid.MacID,
			Type:         constants.ActivityOutdoorType,
			DateCreated:  GetNowTime(),
			LastUpdated:  GetNowTime(),
		}).Error; err != nil {
			logError(errors.Wrap(err, "Error on create outdoor activity record"))
			return err
		}

	} else {
		for _, a := range todayActivity {
			if a.Type == constants.ActivityIndoorType {
				a.Steps += indoorActivity.Steps
				a.ReceivedDate = timeWithZone
				a.ReceivedTime = indoorActivity.TimeLong
				a.LastUpdated = GetNowTime()

			} else {
				a.Steps += outdoorActivity.Steps
				a.ReceivedDate = timeWithZone
				a.ReceivedTime = outdoorActivity.TimeLong
				a.LastUpdated = GetNowTime()
			}

			if err := db.Model(&model.Activity{}).Update(&a).Error; err != nil {
				logError(errors.Wrap(err, "Error on save activity record"))
				return err
			}
		}
	}

	return nil
}

func GetActivity(c *gin.Context) {
	user := GetSignedInUser(c)

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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("One of parameter is missing: %#v\n", activityRequest),
		})
		return
	}

	var periodDate *time.Time
	switch period {
	case "DAILY":
		periodDate = getTodayDate()
		break
	case "WEEKLY":
		periodDate = getBeginningOfWeek()
		break
	case "MONTHLY":
		periodDate = getBeginningOfMonth()
	case "YEARLY":
		periodDate = getBeginningOfYear()
		break
	default:
		err = errors.New(fmt.Sprintf("Can't recognize the period: %s", period))
	}

	db := database.NewGORM()
	defer db.Close()
	var activities []model.Activity
	if err := db.Joins("JOIN kids ON kids.id = activity.kid_id").Where("kids.id = ? AND kids.parent_id = ? AND activity.received_Date > ?", activityRequest.KidID, user.ID, &periodDate).Find(&activities).Error; err != nil {
		logError(errors.Wrap(err, "Error on retrieve Activity"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error on retriving activity: %#v\n", activityRequest),
			"error":   err,
		})
		return
	}

	//TODO: It's temp solution: Task: https://app.asana.com/0/33043844747220/308456358881086
	for i, activity := range activities {
		newSteps := float32(activity.Steps) * 0.7
		activities[i].Steps = int64(newSteps)
	}

	c.JSON(http.StatusOK, gin.H{
		"activities": activities,
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

	startTimeLong, err := strconv.ParseInt(startTimeString, 10, 64)
	if err != nil {
		logError(errors.Wrap(err, "Error on parse string to int - GetActivityByTime"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on parse string to int",
			"error":   err,
		})
		return
	}

	endTimeLong, err := strconv.ParseInt(endTimeString, 10, 64)
	if err != nil {
		logError(errors.Wrap(err, "Error on parse string to int - GetActivityByTime"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on parse string to int",
			"error":   err,
		})
		return
	}

	kidID, err := strconv.ParseInt(kidIdString, 10, 64)
	if err != nil {
		logError(errors.Wrap(err, "Error on parse string to int - GetActivityByTime"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on parse string to int",
			"error":   err,
		})
		return
	}

	start := time.Unix(startTimeLong, 0)
	end := time.Unix(endTimeLong, 0)

	user := GetSignedInUser(c)

	db := database.NewGORM()
	defer db.Close()

	var activities []model.Activity

	if err := db.Joins("JOIN kids ON kids.id = activity.kid_id").Where("kids.id = ? AND kids.parent_id = ? AND (activity.received_Date between ? and ?)", kidID, user.ID, start, end).Find(&activities).Error; err != nil {
		logError(errors.Wrap(err, "Error on getting activities"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on getting activities",
			"error":   err,
		})
		return
	}

	//TODO: It's temp solution: Task: https://app.asana.com/0/33043844747220/308456358881086
	for i, activity := range activities {
		newSteps := float32(activity.Steps) * 0.7
		activities[i].Steps = int64(newSteps)
	}

	c.JSON(http.StatusOK, gin.H{
		"activities": activities,
	})
	LogUserActivity(db, &user, "Get Activity By Time", nil)
}

func getTodayDate() *time.Time {
	now := time.Now()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, now.Location())

	return &today
}

func getBeginningOfWeek() *time.Time {
	now := time.Now()

	days := int(now.Weekday())
	if days == 0 {
		days = 7
	}

	now = now.AddDate(0, 0, -days+1)

	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, now.Location())

	return &today
}

func getBeginningOfMonth() *time.Time {
	now := time.Now()
	year, month, _ := now.Date()
	today := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())

	return &today
}

func getBeginningOfYear() *time.Time {
	now := time.Now()
	year, _, _ := now.Date()
	today := time.Date(year, 1, 1, 0, 0, 0, 0, now.Location())

	return &today
}

func GetActivityList(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	db := database.NewGORM()
	defer db.Close()

	var activity []model.Activity
	if err := db.Where("kid_id = ?", c.Param("kidId")).Find(&activity).Error; err != nil {
		logError(errors.Wrap(err, "Error on getting activities"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on getting activities",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, activity)
}
