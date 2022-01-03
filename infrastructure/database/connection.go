package database

import (
	"fmt"

	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/database/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func ConnectionDB() *gorm.DB {

	db, err := gorm.Open("sqlite3", "lumbre.db")

	if err != nil {
		defer db.Close()
		panic("failed to connect database")

	}

	fmt.Println("Connection Opened to Database")

	db.AutoMigrate(&models.ParkVague{})

	fmt.Println("Database Migrated")

	return db

}
