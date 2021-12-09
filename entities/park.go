package entities

import "errors"

type Park struct {
	Name string
}

type ParkVague struct {
	Park  Park
	Vague int32
	Limit int32
}

type ParkVagueType struct {
	Type  string
	Price float32
	Time  int32
}

func NewPark() *ParkVague {
	return &ParkVague{}
}

func (park *ParkVague) AvailableVacancies(vague int) {

}

func (park *ParkVague) IsValidLimitPark() error {
	if park.Limit < 5 {
		return errors.New("very lower limit")
	}

	if park.Vague > park.Limit {
		return errors.New("vague upper limit")
	}
	return nil
}
