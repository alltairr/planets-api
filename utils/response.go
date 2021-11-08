package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseError(w http.ResponseWriter, statusCode int, erro error) {
	ResponseJSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}

func ResponseJSON(w http.ResponseWriter, statusCode int, payload interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if payload == nil {
		return
	}

	if erro := json.NewEncoder(w).Encode(payload); erro != nil {
		log.Fatal(erro)
	}
}
