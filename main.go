package main

import (
	"log"

	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/api/rest/routes"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/database"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/injection"
	"github.com/labstack/echo/v4"
)

func main() {

	server := echo.New()

	db := database.ConnectionDB()
	i := injection.NewRegistry(db)

	routes.ParkRoute(server, i.NewAppController())

	if err := server.Start(":8000"); err != nil {
		log.Println("Not Running Server A...", err.Error())
	}
}
