package handlers

import (
	"encoding/json"
	"fmt"
	"interview_restapi/models"
	"math"
	"net/http"
	"strconv"
)

func (h Handler) GetMovies(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var movie []models.Movie
	pageParams := request.URL.Query()["page"]
	recPerPageParams := request.URL.Query()["record"]

	var page string
	var total int64
	var perPage string

	if len(pageParams) == 0 {
		page = "1"
	} else {
		page = pageParams[0]
	}
	if len(recPerPageParams) == 0 {
		perPage = "10"
	} else {
		perPage = recPerPageParams[0]
	}
	parsedPage, _ := strconv.Atoi(page)
	parsedPerPage, _ := strconv.Atoi(perPage)

	all := "SELECT * FROM movies"
	cmdPage := fmt.Sprintf("%s LIMIT %d OFFSET %d", all, parsedPerPage, (parsedPage-1)*parsedPerPage)

	h.DB.Raw(all).Count(&total)

	h.DB.Raw(cmdPage).Scan(&movie)

	//log.Println(parsedPage, total)

	writer.WriteHeader(http.StatusOK)
	if len(movie) > 0 {
		err := json.NewEncoder(writer).Encode(map[string]any{
			"data":            movie,
			"current_page":    parsedPage,
			"data_count":      total,
			"record_per_page": parsedPerPage,
			"num_of_pages":    math.Ceil(float64(total/int64(parsedPerPage))) + 1, //TODO FIND A WAY TO DEAL WITH REMAINDER FOR ACCURACY
		})
		if err != nil {
			return
		}
	} else {
		err := json.NewEncoder(writer).Encode(map[string]any{
			"data":            []models.Movie{},
			"current_page":    parsedPage,
			"data_count":      total,
			"record_per_page": parsedPerPage,
			"num_of_pages":    math.Ceil(float64(total/int64(parsedPerPage))) + 1, //TODO FIND A WAY TO DEAL WITH REMAINDER FOR ACCURACY
		})
		if err != nil {
			return
		}
	}
}
