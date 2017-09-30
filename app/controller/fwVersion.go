package controller

import (
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/app/database"
	"github.com/kidsdynamic/childrenlab_v2/app/model"
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

	var currentVersion model.FwFile

	if err := db.Order("id desc").First(&currentVersion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on retriving list",
			"error":   err,
		})
		return
	}

	//TODO: For NOW
	currentVersion.Version = ""

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
