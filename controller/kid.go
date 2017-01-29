package controller

import (
	"net/http"

	"log"

	"fmt"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

func AddKid(c *gin.Context) {
	user := GetSignedInUser(c)

	var request model.KidRequest

	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Reqeust",
			"error":   err,
		})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	var kid model.Kid

	db.Where("mac_id = ?", request.MacID).First(&kid)

	if kid.MacID != "" {
		c.JSON(http.StatusConflict, gin.H{
			"message": "The device is already registered",
		})
		return
	}

	kid.MacID = request.MacID
	kid.FirstName = request.FirstName
	kid.LastName = request.LastName
	kid.ParentID = user.ID
	kid.DateCreated = time.Now()

	if err := db.Save(&kid).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when insert kid data",
			"error":   err,
		})
		return
	}

	var device model.Device
	device.MacID = kid.MacID
	device.DateCreated = time.Now()

	if err := db.Save(&device).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error when insert device data",
			"error":   err,
		})
		return
	}

	if err := db.Preload("Parent").Find(&kid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error when Preload parent device data",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, kid)

}

func UpdateKid(c *gin.Context) {
	var request model.UpdateKidRequest

	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	fmt.Printf("Kid Update Request: %#v", request)

	db := database.NewGORM()
	defer db.Close()

	user := GetSignedInUser(c)
	kid, err := GetKidByUserIdAndKidId(db, user.ID, request.ID)

	if err != nil {
		fmt.Printf("Can't find kid. %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can't find kid",
			"error":   err,
		})
		return
	}

	if err := db.Model(&kid).Updates(request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retreive updated user information",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"kid": kid,
	})

}
