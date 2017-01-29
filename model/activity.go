package model

import "time"

type ActivityRawData struct {
	ID          int64  `json:"id" gorm:"AUTO_INCREMENT"`
	Indoor      string `json:"indoorActivity" binding:"required"`
	Outdoor     string `json:"outdoorActivity" binding:"required"`
	Time        int64  `json:"time" binding:"required"`
	MacID       string `json:"macId" binding:"required"`
	UserID      int64
	DateCreated time.Time
	LastUpdated time.Time
}

func (ActivityRawData) TableName() string {
	return "activity_raw"
}

type ActivityInsight struct {
	Date     time.Time
	TimeLong int64
	Steps    int64
}

type Activity struct {
	ID           int64     `json:"id" gorm:"AUTO_INCREMENT"`
	MacID        string    `json:"macId" gorm:"not null"`
	Kid          Kid       `json:"-"`
	KidID        int64     `json:"kidId" gorm:"not null"`
	Type         string    `json:"type" gorm:"not null"`
	Steps        int64     `json:"steps"`
	Distance     int64     `json:"distance"`
	ReceivedDate time.Time `json:"receivedDate"`
	ReceivedTime int64     `json:"-"`
	DateCreated  time.Time
	LastUpdated  time.Time
}

type ActivityRequest struct {
	KidID  int64  `json:"kidId" binding:"required"`
	Period string `json:"period" binding:"required"`
}
