package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVehicleofheavy(t *testing.T) {
	_, err := NewTypeVehicle("pesado")
	assert.Nil(t, err, err)
}

func TestNewVehicleoflight(t *testing.T) {
	_, err := NewTypeVehicle("ligeiro")
	assert.Nil(t, err, err)
}
