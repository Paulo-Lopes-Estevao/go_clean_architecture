package interfaces

import "github.com/Paulo-Lopes-Estevao/go_clean_architecture/entities/repository"

type RepositoryFactoryInterface interface {
	RegistryPark() repository.ParkRepositoryInterface
}
