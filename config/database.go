package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_NAME = "golang_mysql"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return DB
}
