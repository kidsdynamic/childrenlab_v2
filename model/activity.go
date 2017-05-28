package model

import "time"

type ActivityRawData struct {
	ID             int64     `json:"id" gorm:"AUTO_INCREMENT"`
	Indoor         string    `json:"indoorActivity" binding:"required"`
	Outdoor        string    `json:"outdoorActivity" binding:"required"`
	Time           int64     `json:"time" binding:"required"`
	MacID          string    `json:"macId" binding:"required"`
	TimeZoneOffset int64     `json:"timeZoneOffset"`
	UserID         int64     `json:"userId"`
	IndoorSteps    int64     `json:"indoorSteps" gorm:"default:NULL"`
	OutdoorSteps   int64     `json:"outdoorSteps"`
	DateCreated    time.Time `json:"dateCreated"`
	LastUpdated    time.Time `json:"lastUpdated"`
}

func (ActivityRawData) TableName() string {
	return "activity_raw"
}

type ActivityInsight struct {
	Date     time.Time
	TimeLong int64
	TimeZone int64
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

func (Activity) TableName() string {
	return "activity"
}

type ActivityRequest struct {
	KidID  int64  `json:"kidId" binding:"required"`
	Period string `json:"period" binding:"required"`
}
