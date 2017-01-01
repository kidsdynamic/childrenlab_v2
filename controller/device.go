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

	db := database.New()
	defer db.Close()

	kid, err := GetKidByMacID(db, macID)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	kidParent, err := GetUserByID(db, kid.ParentID)

	if err != nil {
		fmt.Printf("Can't find kid's parent. %#v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't kid's parent",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"kid":  kid,
		"user": kidParent,
	})
}
