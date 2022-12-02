package repository

import (
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/adapter/park/repository/interfaces"
	repo "github.com/Paulo-Lopes-Estevao/go_clean_architecture/entities/repository"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/database/repository"
	"github.com/jinzhu/gorm"
)

type RepositoryFactory struct {
	DB *gorm.DB
}

func NewRepositoryFactory(db *gorm.DB) interfaces.RepositoryFactoryInterface {
	return &RepositoryFactory{DB: db}
}

func (f RepositoryFactory) RegistryPark() repo.ParkRepositoryInterface {
	return repository.NewRegistreParkRepository(f.DB)
}
