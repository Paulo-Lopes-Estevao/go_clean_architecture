package database

import (
	"fmt"

	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/database/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func ConnectionDB() {

	db, err := gorm.Open("sqlite3", "lumbre.db")

	if err != nil {

		panic("failed to connect database")

	}

	defer db.Close()

	fmt.Println("Connection Opened to Database")

	db.AutoMigrate(&models.ParkVague{})

	fmt.Println("Database Migrated")

}
