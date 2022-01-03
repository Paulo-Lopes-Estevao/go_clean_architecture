package repository_test

import (
	"testing"

	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/adapter/park/repository"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestCreatePark(t *testing.T) {
	db := database.ConnectionDB()
	repositories := repository.NewRepositoryFactory(db)
	err := repositories.CreateRegistryPark().
		Registre("Viana Park", 10, 15, true)

	if err != nil {
		assert.Nil(t, err.Error(), "Não criado")
	}

	assert.Nil(t, err)

}
