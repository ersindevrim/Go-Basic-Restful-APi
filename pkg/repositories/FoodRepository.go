package repositories

import (
	"Go-Basic-Restful-APi/pkg/models"
	"database/sql"
	"sync"
)

var mutex sync.Mutex

const (
	psqlconn = "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
)

func GetFood(id int) models.Food {
	db, _ := sql.Open("postgres", psqlconn)
	defer db.Close()

	var food models.Food
	db.QueryRow(`SELECT "Id","Name","Desc","Photo" FROM "Food" WHERE "Id" = $1`, id).Scan(&food.Id, &food.Name, &food.Desc, &food.Photo)

	return food
}

func GetAllFoods() []models.Food {
	db, _ := sql.Open("postgres", psqlconn)
	defer db.Close()

	rows, _ := db.Query(`SELECT "Id","Name","Desc","Photo" FROM "Food"`)

	foods := []models.Food{}

	for rows.Next() {
		var food models.Food
		rows.Scan(&food.Id, &food.Name, &food.Desc, &food.Photo)

		foods = append(foods, food)
	}

	return foods
}

func AddFood(food models.Food, wg *sync.WaitGroup) {
	mutex.Lock()
	db, _ := sql.Open("postgres", psqlconn)

	defer db.Close()

	insertDynStmt := `INSERT INTO "Food"("Name", "Desc","Photo") values($1, $2, $3)`
	db.Exec(insertDynStmt, food.Name, food.Desc, food.Photo)
	mutex.Unlock()
	wg.Done()
}

func UpdateFood(id int, newFood models.Food) models.Food {
	selectedFood := GetFood(id)

	if selectedFood.Id != 0 {
		selectedFood.Desc = newFood.Desc
		selectedFood.Name = newFood.Name
		selectedFood.Photo = newFood.Photo

		db, _ := sql.Open("postgres", psqlconn)
		defer db.Close()

		updateStmt := `update "Food" set "Name"=$1, "Desc"=$2, "Photo" = $3 where "Id"=$4`
		_, e := db.Exec(updateStmt, newFood.Name, newFood.Desc, newFood.Photo, id)

		if e != nil {
			return models.Food{}
		}

		return selectedFood
	}

	return models.Food{}
}
