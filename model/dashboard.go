package model

type Dashboard struct {
	TotalUserCount     int64                 `json:"totalUserCount"`
	Signup             []SignupCountByDate   `json:"signup"`
	TotalActivityCount int64                 `json:"totalActivityCount"`
	Activity           []ActivityCountByDate `json:"activity"`
}

type SignupCountByDate struct {
	SignupCount int64  `json:"signupCount"`
	Date        string `json:"date"`
}

type ActivityCountByDate struct {
	ActivityCount int64  `json:"activityCount"`
	UserCount     int64  `json:"userCount"`
	Date          string `json:"date"`
}
