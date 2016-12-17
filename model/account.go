package model

import (
	"database/sql"
	"time"
)

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenRequest struct {
	Email string `json:"email" binding:"required"`
	Token string `json:"token" binding:"required"`
}

type AccessToken struct {
	ID          int64     `db:"id"`
	Email       string    `db:"email"`
	Token       string    `db:"token"`
	LastUpdated time.Time `db:"last_updated"`
}

type User struct {
	ID          int64          `db:"id"`
	Email       string         `db:"email"`
	FirstName   sql.NullString `db:"first_name"`
	LastName    sql.NullString `db:"last_name"`
	LastUpdated string         `db:"last_updated"`
	DateCreated time.Time      `db:"date_created"`
	ZipCode     sql.NullString `db:"zip_code"`
}

type Register struct {
	Email       string `json:"email" db:"email" binding:"required"`
	Password    string `json:"password" db:"password" binding:"required"`
	FirstName   string `json:"firstName" db:"first_name" binding:"required"`
	LastName    string `json:"lastName" db:"last_name" binding:"required"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number"`
	ZipCode     string `json:"zipCode" db:"zip_code"`
}
