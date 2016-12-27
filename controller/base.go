package controller

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
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
	"github.com/jmoiron/sqlx"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

const (
	SignedUserKey = "SignedUser"
	S3ProfilePath = "userProfile"
	TimeLayout    = "2006-01-02T15:04:05"
)

func EncryptPassword(password string) string {
	h := sha256.New()
	io.WriteString(h, password)
	fmt.Printf("\n%x\n", h.Sum(nil))

	return fmt.Sprintf("%x", h.Sum(nil))

}

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
	log.Printf("TOKEN: %s", authToken)
	if authToken == "" {
		c.JSON(http.StatusForbidden, gin.H{})
		c.Abort()
		return
	}

	db := database.New()
	defer db.Close()

	var user model.User
	err := db.Get(&user, "SELECT u.id, u.email, COALESCE(first_name, '') as first_name, COALESCE(last_name, '') as last_name "+
		", u.date_created, COALESCE(zip_code, '') as zip_code, u.last_updated FROM user u join "+
		"authentication_token a ON u.email = a.email WHERE token = ?", authToken)

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

func GetUserByID(db *sqlx.DB, id int64) (model.User, error) {
	var user model.User

	err := db.Get(&user, "SELECT id, email, COALESCE(first_name, '') as first_name, COALESCE(last_name, '') as last_name "+
		", date_created, COALESCE(zip_code, '') as zip_code, last_updated, COALESCE(phone_number, '') as phone_number FROM user WHERE id = ?", id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func GetKidByUserIdAndKidId(db *sqlx.DB, userId, kidId int64) (model.Kid, error) {
	var kid model.Kid
	err := db.Get(&kid, "SELECT k.id, COALESCE(k.first_name, '') as first_name, COALESCE(k.last_name, '') as last_name, "+
		"k.date_created FROM user u JOIN kids k ON u.id = k.parent_id  WHERE u.id = ? AND k.id = ?", userId, kidId)

	if err != nil {
		return kid, err
	}

	return kid, nil
}

func UploadFileToS3(file *os.File, fileName, bucketName string) error {

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
		Bucket:        aws.String(bucketName),
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
	db := database.New()
	defer db.Close()
	var kids []model.Kid

	err := db.Select(&kids, "SELECT kids.id, first_name, last_name, kids.date_created, mac_id FROM kids JOIN device on"+
		" device.kid_id = kids.id WHERE parent_id = ?", user.ID)

	return kids, err
}

func GetDeviceByMacID(db *sqlx.DB, macId string) (model.Device, error) {
	var device model.Device
	err := db.Get(&device, "SELECT id, mac_id, date_created FROM device WHERE mac_id = ?", macId)

	return device, err
}
