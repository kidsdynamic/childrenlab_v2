package model

import (
	"database/sql"
	"time"
)

type User struct {
	Email       string         `db:"email"`
	FirstName   sql.NullString `db:"first_name"`
	LastName    sql.NullString `db:"last_name"`
	Password    string         `db:"password"`
	LastUpdated string         `db:"last_updated"`
	DateCreated time.Time      `db:"date_created"`
	ZipCode     sql.NullString `db:"zip_code"`
}
