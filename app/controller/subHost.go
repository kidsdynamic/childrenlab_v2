package controller

import (
	"net/http"

	"fmt"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kidsdynamic/childrenlab_v2/app/database"
	"github.com/kidsdynamic/childrenlab_v2/app/model"
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

	var exists bool = false
	row := db.Raw("SELECT EXISTS(SELECT id FROM user WHERE id = ? LIMIT 1)", requestSubHostReq.HostID).Row()
	row.Scan(&exists)
	fmt.Printf("Existing: %#v", exists)
	if !exists {
		logError(errors.Wrapf(errors.New("The host not found"), "Error on saving sub host ID: %d", requestSubHostReq.HostID))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The hot not found",
		})
		return
	}

	var subHost model.SubHost

	if err := db.Where("request_from_id = ? AND request_to_id = ?", user.ID, requestSubHostReq.HostID).First(&subHost).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logError(errors.Wrapf(err, "Error on retriving sub host request: %#v", requestSubHostReq))
		}

	}

	if subHost.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{
			"message": "The request is already exist",
		})
		return
	}

	subHost.DateCreated = GetNowTime()
	subHost.RequestFromID = user.ID
	subHost.RequestToID = requestSubHostReq.HostID
	subHost.LastUpdated = GetNowTime()

	if err := db.Save(&subHost).Error; err != nil {
		logError(errors.Wrapf(err, "Error on saving sub host: %#v", subHost))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when inserting request",
			"error":   err,
		})
		return
	}

	if err := db.Model(&model.SubHost{}).Where("id = ?", subHost.ID).Preload("RequestFrom").Preload("RequestTo").First(&subHost).Error; err != nil {
		logError(errors.Wrapf(err, "Error on retriving sub host: %#v", subHost))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when inserting request",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, subHost)
}

func AcceptRequest(c *gin.Context) {
	var acceptRequest model.AcceptSubHostRequest

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
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "The user doesn't have permission to accept the request",
		})

		return
	}

	var kids []model.Kid

	if err := db.Joins("JOIN user ON user.id = kids.parent_id").Where("kids.id in (?) AND kids.parent_id = ?", acceptRequest.KidID, user.ID).Find(&kids).Error; err != nil {
		logError(errors.Wrapf(err, "Error on update subhost request status: %#v", acceptRequest))
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
		logError(errors.Wrapf(err, "Error on saving subhost: %#v", subHost))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on updating status",
			"error":   err,
		})

		return
	}

	var updatedSubHost model.SubHost
	if err := db.Model(&updatedSubHost).Where("id = ?", subHost.ID).Preload("Kids").Preload("RequestFrom").Preload("RequestTo").First(&updatedSubHost).Error; err != nil {
		logError(errors.Wrapf(err, "Error on retriving sub host: %#v", updatedSubHost))
	}

	c.JSON(http.StatusOK, updatedSubHost)
}

func DeleteRequest(c *gin.Context) {
	subHostId := c.Query("subHostId")
	if subHostId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "One of parameter is missing",
		})
		return
	}
	user := GetSignedInUser(c)

	db := database.NewGORM()
	defer db.Close()

	var subHost model.SubHost

	if err := db.Model(&subHost).Where("id = ? AND request_from_id = ?", subHostId, user.ID).Preload("Kids").First(&subHost).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{})
			return
		} else {
			logError(errors.Wrapf(err, "Error on retriving subhost: %#v", subHost))
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error on retrieve subhost",
				"error":   err,
			})
		}
	}

	if err := db.Where("sub_host_id = ?", subHost.ID).Delete(model.SubHostKid{}).Error; err != nil {
		logError(errors.Wrapf(err, "Error on deleting subhost kids: %#v", subHost))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on updating status",
			"error":   err,
		})

		return
	}

	if err := db.Delete(&subHost).Error; err != nil {
		logError(errors.Wrapf(err, "Error on delete sub host: %#v", subHost))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on updating status",
			"error":   err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

func DenyRequest(c *gin.Context) {
	var request model.DenyRequest

	if err := c.BindJSON(&request); err != nil {
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

	db.Model(&subHost).Preload("RequestFrom").Preload("RequestTo").Preload("Kids").Where("id = ? AND request_to_id = ?", request.SubHostID, user.ID).First(&subHost)

	if subHost.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "The user doesn't have permission to accept the request",
		})

		return
	}

	if err := db.Where("sub_host_id = ?", subHost.ID).Delete(model.SubHostKid{}).Error; err != nil {
		logError(errors.Wrapf(err, "Error on delete subhost kid request: %#v", subHost))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on updating status",
			"error":   err,
		})

		return
	}

	if err := db.Delete(&subHost).Error; err != nil {
		logError(errors.Wrapf(err, "Error on deleting sub host: %#v", subHost))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on updating status",
			"error":   err,
		})

		return
	}

	c.JSON(http.StatusOK, subHost)
}

