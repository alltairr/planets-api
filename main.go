package main

import (
	"fmt"
	"log"
	"net/http"
	"planets-api/config"
	"planets-api/controllers"
	"planets-api/repositories"
	"planets-api/routers"
)

func init() {
	config.Initialize()
}

func main() {

	dbClient := repositories.NewDBClient()

	repositories := repositories.NewRepositories(dbClient)

	controllers := controllers.NewController(repositories)

	route := routers.NewRouter(controllers)

	fmt.Printf("Running in port %d\r\n", config.PORT)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), route))
}
