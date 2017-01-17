package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/database"
)

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
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"kid": kid,
	})
}