func RemoveSubHostKid(c *gin.Context) {
	var request model.RemoveSubHostRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "One of parameter is missing",
			"error":   err,
		})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	user := GetSignedInUser(c)

	kidIds := []int64{
		request.KidID,
	}
	if !HasPermissionToKid(db, &user, kidIds) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "The user doesn't have permission to remove kid",
		})
		return
	}

	var subHost model.SubHost
	if err := db.Where("id = ? AND (request_to_id = ? or request_from_id = ?) AND status = ?", request.SubHostID, user.ID, user.ID, SubHostStatusAccepted).First(&subHost).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "The subhost is not exists",
			})
			return
		} else {
			logError(errors.Wrapf(err, "Error on retrive subhost: %#v", request))
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error occur",
				"error":   err,
			})

			return
		}
	}

	if err := db.Where("sub_host_id = ? and kid_id = ?", request.SubHostID, request.KidID).Delete(&model.SubHostKid{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "The subhost kid is not exists",
			})
			return
		} else {
			logError(errors.Wrapf(err, "Error on delete sub host kid: %#v", request))
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error occur",
				"error":   err,
			})

			return
		}
	}

	if err := db.Model(&model.SubHost{}).Where("id = ?", subHost.ID).Preload("RequestFrom").Preload("RequestTo").Preload("Kids").First(&subHost).Error; err != nil {
		logError(errors.Wrapf(err, "Error on retriving sub host: %#v", subHost))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occur",
			"error":   err,
		})

		return
	}

	//If the request doesn't have any kid, remove the request
	if len(subHost.Kids) == 0 {
		if err := db.Delete(&subHost).Error; err != nil {
			logError(errors.Wrapf(err, "Error on deleting sub host: %#v", subHost))
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error on updating status",
				"error":   err,
			})

			return
		}
	}

	c.JSON(http.StatusOK, subHost)
}

func SubHostList(c *gin.Context) {
	db := database.NewGORM()
	defer db.Close()

	query := c.Query("status")
	var requestTo []model.SubHost
	var requestFrom []model.SubHost
	var err error

	user := GetSignedInUser(c)
	if query == "" {
		err = db.Where("request_to_id = ?", user.ID).Preload("RequestFrom").Preload("RequestTo").Preload("Kids").Find(&requestFrom).Error
	} else {
		err = db.Where("request_to_id = ? AND status = ?", user.ID, query).Preload("RequestFrom").Preload("RequestTo").Preload("Kids").Find(&requestFrom).Error
	}

	if err != nil {
		fmt.Printf("Error on Sub Host List. %#v", err)
	}

	if query == "" {
		err = db.Where("request_from_id = ?", user.ID).Preload("RequestFrom").Preload("RequestTo").Preload("Kids").Find(&requestTo).Error
	} else {
		err = db.Where("request_from_id = ? AND status = ?", user.ID, query).Preload("RequestFrom").Preload("RequestTo").Preload("Kids").Find(&requestTo).Error
	}

	if err != nil {
		logError(errors.Wrapf(err, "Error on Sub Host List: %#v", query))
	}

	c.JSON(http.StatusOK, gin.H{
		"requestFrom": requestFrom,
		"requestTo":   requestTo,
	})

}
