package controller

import (
	"log"
	"net/http"

	"time"

	"strconv"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

	if len(request.KidID) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Kid ID is required",
		})
		return
	}

	user := GetSignedInUser(c)

	db := database.NewGORM()
	defer db.Close()

	if !HasPermissionToKid(db, &user, request.KidID) {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "You don't have permission to do it",
		})
		return
	}

	var event model.Event

	event.PushTimeUTC = request.Start.Add(time.Duration(-request.TimezoneOffset) * time.Minute)

	event.Status = EVENT_OPEN
	event.User = user
	event.Alert = request.Alert
	event.City = request.City
	event.State = request.State
	event.Repeat = request.Repeat
	event.Name = request.Name
	event.Start = request.Start
	event.End = request.End
	event.Color = request.Color
	event.Description = request.Description
	event.DateCreated = GetNowTime()
	event.LastUpdated = GetNowTime()
	event.TimezoneOffset = request.TimezoneOffset

	var kids []model.Kid
	if err := db.Model(model.Kid{}).Where("id in (?)", request.KidID).Find(&kids).Error; err != nil {
		log.Printf("Error on retrieve Kid. %#v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on retrieve Kid",
			"error":   err.Error(),
		})
		return
	}

	event.Kid = kids

	var todos []model.Todo

	for _, todoReq := range request.Todo {
		var todo model.Todo

		todo.Status = TODO_PENDING
		todo.Text = todoReq
		todo.LastUpdated = GetNowTime()
		todo.DateCreated = GetNowTime()

		todos = append(todos, todo)
	}

	event.Todo = todos

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

	user := GetSignedInUser(c)

	if err := db.Where("id = ? and user_id = ?", eventRequest.ID, user.ID).Preload("User").Preload("Kid").Preload("Todo").First(&event).Error; err != nil {
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
		todo.LastUpdated = GetNowTime()
		todo.DateCreated = GetNowTime()

		todos = append(todos, todo)
	}

	if len(todos) > 0 {
		if err := db.Delete(model.Todo{}, "event_id = ?", event.ID).Error; err != nil {
			log.Printf("Error on Deleting todo. %#v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error on Delete todo",
				"error":   err.Error(),
			})
			return
		}
		event.Todo = todos
	}

	event.Color = eventRequest.Color
	event.Alert = eventRequest.Alert
	event.Description = eventRequest.Description
	event.Start = eventRequest.Start
	event.End = eventRequest.End
	event.Name = eventRequest.Name
	event.Repeat = eventRequest.Repeat
	event.PushTimeUTC = eventRequest.Start.Add(time.Duration(-eventRequest.TimezoneOffset) * time.Minute)

	if err := db.Model(&model.Event{}).Updates(&event).Error; err != nil {
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

	eventID, err := strconv.ParseInt(eventIDString, 10, 16)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The event id has to be a number",
		})
		return
	}

	db := database.NewGORM()
	defer db.Close()

	var event model.Event

	if err := db.Where("id = ?", eventID).Preload("User").First(&event).Error; err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error when retriving event",
				"error":   err.Error,
			})
			return
		}
	}

	user := GetSignedInUser(c)

	if user.ID != event.User.ID {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "You don't have permission to do it",
		})
		return
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

	if len(event.Todo) > 0 {
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

	if err := db.Delete(&model.EventKid{}, "event_id = ?", eventID).Error; err != nil {
		if err != nil {
			fmt.Printf("Error on deleting event_kid. %#v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error when deleting event_kid",
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
		"event": events,
	})
}

func RetrieveAllEventWithTodoByUser(c *gin.Context) {
	user := GetSignedInUser(c)

	db := database.NewGORM()
	defer db.Close()
	events := []model.Event{}

	var kidsID []model.UserKidIDs
	if err := db.Table("kids").Select("id").Where("parent_id = ?", user.ID).Find(&kidsID).Error; err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retrieving User's kid",
			"error":   err,
		})
		return
	}

	//Find all of events that belong to User's Kid
	if len(kidsID) > 0 {
		if err := db.Model(model.Event{}).Joins("JOIN event_kid ON event.id = event_kid.event_id").Where("event_kid.kid_id in (?)", toString(kidsID)).Group("event.id").Preload("User").Preload("Kid").Preload("Todo").Find(&events).Error; err != nil {

			fmt.Printf("Error on retriving events. %#v", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something wrong when retrieving events",
				"error":   err,
			})
			return
		}
	}

	var otherKidID []model.UserKidIDs
	if err := db.Table("sub_host_kid").Joins("JOIN sub_host ON sub_host.id = sub_host_kid.sub_host_id").Select("kid_id as id").Where("request_from_id = ?", user.ID).Find(&otherKidID).Error; err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retrieving User's kid",
			"error":   err,
		})
		return
	}

	if len(otherKidID) > 0 {
		var otherkidsEvent model.Event
		//Find all of events that belong to Other host's kid
		if err := db.Model(model.Event{}).Joins("JOIN event_kid ON event.id = event_kid.event_id").Where("event_kid.kid_id in (?)", toString(otherKidID)).Preload("User").Preload("Kid").Preload("Todo").Find(&otherkidsEvent).Error; err != nil && err != gorm.ErrRecordNotFound {

			fmt.Printf("Error on retriving events. %#v", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something wrong when retrieving events",
				"error":   err,
			})
			return
		}

		if otherkidsEvent.ID != 0 {
			removeUnacceptableKid(db, &user, &otherkidsEvent)
			events = append(events, otherkidsEvent)
		}

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

func RetrieveEventsByKid(c *gin.Context) {
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

	var kidIDs []int64
	kidIDs = append(kidIDs, kidID)
	if !HasPermissionToKid(db, &user, kidIDs) {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "You don't have permission to access",
		})
		return
	}

	var events []model.Event

	if err := db.Model(model.Event{}).Joins("JOIN event_kid ON event.id = event_kid.event_id").Where("event_kid.kid_id = ?", kidID).Preload("User").Preload("Todo").Find(&events).Error; err != nil {

		fmt.Printf("Error on retriving events. %#v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retrieving events",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, events)
}

func toString(kidsID []model.UserKidIDs) []int64 {
	var ids []int64
	for _, id := range kidsID {
		ids = append(ids, id.ID)
	}

	return ids
}

func removeUnacceptableKid(db *gorm.DB, user *model.User, event *model.Event) {
	var removedCount int = 0
	for key, kid := range event.Kid {
		var exists bool
		row := db.Raw("SELECT EXISTS(SELECT id FROM sub_host s JOIN sub_host_kid sk ON s.id = sk.sub_host_id WHERE s.request_from_id = ? and sk.kid_id = ? and s.status = ? LIMIT 1)", user.ID, kid.ID, SubHostStatusAccepted).Row()

		row.Scan(&exists)
		if !exists {
			if len(event.Kid) > 0 {
				kids := event.Kid
				kids = append(kids[:key-removedCount], kids[key+1-removedCount:]...)
				event.Kid = kids
			}
			removedCount++

		}

	}
}
