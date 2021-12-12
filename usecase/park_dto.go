package usecase

type ParkDtoInput struct {
	Name  string
	Vague int32
	Limit int32
	Type  string
	Price float64
	Time  int32
}

type ParkDtoOutput struct {
	Status       bool
	ErrorMessage string
}
