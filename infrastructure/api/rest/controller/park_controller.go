package controller

import (
	"errors"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/api/rest/controller/interfaces"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/api/rest/presenter"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/usecase/dto"
	interfaceUsecase "github.com/Paulo-Lopes-Estevao/go_clean_architecture/usecase/interfaces"
)

type ParkController struct {
	parkUsecase interfaceUsecase.ParkUsecaseInterface
}

func NewParkController(pUsecase interfaceUsecase.ParkUsecaseInterface) interfaces.IParkController {
	return &ParkController{parkUsecase: pUsecase}
}

func (c *ParkController) Welcome(ctx presenter.PresenterContext) error {

	return ctx.JSON(http.StatusOK, "... Welcome API REST")
}

var input dto.ParkDtoInput

func (c *ParkController) AddPark(ctx presenter.PresenterContext) error {

	if err := ctx.Bind(&input); !errors.Is(err, nil) {
		return ctx.JSON(http.StatusBadRequest, presenter.ResponseData(map[string]interface{}{"error": err.Error()}))
	}

	output, err := c.parkUsecase.RegistreNewPark(&input)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusCreated, output)
}
