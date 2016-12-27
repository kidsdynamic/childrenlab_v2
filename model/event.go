package model

import "time"

type Event struct {
	ID             int64     `json:"id" db:"id"`
	UserID         int64     `json:"userId" db:"user_id"`
	KidID          int64     `json:"kidId" db:"kid_id"`
	Name           string    `json:"name" db:"event_name"`
	Start          time.Time `json:"startDate" db:"start_date"`
	End            time.Time `json:"endDate" db:"end_date"`
	Color          string    `json:"color" db:"color"`
	Status         string    `json:"status" db:"status"`
	Description    string    `json:"description" db:"description"`
	Alert          int64     `json:"alert" db:"alert"`
	City           string    `json:"city" db:"city"`
	State          string    `json:"state" db:"state"`
	Repeat         string    `json:"repeat" db:"event_repeat"`
	TimezoneOffset int64     `json:"timezoneOffset" db:"timezone_offset"`
	DateCreated    time.Time `json:"dateCreated" db:"date_created"`
	LastUpdated    time.Time `json:"lastUpdated" db:"last_updated"`
	Todo           []Todo    `json:"todo,omitempty" `
}

type Todo struct {
	ID          int64     `json:"id" db:"id"`
	Text        string    `json:"text" db:"text"`
	Status      string    `json:"status" db:"status"`
	DateCreated time.Time `json:"dateCreated" db:"date_created"`
	LastUpdated time.Time `json:"lastUpdated" db:"last_updated"`
}

type EventRequest struct {
	UserID         int64       `db:"user_id"`
	KidID          int64       `json:"kidId" db:"kid_id" binding:"required"`
	Name           string      `json:"name" db:"event_name" binding:"required"`
	Status         string      `json:"status" db:"status"`
	Start          interface{} `json:"startDate" db:"start_date" binding:"required"`
	End            interface{} `json:"endDate" db:"end_date" binding:"required"`
	Color          string      `json:"color" db:"color" binding:"required"`
	Description    string      `json:"description" db:"description"`
	Alert          int64       `json:"alert" db:"alert"`
	City           string      `json:"city" db:"city"`
	State          string      `json:"state" db:"state"`
	Repeat         string      `json:"repeat" db:"event_repeat"`
	TimezoneOffset int64       `json:"timezoneOffset" db:"timezone_offset" binding:"required"`
	Todo           []string    `json:"todo"`
}

type UpdateEventRequest struct {
	ID             int64       `json:"eventId" db:"id" binding:"required"`
	Name           string      `json:"name" db:"event_name" binding:"required"`
	Start          interface{} `json:"startDate" db:"start_date" binding:"required"`
	End            interface{} `json:"endDate" db:"end_date" binding:"required"`
	Color          string      `json:"color" db:"color" binding:"required"`
	Description    string      `json:"description" db:"description"`
	Alert          int64       `json:"alert" db:"alert"`
	City           string      `json:"city" db:"city"`
	State          string      `json:"state" db:"state"`
	Repeat         string      `json:"repeat" db:"event_repeat"`
	TimezoneOffset int64       `json:"timezoneOffset" db:"timezone_offset" binding:"required"`
	Todo           []string    `json:"todo"`
}

type DeleteEventRequest struct {
	EventID int64 `json:"eventId" db:"id" binding:"required"`
}

type GetEventRequest struct {
	Period string `json:"period" binding:"required"`
	Date   string `json:"date" binding:"required"`
}
