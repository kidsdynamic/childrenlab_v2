package controller

import (
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"fmt"
	"io"
	"log"

	"net/http"

	"bytes"

	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

const (
	SignedUserKey = "SignedUser"
	S3ProfilePath = "userProfile"
)

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%x", b))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func checkInsertResult(result sql.Result) bool {
	_, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func getInsertedID(result sql.Result) int64 {
	ID, err := result.LastInsertId()

	if err != nil {
		log.Println(err)
		return -1
	}

	return ID
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
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{})
		c.Abort()
		return
	}

	log.Printf("\nLogged in user: %#v\n", user)
	c.Set(SignedUserKey, user)

	c.Next()

}

func GetSignedInUser(c *gin.Context) *model.User {
	signedUser, ok := c.Get(SignedUserKey)

	if !ok {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "can't find login user",
		})
		c.Abort()
		return nil
	}

	user := signedUser.(model.User)
	return &user
}

func GetUserByID(db *gorm.DB, id int64) (model.User, error) {
	var user model.User

	err := db.Where("id = ?", id).First(&user).Error

	return user, err
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

func UploadFileToS3(file *os.File, fileName string) error {

	ss, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	_, err = ss.Config.Credentials.Get()
	if err != nil {
		log.Fatal(err)
	}
	svc := s3.New(session.New(&aws.Config{}))

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()

	buffer := make([]byte, size)
	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	uploadResult, err := svc.PutObject(&s3.PutObjectInput{
		Body:          fileBytes,
		Bucket:        aws.String(model.AwsConfig.Bucket),
		Key:           aws.String(fmt.Sprintf("/userProfile/%s", fileName)),
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
		ACL:           aws.String("public-read"),
	})
	if err != nil {
		log.Printf("Failed to upload data to %s\n", err)
		return err
	}

	log.Printf("Response: %s\n", awsutil.StringValue(uploadResult))

	return nil

}

func GetKidsByUser(user *model.User) ([]model.Kid, error) {
	db := database.NewGORM()
	defer db.Close()
	var kids []model.Kid

	//err := db.Select(&kids, "SELECT id, first_name, last_name, mac_id, kids.date_created, mac_id, profile FROM kids WHERE parent_id = ?", user.ID)

	err := db.Where("parent_id = ?", user.ID).Find(&kids).Error

	return kids, err
}

func GetDeviceByMacID(db *sqlx.DB, macId string) (model.Device, error) {
	var device model.Device
	err := db.Get(&device, "SELECT id, mac_id, date_created FROM device WHERE mac_id = ?", macId)

	return device, err
}

func GetUserRole(db *gorm.DB) model.Role {
	var role model.Role
	if err := db.Where("authority = ?", "ROLE_USER").First(&role).Error; err != nil {
		panic(err)
	}

	return role
}
