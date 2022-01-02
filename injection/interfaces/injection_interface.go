package interfaces

import (
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/adapter/park/controllers/interfaces"
)

type ControllerAdapter struct {
	Park interface {
		interfaces.ParkControllerInterface
	}
}

type InjectionInterface interface {
	NewAppController() ControllerAdapter
}
