package patterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {

	bikeBuilder := getBuilder(BikeType)
	carBuilder := getBuilder(CarType)

	buildDirector := newDirector(bikeBuilder)

	bike := buildDirector.buildVehicle()
	assert.IsType(t, Vehicle{}, bike)

	buildDirector.setBuilder(carBuilder)

	car := buildDirector.buildVehicle()
	assert.IsType(t, Vehicle{}, car)

}
