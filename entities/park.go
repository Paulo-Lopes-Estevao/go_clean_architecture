package entities

import (
	"errors"
)

type Park struct {
	Name string
}

type ParkVague struct {
	Park  Park
	Vague int32
	Limit int32
}

func NewPark() *ParkVague {
	return &ParkVague{}
}

func (park *ParkVague) getVerifyVague() error {
	if park.Vague <= 1 {
		return errors.New("exceeded the limit")
	}
	return nil
}

func (park *ParkVague) AvailableVacancies() (int32, error) {
	err := park.getVerifyVague()
	if err != nil {
		return 0, err
	}
	result := park.Vague - 1
	return result, nil
}

func (park *ParkVague) IsValidLimitPark() error {
	if park.Limit < 5 {
		return errors.New("very lower limit")
	}

	if park.Vague > park.Limit {
		return errors.New("vague upper limit")
	}

	if park.Vague <= 2 {
		return errors.New("vague inválid")
	}
	return nil
}
