package controller

import (
	"net/http"

	"log"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

func Login(c *gin.Context) {
	var json model.Login

	if c.BindJSON(&json) != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	log.Printf("\nEmail: %s, Password:%s, Line: %d\n", json.Email, json.Password, log.LstdFlags)
	db := database.NewGORM()
	defer db.Close()
	var user model.User

	json.Password = database.EncryptPassword(json.Password)

	db.Where("email = ? AND password = ?", json.Email, json.Password).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	log.Printf("\nUser login request. User: %#v\n", user)

	accessToken := model.AccessToken{
		Email:       user.Email,
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

func storeToken(db *gorm.DB, accessToken model.AccessToken) error {
	var existToken model.AccessToken
	db.Where("email = ?", accessToken.Email).First(&existToken)

	var err error
	if existToken.ID != 0 {
		existToken.LastUpdated = accessToken.LastUpdated
		existToken.Token = accessToken.Token
		existToken.AccessCount += 1
		err = db.Save(&existToken).Error
	} else {
		err = db.Create(&accessToken).Error

	}

	return err

}

func Register(c *gin.Context) {
	var userRequest model.RegisterRequest
	if err := c.BindJSON(&userRequest); err != nil {
		log.Printf("Register Error: %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing some of required paramters.",
			"error":   err,
		})
		return
	}
	log.Printf("\nEmail: %s, Password:%s, Line: %d\n", userRequest.Email, userRequest.Password, log.LstdFlags)
	db := database.NewGORM()
	defer db.Close()

	var user model.User

	db.Where("email = ?", userRequest.Email).First(&user)

	if user.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{
			"message": "The email is already registered",
		})
		return
	}

	userRequest.Password = database.EncryptPassword(userRequest.Password)

	//Set default language
	if userRequest.Language == "" {
		userRequest.Language = "en"
	}

	//set user role
	role := GetUserRole(db)
	user.Role = role
	user.Email = userRequest.Email
	user.Password = userRequest.Password
	user.DateCreated = GetNowTime()
	user.LastUpdated = GetNowTime()
	user.FirstName = userRequest.FirstName
	user.LastName = userRequest.LastName
	user.PhoneNumber = userRequest.PhoneNumber
	user.ZipCode = userRequest.ZipCode
	user.Language = userRequest.Language

	if err := db.Create(&user).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when insert User to database",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, user)

}

type LanguageRequest struct {
	Language string `json:"language"`
}

func UpdateLanguage(c *gin.Context) {
	var languageRequest LanguageRequest

	if err := c.BindJSON(&languageRequest); err != nil {
		log.Printf("Register Error: %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing some of required paramters.",
			"error":   err,
		})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	user := GetSignedInUser(c)

	if err := db.Model(&model.User{}).Where("id = ?", user.ID).UpdateColumn("language", languageRequest.Language).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when update user language",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func IsTokenValid(c *gin.Context) {
	var tokenRequest model.TokenRequest

	tokenRequest.Email = c.Query("email")
	tokenRequest.Token = c.Query("token")

	var existToken model.AccessToken
	db := database.NewGORM()
	defer db.Close()

	err := db.Where("email = ? AND token = ?", tokenRequest.Email, tokenRequest.Token).First(&existToken).Error

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

func UpdateProfile(c *gin.Context) {
	var request model.ProfileUpdateRequest

	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	fmt.Printf("Profile Update Request: %#v", request)

	signedInUser := GetSignedInUser(c)

	db := database.NewGORM()
	defer db.Close()

	var user model.User
	if err := db.Where("id = ?", signedInUser.ID).First(&user).Error; err != nil {
		log.Printf("Error on retrieve user from udpate Profile. Error: %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if err := db.Model(&user).Updates(request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retreive updated user information",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UserProfile(c *gin.Context) {
	user := GetSignedInUser(c)

	kids, err := GetKidsByUser(user)
	if err != nil {
		fmt.Printf("Kids error: %#v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error when retrieve kids",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
		"kids": kids,
	})

}

func IsEmailAvailableToRegister(c *gin.Context) {
	email := c.Query("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	var user model.User
	db.Where("email = ?", email).First(&user)

	if user.Email != "" {
		c.JSON(http.StatusConflict, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func FindUserByEmail(c *gin.Context) {
	email := c.Query("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	var user model.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error on finding user by email",
				"error":   err,
			})

		}
		return
	}

	c.JSON(http.StatusOK, user)
}

type PushNotificationID struct {
	RegistrationId string
}

func UpdateIOSRegistrationId(c *gin.Context) {
	var ios PushNotificationID

	err := c.BindJSON(&ios)

	if err != nil {
		log.Printf("Error on UpdateIosRegistrationId: Error: %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	user := GetSignedInUser(c)

	user.RegistrationID = ios.RegistrationId

	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong on server side",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateAndroidRegistrationId(c *gin.Context) {
	var android PushNotificationID

	err := c.BindJSON(&android)

	if err != nil {
		log.Printf("Error on UpdateIosRegistrationId: Error: %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	user := GetSignedInUser(c)

	user.AndroidRegistrationToken = android.RegistrationId

	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong on server side",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
