package interfaces

import (
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/usecase/dto"
)

type ParkUsecaseInterface interface {
	RegistreNewPark(input *dto.ParkDtoInput) (dto.ParkDtoOutput, error)
}
