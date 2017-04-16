package controller

import (
	"net/http"

	"github.com/kidsdynamic/childrenlab_v2/database"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/model"
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

	db.Table("user").Joins("JOIN role ON user.role_id = role.id").Where("email = ? AND password = ? and authority = ?", adminLogin.Name, adminLogin.Password, model.ROLE_ADMIN).First(&admin)

	if admin.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	log.Printf("\nUser login request. Admin: %#v\n", admin)

	accessToken := model.AccessToken{
		Email:       admin.Email,
		Token:       randToken(),
		LastUpdated: GetNowTime(),
	}

	err := storeToken(db, accessToken)

	if err != nil {
		log.Printf("Store token fail!!!! ERROR: %#v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Store token failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username":     accessToken.Email,
		"access_token": accessToken.Token,
	})
}

func GetAllUser(c *gin.Context) {
	db := database.NewGORM()
	defer db.Close()

	var userList []model.User
	if err := db.Joins("JOIN role ON role.id = user.role_id").Where("role.authority = 'ROLE_USER'").Order("date_created desc").Limit(50).Find(&userList).Error; err != nil {
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
	db.Table("user").Joins("JOIN role ON user.role_id = role.id").Where("email = ? AND password = ? and authority = ?", request.Name, request.Password, model.ROLE_ADMIN).First(&admin)

	if admin.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{
			"message": "The admin name exists",
		})
		return
	}

	var role model.Role
	if err := db.Where("authority = ?", model.ROLE_ADMIN).First(&role).Error; err != nil {
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on creating admin user",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
