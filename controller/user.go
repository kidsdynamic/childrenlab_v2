package controller

import (
	"net/http"

	"log"

	"fmt"

	"time"

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
		LastUpdated: time.Now(),
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

	//set user role
	role := GetUserRole(db)
	user.Role = role
	user.Email = userRequest.Email
	user.Password = userRequest.Password
	user.DateCreated = time.Now()
	user.FirstName = userRequest.FirstName
	user.LastName = userRequest.LastName
	user.PhoneNumber = userRequest.PhoneNumber
	user.ZipCode = userRequest.ZipCode

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

func IsTokenValid(c *gin.Context) {
	var tokenRequest model.TokenRequest

	if c.BindJSON(&tokenRequest) != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	var existToken model.AccessToken
	db := database.New()
	defer db.Close()
	err := db.Get(&existToken, "SELECT email, token, last_updated FROM authentication_token WHERE email = ? AND token = ?",
		tokenRequest.Email,
		tokenRequest.Token)
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

	db := database.New()
	defer db.Close()

	tx := db.MustBegin()
	if request.FirstName != "" {
		tx.MustExec("UPDATE user SET first_name = ? WHERE id = ?", request.FirstName, signedInUser.ID)
	}

	if request.LastName != "" {
		tx.MustExec("UPDATE user SET last_name = ? WHERE id = ?", request.LastName, signedInUser.ID)
	}

	if request.PhoneNumber != "" {
		tx.MustExec("UPDATE user SET phone_number = ? WHERE id = ?", request.PhoneNumber, signedInUser.ID)
	}

	if request.ZipCode != "" {
		tx.MustExec("UPDATE user SET zip_code = ? WHERE id = ?", request.ZipCode, signedInUser.ID)
	}

	tx.MustExec("UPDATE user SET last_updated = NOW() WHERE id = ?", signedInUser.ID)
	tx.Commit()

	user, err := GetUserByID(db, signedInUser.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retreive updated user information",
		})
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

	db := database.New()
	defer db.Close()

	var exist bool
	if err := db.Get(&exist, "SELECT EXISTS(SELECT id FROM user WHERE email = ? LIMIT 1)", email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong on server side",
			"error":   err,
		})
		return
	}

	if exist {
		c.JSON(http.StatusConflict, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

type iOS struct {
	RegistrationId string
}

func UpdateIOSRegistrationId(c *gin.Context) {
	var ios iOS

	err := c.BindJSON(&ios)

	if err != nil {
		log.Printf("Error on UpdateIosRegistrationId: Error: %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	db := database.New()
	defer db.Close()

	user := GetSignedInUser(c)

	if _, err := db.Exec("UPDATE user SET registration_id = ? WHERE id = ?", ios.RegistrationId, user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong on server side",
			"error":   err,
		})
		return
	}

	updatedUser, _ := GetUserByID(db, user.ID)

	c.JSON(http.StatusOK, updatedUser)
}
