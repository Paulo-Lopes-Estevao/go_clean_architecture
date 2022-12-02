package server

import (
	"log"

	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/api/dependency/injection"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ServerRestApi() {
	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	db := database.ConnectionDB()
	defer db.Close()

	injection.Injection(db, server)

	if err := server.Start(":8000"); err != nil {
		log.Println("Not Running Server A...", err.Error())
	}
}
