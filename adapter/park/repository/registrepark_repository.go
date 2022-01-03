package repository

import (
	repo "github.com/Paulo-Lopes-Estevao/go_clean_architecture/entities/repository"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/database/repository"
	"gorm.io/gorm"
)

type RepositoryFactory struct {
	DB *gorm.DB
}

func NewRepositoryFactory(db *gorm.DB) *RepositoryFactory {
	return &RepositoryFactory{DB: db}
}

func (f RepositoryFactory) CreateRegistryPark() repo.ParkRepositoryInterface {
	return repository.NewRegistreParkRepository(f.DB)
}
