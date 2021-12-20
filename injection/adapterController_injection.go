package injection

import (
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/adapter/park/controllers/interfaces"
)

type ControllerAdapter struct {
	Park interface {
		interfaces.ParkControllerInterface
	}
}
