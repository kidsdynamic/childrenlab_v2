package controller

import (
	"net/http"

	"github.com/kidsdynamic/childrenlab_v2/app/database"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/app/model"
)

func AdminLogin(c *gin.Context) {
	var adminLogin model.AdminLogin

	if c.BindJSON(&adminLogin) != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	db := database.NewGORM()
	defer db.Close()
	var admin model.User

	adminLogin.Password = database.EncryptPassword(adminLogin.Password)

	db.Table("user").Joins("JOIN role ON user.role_id = role.id").Where("email = ? AND password = ? and "+
		"(authority = ? or authority = ?)", adminLogin.Name, adminLogin.Password, model.ROLE_ADMIN, model.ROLE_SUPER_ADMIN).Preload("Role").First(&admin)

	if admin.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	accessToken := model.AccessToken{
		Email:       admin.Email,
		Token:       randToken(),
		LastUpdated: GetNowTime(),
	}

	err := storeToken(db, accessToken)

	if err != nil {
		logError(errors.Wrap(err, "Store token fail!!!! ERROR from admin"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Store token failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username":     accessToken.Email,
		"role":         admin.Role.Authority,
		"access_token": accessToken.Token,
	})
}

func GetAllUser(c *gin.Context) {
	db := database.NewGORM()
	defer db.Close()

	var userList []model.User
	if err := db.Joins("JOIN role ON role.id = user.role_id").Where("role.authority = 'ROLE_USER'").Order("date_created desc").Limit(50).Find(&userList).Error; err != nil {
		logError(errors.Wrap(err, "Error on retriving user list from admin"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on retriving user list",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, userList)
}

func GetActivityListForAdmin(c *gin.Context) {
	db := database.NewGORM()
	defer db.Close()

	var activity []model.Activity
	if err := db.Where("kid_id = ?", c.Param("kidId")).Limit(100).Find(&activity).Error; err != nil {
		logError(errors.Wrap(err, "Error on getting activities from Admin"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on getting activities",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func GetActivityRawForAdmin(c *gin.Context) {
	db := database.NewGORM()
	defer db.Close()

	var activityRaw []model.ActivityRawData
	if err := db.Where("mac_id = ?", c.Param("macId")).Order("id desc").Limit(100).Find(&activityRaw).Error; err != nil {
		logError(errors.Wrap(err, "Error on getting activities from Admin"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on getting activities",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, activityRaw)
}

// Only for internal use
func CreateAdminUser(c *gin.Context) {
	var request model.AdminSignUpRequest

	if c.BindJSON(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	db := database.NewGORM()
	defer db.Close()
	var admin model.User

	request.Password = database.EncryptPassword(request.Password)
	err := db.Table("user").Joins("JOIN role ON user.role_id = role.id").Where("email = ? AND password = ? and authority = ?", request.Name, request.Password, model.ROLE_ADMIN).First(&admin).Error

	if err != nil {
		logError(errors.Wrap(err, "Error on create admin user"))
	}

	if admin.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{
			"message": "The admin name exists",
		})
		return
	}

	var role model.Role
	if err := db.Where("authority = ?", model.ROLE_ADMIN).First(&role).Error; err != nil {
		logError(errors.Wrap(err, "Error on getting admin role"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on getting admin role",
			"error":   err,
		})
		return
	}

	adminUser := model.User{
		Email:     request.Name,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Role:      role,
	}

	if err := db.Create(&adminUser).Error; err != nil {
		logError(errors.Wrap(err, "Error on creating admin user"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on creating admin user",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Dashboard(c *gin.Context) {
	db := database.NewGORM()
	defer db.Close()

	var dashboard model.Dashboard

	var signupCounts []model.SignupCountByDate
	if rows, err := db.Raw("select count(*) as signup, DATE_FORMAT(DATE(date_created), '%Y/%m/%d') as date from user u JOIN role r ON u.role_id = r.id where date_created != 0000-00-00" +
		" and r.`authority` = 'ROLE_USER' AND date_created >= '2017-04-24' group by date order by date desc LIMIT 20").Rows(); err != nil {
		if err != nil {
			logError(errors.Wrap(err, "Error on retrieve signup dashboard from Admin"))
		}
	} else {
		defer rows.Close()
		for rows.Next() {
			var signup model.SignupCountByDate
			rows.Scan(&signup.SignupCount, &signup.Date)
			signupCounts = append(signupCounts, signup)
		}
	}
	dashboard.Signup = signupCounts

	var activityCount []model.ActivityCountByDate
	if rows, err := db.Raw("select count(*), DATE_FORMAT(DATE(date_created), '%Y/%m/%d') as date, count(DISTINCT(user_id)) as userCount, sum(indoor_steps), sum(outdoor_steps) " +
		"from activity_raw group by date order by date desc LIMIT 20").Rows(); err != nil {
		logError(errors.Wrap(err, "Error on retrieve activity dashboard from Admin"))
	} else {
		defer rows.Close()
		for rows.Next() {
			var activity model.ActivityCountByDate
			rows.Scan(&activity.ActivityCount, &activity.Date, &activity.UserCount, &activity.IndoorSteps, &activity.OutdoorSteps)
			activityCount = append(activityCount, activity)
		}
	}

	if err := db.Table("user").Joins("JOIN role ON user.role_id = role.id where date_created != 0000-00-00 and authority = 'ROLE_USER' AND email not like '%kidsdynamic.com'").Count(&dashboard.TotalUserCount).Error; err != nil {
		logError(errors.Wrap(err, "Error on retrieve signup user dashboard from Admin"))
	}

	if err := db.Table("activity_raw").Count(&dashboard.TotalActivityCount).Error; err != nil {
		logError(errors.Wrap(err, "Error on getting activity raw count from Admin"))
	}

	dashboard.Activity = activityCount

	// Activity count depend on activity (event) date - https://app.asana.com/0/33043844747220/349867637183239
	var activityCountOnEventDate []model.ActivityCountByDate
	if rows, err := db.Raw("select count(*), DATE_FORMAT(FROM_UNIXTIME(time), '%Y/%m/%d') as date, count(DISTINCT(user_id)) as userCount, sum(indoor_steps), sum(outdoor_steps) " +
		"from activity_raw group by date order by date desc LIMIT 20").Rows(); err != nil {
		logError(errors.Wrap(err, "Error on retrieve activity dashboard from Admin"))
	} else {
		defer rows.Close()
		for rows.Next() {
			var activity model.ActivityCountByDate
			rows.Scan(&activity.ActivityCount, &activity.Date, &activity.UserCount, &activity.IndoorSteps, &activity.OutdoorSteps)
			activityCountOnEventDate = append(activityCountOnEventDate, activity)
		}
	}
	dashboard.ActivityByEventDate = activityCountOnEventDate

	c.JSON(http.StatusOK, dashboard)
}

func DeleteMacID(c *gin.Context) {
	macID := c.Query("macId")

	db := database.NewGORM()
	if err := db.Where("mac_id = ?", macID).Delete(&model.Kid{}).Error; err != nil {
		logError(errors.Wrapf(err, "Error on deleting kid. Mac ID: %s", macID))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when deleting kid from database",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
