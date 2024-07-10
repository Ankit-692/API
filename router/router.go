package router

import (
	"github.com/Ankit-692/API/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/api/movie", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovies).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteONeMovie).Methods("DELETE")
	router.HandleFunc("/api/movie/delete-all", controller.DeleteAll).Methods("POST")

	return router
}