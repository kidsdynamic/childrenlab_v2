package model

import "time"

type SubHost struct {
	ID            int64     `json:"id" gorm:"AUTO_INCREMENT"`
	RequestFrom   User      `json:"requestFromUser"`
	RequestFromID int64     `json:"-"`
	RequestTo     User      `json:"requestToUser"`
	RequestToID   int64     `json:"-"`
	Status        string    `json:"status" gorm:"default:'PENDING'"`
	DateCreated   time.Time `json:"createdDate"`
	LastUpdated   time.Time `json:"lastUpdated"`
	Kids          []Kid     `json:"kids,omitempty" gorm:"many2many:sub_host_kid"`
}

/*type SubHostKid struct {
	ID        int64 `json:"-" gorm:"AUTO_INCREMENT"`
	Kid       *Kid
	KidID     int64
	SubHostID int64
}*/

type RequestSubHostWithMacIDRequest struct {
	MacID    string `json:"macId" binding:"required"`
	HostID   int64  `json:"hostId" db:"request_to_id" binding:"required"`
	UserID   int64  `db:"request_from_id"`
	Status   string `db:"status"`
	DeviceID int64  `db:"device_id"`
}

type RequestSubHostToUser struct {
	HostID int64 `json:"hostId" db:"request_to_id" binding:"required"`
}

type UpdateSubHostRequest struct {
	SubHostID int64   `json:"subHostId" binding:"required"`
	KidID     []int64 `json:"kidId"`
}

func (SubHost) TableName() string {
	return "sub_host"
}
