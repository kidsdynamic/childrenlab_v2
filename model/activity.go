package model

import "time"

type ActivityRawData struct {
	ID      int64  `json:"id" db:"id"`
	Indoor  string `json:"indoorActivity" db:"indoor_activity" binding:"required"`
	Outdoor string `json:"outdoorActivity" db:"outdoor_activity" binding:"required"`
	Time    int64  `json:"time" db:"time" binding:"required"`
	MacID   string `json:"macId" db:"mac_id" binding:"required"`
}

type ActivityInsight struct {
	Time  time.Time
	Steps int64
}

type ActivityRawDataRequest struct {
	Indoor  string `json:"indoorActivity"`
	Outdoor string `json:"outdoorActivity"`
	Time    int64  `json:"time"`
	MacID   string `json:"macId"`
}

type Activity struct {
	ID           int64     `json:"id" db:"id"`
	MacID        string    `json:"macId" db:"mac_id"`
	Type         string    `json:"type" db:"type"`
	Steps        int64     `json:"steps" db:"steps"`
	Distance     int64     `json:"distance" db:"distance"`
	DateCreated  time.Time `json:"dateCreated" db:"date_created"`
	ReceivedDate time.Time `json:"receivedDate" db:"received_date"`
}
