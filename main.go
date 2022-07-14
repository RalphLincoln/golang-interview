package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"interview_restapi/db"
	"interview_restapi/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//INIT ROUTER
	r := mux.NewRouter()

	// DB CONNECTION
	DB := db.InitDB()

	//HANDLER TO SEND DB CONNECTIONS TO ENDPOINTS
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

	port := os.Getenv("PORT")

	url := fmt.Sprintf(":%s", "3000")

	//SERVER SET UP
	srv := &http.Server{
		Handler:      r,
		Addr:         url,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("API is running!")
	log.Println(port, "<=port", url, "<=url")
	log.Fatal(srv.ListenAndServe())
}
