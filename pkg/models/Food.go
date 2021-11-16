package models

// This text will appear as description of your response body.
// swagger:response FoodResponse
type Food struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Photo string `json:"photo"`
}
