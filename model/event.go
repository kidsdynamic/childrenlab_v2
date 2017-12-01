package model

import "time"

type Event struct {
	ID             int64     `json:"id" gorm:"AUTO_INCREMENT"`
	User           User      `json:"user" gorm:"not null"`
	UserID         int64     `json:"-"`
	Kid            []Kid     `json:"kid" gorm:"many2many:event_kid;"`
	Name           string    `json:"name" gorm:"not null"`
	PushTimeUTC    time.Time `json:"-"`
	Start          time.Time `json:"startDate" gorm:"not null"`
	End            time.Time `json:"endDate" gorm:"not null"`
	Color          string    `json:"color" gorm:"not null"`
	Status         string    `json:"status" gorm:"not null"`
	Description    string    `json:"description"`
	Alert          int64     `json:"alert"`
	City           string    `json:"-"`
	State          string    `json:"-"`
	Repeat         string    `json:"repeat"`
	TimezoneOffset int64     `json:"timezoneOffset"`
	DateCreated    time.Time `json:"dateCreated"`
	LastUpdated    time.Time `json:"lastUpdated"`
	Todo           []Todo    `json:"todo,omitempty" `
}

type EventKid struct {
	EventID int64 `gorm:"primary_key:true"`
	KidID   int64 `gorm:"primary_key:true"`
}

type Todo struct {
	ID          int64     `json:"id" gorm:"AUTO_INCREMENT"`
	Text        string    `json:"text"`
	Status      string    `json:"status" gorm:"not null"`
	DateCreated time.Time `json:"dateCreated" gorm:"not null"`
	LastUpdated time.Time `json:"lastUpdated" gorm:"not null"`
	EventID     int64     `json:"-"`
}

func (Event) TableName() string {
	return "event"
}

type EventRequest struct {
	UserID         int64     `db:"user_id"`
	KidID          []int64   `json:"kidId" binding:"required"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	Start          time.Time `json:"startDate"`
	End            time.Time `json:"endDate"`
	Color          string    `json:"color"`
	Description    string    `json:"description"`
	Alert          int64     `json:"alert"`
	City           string    `json:"city"`
	State          string    `json:"state"`
	Repeat         string    `json:"repeat"`
	TimezoneOffset int64     `json:"timezoneOffset"`
	Todo           []string  `json:"todo"`
}

type UpdateEventRequest struct {
	ID             int64     `json:"eventId" binding:"required"`
	Name           string    `json:"name" binding:"required"`
	Start          time.Time `json:"startDate" binding:"required"`
	End            time.Time `json:"endDate" binding:"required"`
	Color          string    `json:"color" binding:"required"`
	Description    string    `json:"description"`
	Alert          int64     `json:"alert"`
	City           string    `json:"city"`
	State          string    `json:"state"`
	Repeat         string    `json:"repeat"`
	TimezoneOffset int64     `json:"timezoneOffset"`
	Todo           []string  `json:"todo"`
	KidsID         []int64   `json:"kidId" binding:"required"`
}

type DeleteEventRequest struct {
	EventID int64 `json:"eventId" db:"id" binding:"required"`
}

type GetEventRequest struct {
	Period string `json:"period" binding:"required"`
	Date   string `json:"date" binding:"required"`
}
