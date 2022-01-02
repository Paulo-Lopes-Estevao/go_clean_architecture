package routes

import (
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/injection/interfaces"
	"github.com/labstack/echo/v4"
)

func ParkRoute(e *echo.Echo, injection interfaces.ControllerAdapter) *echo.Echo {

	e.GET("/", func(context echo.Context) error { return injection.Park.Welcome(context) })
	e.POST("/park", func(context echo.Context) error { return injection.Park.AddPark(context) })

	return e

}
