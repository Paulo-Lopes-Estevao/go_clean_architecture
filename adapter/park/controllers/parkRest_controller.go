package controllers

import (
	"errors"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/adapter/park/controllers/interfaces"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/adapter/park/presenter"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/usecase/dto"
	interfaceUsecase "github.com/Paulo-Lopes-Estevao/go_clean_architecture/usecase/interfaces"
)

type ParkController struct {
	parkUsecase interfaceUsecase.ParkUsecaseInterface
}

func NewParkController(pUsecase interfaceUsecase.ParkUsecaseInterface) interfaces.ParkControllerInterface {
	return &ParkController{parkUsecase: pUsecase}
}

func (c *ParkController) Welcome(ctx presenter.ParkRestPresenterContext) error {

	return ctx.JSON(http.StatusOK, "... Welcome API REST")
}

var input dto.ParkDtoInput

func (c *ParkController) AddPark(ctx presenter.ParkRestPresenterContext) error {

	if err := ctx.Bind(&input); !errors.Is(err, nil) {
		ctx.JSON(http.StatusBadRequest, presenter.ResponseData(map[string]interface{}{"error": err.Error()}))
	}

	output, err := c.parkUsecase.RegistreNewPark(&input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, output)
}
