package route

import (
	"net/http"
	"planets-api/controllers"
)

func CreatePlanetsRouters(ctr controllers.PlanetController) []Route {
	return []Route{
		{
			URI:     "/planets",
			Method:  http.MethodGet,
			Handler: ctr.GetPlanets,
		},
		{
			URI:     "/planets/{planetId}",
			Method:  http.MethodGet,
			Handler: ctr.GetPlanet,
		},
		{
			URI:     "/planets",
			Method:  http.MethodPost,
			Handler: ctr.CreatePlanet,
		},
		{
			URI:     "/planets/{planetId}",
			Method:  http.MethodPut,
			Handler: ctr.UpdatePlanet,
		},
		{
			URI:     "/planets/{planetId}",
			Method:  http.MethodDelete,
			Handler: ctr.DeletePlanet,
		},
	}
}
