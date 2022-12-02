package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVehicleofheavy(t *testing.T) {
	_, err := NewTypeVehicle("heavy")
	assert.Nil(t, err, err)
}

func TestNewVehicleoflight(t *testing.T) {
	_, err := NewTypeVehicle("slight")
	assert.Nil(t, err, err)
}
