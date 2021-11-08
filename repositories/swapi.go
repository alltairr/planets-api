package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"planets-api/config"
	"planets-api/models"
	"time"
)

type swapiRepository struct {
}

var _ SWAPIRepository = &swapiRepository{}

type SWAPIRepository interface {
	GetPlanetByName(name string) (models.PlanetSW, error)
}

func NewSWAPIRepository() *swapiRepository {
	return &swapiRepository{}
}

func (swapi *swapiRepository) GetPlanetByName(name string) (models.PlanetSW, error) {

	httpClient := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, erro := httpClient.Get(config.SWAPI_BASE_URL + "/planets?search=" + name)
	if erro != nil {
		fmt.Printf("Error %s", erro)
		return models.PlanetSW{}, erro
	}
	defer resp.Body.Close()

	var planet models.PlanetSW

	jsonResponse, erro := ioutil.ReadAll(resp.Body)
	if erro != nil {
		fmt.Printf("Error %s", erro)
		return planet, erro
	}

	var resultsPaged models.Paged
	if erro = json.Unmarshal(jsonResponse, &resultsPaged); erro != nil {
		return planet, erro
	}

	planet = resultsPaged.Results[0]

	return planet, nil
}
