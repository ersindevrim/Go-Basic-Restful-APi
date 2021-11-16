package handlers

import (
	"Go-Basic-Restful-APi/pkg/mocks"
	"Go-Basic-Restful-APi/pkg/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func GetAllFoods(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(mocks.Foods)
}

func GetFood(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	selectedID, _ := strconv.Atoi(params["id"])

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, _ := sql.Open("postgres", psqlconn)
	defer db.Close()

	var food models.Food
	db.QueryRow(`SELECT "Id","Name","Desc","Photo" FROM "Food" WHERE "Id" = $1`, selectedID).Scan(&food.Id, &food.Name, &food.Desc, &food.Photo)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(food)
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

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, _ := sql.Open("postgres", psqlconn)

	defer db.Close()

	insertDynStmt := `insert into "Food"("Name", "Desc","Photo") values($1, $2, $3)`
	insertedFood, _ := db.Exec(insertDynStmt, food.Name, food.Desc, food.Photo)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(insertedFood)
}
