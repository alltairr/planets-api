package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	MONGO_URI      = ""
	MONGO_DATABASE = ""
	PORT           = 0
	SWAPI_BASE_URL = ""
)

//Initialize config
func Initialize() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	MONGO_URI = os.Getenv("MONGO_URI")
	if MONGO_URI == "" {
		log.Fatal("Variable of Environment MONGO_URI not informed")
	}

	MONGO_DATABASE = os.Getenv("MONGO_DATABASE")
	if MONGO_DATABASE == "" {
		log.Fatal("Variable of Environment MONGO_DATABASE not informed")
	}

	SWAPI_BASE_URL = os.Getenv("SWAPI_BASE_URL")
	if SWAPI_BASE_URL == "" {
		log.Fatal("Variable of Environment SWAPI_BASE_URL not informed")
	}

	PORT, erro = strconv.Atoi(os.Getenv("PORT"))

	if erro != nil {
		PORT = 9000
	}
}
