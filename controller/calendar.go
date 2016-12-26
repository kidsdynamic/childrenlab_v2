package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

func AddCalendarEvent(c *gin.Context) {
	var eventRequest model.AddEventRequest

	if err := c.BindJSON(&eventRequest); err != nil {
		log.Printf("Error on Add Event. Bind with json. %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing some parameters",
			"error":   err,
		})
		return
	}

	user := GetSignedInUser(c)
	eventRequest.UserID = user.ID
	eventRequest.Status = "OPEN"

	db := database.New()
	defer db.Close()

	result, err := db.NamedExec("INSERT INTO calendar_event (event_name, kid_id, start_date, end_date, color, description, "+
		"alert, city, state, event_repeat, timezone_offset, date_created, last_updated, user_id, status) VALUES "+
		"(:event_name, :kid_id, :start_date, :end_date, :color, :description, :alert, :city, :state, :event_repeat, :timezone_offset, NOW(), NOW(), :user_id, :status)",
		eventRequest)

	if err != nil {
		log.Printf("Error on Add Event. %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on adding the event to database",
			"error":   err.Error(),
		})
		return
	}

	insertedEventID := getInsertedID(result)

	tx := db.MustBegin()
	for _, value := range eventRequest.Todo {
		tx.MustExec("INSERT INTO todo_list (text, status, event_id, date_created, last_updated) VALUES (?, ?, ?, Now(), Now())",
			value, "PENDING", insertedEventID)
	}
	tx.Commit()

	var event model.Event

	err = db.Get(&event, "SELECT id, user_id, event_name, status, kid_id, start_date, end_date, color, COALESCE(description, '') as description, "+
		"alert, city, state, COALESCE(event_repeat, '') as event_repeat, timezone_offset, date_created, last_updated FROM calendar_event WHERE id = ?",
		insertedEventID)

	if err != nil {
		log.Printf("Error on retrieve inserted Event. %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on retrieve inserted Event",
			"error":   err.Error(),
		})
		return
	}

	err = db.Select(&event.Todo, "SELECT id, date_created, last_updated, status, text FROM todo_list WHERE event_id = ?", insertedEventID)

	if err != nil {
		log.Printf("Error on retrieve todos. %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on retrieve todos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}
