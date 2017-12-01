package controller

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"log"

	"net/http"

	"time"

	"net/smtp"
	"strconv"

	"strings"

	"github.com/pkg/errors"

	"encoding/json"

	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kidsdynamic/childrenlab_v2/config"
	"github.com/kidsdynamic/childrenlab_v2/constants"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

const (
	SignedUserKey = "SignedUser"
)

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%x", b))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Auth(c *gin.Context) {
	authToken := c.Request.Header.Get("x-auth-token")
	if authToken == "" {
		c.JSON(http.StatusForbidden, gin.H{})
		c.Abort()
		return
	}

	db := database.NewGORM()
	defer db.Close()

	var user model.User

	err := db.Table("user").Joins("JOIN authentication_token a ON user.email = a.email").Where("a.token = ?", authToken).Find(&user).Error

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{})
		c.Abort()
		return
	}

	c.Set(SignedUserKey, user)

	c.Next()

}

func GetSignedInUser(c *gin.Context) model.User {
	var user model.User
	signedUser, ok := c.Get(SignedUserKey)

	if !ok {
		return user
	}

	user = signedUser.(model.User)
	return user
}

func GetKidByUserIdAndKidId(db *gorm.DB, userId, kidId int64) (model.Kid, error) {
	var kid model.Kid

	err := db.Where("parent_id = ? AND id = ?", userId, kidId).Preload("Parent").Find(&kid).Error
	return kid, err
}

func GetKidByMacID(db *gorm.DB, macID string) (model.Kid, error) {
	var kid model.Kid

	err := db.Where("mac_id = ?", macID).Preload("Parent").First(&kid).Error
	return kid, err
}

func GetKidsByUser(user model.User) ([]model.Kid, error) {
	db := database.NewGORM()
	defer db.Close()
	var kids []model.Kid

	err := db.Where("parent_id = ?", user.ID).Find(&kids).Error
	if err == gorm.ErrRecordNotFound {
		return kids, nil
	}

	return kids, err
}

func GetUserRole(db *gorm.DB) model.Role {
	var role model.Role
	if err := db.Where("authority = ?", "ROLE_USER").First(&role).Error; err != nil {
		panic(err)
	}

	return role
}

func GetNowTime() time.Time {
	now := time.Now()

	s := fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02dZ", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	t, err := time.Parse(constants.TimeLayout, s)

	if err != nil {
		logError(errors.Wrap(err, "Error on get now time"))

	}

	return t
}

func HasPermissionToKid(db *gorm.DB, user *model.User, kidID []int64) bool {

	var exists bool = true
	for _, id := range kidID {
		row := db.Raw("SELECT EXISTS(SELECT id FROM kids WHERE id = ? and parent_id = ? LIMIT 1)", id, user.ID).Row()

		row.Scan(&exists)

		if !exists {
			subhostRow := db.Raw("SELECT EXISTS(SELECT id FROM sub_host s JOIN sub_host_kid sk ON s.id = sk.sub_host_id WHERE s.request_from_id = ? and sk.kid_id = ? and s.status = ? LIMIT 1)", user.ID, id, SubHostStatusAccepted).Row()

			subhostRow.Scan(&exists)
			if !exists {
				return false
			}
		}
	}

	return exists
}

func LogUserActivity(db *gorm.DB, user *model.User, action string, macID *string) {
	/*	logAction := &model.LogUserAction{
			User:        user,
			UserID:      user.ID,
			MacID:       macID,
			Action:      action,
			DateCreated: time.Now(),
			LastUpdated: time.Now(),
		}

		if err := db.Create(logAction).Error; err != nil {
			logError(errors.Wrap(err, "Error on the log user action"))
			return
		}*/
}

func logError(err error) {
	log.Printf("Error occur: \n%+v", err)

	if config.ServerConfig.Debug != true {
		emailUser := &EmailUser{
			Username:    config.ServerConfig.EmailAuthName,
			Password:    config.ServerConfig.EmailAuthPassword,
			EmailServer: config.ServerConfig.EmailServer,
			Port:        config.ServerConfig.EmailPort,
		}
		body := fmt.Sprintf("%+v", err)
		body = strings.Replace(body, "\n", "<br/>", -1)
		sendMail(emailUser, config.ServerConfig.ErrorLogEmail, fmt.Sprintf("Server Error: %s", config.ServerConfig.BaseURL), body)
	}

}

type IPResp struct {
	Status      string
	Country     string
	CountryCode string
	Region      string
	regionName  string
	City        string
	Zip         string
	TimeZone    string
}

func getDetailFromIP(ip string) *IPResp {
	response, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return nil
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logError(errors.Wrap(err, "Error on reading IP API response body"))
		return nil

	}
	var ipRes *IPResp
	if err := json.Unmarshal(body, &ipRes); err != nil {
		logError(err)
		return nil
	}

	return ipRes

}

type EmailUser struct {
	Username    string
	Password    string
	EmailServer string
	Port        int
}

func sendMail(emailUser *EmailUser, toEmail, subject, message string) error {

	auth := smtp.PlainAuth(
		"Swing",
		emailUser.Username,
		emailUser.Password,
		emailUser.EmailServer,
	)

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf("TO: %s\r\n"+
		"Subject: %s\r\n%s"+
		"\r\n%s", toEmail, subject, mime, message)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		emailUser.EmailServer+":"+strconv.Itoa(emailUser.Port),
		auth,
		emailUser.Username,
		[]string{toEmail},
		[]byte(body),
	)
	return err
}
