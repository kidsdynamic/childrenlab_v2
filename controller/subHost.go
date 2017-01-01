package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

const (
	SubHostStatusPending  = "PENDING"
	SubHostStatusAccepted = "ACCEPTED"
	SubHostStatusDenied   = "DENIED"
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
		"VALUES (:device_id, :request_from_id, :request_to_id, :status, Now(), Now())",
		requestSubHostReq)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when inserting request",
		})
		return
	}
	subHostRequest, err := GetSubHostRequestByID(db, getInsertedID(result))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on getting inerted row",
			"error":   err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"SubHostRequest": subHostRequest,
	})
}

func GetSubHostRequestByID(db *sqlx.DB, ID int64) (model.SubHostRequest, error) {
	var subHostRequest model.SubHostRequest

	err := db.Get(&subHostRequest, "SELECT s.id, d.mac_id, s.request_from_id, s.request_to_id, s.status, s.date_created, s.last_updated "+
		"FROM sub_host_request s JOIN device d ON s.device_id = d.id WHERE s.id = ? LIMIT 1", ID)

	if err != nil {
		return subHostRequest, err
	}

	return subHostRequest, nil
}

func ValidateHostRequest(db *sqlx.DB, subHostID, hostID int64) (bool, error) {
	var exists bool

	if err := db.Get(&exists, "SELECT EXISTS(SELECT id FROM sub_host_request WHERE "+
		"request_to_id = ? AND id = ? LIMIT 1)", hostID, subHostID); err != nil {
		return false, err
	}

	return exists, nil
}

func IsRequestExists(db *sqlx.DB, req *model.RequestSubHostRequest) (bool, error) {
	var exists bool
	if err := db.Get(&exists, "SELECT EXISTS(SELECT s.id FROM sub_host_request s JOIN device d ON s.device_id = d.id WHERE request_from_id = ? AND "+
		"request_to_id = ? AND d.mac_id = ? LIMIT 1)", req.UserID, req.HostID, req.MacID); err != nil {
		return false, err
	}

	return exists, nil
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

	db := database.New()
	defer db.Close()

	valid, err := ValidateHostRequest(db, acceptRequest.RequestID, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on checking user's permission",
			"error":   err,
		})
		return
	}

	if !valid {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "The user doesn't have permission to accept the request",
		})

		return
	}

	result := db.MustExec("UPDATE sub_host_request SET status = ?, last_updated = NOW() WHERE id = ?", SubHostStatusAccepted, acceptRequest.RequestID)

	if success := checkInsertResult(result); !success {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on updating status",
			"error":   err,
		})

		return
	}
	subHost, err := GetSubHostRequestByID(db, acceptRequest.RequestID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on retriving subhost",
			"error":   err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"SubHostRequest": subHost,
	})
}

func DenyRequest(c *gin.Context) {
	var acceptRequest model.AcceptSubHostRequest

	if err := c.BindJSON(&acceptRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "One of parameter is missing",
			"error":   err,
		})
		return
	}

	user := GetSignedInUser(c)

	db := database.New()
	defer db.Close()

	valid, err := ValidateHostRequest(db, acceptRequest.RequestID, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on checking user's permission",
			"error":   err,
		})
		return
	}

	if !valid {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "The user doesn't have permission to accept the request",
		})

		return
	}

	result := db.MustExec("UPDATE sub_host_request SET status = ?, last_updated = NOW() WHERE id = ?", SubHostStatusDenied, acceptRequest.RequestID)

	if success := checkInsertResult(result); !success {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on updating status",
			"error":   err,
		})

		return
	}
	subHost, err := GetSubHostRequestByID(db, acceptRequest.RequestID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on retriving subhost",
			"error":   err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"SubHostRequest": subHost,
	})
}

func SubHostList(c *gin.Context) {
	db := database.New()
	defer db.Close()

	query := c.DefaultQuery("status", SubHostStatusPending)
	var subHostRequest []model.SubHostRequest

	user := GetSignedInUser(c)
	_ = db.Select(&subHostRequest, "SELECT s.id, d.mac_id, s.request_from_id, s.request_to_id, s.status, s.date_created, s.last_updated "+
		"FROM sub_host_request s JOIN device d ON s.device_id = d.id WHERE s.request_to_id = ? AND status = ?", user.ID, query)

	c.JSON(http.StatusOK, gin.H{
		"subHost": subHostRequest,
	})
}
