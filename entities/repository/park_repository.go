package repository

type ParkRepositoryInterface interface {
	Registre(name string, vague int32, limit int32, status bool) error
}
