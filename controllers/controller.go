package controllers

import "planets-api/repositories"

type Controllers struct {
	PlanetsCtr PlanetController
}

func NewController(rp *repositories.Repositories) *Controllers {
	return &Controllers{
		PlanetsCtr: NewPlanetController(rp.Planets, rp.SWAPI),
	}
}
