package injection

import (
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/api/rest/controller"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/api/rest/controller/interfaces"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/api/rest/repository"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/api/rest/routes"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/usecase"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func parkInjectionDependency(db *gorm.DB) interfaces.IParkController {
	park_implement_repository := repository.NewRepositoryFactory(db).RegistryPark()
	park_implement_usecase := usecase.NewParkUsecase(park_implement_repository)
	park_implement_controller := controller.NewParkController(park_implement_usecase)
	return park_implement_controller
}

func Injection(db *gorm.DB, e *echo.Echo) *echo.Echo {
	park := parkInjectionDependency(db)
	echo := routes.NewRouter(e, park)
	return echo.ParkRouter()
}
