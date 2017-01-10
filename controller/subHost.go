package controller

import (
	"net/http"

	"fmt"

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

	result, err := db.NamedExec("INSERT INTO sub_host (request_from_id, request_to_id, status, date_created, last_updated) "+
		"VALUES (:request_from_id, :request_to_id, :status, Now(), Now())",
		requestSubHostReq)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when inserting request",
		})
		return
	}
	subHostRequest := getSubHostByID(db, getInsertedID(result))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on getting inerted row",
			"error":   err,
		})
	}

	c.JSON(http.StatusOK, subHostRequest)
}
func ValidateHostRequest(db *sqlx.DB, subHostID, hostID int64) (bool, error) {
	var exists bool

	if err := db.Get(&exists, "SELECT EXISTS(SELECT id FROM sub_host WHERE "+
		"request_to_id = ? AND id = ? LIMIT 1)", hostID, subHostID); err != nil {
		return false, err
	}

	return exists, nil
}

func IsRequestExists(db *sqlx.DB, req *model.RequestSubHostToUser) (bool, error) {
	var exists bool
	if err := db.Get(&exists, "SELECT EXISTS(SELECT s.id FROM sub_host s WHERE request_from_id = ? AND "+
		"request_to_id = ? LIMIT 1)", req.UserID, req.HostID); err != nil {
		return false, err
	}

	return exists, nil
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

	db := database.New()
	defer db.Close()

	valid, err := ValidateHostRequest(db, acceptRequest.SubHostID, user.ID)

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

	result := db.MustExec("UPDATE sub_host SET status = ?, last_updated = NOW() WHERE id = ?", SubHostStatusAccepted, acceptRequest.SubHostID)

	if success := checkInsertResult(result); !success {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on updating status",
			"error":   err,
		})

		return
	}

	tx := db.MustBegin()
	for _, kidID := range acceptRequest.KidID {
		if !isKidExistsInSubHost(db, acceptRequest.SubHostID, kidID) {
			tx.MustExec("INSERT INTO sub_host_kids (sub_host_kid_id, kids_id) VALUES (?, ?)", acceptRequest.SubHostID, kidID)
		}

	}
	tx.Commit()

	subHost := getSubHostByID(db, acceptRequest.SubHostID)

	c.JSON(http.StatusOK, subHost)
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

	db := database.New()
	defer db.Close()

	valid, err := ValidateHostRequest(db, updateRequest.SubHostID, user.ID)

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

	result := db.MustExec("UPDATE sub_host SET status = ?, last_updated = NOW() WHERE id = ?", SubHostStatusDenied, updateRequest.SubHostID)

	if success := checkInsertResult(result); !success {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on updating status",
			"error":   err,
		})

		return
	}
	subHost := getSubHostByID(db, updateRequest.SubHostID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on retriving subhost",
			"error":   err,
		})

		return
	}

	c.JSON(http.StatusOK, subHost)
}

func SubHostList(c *gin.Context) {
	db := database.New()
	defer db.Close()

	query := c.Query("status")
	var subHosts []model.SubHost

	user := GetSignedInUser(c)
	if query == "" {
		_ = db.Select(&subHosts, "SELECT s.id, s.request_from_id, s.request_to_id, s.status, s.date_created, s.last_updated "+
			"FROM sub_host s WHERE s.request_to_id = ?", user.ID)
	} else {
		_ = db.Select(&subHosts, "SELECT s.id, s.request_from_id, s.request_to_id, s.status, s.date_created, s.last_updated "+
			"FROM sub_host s WHERE s.request_to_id = ? AND status = ?", user.ID, query)
	}

	for key, subHost := range subHosts {
		kids := getSubHostKid(db, subHost.ID)
		subHosts[key].Kid = kids
	}

	c.JSON(http.StatusOK, subHosts)
}

func getSubHostByID(db *sqlx.DB, ID int64) model.SubHost {
	var subHost model.SubHost
	_ = db.Get(&subHost, "SELECT s.id, s.request_from_id, s.request_to_id, s.status, s.date_created, s.last_updated "+
		"FROM sub_host s WHERE s.id = ?", ID)
	fmt.Printf("SELECT s.id, s.request_from_id, s.request_to_id, s.status, s.date_created, s.last_updated FROM sub_host s WHERE s.id = %d\n", ID)
	kids := getSubHostKid(db, ID)
	fmt.Println(kids)
	subHost.Kid = kids

	return subHost

}

func isKidExistsInSubHost(db *sqlx.DB, subHostID, kidID int64) bool {
	var exist bool
	if err := db.Get(&exist, "SELECT EXISTS(SELECT sub_host_kid_id FROM sub_host_kids WHERE sub_host_kid_id = ? AND kids_id = ? LIMIT 1)", subHostID, kidID); err != nil {
		panic(err)
		return false
	}

	return exist
}

func getSubHostKid(db *sqlx.DB, subHostID int64) []model.Kid {
	var kids []model.Kid
	err := db.Select(&kids, "SELECT k.id,  COALESCE(k.first_name, '') as first_name, COALESCE(k.last_name, '') as last_name"+
		", COALESCE(k.mac_id, '') as mac_id, COALESCE(k.profile, '') as profile, parent_id, k.date_created FROM kids k JOIN sub_host_kids s ON k.id = s.kids_id WHERE s.sub_host_kid_id = ?", subHostID)

	if err != nil {
		panic(err)
		return kids
	}
	return kids

}
