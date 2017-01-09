package model

import "time"

type SubHost struct {
	ID            int64     `json:"id" db:"id"`
	MacID         string    `json:"macId" db:"mac_id"`
	RequestFromID int64     `json:"requestFromID" db:"request_from_id"`
	RequestToID   int64     `json:"requestToID" db:"request_to_id"`
	Status        string    `json:"status" db:"status"`
	DateCreated   time.Time `json:"createdDate" db:"date_created"`
	LastUpdated   time.Time `json:"lastUpdated" db:"last_updated"`
	Kid           []Kid     `json:"kid"`
}

type RequestSubHostWithMacIDRequest struct {
	MacID    string `json:"macId" binding:"required"`
	HostID   int64  `json:"hostId" db:"request_to_id" binding:"required"`
	UserID   int64  `db:"request_from_id"`
	Status   string `db:"status"`
	DeviceID int64  `db:"device_id"`
}

type RequestSubHostToUser struct {
	HostID int64  `json:"hostId" db:"request_to_id" binding:"required"`
	UserID int64  `db:"request_from_id"`
	Status string `db:"status"`
}

type UpdateSubHostRequest struct {
	SubHostID int64   `json:"subHostId" binding:"required"`
	KidID     []int64 `json:"kidId"`
}
