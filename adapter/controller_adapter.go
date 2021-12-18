package adapter

import (
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/interface/park/controllers"
)

type ControllerAdapter struct {
	Park interface {
		controllers.ParkControllerInterface
	}
}
