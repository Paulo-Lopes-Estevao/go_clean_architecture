package presenter

type ParkPresenter struct {
	Name  string  `json:"name"`
	Vague int32   `json:"vague"`
	Limit int32   `json:"limit"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
	Time  int32   `json:"time"`
}
