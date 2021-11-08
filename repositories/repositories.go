package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	Planets PlanetsRepository
	SWAPI   SWAPIRepository
	db      *mongo.Client
}

func NewRepositories(db *mongo.Client) *Repositories {

	return &Repositories{
		Planets: NewPlanetsRepository(db),
		SWAPI:   NewSWAPIRepository(),
		db:      db,
	}
}
