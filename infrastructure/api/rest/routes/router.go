package routes

import (
	"github.com/labstack/echo/v4"
)

func (r *router) ParkRouter() *echo.Echo {

	r.echo.GET("", func(context echo.Context) error { return r.IParkController.Welcome(context) })
	r.echo.POST("/park", func(context echo.Context) error { return r.IParkController.AddPark(context) })

	return r.echo

}
