package tests

import (
	"Go-Basic-Restful-APi/pkg/mocks"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
)

const (
	psqlconn = "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
)

func TestSQLCon(t *testing.T) {
	db, ex := sql.Open("postgres", psqlconn)

	if ex != nil {
		t.Errorf(ex.Error())
	}

	defer db.Close()
}

// This is example of single argument usage
func ExampleGetAllFoods() {
	fmt.Println(mocks.Foods)
	// Output: Returns All Foods From DB
}
