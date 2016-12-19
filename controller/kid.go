package controller

import (
	"net/http"

	"log"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

func AddKid(c *gin.Context) {
	user := GetSignedInUser(c)

	var request model.KidRequest

	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Reqeust",
			"error":   err,
		})
		return
	}

	db := database.New()
	defer db.Close()

	//check if device exist
	var exist bool
	if err := db.Get(&exist, "SELECT EXISTS(SELECT id FROM device WHERE mac_id = ? LIMIT 1)", request.MacID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong on server side",
			"error":   err,
		})
		return
	}

	if exist {
		c.JSON(http.StatusConflict, gin.H{
			"message": "The device is already registered",
		})
		return
	}

	result, err := db.Exec("INSERT INTO kids (first_name, last_name, parent_id, date_created, last_updated)"+
		" VALUES (?, ?, ?, Now(), Now())", request.FirstName, request.LastName, user.ID)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when insert kid data",
			"error":   err,
		})
		return
	}

	kidId, err := result.LastInsertId()

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error when insert kid data",
			"error":   err,
		})
		return
	}

	result = db.MustExec("INSERT INTO device (kid_id, mac_id, user_id, date_created, last_updated) VALUES "+
		"(?, ?, ?, Now(), Now())", kidId, request.MacID, user.ID)

	if !checkInsertResult(result) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error when insert device data",
		})
		return
	}

	kids, err := GetKidsByUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error when retrieve kids",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"kids": kids,
	})

}

func UpdateKid(c *gin.Context) {
	var request model.UpdateKidRequest

	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	fmt.Printf("Kid Update Request: %#v", request)

	db := database.New()
	defer db.Close()

	user := GetSignedInUser(c)
	kid, err := GetKidByUserIdAndKidId(db, user.ID, request.ID)

	if err != nil {
		fmt.Printf("Can't find kid. %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can't find kid",
			"error":   err,
		})
		return
	}

	tx := db.MustBegin()
	if request.FirstName != "" {
		tx.MustExec("UPDATE kids SET first_name = ? WHERE id = ?", request.FirstName, kid.ID)
	}

	if request.LastName != "" {
		tx.MustExec("UPDATE kids SET last_name = ? WHERE id = ?", request.LastName, kid.ID)
	}

	/*	if request.MacID != "" {
		tx.MustExec("UPDATE kids SET phone_number = ? WHERE id = ?", request.PhoneNumber, kid.ID)
	}*/

	tx.MustExec("UPDATE user SET last_updated = NOW() WHERE id = ?", kid.ID)
	tx.Commit()

	kid, err = GetKidByUserIdAndKidId(db, user.ID, request.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retreive updated user information",
		})
	}

	if err != nil {
		fmt.Printf("Can't find kid. %#v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't find kid",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"kid": kid,
	})

}
