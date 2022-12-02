package injection

import (
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/adapter/park/controllers"
	interfacesController "github.com/Paulo-Lopes-Estevao/go_clean_architecture/adapter/park/controllers/interfaces"
	repo "github.com/Paulo-Lopes-Estevao/go_clean_architecture/adapter/park/repository"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/entities/repository"
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
	return repo.NewRepositoryFactory(i.DB).RegistryPark()
}
