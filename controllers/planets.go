package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"planets-api/models"
	"planets-api/repositories"
	"planets-api/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type planetController struct {
	Repository repositories.PlanetsRepository
	SWAPI      repositories.SWAPIRepository
}

var _ PlanetController = &planetController{}

func NewPlanetController(planetRepository repositories.PlanetsRepository, swapiRepository repositories.SWAPIRepository) *planetController {
	return &planetController{
		Repository: planetRepository,
		SWAPI:      swapiRepository,
	}
}

type PlanetController interface {
	GetPlanet(w http.ResponseWriter, r *http.Request)
	GetPlanets(w http.ResponseWriter, r *http.Request)
	CreatePlanet(w http.ResponseWriter, r *http.Request)
	UpdatePlanet(w http.ResponseWriter, r *http.Request)
	DeletePlanet(w http.ResponseWriter, r *http.Request)
}

func (ctrl *planetController) GetPlanet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	planetID := params["planetId"]

	if !primitive.IsValidObjectID(planetID) {
		utils.ResponseError(w, http.StatusBadRequest, errors.New("Id not invalid: "+planetID))
		return
	}

	planet, erro := ctrl.Repository.GetById(planetID)
	if erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}
	utils.ResponseJSON(w, http.StatusOK, planet)
}

func (ctrl *planetController) GetPlanets(w http.ResponseWriter, r *http.Request) {

	var erro error
	var planets []models.Planet

	searchQuery := r.FormValue("search")

	if searchQuery != "" {
		planets, erro = ctrl.Repository.GetByName(searchQuery)
	} else {
		planets, erro = ctrl.Repository.GetAll()
	}

	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, planets)
}

func (ctrl *planetController) CreatePlanet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	jsonRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		utils.ResponseError(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var planet models.Planet
	if erro = json.Unmarshal(jsonRequest, &planet); erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	if erro := planet.IsValid(); erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	planetSwapi, erro := ctrl.SWAPI.GetPlanetByName(planet.Name)
	if erro != nil {
		planet.ViewedQuantity = 0
	} else {
		planet.ViewedQuantity = len(planetSwapi.Films)
	}

	erro = ctrl.Repository.Create(&planet)
	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, planet)
}

func (ctrl *planetController) UpdatePlanet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	params := mux.Vars(r)
	planetID := params["planetId"]

	if !primitive.IsValidObjectID(planetID) {
		utils.ResponseError(w, http.StatusBadRequest, errors.New("Id not invalid: "+planetID))
		return
	}

	planetExist, erro := ctrl.Repository.GetById(planetID)
	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	if planetExist == (models.Planet{}) {
		utils.ResponseError(w, http.StatusNotFound, nil)
		return
	}

	jsonRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		utils.ResponseError(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var planet models.Planet
	if erro = json.Unmarshal(jsonRequest, &planet); erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	if erro := planet.IsValid(); erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	planetSwapi, erro := ctrl.SWAPI.GetPlanetByName(planet.Name)
	if erro != nil {
		planet.ViewedQuantity = 0
	} else {
		planet.ViewedQuantity = len(planetSwapi.Films)
	}

	erro = ctrl.Repository.Update(planetID, planet)
	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, nil)
}

func (ctrl *planetController) DeletePlanet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	planetID := params["planetId"]

	if !primitive.IsValidObjectID(planetID) {
		utils.ResponseError(w, http.StatusBadRequest, errors.New("Id not invalid: "+planetID))
		return
	}

	planetExist, erro := ctrl.Repository.GetById(planetID)
	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	if planetExist == (models.Planet{}) {
		utils.ResponseError(w, http.StatusNotFound, nil)
		return
	}

	erro = ctrl.Repository.Delete(planetID)
	if erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, nil)
}
