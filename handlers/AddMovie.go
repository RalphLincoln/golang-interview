package handlers

import (
	"encoding/json"
	"github.com/leonelquinteros/gorand"
	"interview_restapi/models"
	"log"
	"net/http"
)

func (h Handler) AddMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var movie models.Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)

	log.Println(request.Body, "BODY", movie)

	out, err := gorand.GetAlphaNumString(24)
	if err != nil {
		log.Fatal(err)
	}

	movie.ID = out

	if result := h.DB.Create(&movie); result.Error != nil {
		log.Fatalln(result.Error)
	}

	writer.WriteHeader(http.StatusCreated)
	errs := json.NewEncoder(writer).Encode("Movie added successfully")
	if errs != nil {
		return
	}
}
