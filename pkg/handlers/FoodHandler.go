package handlers

import (
	"Go-Basic-Restful-APi/pkg/mocks"
	"Go-Basic-Restful-APi/pkg/models"
	"Go-Basic-Restful-APi/pkg/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func GetAllFoods(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")

	foods := repositories.GetAllFoods()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(foods)
}

func GetFood(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	selectedID, _ := strconv.Atoi(params["id"])

	food := repositories.GetFood(selectedID)

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(food)
}

func UpdateFood(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	params := mux.Vars(request)
	body, err := ioutil.ReadAll(request.Body)
	selectedID, _ := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalln(err)
	}

	var selectedFood models.Food
	json.Unmarshal(body, &selectedFood)

	food := repositories.UpdateFood(selectedID, selectedFood)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(food)
}

func DeleteFood(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var selectedFood models.Food
	json.Unmarshal(body, &selectedFood)

	for i, food := range mocks.Foods {
		if food.Id == selectedFood.Id {
			mocks.Foods = append(mocks.Foods[:i], mocks.Foods[i+1:]...)
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode("Deleted")
			break
		}
	}
}

//Wait Gorup Example
func AddFood(writer http.ResponseWriter, request *http.Request) {

	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var food models.Food
	json.Unmarshal(body, &food)

	var wg sync.WaitGroup
	wg.Add(1)
	repositories.AddFood(food, &wg)
	wg.Wait() // We are going to wait for waitGroup complete its job in repository.

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode("Succesfully Inserted")
}
