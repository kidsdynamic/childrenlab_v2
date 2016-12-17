package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

var DatabaseInfo model.Database

func New() *sqlx.DB {
	db := sqlx.MustConnect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		DatabaseInfo.User, DatabaseInfo.Password, DatabaseInfo.IP, DatabaseInfo.Name))

	return db
}
