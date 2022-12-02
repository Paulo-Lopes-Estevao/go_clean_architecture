package usecase

import (
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/entities"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/entities/repository"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/usecase/dto"
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/usecase/interfaces"
)

type parkUsecase struct {
	repository repository.ParkRepositoryInterface
}

func NewParkUsecase(repo repository.ParkRepositoryInterface) interfaces.ParkUsecaseInterface {
	return &parkUsecase{repository: repo}
}

func (usecase *parkUsecase) RegistreNewPark(input *dto.ParkDtoInput) (*dto.ParkDtoOutput, error) {
	park := entities.NewPark()
	park.Park.Name = input.Name
	park.Vague = input.Vague
	park.Limit = input.Limit
	err := park.IsValidLimitPark()
	if err != nil {
		return usecase.parkNotActived(park, err)
	}
	return usecase.parkActived(park)

}

func (usecase *parkUsecase) parkActived(park *entities.ParkVague) (*dto.ParkDtoOutput, error) {
	err := usecase.repository.Registre(park.Park.Name, park.Limit, park.Vague, true)
	if err != nil {
		return nil, err
	}

	parkOutput := &dto.ParkDtoOutput{
		Status:       true,
		ErrorMessage: "",
	}

	return parkOutput, nil
}

func (usecase *parkUsecase) parkNotActived(park *entities.ParkVague, parkinvalid error) (*dto.ParkDtoOutput, error) {
	err := usecase.repository.Registre(park.Park.Name, park.Limit, park.Vague, false)
	if err != nil {
		return nil, err
	}

	parkOutput := &dto.ParkDtoOutput{
		Status:       false,
		ErrorMessage: parkinvalid.Error(),
	}

	return parkOutput, nil
}
