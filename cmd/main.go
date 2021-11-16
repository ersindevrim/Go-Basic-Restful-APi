package main

import (
	"Go-Basic-Restful-APi/pkg/handlers"
	"log"
	"net/http"

	_ "github.com/pdrum/swagger-automation/docs"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// swagger:route GET /foods Foods GetAllFoods
	// Returns all foods.
	// responses:
	//   200: FoodResponse

	// Returns Foods.
	// swagger:response FoodResponse
	router.HandleFunc("/foods", handlers.GetAllFoods).Methods("GET")

	// swagger:route POST /foods Foods CreateFood
	// Creates new foods.
	// responses:
	//   200

	// Returns string.
	// swagger:response string
	router.HandleFunc("/foods", handlers.AddFood).Methods("POST")

	// swagger:route Get /foods/{id} Foods GetFoodById
	// Returns spesific food
	// responses:
	//   200: FoodResponse

	// This endpoint returns requested food
	// swagger:response FoodResponse
	router.HandleFunc("/foods/{id}", handlers.GetFood).Methods("GET")

	// swagger:route PUT /foods/{id} Foods UpdateFoodById
	// Updates spesific food
	// responses:
	//   200: FoodResponse

	// This endpoint returns updated food
	// swagger:response FoodResponse
	router.HandleFunc("/foods/{id}", handlers.UpdateFood).Methods("PUT")

	// swagger:route DELETE /foods/{id} Foods DeleteFoodById
	// Updates spesific food
	// responses:
	//   200

	// This endpoint returns string
	// swagger:response string
	router.HandleFunc("/foods/{id}", handlers.UpdateFood).Methods("DELETE")

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
