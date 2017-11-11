package model

import "time"

type FwFile struct {
	ID           int64     `json:"id"`
	Version      string    `json:"version"`
	FileAURL     string    `json:"fileAUrl"`
	FileBURL     string    `json:"fileBUrl"`
	UploadedDate time.Time `json:"uploadedDate"`
	Active       bool      `json:"active",gorm:"default:false"`
}
