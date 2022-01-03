package repository

import (
	"fmt"

	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/database/models"
	"gorm.io/gorm"
)

type repositorydb struct {
	Db *gorm.DB
}

func NewRegistreParkRepository(db *gorm.DB) *repositorydb {
	return &repositorydb{Db: db}
}

func (repodb *repositorydb) Registre(name string, vague int32, limit int32, status bool) error {

	data := models.ParkVague{
		Name:  name,
		Vague: vague,
		Limit: limit,
	}
	fmt.Println(&data, status)
	err := repodb.Db.Create(data).Error

	if err != nil {
		return err
	}
	return nil
}