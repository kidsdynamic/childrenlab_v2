package controller

import (
	"net/http"

	"log"

	"database/sql"

	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
	"github.com/pkg/errors"
)

type FWVersionModel struct {
	MacID           string `json:"macId"`
	FirmwareVersion string `json:"firmwareVersion"`
}

func GetCurrentFWVersionAndLink(c *gin.Context) {
	db := database.NewGORM()
	defer db.Close()

	macID := c.Param("macId")
	if macID == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	var deviceInitVersion model.InitialDeviceFirmware
	deviceVersion := "KDV0005-A"
	if err := db.Where("mac_id = ?", macID).First(&deviceInitVersion).Error; err == nil {
		deviceVersion = deviceInitVersion.FirmwareVersion
	}

	r, _ := regexp.Compile("-.*")
	deviceVersion = r.FindString(deviceVersion)
	fmt.Println(deviceVersion)

	var currentVersion model.FwFile

	if err := db.Where("active = true and version like ?", fmt.Sprintf("%%%s%%", deviceVersion)).Order("id desc").First(&currentVersion).Error; err != nil {
		if err != sql.ErrNoRows {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error on retriving list",
				"error":   err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, currentVersion)
}

func UpdateDeviceFWVersion(c *gin.Context) {

	var versionModel FWVersionModel
	if err := c.BindJSON(&versionModel); err != nil {
		log.Printf("Error on Version data. Bind with json. %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	var kid model.Kid
	if err := db.Where("mac_id = ?", versionModel.MacID).Find(&kid).Error; err != nil {
		logError(errors.Wrapf(err, "Can't find kid from Mac ID: %s", versionModel.MacID))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can't find device",
			"error":   err,
		})
		return
	}

	if err := db.Model(&model.Kid{}).Where("mac_id = ?", versionModel.MacID).Update("firmware_version", versionModel.FirmwareVersion).Error; err != nil {
		logError(errors.Wrapf(err, "Error on update device firmware version. Kid: %#v", kid))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when updating device firmware version",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})

}
