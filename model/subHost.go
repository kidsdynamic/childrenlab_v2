package model

import "time"

type SubHost struct {
	ID            int64     `json:"id" gorm:"AUTO_INCREMENT;primary_key:true"`
	RequestFrom   User      `json:"requestFromUser"`
	RequestFromID int64     `json:"-"`
	RequestTo     User      `json:"requestToUser"`
	RequestToID   int64     `json:"-"`
	Status        string    `json:"status" gorm:"default:'PENDING'"`
	DateCreated   time.Time `json:"createdDate"`
	LastUpdated   time.Time `json:"lastUpdated"`
	Kids          []Kid     `json:"kids,omitempty" gorm:"many2many:sub_host_kid;"`
}

type SubHostKid struct {
	SubHostID int64 `gorm:"primary_key:true"`
	KidID     int64 `gorm:"primary_key:true"`
}

type RequestSubHostWithMacIDRequest struct {
	MacID    string `json:"macId" binding:"required"`
	HostID   int64  `json:"hostId" db:"request_to_id" binding:"required"`
	UserID   int64  `db:"request_from_id"`
	Status   string `db:"status"`
	DeviceID int64  `db:"device_id"`
}

type RequestSubHostToUser struct {
	HostID int64 `json:"hostId" binding:"required"`
}

type DenyRequest struct {
	SubHostID int64 `json:"subHostId" binding:"required"`
}

type AcceptSubHostRequest struct {
	SubHostID int64   `json:"subHostId" binding:"required"`
	KidID     []int64 `json:"kidId"`
}

type RemoveSubHostRequest struct {
	SubHostID int64 `json:"subHostId" binding:"required"`
	KidID     int64 `json:"kidId"  binding:"required"`
}

type AddKidToSubHost struct {
	SubHostID int64 `json:"subHostId" binding:"required"`
	KidID     int64 `json:"kidId" binding:"required"`
}

func (SubHost) TableName() string {
	return "sub_host"
}
