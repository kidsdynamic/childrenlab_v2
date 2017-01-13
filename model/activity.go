package model

import "time"

const TimeLayout = "2006-01-02T15:04:05Z"

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
	KidID        string    `json:"kidId" db:"kid_id"`
	Type         string    `json:"type" db:"type"`
	Steps        int64     `json:"steps" db:"steps"`
	Distance     int64     `json:"distance" db:"distance"`
	ReceivedDate time.Time `json:"receivedDate" db:"received_date"`
}

type ActivityRequest struct {
	KidID  int64  `json:"kidId" binding:"required"`
	Period string `json:"period" binding:"required"`
}

/*
func (t *Activity) MarshalJSON() ([]byte, error) {
	type Alias Activity
	return json.Marshal(&struct {
		*Alias
		ReceivedDate string `json:"receivedDate"`
	}{
		Alias:        (*Alias)(t),
		ReceivedDate: t.ReceivedDate.Format(TimeLayout),
	})
}
*/
