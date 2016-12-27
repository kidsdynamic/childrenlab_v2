package controller

import (
	"log"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

func AddCalendarEvent(c *gin.Context) {
	var eventRequest model.EventRequest

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

	//err = getEventWithTodoList(db, event, insertedEventID)
	err = db.Get(&event, "SELECT id, user_id, event_name, status, kid_id, start_date, end_date, color, COALESCE(description, '') as description, "+
		"alert, city, state, COALESCE(event_repeat, '') as event_repeat, timezone_offset, date_created, last_updated FROM calendar_event WHERE id = ?",
		insertedEventID)

	if err != nil {
		log.Printf("Error on retrieve inserted Event. %#v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on retrieve inserted Event",
			"error":   err.Error(),
		})
		return
	}

	err = db.Select(&event.Todo, "SELECT id, date_created, last_updated, status, text FROM todo_list WHERE event_id = ?", event.ID)

	if err != nil {
		log.Printf("Error on retrieve todos. %#v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on retrieve inserted Event",
			"error":   err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"event": event,
	})

}

func UpdateCalendarEvent(c *gin.Context) {
	var eventRequest model.UpdateEventRequest

	if err := c.BindJSON(&eventRequest); err != nil {
		log.Printf("Error on Add Event. Bind with json. %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing some parameters",
			"error":   err,
		})
		return
	}

	//user := GetSignedInUser(c)

	db := database.New()
	defer db.Close()

	_, err := db.NamedExec("UPDATE calendar_event SET event_name = :event_name, start_date = :start_date, end_date = :end_date, color = :color, "+
		"description = :description, alert = :alert, city = :city, state = :state, event_repeat = :event_repeat, timezone_offset = :timezone_offset, "+
		"last_updated = Now() WHERE id = :id", eventRequest)

	if err != nil {
		log.Printf("Error on Updat event. %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on updating event",
			"error":   err.Error(),
		})
		return
	}

	//Remove all of todos under the event
	_ = db.MustExec("DELETE FROM todo_list WHERE event_id = ?", eventRequest.ID)

	//Insert all of todos again
	tx := db.MustBegin()
	for _, value := range eventRequest.Todo {
		tx.MustExec("INSERT INTO todo_list (text, status, event_id, date_created, last_updated) VALUES (?, ?, ?, Now(), Now())",
			value, "PENDING", eventRequest.ID)
	}
	tx.Commit()

	var event model.Event
	err = db.Get(&event, "SELECT id, user_id, event_name, status, kid_id, start_date, end_date, color, COALESCE(description, '') as description, "+
		"alert, city, state, COALESCE(event_repeat, '') as event_repeat, timezone_offset, date_created, last_updated FROM calendar_event WHERE id = ?",
		eventRequest.ID)

	if err != nil {
		log.Printf("Error on retrieve inserted Event. %#v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on retrieve inserted Event",
			"error":   err.Error(),
		})
		return
	}

	err = db.Select(&event.Todo, "SELECT id, date_created, last_updated, status, text FROM todo_list WHERE event_id = ?", event.ID)

	if err != nil {
		log.Printf("Error on retrieve todos. %#v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on retrieve inserted Event",
			"error":   err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"event": event,
	})

}

func DeleteEvent(c *gin.Context) {
	var deleteEventRequest model.DeleteEventRequest

	if err := c.BindJSON(&deleteEventRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing some parameters",
			"error":   err,
		})
		return
	}

	db := database.New()
	defer db.Close()

	user := GetSignedInUser(c)
	result := db.MustExec("DELETE FROM calendar_event WHERE id = ? AND user_id = ?", deleteEventRequest.EventID, user.ID)
	if !checkInsertResult(result) {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	_ = db.MustExec("DELETE FROM todo_list WHERE event_id = ?", deleteEventRequest.EventID)

	c.JSON(http.StatusOK, gin.H{})
}

func GetCalendarEvent(c *gin.Context) {
	var getEventRequest model.GetEventRequest
	getEventRequest.Period = c.Query("period")
	getEventRequest.Date = c.Query("date")

	if getEventRequest.Period == "" || getEventRequest.Date == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing some parameters",
		})
		return
	}

	db := database.New()
	defer db.Close()
	var events []model.Event

	t, err := time.Parse(TimeLayout, getEventRequest.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Date formate is wrong",
			"error":   err,
		})
		return
	}

	log.Printf("Get Calendar Event Date: %s", t)

	switch getEventRequest.Period {
	case "DAY":

		err = db.Select(&events, "SELECT id, user_id, event_name, status, kid_id, start_date, end_date, color, COALESCE(description, '') as description, "+
			"alert, COALESCE(city, '') as city, COALESCE(state, '') as state, COALESCE(event_repeat, '') as event_repeat, timezone_offset, date_created, last_updated FROM calendar_event WHERE "+
			"YEAR(start_date) = ? AND MONTH(start_date) = ? AND DAY(start_date) = ?", t.Year(), t.Month(), t.Day())
		break
	case "MONTH":
		err = db.Select(&events, "SELECT id, user_id, event_name, status, kid_id, start_date, end_date, color, COALESCE(description, '') as description, "+
			"alert, COALESCE(city, '') as city, COALESCE(state, '') as state, COALESCE(event_repeat, '') as event_repeat, timezone_offset, date_created, last_updated FROM calendar_event WHERE "+
			"YEAR(start_date) = ? AND MONTH(start_date) = ?", t.Year(), t.Month())
		break
	}
	if err != nil {
		log.Printf("Error on retriving calendar Event. %#v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retrieving calendar event",
			"error":   err,
		})
		return
	}

	//Retrieve todos
	if len(events) > 0 {
		for key, value := range events {
			events[key].Todo, err = retrieveTodosByEventID(db, value.ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Something wrong when retrieving todos",
					"error":   err,
				})
				return
				break
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func retrieveTodosByEventID(db *sqlx.DB, eventID int64) ([]model.Todo, error) {
	var todoList []model.Todo

	err := db.Select(&todoList, "SELECT id, date_created, last_updated, status, text FROM todo_list WHERE event_id = ?", eventID)

	if err != nil {
		return todoList, err
	}

	return todoList, nil
}
