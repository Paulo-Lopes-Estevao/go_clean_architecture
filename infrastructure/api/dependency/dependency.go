package dependency

import "github.com/Paulo-Lopes-Estevao/go_clean_architecture/infrastructure/api/rest/controller/interfaces"



type UserDependency struct {
	IParkController        interfaces.IParkController
}