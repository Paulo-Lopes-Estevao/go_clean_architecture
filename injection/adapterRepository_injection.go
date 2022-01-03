package injection

import (
	"github.com/Paulo-Lopes-Estevao/go_clean_architecture/injection/interfaces"
	"github.com/jinzhu/gorm"
)

type repositoryAdapter struct {
	DB *gorm.DB
}

func NewRegistry(db *gorm.DB) interfaces.InjectionInterface {
	return &repositoryAdapter{DB: db}
}

func (i *repositoryAdapter) NewAppController() interfaces.ControllerAdapter {
	return interfaces.ControllerAdapter{
		Park: i.NewRestController(),
	}
}
