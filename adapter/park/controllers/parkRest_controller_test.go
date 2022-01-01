package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	repository_mock "github.com/Paulo-Lopes-Estevao/go_clean_architecture/entities/repository/mock"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/usecase"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/likexian/gokit/assert"
)

func TestWelcomeController(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(req, recorder)

	controllerMock := gomock.NewController(t)
	defer controllerMock.Finish()

	repositoryMockPark := repository_mock.NewMockParkRepository(controllerMock)

	usecasepark := usecase.NewParkUsecase(repositoryMockPark)

	controllerpark := NewParkController(usecasepark)

	err := controllerpark.Welcome(ctx)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

}

func TestRegitreParkController(t *testing.T) {
	e := echo.New()

	input := usecase.ParkDtoInput{
		Name:  "Viana Park",
		Limit: 20,
		Vague: 15,
	}

	value, erro := json.Marshal(input)

	if erro != nil {
		assert.Nil(t, erro, "Não converteu para json")
	}

	req := httptest.NewRequest(http.MethodPost, "/park", bytes.NewBuffer(value))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(req, recorder)

	controllerMock := gomock.NewController(t)
	defer controllerMock.Finish()

	repositoryMockPark := repository_mock.NewMockParkRepository(controllerMock)

	usecasepark := usecase.NewParkUsecase(repositoryMockPark)

	controllerpark := NewParkController(usecasepark)

	err := controllerpark.AddPark(ctx)

	if err != nil {
		assert.Equal(t, http.StatusCreated, recorder.Code, "error status code != 201")
	}

}
