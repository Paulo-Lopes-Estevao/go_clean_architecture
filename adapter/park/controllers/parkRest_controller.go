package controllers

import (
	"errors"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/adapter/park/controllers/interfaces"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/adapter/park/presenter"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/entities/repository"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/usecase"
)

type ParkController struct {
	parkUsecase usecase.ParkUsecaseInterface
	repository  repository.ParkRepositoryInterface
}

func NewParkController(pUsecase usecase.ParkUsecaseInterface) interfaces.ParkControllerInterface {
	return &ParkController{parkUsecase: pUsecase}
}

func (c *ParkController) Welcome(ctx presenter.ParkRestPresenterContext) error {

	return ctx.JSON(http.StatusOK, "... Welcome API REST")
}

func (c *ParkController) AddPark(ctx presenter.ParkRestPresenterContext) error {
	input := &usecase.ParkDtoInput{}

	if err := ctx.Bind(input); !errors.Is(err, nil) {
		ctx.JSON(http.StatusBadRequest, presenter.ResponseData(map[string]interface{}{"error": err.Error()}))
	}

	usecasepark := usecase.NewParkUsecase(c.repository)
	output, _ := usecasepark.RegistreNewPark(*input)

	return ctx.JSON(http.StatusCreated, output)
}
