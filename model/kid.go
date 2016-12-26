package model

import "time"

type Kid struct {
	ID          int64     `json:"id" db:"id"`
	FirstName   string    `json:"firstName" db:"first_name"`
	LastName    string    `json:"lastName" db:"last_name"`
	DateCreated time.Time `json:"dateCreated" db:"date_created"`
	MacID       string    `json:"macId" db:"mac_id"`
	Profile     string    `json:"profile" db:"profile"`
}

type Device struct {
	ID          int64     `json:"id" db:"id"`
	MacID       string    `json:"macId" db:"mac_id"`
	DateCreated time.Time `json:"dateCreated" db:"date_created"`
}

type KidRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	MacID     string `json:"macId" binding:"required"`
}

type UpdateKidRequest struct {
	ID        int64  `json:"kidId" binding:"required"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	MacID     string `json:"macId"`
}
