package main

import (
	"log"

	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/api/rest/routes"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/database"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/injection"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func main() {

	server := echo.New()

	database.ConnectionDB()
	i := injection.NewRegistry(&gorm.DB{})

	routes.ParkRoute(server, i.NewAppController())

	if err := server.Start(":8000"); err != nil {
		log.Println("Not Running Server A...", err.Error())
	}
}
