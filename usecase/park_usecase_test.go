package usecase_test

import (
	"testing"

	repository_mock "github.com/Paulo-Lopes-Estevao/go_clean_architecture/entities/repository/mock"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewParkActived(t *testing.T) {
	input := usecase.ParkDtoInput{
		Name:  "Viana Park",
		Limit: 20,
		Vague: 15,
	}

	outputExpected := usecase.ParkDtoOutput{
		Status:       true,
		ErrorMessage: "",
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	repositoryMockPark := repository_mock.NewMockParkRepository(controller)
	repositoryMockPark.EXPECT().
		Registre(input.Name, input.Limit, input.Vague, outputExpected.Status).
		Return(nil)

	usecasePark := usecase.NewParkUsecase(repositoryMockPark)

	outputPark, err := usecasePark.RegistreNewPark(input)
	assert.Nil(t, err, err)
	assert.Equal(t, outputExpected, outputPark)

}

func TestNewParkNotActivedLowerLimit(t *testing.T) {
	input := usecase.ParkDtoInput{
		Name:  "Viana Park",
		Limit: 4,
		Vague: 15,
	}

	outputExpected := usecase.ParkDtoOutput{
		Status:       false,
		ErrorMessage: "very lower limit",
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	repositoryMockPark := repository_mock.NewMockParkRepository(controller)
	repositoryMockPark.EXPECT().
		Registre(input.Name, input.Limit, input.Vague, outputExpected.Status).
		Return(nil)

	usecasePark := usecase.NewParkUsecase(repositoryMockPark)

	outputPark, err := usecasePark.RegistreNewPark(input)
	assert.Nil(t, err, err)
	assert.Equal(t, outputExpected, outputPark)

}
