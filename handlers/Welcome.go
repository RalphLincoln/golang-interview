package handlers

import (
	"encoding/json"
	"net/http"
)

func WelcomeHome(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode("Welcome to my api. Have fun checking it out.")
	if err != nil {
		return
	}
}
