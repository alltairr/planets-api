package route

import (
	"net/http"
	"planets-api/controllers"

	"github.com/gorilla/mux"
)

type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func Configure(router *mux.Router, crt *controllers.Controllers) *mux.Router {
	routers := CreatePlanetsRouters(crt.PlanetsCtr)

	for _, route := range routers {
		router.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	return router
}
