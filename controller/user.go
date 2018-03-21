package controller

import (
	"net/http"

	"log"

	"fmt"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kidsdynamic/childrenlab_v2/config"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

func Login(c *gin.Context) {
	var json model.Login

	if c.BindJSON(&json) != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	db := database.NewGORM()
	defer db.Close()
	var user model.User

	json.Password = database.EncryptPassword(json.Password)

	db.Where("email = ? AND password = ?", json.Email, json.Password).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	accessToken := model.AccessToken{
		Email:       user.Email,
		Token:       randToken(),
		LastUpdated: GetNowTime(),
	}

	err := storeToken(db, accessToken)

	if err != nil {
		logError(errors.Wrapf(err, "Error on login: %#v", json))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Store token failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username":     accessToken.Email,
		"access_token": accessToken.Token,
	})

	if user.SignUpIP == "" {
		ipDetail := getDetailFromIP(c.ClientIP())
		if ipDetail != nil && ipDetail.Country != "" {
			user.SignUpCountryCode = ipDetail.CountryCode
			user.SignUpIP = c.ClientIP()
		}
	}

	LogUserActivity(db, &user, "Logged in", nil)
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

	ipDetail := getDetailFromIP(c.ClientIP())
	if ipDetail != nil && ipDetail.Country != "" {
		user.SignUpCountryCode = ipDetail.CountryCode
		user.SignUpIP = c.ClientIP()
	}

	if err := db.Create(&user).Error; err != nil {
		logError(errors.Wrapf(err, "Error on creating user: %#v", user))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when insert User to database",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, user)
	LogUserActivity(db, &user, fmt.Sprintf("Register (%d)", user.ID), nil)
}

type LanguageRequest struct {
	Language string `json:"language"`
}

func UpdateLanguage(c *gin.Context) {
	var languageRequest LanguageRequest

	if err := c.BindJSON(&languageRequest); err != nil {

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
		logError(errors.Wrapf(err, "Error on update language: %#v", languageRequest))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when update user language",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

	LogUserActivity(db, &user, fmt.Sprintf("User - Update Language (%d)", user.ID), nil)
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

	var user model.User
	db.Table("user").Joins("JOIN authentication_token a ON user.email = a.email").Where("a.token = ?", tokenRequest.Token).Find(&user)

	LogUserActivity(db, &user, "User - Is Token Valid (%d)", nil)

}

func UpdateProfile(c *gin.Context) {
	var request model.ProfileUpdateRequest

	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	signedInUser := GetSignedInUser(c)

	db := database.NewGORM()
	defer db.Close()

	var user model.User
	if err := db.Where("id = ?", signedInUser.ID).First(&user).Error; err != nil {
		logError(errors.Wrapf(err, "Error on retrieve user from udpate Profile: %#v", signedInUser))
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if err := db.Model(&user).Updates(request).Error; err != nil {
		logError(errors.Wrapf(err, "Error on updating profile: %#v", request))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retreive updated user information",
		})
		return
	}

	c.JSON(http.StatusOK, user)

	LogUserActivity(db, &user, fmt.Sprintf("User - Update Profile (%d)", user.ID), nil)
}

func UserProfile(c *gin.Context) {
	user := GetSignedInUser(c)

	kids, err := GetKidsByUser(user)
	if err != nil {
		logError(errors.Wrapf(err, "Error on retrive kids by user: %#v", user))
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

	ipDetail := getDetailFromIP(c.ClientIP())
	message := "System upgrading now. Please login again at a later time. Thank you"
	if ipDetail != nil && ipDetail.Country != "" {
		if ipDetail.CountryCode == "CN" || ipDetail.CountryCode == "TW" {
			message = "系統更新中。請稍晚再重新登入。"
		} else if ipDetail.CountryCode == "JP" {
			message = "現在システムの更新中です。本日（3月21日）の午後８時以降に再度ログインしていただけますよう、お願い申し上げます。"
		} else if ipDetail.CountryCode == "RU" {
			message = "Обновление системы сейчас. Повторите попытку позже. Спасибо"
		} else if ipDetail.CountryCode == "ES" {
			message = "Actualización del sistema ahora. Por favor inicie sesión nuevamente en otro momento. Gracias"
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
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
			logError(errors.Wrapf(err, "Error on finding user by email: %#v", email))
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
	RegistrationId string `json:"registrationId"`
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
		logError(errors.Wrapf(err, "Error on saving user registration ID: %#v", ios))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong on server side",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, user)

	LogUserActivity(db, &user, fmt.Sprintf("User - Update IOS Registration ID (%d)", user.ID), nil)
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

	if err := db.Model(&user).Update("android_registration_token", user.AndroidRegistrationToken).Error; err != nil {
		logError(errors.Wrapf(err, "Error on update android registration ID: %#v", android))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong on server side",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, user)

	LogUserActivity(db, &user, fmt.Sprintf("User - Update Android ID (%d)", user.ID), nil)
}

func SendResetPasswordEmail(c *gin.Context) {

	db := database.NewGORM()
	defer db.Close()

	var user model.User

	authToken := c.Request.Header.Get("x-auth-token")
	if authToken != "" {
		fmt.Println(authToken)
		if err := db.Model(&user).Joins("JOIN authentication_token a ON user.email = a.email").Where("a.token = ?", authToken).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		fmt.Printf("Email: %#v", user)

	} else {
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		if user.Email != "" {
			if err := db.Model(&user).Where("email = ?", user.Email).First(&user).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{})
				return
			}
		}
	}

	token := randToken()

	if err := db.Model(&user).Where("id = ?", user.ID).Update("reset_password_token", token).Error; err != nil {
		logError(errors.Wrapf(err, "Error on update reset password token: %#v", token))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on updating password token",
			"error":   err,
		})
		return
	}

	emailUser := &EmailUser{
		Username:    config.ServerConfig.EmailAuthName,
		Password:    config.ServerConfig.EmailAuthPassword,
		EmailServer: config.ServerConfig.EmailServer,
		Port:        config.ServerConfig.EmailPort,
	}

	htmlBody := `
		<html>
		<body style="margin: 0; padding: 0; padding: 50px 20px; text-align: center; background-color: #C4ECF6; color: #FF6A23;">
		 <table border="0" cellpadding="0" cellspacing="0" width="100%%">
		  <tr>
		   <td>
		    <h2>Reset Swing password</h2>
		   </td>
		  </tr>
		  <tr>
		    <td>
		      <h2>Hi %s %s,</h2>
		      <h3>You have recently requested to reset your Cacoo password. Set a new password here:</h3>
		    </td>
		  </tr>
		  <tr>
		    <td align="center" style="height: 50px">
		    	<a href="%s" style="text-decoration: none;line-height: 100%%; background: #FD733D; color: white; font-family: Open Sans,Helvetica,Arial,sans-serif; font-size: 20px; font-weight: bold; text-transform: none; margin: 0px;padding: 10px 50px; border-radius: 9px;">
		    	  Reset Password
		    	</a>
		    </td>
		  </tr>
		  <tr>
		    <td style="margin-top: 20px; text-align: right;">
		      <p>Team Swing</p>
		    </td>
		  </tr>

		 </table>
		</body>
		</html>
	`
	resetPasswordURL := fmt.Sprintf("%s/v1/user/resetPasswordPage?token=%s&email=%s", config.ServerConfig.BaseURL, token, user.Email)
	emailBody := fmt.Sprintf(htmlBody, user.FirstName, user.LastName, resetPasswordURL)

	if err := sendMail(emailUser, user.Email, "Reset your Swing password", emailBody); err != nil {
		logError(errors.Wrap(err, "Sending Reset Password Error"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Please try again later",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

func ResetPasswordPage(c *gin.Context) {
	token := c.Query("token")
	email := c.Query("email")

	if token == "" || email == "" {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	db := database.NewGORM()
	defer db.Close()

	var user model.User
	if err := db.Model(&model.User{}).Where("reset_password_token = ? and email = ?", token, email).First(&user).Error; err != nil {
		logError(errors.Wrapf(err, "Error on reset password page: %#v", email))
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	c.HTML(http.StatusOK, "reset_password", gin.H{
		"user":         user,
		"errorMessage": "",
	})
}

func ResetPassword(c *gin.Context) {
	token := c.PostForm("token")
	email := c.PostForm("email")
	password := c.PostForm("password")

	if token == "" || email == "" || password == "" {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	db := database.NewGORM()
	defer db.Close()

	var user model.User
	if err := db.Model(&model.User{}).Where("reset_password_token = ? and email = ?", token, email).First(&user).Error; err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	if len(password) < 6 {
		c.HTML(http.StatusBadRequest, "reset_password", gin.H{
			"user":         user,
			"errorMessage": "The password has to be longer than <strong>6</strong> characters",
		})
		return
	}

	password = database.EncryptPassword(password)

	if err := db.Exec("UPDATE user SET password = ?, reset_password_token = null WHERE email = ? and reset_password_token = ?", password, email, token).Error; err != nil {
		logError(errors.Wrapf(err, "Error on reset user password : %#v %#v", email, password))
		c.HTML(http.StatusBadRequest, "reset_password", gin.H{
			"user":         user,
			"errorMessage": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "reset_password_success", nil)
}

func GetUserByEmail(c *gin.Context) {
	email := c.Query("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email is missing",
			"error":   "",
		})
		return
	}

	db := database.NewGORM()
	defer db.Close()
	var user model.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "No record",
			})
			return
		} else {
			logError(errors.Wrapf(err, "Error on search user by email : %s", email))
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error on search email",
				"error":   err,
			})
			return
		}

	}

	c.JSON(http.StatusOK, user)
}

func UpdatePassword(c *gin.Context) {

	var newPassword model.UpdatePasswordReq
	if err := c.BindJSON(&newPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The password has to be longer than 6 characters",
		})
		return
	}

	if newPassword.NewPassword == "" || len(newPassword.NewPassword) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The password has to be longer than 6 characters",
		})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	user := GetSignedInUser(c)
	hashedPassword := database.EncryptPassword(newPassword.NewPassword)

	if err := db.Exec("UPDATE user SET password = ?, reset_password_token = null WHERE email = ?", hashedPassword, user.Email).Error; err != nil {
		logError(errors.Wrapf(err, "Error on reset user password : %s", user.Email))
		c.HTML(http.StatusBadRequest, "reset_password", gin.H{
			"user":         user,
			"errorMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Logout(c *gin.Context) {
	user := GetSignedInUser(c)

	db := database.NewGORM()
	defer db.Close()

	var auth model.AccessToken

	if err := db.Where("email = ?", user.Email).First(&auth).Error; err != nil {
		logError(errors.Wrapf(err, "Error on retrieve access auth by email : %s", user.Email))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on retrieve auth",
			"error":   err,
		})
		return
	}

	if err := db.Delete(&auth).Error; err != nil {
		logError(errors.Wrapf(err, "Error on delete access auth by email : %s", user.Email))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on delete auth",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func MyCountryCode(c *gin.Context) {
	user := GetSignedInUser(c)

	db := database.NewGORM()
	defer db.Close()

	var countryCode []string
	if err := db.Model(&model.User{}).Where("id = ?", user.ID).Pluck("sign_up_country_code", &countryCode).Error; err != nil {
		logError(errors.Wrapf(err, "Error on access user country code : %s", user.Email))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on retrieve country code",
			"error":   err,
		})
		return
	}

	if len(countryCode) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"countryCode": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"countryCode": countryCode[0],
	})
}
