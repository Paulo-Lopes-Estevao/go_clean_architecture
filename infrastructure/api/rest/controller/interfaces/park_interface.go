package interfaces

import "github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/api/rest/presenter"


type IParkController interface {
	AddPark(ctx presenter.PresenterContext) error
	Welcome(ctx presenter.PresenterContext) error
}
