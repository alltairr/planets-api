package routers

import (
	"planets-api/controllers"
	"planets-api/routers/route"

	"github.com/gorilla/mux"
)

func NewRouter(crt *controllers.Controllers) *mux.Router {
	router := mux.NewRouter()
	return route.Configure(router, crt)
}
