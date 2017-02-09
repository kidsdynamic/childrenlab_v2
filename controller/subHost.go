package controller

import (
	"net/http"

	"fmt"

	"time"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

const (
	SubHostStatusPending  = "PENDING"
	SubHostStatusAccepted = "ACCEPTED"
	SubHostStatusDenied   = "DENIED"
)

func RequestSubHostToUser(c *gin.Context) {
	var requestSubHostReq model.RequestSubHostToUser

	if err := c.BindJSON(&requestSubHostReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "One of parameter is missing",
			"error":   err,
		})
		return
	}

	user := GetSignedInUser(c)

	db := database.NewGORM()
	defer db.Close()

	var subHost model.SubHost

	db.Where("request_from_id = ? AND request_to_iD = ?", user.ID, requestSubHostReq.HostID).First(&subHost)

	if subHost.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{
			"message": "The request is already exist",
		})
		return
	}

	subHost.DateCreated = time.Now()
	subHost.RequestFromID = user.ID
	subHost.RequestToID = requestSubHostReq.HostID
	subHost.LastUpdated = time.Now()

	if err := db.Save(&subHost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when inserting request",
			"error":   err,
		})
		return
	}

	if err := db.Preload("RequestFrom").Preload("RequestTo").First(&subHost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when inserting request",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, subHost)
}

func AcceptRequest(c *gin.Context) {
	var acceptRequest model.UpdateSubHostRequest

	if err := c.BindJSON(&acceptRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "One of parameter is missing",
			"error":   err,
		})
		return
	}

	user := GetSignedInUser(c)

	db := database.NewGORM()
	defer db.Close()

	var subHost model.SubHost
	db.Where("Request_to_id = ? AND id = ?", user.ID, acceptRequest.SubHostID).Preload("RequestFrom").Preload("RequestTo").First(&subHost)

	if subHost.ID == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "The user doesn't have permission to accept the request",
		})

		return
	}

	var kids []model.Kid

	if err := db.Joins("JOIN user ON user.id = kids.parent_id").Where("kids.id in (?) AND kids.parent_id = ?", acceptRequest.KidID, user.ID).Find(&kids).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on updating status",
			"error":   err,
		})

		return
	}

	for index, kid := range kids {
		var exists bool
		row := db.Raw("SELECT EXISTS(SELECT sub_host_id FROM sub_host_kid WHERE sub_host_id = ? AND kid_id = ? LIMIT 1)", subHost.ID, kid.ID).Row()
		row.Scan(&exists)
		if exists {
			if len(kids) > index+1 {
				kids = append(kids[:index], kids[index+1:]...)
			}

		}

	}
	if len(kids) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can't find kid or the kid is already in the sub host",
		})
		return

	}

	subHost.Kids = kids
	subHost.Status = SubHostStatusAccepted

	if err := db.Save(&subHost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on updating status",
			"error":   err,
		})

		return
	}

	var updatedSubHost model.SubHost
	if err := db.Model(&updatedSubHost).Preload("Kids").Preload("RequestFrom").Preload("RequestTo").Where("id = ?", subHost.ID).First(&updatedSubHost).Error; err != nil {
		fmt.Printf("ERror on retrieve subhost. Error: %#v", err)
	}

	c.JSON(http.StatusOK, updatedSubHost)
}

func DenyRequest(c *gin.Context) {
	var updateRequest model.UpdateSubHostRequest

	if err := c.BindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "One of parameter is missing",
			"error":   err,
		})
		return
	}

	user := GetSignedInUser(c)

	db := database.NewGORM()
	defer db.Close()

	var subHost model.SubHost

	db.Model(&subHost).Preload("RequestFrom").Preload("RequestTo").Preload("kids").Where("id = ? AND request_to_id = ?", updateRequest.SubHostID, user.ID).First(&subHost)

	if subHost.ID == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "The user doesn't have permission to accept the request",
		})

		return
	}

	subHost.Status = SubHostStatusDenied
	subHost.LastUpdated = time.Now()

	if err := db.Save(&subHost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on updating status",
			"error":   err,
		})

		return
	}

	c.JSON(http.StatusOK, subHost)
}

func SubHostList(c *gin.Context) {
	db := database.NewGORM()
	defer db.Close()

	query := c.Query("status")
	var subHosts []model.SubHost
	var err error

	user := GetSignedInUser(c)
	if query == "" {
		err = db.Where("request_to_id = ?", user.ID).Preload("RequestFrom").Preload("RequestTo").Preload("Kids").Find(&subHosts).Error
	} else {
		err = db.Where("request_to_id = ? AND status = ?", user.ID, query).Preload("RequestFrom").Preload("RequestTo").Preload("Kids").Find(&subHosts).Error
	}

	/*
		for _, subHost := range subHosts {
			var kidIDs []int64
			for _, kid := range subHost.Kids {
				kidIDs = append(kidIDs, kid.KidID)
			}
			if len(kidIDs) > 0 {
			}
		}
	*/

	if err != nil {
		fmt.Printf("Error on Sub Host List. %#v", err)
	}
	c.JSON(http.StatusOK, subHosts)
}

func HasPermission(c *gin.Context) {
	kidIDString := c.Query("kidId")
	kidID, err := strconv.ParseInt(kidIDString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when parse kid ID to int",
			"error":   err,
		})
		return
	}
	user := GetSignedInUser(c)

	db := database.NewGORM()
	defer db.Close()

	fmt.Println(HasPermissionToKid(db, user, kidID))
}
