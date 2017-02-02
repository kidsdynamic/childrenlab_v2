package model

import "time"

type Kid struct {
	ID          int64     `json:"id" gorm:"AUTO_INCREMENT;primary_key:true"`
	Name        string    `json:"name"`
	DateCreated time.Time `json:"dateCreated"`
	MacID       string    `json:"macId"`
	Profile     string    `json:"profile"`
	Parent      *User     `json:"parent,omitempty"`
	ParentID    int64     `json:"-"`
}

type Device struct {
	ID          int64     `json:"id" gorm:"AUTO_INCREMENT"`
	MacID       string    `json:"macId"`
	DateCreated time.Time `json:"dateCreated"`
}

type KidRequest struct {
	Name    string `json:"name" binding:"required"`
	ZipCode string `json:"zipCode"`
	MacID   string `json:"macId" binding:"required"`
}

type UpdateKidRequest struct {
	ID        int64  `json:"kidId" binding:"required"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	MacID     string `json:"macId"`
}

func (Kid) TableName() string {
	return "kids"
}

func (Device) TableName() string {
	return "device"
}
