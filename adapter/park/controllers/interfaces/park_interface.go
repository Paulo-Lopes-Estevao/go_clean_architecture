package interfaces

import "github.com/Paulo-Lopes-Estevao/go_clean_architecture/adapter/park/presenter"

type ParkControllerInterface interface {
	AddPark(ctx presenter.ParkRestPresenterContext) error
	Welcome(ctx presenter.ParkRestPresenterContext) error
}
