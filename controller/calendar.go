package controller

import (
	"log"
	"net/http"

	"time"

	"strconv"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kidsdynamic/childrenlab_v2/constants"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

const (
	TODO_PENDING = "PENDING"
	TODO_DONE    = "DONE"
	EVENT_OPEN   = "OPEN"
	EVENT_PASSED = "PASSED"
)

func AddCalendarEvent(c *gin.Context) {
	var request model.EventRequest

	if err := c.BindJSON(&request); err != nil {
		log.Printf("Error on Add Event. Bind with json. %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing some parameters",
			"error":   err,
		})
		return
	}

	user := GetSignedInUser(c)

	var event model.Event

	event.KidID = request.KidID
	event.Status = EVENT_OPEN
	event.UserID = user.ID
	event.Alert = request.Alert
	event.City = request.City
	event.State = request.State
	event.Repeat = request.Repeat
	event.Name = request.Name
	event.Start = request.Start
	event.End = request.End
	event.Color = request.Color
	event.DateCreated = GetNowTime()
	event.LastUpdated = GetNowTime()
	event.TimezoneOffset = request.TimezoneOffset

	var todos []model.Todo

	for _, todoReq := range request.Todo {
		var todo model.Todo

		todo.Status = TODO_PENDING
		todo.Text = todoReq
		todo.LastUpdated = time.Now()
		todo.DateCreated = time.Now()

		todos = append(todos, todo)
	}

	event.Todo = todos

	db := database.NewGORM()
	defer db.Close()

	if err := db.Create(&event).Error; err != nil {
		log.Printf("Error on Add Event. %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on adding the event to database",
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

	db := database.NewGORM()
	defer db.Close()

	var event model.Event

	if err := db.Where("id = ?", eventRequest.ID).First(&event).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can't find the event from database",
			"error":   err.Error(),
		})
		return
	}

	var todos []model.Todo

	for _, todoReq := range eventRequest.Todo {
		var todo model.Todo

		todo.Status = TODO_PENDING
		todo.Text = todoReq
		todo.LastUpdated = time.Now()
		todo.DateCreated = time.Now()

		todos = append(todos, todo)
	}

	if err := db.Delete(model.Todo{}, "event_id = ?", event.ID).Error; err != nil {
		log.Printf("Error on Deleting todo. %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on Delete todo",
			"error":   err.Error(),
		})
		return
	}

	event.Todo = todos

	if err := db.Model(&event).Updates(&eventRequest).Error; err != nil {
		log.Printf("Error on Updat event. %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on updating event",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"event": event,
	})

}

func DeleteEvent(c *gin.Context) {
	eventIDString := c.Query("eventId")

	if eventIDString == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing some parameters",
		})
		return
	}

	eventID, err := strconv.ParseInt(eventIDString, 10, 6)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The event id has to be a number",
		})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	var event model.Event

	if err := db.Where("id = ?", eventID).Preload("Todo").First(&event).Error; err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error when retriving event",
				"error":   err.Error,
			})
			return
		}
	}

	if len(event.Todo) > 0 {
		fmt.Println("TOdo > 0")
		if err := db.Delete(&model.Todo{}, "event_id = ?", eventID).Error; err != nil {
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Error when deleting todo",
					"error":   err.Error,
				})
				return
			}
		}

	}

	if err := db.Delete(&model.Event{}, "id = ?", eventID).Error; err != nil {
		if err != nil {
			fmt.Printf("Error on deleting event. %#v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error when deleting event",
				"error":   err.Error,
			})
			return
		}
	}

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

	db := database.NewGORM()
	defer db.Close()
	var events []model.Event

	t, err := time.Parse(constants.TimeLayout, getEventRequest.Date)
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
		err = db.Where("YEAR(start) = ? AND MONTH(start) = ? AND DAY(start) = ?", t.Year(), t.Month(), t.Day()).Preload("Todo").Find(&events).Error
		break
	case "MONTH":

		err = db.Where("YEAR(start) = ? AND MONTH(start) = ?", t.Year(), t.Month()).Preload("Todo").Find(&events).Error
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

	c.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}

func RetrieveAllEventWithTodoByUser(c *gin.Context) {
	user := GetSignedInUser(c)

	db := database.NewGORM()
	defer db.Close()
	var events []model.Event

	if err := db.Where("user_id = ?", user.ID).Preload("Todo").Find(&events).Error; err != nil {

		fmt.Printf("Error on retriving events. %#v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retrieving events",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, events)
}

func retrieveTodosByEventID(db *sqlx.DB, eventID int64) ([]model.Todo, error) {
	var todoList []model.Todo

	err := db.Select(&todoList, "SELECT id, date_created, last_updated, status, text FROM todo_list WHERE event_id = ?", eventID)

	if err != nil {
		return todoList, err
	}

	return todoList, nil
}

type todoDoneRequest struct {
	EventID int64 `json:"eventId" binding:"required"`
	TodoID  int64 `json:"todoId" binding:"required"`
}

func TodoDone(c *gin.Context) {
	var todoDoneRequest todoDoneRequest

	if err := c.BindJSON(&todoDoneRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Date formate is wrong",
			"error":   err,
		})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	if err := db.Model(&model.Todo{}).Where("id = ? AND event_id = ?", todoDoneRequest.TodoID, todoDoneRequest.EventID).Update("status", TODO_DONE).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retrieving todos",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
