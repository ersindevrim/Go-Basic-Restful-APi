package models

type Food struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Photo string `json:"photo"`
}
