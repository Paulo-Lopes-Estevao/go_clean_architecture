package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const DB_USERNAME = "postgres"
const DB_PASSWORD = "postgres"
const DB_NAME = "db_lubre"
const DB_HOST = "localhost"
const DB_PORT = "5432"

func ConnectionDB() *gorm.DB {

	dsn := "host=" + DB_HOST + " user=" + DB_USERNAME + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " port=" + DB_PORT + " sslmode=disable"

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		defer db.Close()
		panic("failed to connect database")

	}

	return db

}
