package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type CaseTestPlanetIsValid struct {
	Planet       Planet
	MessageError string
}

func TestPlanet_IsValid(t *testing.T) {

	casesTestes := []CaseTestPlanetIsValid{
		{Planet: Planet{Name: "", Climate: "Arid", Terrain: "Desert"}, MessageError: "Name not informed"},
		{Planet: Planet{Name: "Tatooine", Climate: "", Terrain: "Desert"}, MessageError: "Climate not informed"},
		{Planet: Planet{Name: "Tatooine", Climate: "Arid", Terrain: ""}, MessageError: "Terrain not informed"},
	}

	for _, caseTest := range casesTestes {
		erro := caseTest.Planet.IsValid()
		assert.EqualError(t, erro, caseTest.MessageError)
	}
}
func TestPlanet_Valid(t *testing.T) {

	planetTest := Planet{Name: "Tatooine", Climate: "Arid", Terrain: "Desert"}

	erro := planetTest.IsValid()
	assert.NoError(t, erro)
}
