package model

import "time"

type FwFile struct {
	ID           int64     `json:"id"`
	Version      string    `json:"version"`
	FileName     string    `json:"fileName"`
	FileURL      string    `json:"fileUrl"`
	UploadedDate time.Time `json:"uploadedDate"`
}
