package usecase

import (
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/entities"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/entities/repository"
)

type ParkUsecaseInterface interface {
	RegistreNewPark(input *ParkDtoInput) (ParkDtoOutput, error)
}

type parkUsecase struct {
	repository repository.ParkRepositoryInterface
}

func NewParkUsecase(repo repository.ParkRepositoryInterface) ParkUsecaseInterface {
	return &parkUsecase{repository: repo}
}

func (usecase *parkUsecase) RegistreNewPark(input *ParkDtoInput) (ParkDtoOutput, error) {
	park := entities.NewPark()
	park.Park.Name = input.Name
	park.Vague = input.Vague
	park.Limit = input.Limit
	err := park.IsValidLimitPark()
	if err != nil {
		return usecase.parkNotActived(input, err)
	}
	return usecase.parkActived(input)

}

func (usecase *parkUsecase) parkActived(park *ParkDtoInput) (ParkDtoOutput, error) {
	err := usecase.repository.Registre(park.Name, park.Limit, park.Vague, true)
	if err != nil {
		return ParkDtoOutput{}, err
	}

	parkOutput := ParkDtoOutput{
		Status:       true,
		ErrorMessage: "",
	}

	return parkOutput, nil
}

func (usecase *parkUsecase) parkNotActived(park *ParkDtoInput, parkinvalid error) (ParkDtoOutput, error) {
	err := usecase.repository.Registre(park.Name, park.Limit, park.Vague, false)
	if err != nil {
		return ParkDtoOutput{}, err
	}

	parkOutput := ParkDtoOutput{
		Status:       false,
		ErrorMessage: parkinvalid.Error(),
	}

	return parkOutput, nil
}
