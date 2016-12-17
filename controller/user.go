package controller

import (
	"net/http"

	"log"

	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

func Login(c *gin.Context) {
	var json model.Login

	if c.BindJSON(&json) != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	log.Printf("\nEmail: %s, Password:%s\n", json.Email, json.Password)
	db := database.New()
	defer db.Close()
	var user model.User

	json.Password = EncryptPassword(json.Password)

	err := db.Get(&user,
		"SELECT email, first_name, last_name, zip_code, last_updated, date_created FROM user WHERE email=? and password=? LIMIT 1",
		json.Email,
		json.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	log.Printf("\nUser login request. User: %#v\n", user)

	accessToken := model.AccessToken{
		Email: user.Email,
		Token: randToken(),
	}

	success := storeToken(db, accessToken)

	if !success {
		log.Println("Store token fail!!!!")
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

func storeToken(db *sqlx.DB, accessToken model.AccessToken) bool {
	var existToken model.AccessToken
	err := db.Get(&existToken, "SELECT email, token, last_updated FROM authentication_token WHERE email = ?", accessToken.Email)

	var result sql.Result
	if err != nil {
		result = db.MustExec("INSERT INTO authentication_token (email, token, date_created, last_updated) VALUES (?,?, Now(), Now())",
			accessToken.Email,
			accessToken.Token)
	} else {
		result = db.MustExec("UPDATE authentication_token SET token = ?, last_updated = NOW(), access_count = access_count + 1 WHERE email = ?",
			accessToken.Token,
			accessToken.Email)

	}

	return checkInsertResult(result)

}

func Register(c *gin.Context) {
	var registerRequest model.Register
	if err := c.BindJSON(&registerRequest); err != nil {
		log.Printf("Register Error: %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing some of required paramters.",
			"error":   err,
		})
		return
	}

	db := database.New()
	defer db.Close()

	var exist bool
	if err := db.Get(&exist, "SELECT EXISTS(SELECT id FROM user WHERE email = ? LIMIT 1)", registerRequest.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong on server side",
			"error":   err,
		})
		return
	}

	if exist {
		c.JSON(http.StatusConflict, gin.H{
			"message": "The email is already registered",
		})
		return
	}

	registerRequest.Password = EncryptPassword(registerRequest.Password)

	result, err := db.NamedExec("INSERT INTO user (email, password, first_name, last_name, phone_number, zip_code, date_created, last_updated) VALUES"+
		" (:email, :password, :first_name, :last_name, :phone_number, :zip_code, Now(), Now())",
		registerRequest)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when insert User to database",
			"error":   err,
		})
		return
	}

	if checkInsertResult(result) {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "The data was not able to write to database. No error",
		})
	}

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
