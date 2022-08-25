package router

import (
	controller "github.com/Ajay123-web/MONGODB/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api" , controller.ServeHome).Methods("GET")
	router.HandleFunc("/api/movies" , controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie" , controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}" , controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}" , controller.DeleteMovie).Methods("DELETE")

	return router
}