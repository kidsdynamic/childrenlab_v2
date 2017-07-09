package controller

import (
	"net/http"

	"log"

	"fmt"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kidsdynamic/childrenlab_v2/app/database"
	"github.com/kidsdynamic/childrenlab_v2/app/model"
)

func AddKid(c *gin.Context) {
	user := GetSignedInUser(c)

	var request model.KidRequest

	if err := c.BindJSON(&request); err != nil {
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
	kid.Name = request.Name
	kid.ParentID = user.ID
	kid.DateCreated = GetNowTime()

	if err := db.Save(&kid).Error; err != nil {
		log.Println(err)
		logError(errors.Wrapf(err, "Error on saving kid: %#v", kid))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when insert kid data",
			"error":   err,
		})
		return
	}

	var device model.Device
	device.MacID = kid.MacID
	device.DateCreated = GetNowTime()

	if err := db.Save(&device).Error; err != nil {
		logError(errors.Wrapf(err, "Error when insert device data Device: %#v", device))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error when insert device data",
			"error":   err,
		})
		return
	}

	if err := db.Preload("Parent").Find(&kid).Error; err != nil {
		logError(errors.Wrapf(err, "Error when Preload parent kidID: %#v", kid))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error when Preload parent device data",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, kid)

	LogUserActivity(db, &user, fmt.Sprintf("Added Kid (%d)", kid.ID), &kid.MacID)
}

func UpdateKid(c *gin.Context) {
	var request model.UpdateKidRequest

	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	user := GetSignedInUser(c)
	kid, err := GetKidByUserIdAndKidId(db, user.ID, request.ID)

	if err != nil {
		logError(errors.Wrapf(err, "Can't find kid: %#v", request.ID))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can't find kid",
			"error":   err,
		})
		return
	}

	if request.Name != "" {
		kid.Name = request.Name

		if err := db.Model(&model.Kid{}).Where("id = ?", kid.ID).Update("name", kid.Name).Error; err != nil {
			logError(errors.Wrapf(err, "Error on update kid name. Kid: %#v", kid))
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something wrong when retreive updated user information",
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"kid": kid,
	})
	LogUserActivity(db, &user, fmt.Sprintf("Update Kid (%d)", kid.ID), &kid.MacID)
}

func DeleteKid(c *gin.Context) {
	kidID := c.Query("kidId")

	user := GetSignedInUser(c)

	db := database.NewGORM()
	defer db.Close()

	var kid model.Kid
	if err := db.Where("id = ? AND parent_id = ?", kidID, user.ID).First(&kid).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can't find kid",
			"error":   err,
		})
		return
	}

	if err := db.Where("id = ?", kid.ID).Delete(&model.Kid{}).Error; err != nil {
		logError(errors.Wrapf(err, "Error on deleting kid. Kid ID: %#v", kidID))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when deleting kid from database",
			"error":   err,
		})
		return
	}
	LogUserActivity(db, &user, fmt.Sprintf("Delete Kid (%d)", kid.ID), &kid.MacID)
}

func GetKidList(c *gin.Context) {
	db := database.NewGORM()
	defer db.Close()

	user := GetSignedInUser(c)

	var kids []model.Kid
	if err := db.Where("parent_id = ?", user.ID).Order("date_created desc").Find(&kids).Error; err != nil {
		logError(errors.Wrapf(err, "Error on retriving kid list: %#v", kids))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retriving kid list",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, kids)

	LogUserActivity(db, &user, "Get Kid List (%d)", nil)
}

func GetAllKidList(c *gin.Context) {
	db := database.NewGORM()
	defer db.Close()

	var kids []model.Kid
	if err := db.Preload("Parent").Order("date_created desc").Find(&kids).Limit(50).Error; err != nil {
		logError(errors.Wrapf(err, "Error on retriving kid list: %#v", kids))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retriving kid list",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, kids)

	user := GetSignedInUser(c)
	LogUserActivity(db, &user, "Get All Kid List", nil)
}

func WhoRegisteredMacID(c *gin.Context) {
	macID := c.Query("macId")

	if macID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "MacID is required",
		})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	kid, err := GetKidByMacID(db, macID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"kid": kid,
	})
}

func UpdateBatteryStatus(c *gin.Context) {
	var batteryStatus model.BatteryStatus

	if err := c.BindJSON(&batteryStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Reqeust",
			"error":   err,
		})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	if err := db.Save(&batteryStatus).Error; err != nil {
		logError(errors.Wrapf(err, "Error when insert battery status: %#v", batteryStatus))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error when insert battery data",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func UpdateKidRevertMacID(c *gin.Context) {
	kidID := c.Query("kidId")
	macID := c.Query("macId")

	if kidID == "" || macID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Kid ID and Mac ID are required",
		})
	}

	db := database.NewGORM()
	defer db.Close()
	if err := fixMacIDReverseIssue(db, macID, kidID); err != nil {
		fmt.Println(err)
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on updating MAC ID",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

func fixMacIDReverseIssue(db *gorm.DB, macID, kidID string) error {
	var revertMacID string
	for i := 0; i < 12; i += 2 {
		revertMacID += fmt.Sprintf("%s%s", string(macID[i+1]), string(macID[i]))
	}
	if err := db.Exec("UPDATE kids SET mac_id = ? WHERE mac_id = ? AND id = ?", macID, revertMacID, kidID).Error; err != nil {
		return err
	}

	if err := db.Exec("UPDATE activity SET mac_id = ? WHERE mac_id = ? AND kid_id = ?", macID, revertMacID, kidID).Error; err != nil {
		return err
	}

	return nil
	// E0E5CF1ED7C2
	// 0E5EFCE17D2C
}
