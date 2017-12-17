package model

import "time"

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenRequest struct {
	Email string `query:"email" binding:"required"`
	Token string `query:"token" binding:"required"`
}

type AccessToken struct {
	ID          int64     `gorm:"AUTO_INCREMENT"`
	Email       string    `gorm:"unique;not null"`
	Token       string    `gorm:"not null"`
	LastUpdated time.Time `gorm:"not null"`
	AccessCount int
}

type ProfileUpdateRequest struct {
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	ZipCode     string `json:"zipCode,omitempty"`
}

type User struct {
	ID                       int64     `json:"id" gorm:"AUTO_INCREMENT"`
	Email                    string    `json:"email" gorm:"unique"`
	Password                 string    `json:"-"`
	FirstName                string    `json:"firstName" gorm:"not null"`
	LastName                 string    `json:"lastName" gorm:"not null"`
	LastUpdated              time.Time `json:"-"`
	DateCreated              time.Time `json:"dateCreated"`
	ZipCode                  string    `json:"zipCode"`
	PhoneNumber              string    `json:"phoneNumber"`
	Profile                  string    `json:"profile"`
	Language                 string    `json:"language"`
	RegistrationID           string    `json:"-"`
	AndroidRegistrationToken string    `json:"-"`
	Role                     Role      `json:"-"`
	RoleID                   int64     `json:"-"`
	ResetPasswordToken       string    `json:"-"`
	SignUpIP                 string    `json:"-"`
	SignUpCountryCode        string    `json:"country"`
}

type Role struct {
	ID        int64  `json:"-" gorm:"AUTO_INCREMENT"`
	Authority string `json:"authority" gorm:"unique"`
}

var ROLE_ADMIN string = "ROLE_ADMIN"
var ROLE_USER string = "ROLE_USER"
var ROLE_SUPER_ADMIN string = "ROLE_SUPER_ADMIN"

type RegisterRequest struct {
	Email       string `json:"email" db:"email" binding:"required"`
	Password    string `json:"password" db:"password" binding:"required"`
	FirstName   string `json:"firstName" db:"first_name" binding:"required"`
	LastName    string `json:"lastName" db:"last_name" binding:"required"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number"`
	ZipCode     string `json:"zipCode" db:"zip_code"`
	Language    string `json:"language"`
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

type UpdatePasswordReq struct {
	NewPassword string `json:"newPassword"`
}
