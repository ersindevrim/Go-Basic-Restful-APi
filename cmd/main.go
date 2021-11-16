package main

import (
	"Go-Basic-Restful-APi/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/foods", handlers.GetAllFoods).Methods("GET")
	router.HandleFunc("/foods", handlers.AddFood).Methods("POST")
	router.HandleFunc("/foods/{id}", handlers.GetFood).Methods("GET")
	router.HandleFunc("/foods/{id}", handlers.UpdateFood).Methods("PUT")
	router.HandleFunc("/foods/{id}", handlers.UpdateFood).Methods("DELETE")

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
