package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"planets-api/models"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SWAPIRepositoryMock struct {
	MockGetPlanetByName func(name string) (models.PlanetSW, error)
}

func (m *SWAPIRepositoryMock) GetPlanetByName(name string) (models.PlanetSW, error) {
	return m.MockGetPlanetByName(name)
}

type PlanetRepositoryMock struct {
	MockGetAll    func() ([]models.Planet, error)
	MockGetById   func(id string) (models.Planet, error)
	MockGetByName func(name string) ([]models.Planet, error)
	MockCreate    func(planet *models.Planet) error
	MockUpdate    func(id string, planet models.Planet) error
	MockDelete    func(id string) error
}

func (m *PlanetRepositoryMock) GetAll() ([]models.Planet, error) {
	return m.MockGetAll()
}
func (m *PlanetRepositoryMock) GetById(id string) (models.Planet, error) {
	return m.MockGetById(id)
}
func (m *PlanetRepositoryMock) GetByName(name string) ([]models.Planet, error) {
	return m.MockGetByName(name)
}
func (m *PlanetRepositoryMock) Create(planet *models.Planet) error {
	return m.MockCreate(planet)
}
func (m *PlanetRepositoryMock) Update(id string, planet models.Planet) error {
	return m.MockUpdate(id, planet)
}
func (m *PlanetRepositoryMock) Delete(id string) error {
	return m.MockDelete(id)
}

func TestGetAllPlanet_Success(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "/planets", nil)
	res := httptest.NewRecorder()

	expectedPlanets := []models.Planet{
		{ID: primitive.NewObjectID(), Name: "Tatooine", Climate: "arid", Terrain: "desert"},
		{ID: primitive.NewObjectID(), Name: "Alderaan", Climate: "temperate", Terrain: "grasslands, mountains"},
	}

	mockSWAPIRepository := &SWAPIRepositoryMock{}

	mockPlanetsRepository := &PlanetRepositoryMock{}
	mockPlanetsRepository.MockGetAll = func() ([]models.Planet, error) {
		return expectedPlanets, nil
	}

	planetController := NewPlanetController(mockPlanetsRepository, mockSWAPIRepository)

	planetController.GetPlanets(res, req)

	planetsResult := []models.Planet{}

	_ = json.Unmarshal(res.Body.Bytes(), &planetsResult)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.EqualValues(t, expectedPlanets, planetsResult)
	assert.EqualValues(t, len(planetsResult), 2)
}

func TestGetAllPlanetWithSearchName_Success(t *testing.T) {

	nameSearch := "Tatooine"

	req, _ := http.NewRequest(http.MethodGet, "/planets?search="+nameSearch, nil)
	res := httptest.NewRecorder()

	expectedPlanets := []models.Planet{
		{ID: primitive.NewObjectID(), Name: "Tatooine", Climate: "arid", Terrain: "desert"},
	}

	mockSWAPIRepository := &SWAPIRepositoryMock{}

	mockPlanetsRepository := &PlanetRepositoryMock{}
	mockPlanetsRepository.MockGetByName = func(name string) ([]models.Planet, error) {
		return expectedPlanets, nil
	}

	planetController := NewPlanetController(mockPlanetsRepository, mockSWAPIRepository)

	planetController.GetPlanets(res, req)

	planetsResult := []models.Planet{}

	_ = json.Unmarshal(res.Body.Bytes(), &planetsResult)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.EqualValues(t, expectedPlanets, planetsResult)
	assert.EqualValues(t, 1, len(planetsResult))
}

func TestCreatePlanet_Success(t *testing.T) {

	var jsonRequest = []byte(`
		{
			"name": "Tatooine",
			"Climate": "Frio",
			"Terrain": "Luz"
		}
	`)

	req, _ := http.NewRequest(http.MethodPost, "/planets", bytes.NewBuffer(jsonRequest))
	res := httptest.NewRecorder()

	expectedPlanet := models.Planet{
		ID:             primitive.NewObjectID(),
		Name:           "Tatooine",
		Climate:        "Frio",
		Terrain:        "Luz",
		ViewedQuantity: 5,
	}

	mockSWAPIRepository := &SWAPIRepositoryMock{}
	mockSWAPIRepository.MockGetPlanetByName = func(name string) (models.PlanetSW, error) {
		var planet = models.PlanetSW{
			Films: []string{
				"Filme 1",
				"Filme 2",
				"Filme 3",
				"Filme 5",
				"Filme 5",
			},
		}

		return planet, nil
	}

	mockPlanetsRepository := &PlanetRepositoryMock{}
	mockPlanetsRepository.MockCreate = func(planet *models.Planet) error {
		planet.ID = expectedPlanet.ID
		return nil
	}

	planetController := NewPlanetController(mockPlanetsRepository, mockSWAPIRepository)

	planetController.CreatePlanet(res, req)

	planetResult := models.Planet{}

	_ = json.Unmarshal(res.Body.Bytes(), &planetResult)

	assert.Equal(t, http.StatusCreated, res.Code)
	assert.EqualValues(t, expectedPlanet, planetResult)
}

