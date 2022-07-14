package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (h Handler) DeleteMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	id := params["id"]

	h.DB.Exec("DELETE FROM movies WHERE id=$1", id)

	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode("Movie Deleted")
	if err != nil {
		return
	}
}
