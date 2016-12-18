package model

import "time"

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

type ProfileUpdateRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	ZipCode     string `json:"zipCode"`
}

type User struct {
	ID          int64     `json:"id" db:"id"`
	Email       string    `json:"email" db:"email"`
	FirstName   string    `json:"firstName" db:"first_name"`
	LastName    string    `json:"lastName" db:"last_name"`
	LastUpdated string    `json:"lastUpdate" db:"last_updated"`
	DateCreated time.Time `json:"dateCreated" db:"date_created"`
	ZipCode     string    `json:"zipCode" db:"zip_code"`
	PhoneNumber string    `json:"phoneNumber" db:"phone_number"`
	Profile     string    `json:"profile" db:"profile"`
}

type Register struct {
	Email       string `json:"email" db:"email" binding:"required"`
	Password    string `json:"password" db:"password" binding:"required"`
	FirstName   string `json:"firstName" db:"first_name" binding:"required"`
	LastName    string `json:"lastName" db:"last_name" binding:"required"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number"`
	ZipCode     string `json:"zipCode" db:"zip_code"`
}
