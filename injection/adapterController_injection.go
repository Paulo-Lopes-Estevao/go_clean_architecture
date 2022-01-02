package injection

import (
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/adapter/park/controllers"
	interfacesController "github.com/Paulo-Lopes-Estevao/go_clean_architecture/adapter/park/controllers/interfaces"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/entities/repository"
	repo "github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/database/repository"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/usecase"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/usecase/interfaces"
)

func (i *repositoryAdapter) NewRestUseCase() interfaces.ParkUsecaseInterface {
	return usecase.NewParkUsecase(i.NewRestRepository())
}

func (i *repositoryAdapter) NewRestController() interfacesController.ParkControllerInterface {
	return controllers.NewParkController(i.NewRestUseCase())
}

func (i *repositoryAdapter) NewRestRepository() repository.ParkRepositoryInterface {
	return repo.NewRegistreParkRepository(i.DB)
}
