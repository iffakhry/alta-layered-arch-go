package util

import (
	"database/sql"
	"fmt"
	"sirclo/restapi/layered/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetDBInstance(config *config.AppConfig) *sql.DB {
	if db == nil {
		path := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)
		dbNewInstance, err := sql.Open(config.Database.Driver, path)

		if err != nil {
			panic(err)
		}

		db = dbNewInstance
	}
	return db
}
