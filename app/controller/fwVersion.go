package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/app/database"
	"github.com/kidsdynamic/childrenlab_v2/app/model"
)

func GetCurrentFWVersionAndLink(c *gin.Context) {
	db := database.NewGORM()
	defer db.Close()

	var currentVersion model.FwFile
	if err := db.Order("id desc").First(&currentVersion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on retriving list",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, currentVersion)
}
