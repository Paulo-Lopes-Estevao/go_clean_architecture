package handler

import (
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/adapter"
	"github.com/labstack/echo/v4"
)

func ParkRoute(e *echo.Echo, ad adapter.ControllerAdapter) *echo.Echo {

	e.POST("/park", func(context echo.Context) error { return ad.Park.AddPark(context) })

	return e

}
