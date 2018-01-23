package model

type InitialDeviceFirmware struct {
	ID              int64  `json:"id"`
	MacId           string `json:"macId"`
	FirmwareVersion string `json:"firmwareVersion"`
	Language        string `json:"language"`
	ProductVersion  int64  `json:"productVersion"`
}
