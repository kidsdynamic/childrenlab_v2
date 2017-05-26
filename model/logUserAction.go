package model

import "time"

type LogUserAction struct {
	ID          int64     `json:"id" gorm:"AUTO_INCREMENT"`
	User        *User     `json:"user" gorm:"not null"`
	UserID      int64     `json:"-"`
	MacID       *string   `json:"mac_id"`
	Action      string    `json:"action"`
	DateCreated time.Time `json:"dateCreated"`
	LastUpdated time.Time `json:"lastUpdated"`
}
