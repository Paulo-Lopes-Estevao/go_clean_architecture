package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegitrePark(t *testing.T) {
	park := NewPark()
	park.Park.Name = "Viana Park"
	park.Limit = 10
	park.Vague = 7
	assert.Nil(t, park.IsValidLimitPark(), park.IsValidLimitPark())
}

func TestParkVagueAvaliable(t *testing.T) {
	park := ParkVague{
		Park:  Park{"Viana Park"},
		Limit: 10,
		Vague: 3,
	}

	result, err := park.AvailableVacancies()
	assert.Less(t, result, park.Vague)
	assert.Nil(t, err, err)
}
