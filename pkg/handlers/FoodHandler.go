package handlers

import (
	"Go-Basic-Restful-APi/pkg/mocks"
	"Go-Basic-Restful-APi/pkg/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllFoods(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(mocks.Foods)
}

func GetFood(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	selectedID, _ := strconv.Atoi(params["id"])

	for _, item := range mocks.Foods {
		if item.Id == selectedID {
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(item)
			break
		}
	}
}

func UpdateFood(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var selectedFood models.Food
	json.Unmarshal(body, &selectedFood)

	for i, food := range mocks.Foods {
		if food.Id == selectedFood.Id {
			food.Name = selectedFood.Name
			food.Desc = selectedFood.Desc
			food.Photo = selectedFood.Photo

			mocks.Foods[i] = food
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(food)
			break
		}
	}
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

func AddFood(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var food models.Food

	json.Unmarshal(body, &food)

	food.Id = len(mocks.Foods) + 1
	mocks.Foods = append(mocks.Foods, food)

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(food)
}
