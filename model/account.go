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
	ID          int64     `gorm:"AUTO_INCREMENT"`
	Email       string    `gorm:"unique;not null"`
	Token       string    `gorm:"not null"`
	LastUpdated time.Time `gorm:"not null"`
	AccessCount int
}

type ProfileUpdateRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	ZipCode     string `json:"zipCode"`
}

type User struct {
	ID             int64     `json:"id" gorm:"AUTO_INCREMENT"`
	Email          string    `json:"email" gorm:"unique"`
	Password       string    `json:"-" gorm:"unique"`
	FirstName      string    `json:"firstName" gorm:"not null"`
	LastName       string    `json:"lastName" gorm:"not null"`
	LastUpdated    time.Time `json:"lastUpdate"`
	DateCreated    time.Time `json:"dateCreated"`
	ZipCode        string    `json:"zipCode"`
	PhoneNumber    string    `json:"phoneNumber"`
	Profile        string    `json:"profile"`
	RegistrationID string    `json:"registrationId,omitempty"`
	Role           Role      `json:"role"`
	RoleID         int64     `json:"-"`
}

type Role struct {
	ID        int64  `json:"-" gorm:"AUTO_INCREMENT"`
	Authority string `json:"authority" gorm:"unique"`
}

func (User) TableName() string {
	return "user"
}

func (Role) TableName() string {
	return "role"
}

func (AccessToken) TableName() string {
	return "authentication_token"
}
