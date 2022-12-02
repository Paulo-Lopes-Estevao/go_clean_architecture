package routes

import (
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/api/rest/controller/interfaces"
	"github.com/labstack/echo/v4"
)

type router struct {
	echo            *echo.Echo
	IParkController interfaces.IParkController
}

func NewRouter(e *echo.Echo, IParkController interfaces.IParkController) *router {
	return &router{
		echo:            e,
		IParkController: IParkController,
	}
}
