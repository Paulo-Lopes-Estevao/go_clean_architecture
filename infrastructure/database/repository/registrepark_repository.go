package repository

import (
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/entities/repository"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/database/models"
	"gorm.io/gorm"
)

type repositorydb struct {
	Db *gorm.DB
}

func NewRegistreParkRepository(db *gorm.DB) repository.ParkRepositoryInterface {
	return &repositorydb{Db: db}
}

func (repodb *repositorydb) Registre(name string, vague int32, limit int32, status bool) error {

	data := models.ParkVague{
		Name:  name,
		Vague: vague,
		Limit: limit,
	}

	err := repodb.Db.Create(&data).Error

	if err != nil {
		return err
	}
	return nil
}
