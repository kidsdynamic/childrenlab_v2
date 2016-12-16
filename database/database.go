package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

func New(database model.Database) *sqlx.DB {
	db := sqlx.MustConnect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", database.User, database.Password, database.IP, database.Name))

	return db
}
