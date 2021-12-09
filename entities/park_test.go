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
