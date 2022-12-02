package entities

import "fmt"

type CarPark struct {
	Type  string
	Price float32
	Time  int32
}

func heavy() *CarPark {
	return &CarPark{
		Type:  "heavy",
		Price: 5.000,
		Time:  24,
	}
}

func light() *CarPark {
	return &CarPark{
		Type:  "light",
		Price: 2.500,
		Time:  24,
	}
}

func NewTypeVehicle(typecar string) (*CarPark, error) {
	if typecar == "heavy" {
		return heavy(), nil
	}

	if typecar == "slight" {
		return light(), nil
	}
	return nil, fmt.Errorf("type of car does not exist")
}
