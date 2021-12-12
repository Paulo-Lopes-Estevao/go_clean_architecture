package repository

type ParkRepositoryInterface interface {
	Registre(Name string, Vague int32, Limit int32, Status bool) error
}
