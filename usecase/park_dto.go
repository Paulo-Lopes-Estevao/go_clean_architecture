package usecase

type ParkDtoInput struct {
	Name  string  `json:"name"`
	Vague int32   `json:"vague"`
	Limit int32   `json:"limit"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
	Time  int32   `json:"time"`
}

type ParkDtoOutput struct {
	Status       bool   `json:"status"`
	ErrorMessage string `json:"error_message"`
}
