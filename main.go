package main

import (
	"github.com/gorilla/mux"
	"interview_restapi/db"
	"interview_restapi/handlers"
	"log"
	"net/http"
	"time"
)

func main() {
	//INIT ROUTER
	r := mux.NewRouter()

	DB := db.InitDB()

	h := handlers.New(DB)

	//ROUTES
	r.HandleFunc("/", handlers.WelcomeHome).Methods(http.MethodGet)

	r.HandleFunc("/movies", h.GetMovies).Methods(http.MethodGet)

	r.HandleFunc("/movie/{id}", h.GetMovie).Methods(http.MethodGet)

	r.HandleFunc("/movies/date", h.GetMovieByDate).Methods(http.MethodGet)

	r.HandleFunc("/movie", h.AddMovie).Methods(http.MethodPost)

	r.HandleFunc("/movie/{id}", h.UpdateMovie).Methods(http.MethodPut)

	r.HandleFunc("/movie/{id}", h.DeleteMovie).Methods(http.MethodDelete)

	r.NotFoundHandler = http.NotFoundHandler()

	r.Use(mux.CORSMethodMiddleware(r))

	//SERVER SET UP
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("API is running!")
	log.Fatal(srv.ListenAndServe())
}