func TestCreatePlanetJsonEmpty_Error(t *testing.T) {

	var jsonRequest = []byte("")

	req, _ := http.NewRequest(http.MethodPost, "/planets", bytes.NewBuffer(jsonRequest))
	res := httptest.NewRecorder()

	mockSWAPIRepository := &SWAPIRepositoryMock{}
	mockPlanetsRepository := &PlanetRepositoryMock{}
	mockPlanetsRepository.MockCreate = func(planet *models.Planet) error {
		return nil
	}

	planetController := NewPlanetController(mockPlanetsRepository, mockSWAPIRepository)

	planetController.CreatePlanet(res, req)

	planetResult := models.Planet{}
	_ = json.Unmarshal(res.Body.Bytes(), &planetResult)

	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func TestCreatePlanetInvalidJson_Error(t *testing.T) {

	var jsonRequest = []byte(`
		{
			"name": "Tatooine",
			"climated": "Frio",
			"terrain": "Luz"
		}
	`)

	req, _ := http.NewRequest(http.MethodPost, "/planets", bytes.NewBuffer(jsonRequest))
	res := httptest.NewRecorder()

	mockSWAPIRepository := &SWAPIRepositoryMock{}
	mockPlanetsRepository := &PlanetRepositoryMock{}
	mockPlanetsRepository.MockCreate = func(planet *models.Planet) error {
		return nil
	}

	planetController := NewPlanetController(mockPlanetsRepository, mockSWAPIRepository)

	planetController.CreatePlanet(res, req)

	planetResult := models.Planet{}
	_ = json.Unmarshal(res.Body.Bytes(), &planetResult)

	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func TestUpdatePlanet_Success(t *testing.T) {

	ID, _ := primitive.ObjectIDFromHex("61896345959f3d16e4549737")
	var jsonRequest = []byte(`
		{
			"id": "61896345959f3d16e4549737",
			"name": "Tatooine",
			"climate": "frio",
			"terrain": "areia"
		}
	`)

	req, _ := http.NewRequest(http.MethodPut, "/planets", bytes.NewBuffer(jsonRequest))
	res := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{
		"planetId": "61896345959f3d16e4549737",
	})

	expectedPlanet := models.Planet{
		ID:             ID,
		Name:           "Tatooine",
		Climate:        "frio",
		Terrain:        "areia",
		ViewedQuantity: 3,
	}

	mockSWAPIRepository := &SWAPIRepositoryMock{}
	mockSWAPIRepository.MockGetPlanetByName = func(name string) (models.PlanetSW, error) {
		var planet = models.PlanetSW{
			Films: []string{
				"Filme 1",
				"Filme 2",
				"Filme 5",
			},
		}
		return planet, nil
	}

	mockPlanetsRepository := &PlanetRepositoryMock{}
	mockPlanetsRepository.MockGetById = func(id string) (models.Planet, error) {
		return expectedPlanet, nil
	}

	mockPlanetsRepository.MockUpdate = func(id string, planet models.Planet) error {
		return nil
	}

	planetController := NewPlanetController(mockPlanetsRepository, mockSWAPIRepository)

	planetController.UpdatePlanet(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestUpdatePlanet_Error(t *testing.T) {

	ID, _ := primitive.ObjectIDFromHex("61896345959f3d16e4549737")
	var jsonRequest = []byte(`
		{
			"id": "61896345959f3d16e4549737",
			"name": "Tatooine",
			"climate": "frio",
			"terrain": "areia"
		}
	`)

	req, _ := http.NewRequest(http.MethodPut, "/planets", bytes.NewBuffer(jsonRequest))
	res := httptest.NewRecorder()

	expectedPlanet := models.Planet{
		ID:             ID,
		Name:           "Tatooine",
		Climate:        "frio",
		Terrain:        "areia",
		ViewedQuantity: 3,
	}

	mockSWAPIRepository := &SWAPIRepositoryMock{}
	mockSWAPIRepository.MockGetPlanetByName = func(name string) (models.PlanetSW, error) {
		var planet = models.PlanetSW{
			Films: []string{
				"Filme 1",
				"Filme 2",
				"Filme 5",
			},
		}
		return planet, nil
	}

	mockPlanetsRepository := &PlanetRepositoryMock{}
	mockPlanetsRepository.MockGetById = func(id string) (models.Planet, error) {
		return expectedPlanet, nil
	}

	mockPlanetsRepository.MockUpdate = func(id string, planet models.Planet) error {
		return nil
	}

	planetController := NewPlanetController(mockPlanetsRepository, mockSWAPIRepository)

	planetController.UpdatePlanet(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func TestDeletePlanet_Success(t *testing.T) {

	ID, _ := primitive.ObjectIDFromHex("61896345959f3d16e4549737")

	req, _ := http.NewRequest(http.MethodDelete, "/planets", nil)
	res := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{
		"planetId": "61896345959f3d16e4549737",
	})

	expectedPlanet := models.Planet{
		ID:             ID,
		Name:           "Tatooine",
		Climate:        "frio",
		Terrain:        "areia",
		ViewedQuantity: 3,
	}

	mockSWAPIRepository := &SWAPIRepositoryMock{}

	mockPlanetsRepository := &PlanetRepositoryMock{}
	mockPlanetsRepository.MockGetById = func(id string) (models.Planet, error) {
		return expectedPlanet, nil
	}

	mockPlanetsRepository.MockDelete = func(id string) error {
		return nil
	}

	planetController := NewPlanetController(mockPlanetsRepository, mockSWAPIRepository)

	planetController.DeletePlanet(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
