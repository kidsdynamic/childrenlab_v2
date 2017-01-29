package controller

import (
	"fmt"
	"io"
	"log"
	"os"

	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

func UploadAvatar(c *gin.Context) {
	user := GetSignedInUser(c)
	file, _, err := c.Request.FormFile("upload")
	fileName := fmt.Sprintf("avatar_%d.jpg", user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "upload parameter is required",
			"error":   err,
		})
		return
	}

	if os.MkdirAll("./tmp", 0755) != nil {

		panic("Unable to create directory for tagfile!")

	}

	out, err := os.OpenFile("./tmp/"+fileName, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Println(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Println(err)
	}

	f, err := os.Open("./tmp/" + fileName)
	if err != nil {
		log.Println(err)
	}

	if err = UploadFileToS3(f, fileName); err == nil {
		db := database.NewGORM()
		defer db.Close()

		if err := db.Model(&model.User{}).Update("profile", fileName).Where("id = ?", user.ID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something wrong when updating profile for the user",
				"error":   err,
			})
		}

		user.Profile = fileName

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	} else {
		fmt.Printf("Error on upload user image to S3. Error: %#v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on upload image to S3",
			"error":   err,
		})
	}

}

func UploadKidAvatar(c *gin.Context) {
	user := GetSignedInUser(c)
	file, _, err := c.Request.FormFile("upload")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "upload parameter is required",
			"error":   err,
		})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	kidId, err := strconv.ParseInt(c.PostForm("kidId"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on parse KidId",
			"error":   err,
		})
		return
	}

	kid, err := GetKidByUserIdAndKidId(db, user.ID, kidId)

	if err != nil {
		log.Printf("Error on get kid from database. Error: %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on Get kid from database",
			"error":   err,
		})
		return
	}

	fileName := fmt.Sprintf("kid_avatar_%d.jpg", kid.ID)
	if err != nil {
		fmt.Printf("err opening file: %s", err)
	}

	if os.MkdirAll("./tmp", 0755) != nil {

		panic("Unable to create directory for tagfile!")

	}

	out, err := os.OpenFile("./tmp/"+fileName, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Println(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Println(err)
	}

	f, err := os.Open("./tmp/" + fileName)
	if err != nil {
		log.Println(err)
	}
	if UploadFileToS3(f, fileName) == nil {

		if err := db.Model(&kid).Update("profile", fileName); err != nil {
			log.Printf("Error on update profile. Error: %#v", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"kid": kid,
		})
	}

}
