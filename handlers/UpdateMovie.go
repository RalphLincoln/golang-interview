package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"interview_restapi/models"
	"net/http"
)

func (h Handler) UpdateMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	id := params["id"]

	var updatedMovie models.Movie
	var movie models.Movie

	_ = json.NewDecoder(request.Body).Decode(&updatedMovie)

	if result := h.DB.First(&movie, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	movie.Title = updatedMovie.Title
	movie.ReleaseDate = updatedMovie.ReleaseDate
	movie.USGross = updatedMovie.USGross
	movie.WorldwideGross = updatedMovie.WorldwideGross
	movie.IMDBVotes = updatedMovie.IMDBVotes
	movie.Director = updatedMovie.Director

	h.DB.Save(&movie)

	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode("Movie Updated")
	if err != nil {
		return
	}
}
