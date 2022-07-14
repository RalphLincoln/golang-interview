package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"interview_restapi/models"
	"net/http"
)

func (h Handler) GetMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id := params["id"]

	var movie models.Movie

	all := fmt.Sprintf("SELECT * FROM movies WHERE id='%s'", id)

	h.DB.Raw(all).Scan(&movie)

	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(movie)
	if err != nil {
		return
	}
}
