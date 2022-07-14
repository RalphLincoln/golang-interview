package handlers

import (
	"encoding/json"
	"fmt"
	"interview_restapi/models"
	"net/http"
)

func (h Handler) GetMovieByDate(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var movie []models.Movie
	startDateParams := request.URL.Query()["start_date"]
	endDateParams := request.URL.Query()["end_date"]

	var startDate string
	var endDate string

	if len(startDateParams) == 0 {
		err := json.NewEncoder(writer).Encode("Please enter start date")
		if err != nil {
			return
		}
	} else {
		startDate = startDateParams[0]
	}
	if len(endDateParams) == 0 {
		err := json.NewEncoder(writer).Encode("Please enter end date")
		if err != nil {
			return
		}
	} else {
		endDate = endDateParams[0]
	}

	all := fmt.Sprintf("SELECT * FROM movies WHERE created_at >= '%s' AND created_at <= '%s'", startDate, endDate)

	h.DB.Raw(all).Scan(&movie)

	writer.WriteHeader(http.StatusOK)
	if len(movie) > 0 {
		err := json.NewEncoder(writer).Encode(movie)
		if err != nil {
			return
		}
	} else {
		err := json.NewEncoder(writer).Encode([]models.Movie{})
		if err != nil {
			return
		}
	}
}
