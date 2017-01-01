package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

func RequestSubHost(c *gin.Context) {
	var requestSubHostReq model.RequestSubHostRequest

	if err := c.BindJSON(&requestSubHostReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "One of parameter is missing",
			"error":   err,
		})
		return
	}

	user := GetSignedInUser(c)
	requestSubHostReq.UserID = user.ID
	requestSubHostReq.Status = SubHostStatusPending

	db := database.New()
	defer db.Close()

	exist, err := IsRequestExists(db, &requestSubHostReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when get IsRequestExists",
			"error":   err,
		})
		return
	}

	if exist {
		c.JSON(http.StatusConflict, gin.H{
			"message": "The request is already exist",
		})
		return
	}

	var device model.Device

	err = db.Get(&device, "SELECT id, mac_id, date_created FROM device WHERE mac_id = ? LIMIT 1", requestSubHostReq.MacID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on finding device",
			"error":   err,
		})
		return
	}

	requestSubHostReq.DeviceID = device.ID

	result, err := db.NamedExec("INSERT INTO sub_host_request (device_id, request_from_id, request_to_id, status, date_created, last_updated) "+
		"VALUES (:device_id, :request_from_id, :request_to_id, 'PENDING', Now(), Now())",
		requestSubHostReq)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when inserting request",
		})
		return
	}

	var subHostRequest model.SubHostRequest

	err = db.Get(&subHostRequest, "SELECT s.id, d.mac_id, s.request_from_id, s.request_to_id, s.status, s.date_created, s.last_updated "+
		"FROM sub_host_request s JOIN device d ON s.device_id = d.id WHERE s.id = ? LIMIT 1", getInsertedID(result))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on getting inerted row",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"SubHostRequest": subHostRequest,
	})
}

func IsRequestExists(db *sqlx.DB, req *model.RequestSubHostRequest) (bool, error) {
	var exists bool
	if err := db.Get(&exists, "SELECT EXISTS(SELECT s.id FROM sub_host_request s JOIN device d ON s.device_id = d.id WHERE request_from_id = ? AND "+
		"request_to_id = ? AND d.mac_id = ? LIMIT 1)", req.UserID, req.HostID, req.MacID); err != nil {
		return false, err
	}

	return exists, nil
}
