package models

type Event struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

var Events = []Event{
	{
		ID: 1,
		Title: "Intro to golang",
		Description: "Come join in golang",
	},
	{
		ID: 2,
		Title: "Intro to Angular",
		Description: "Angular is a web framework",
	},
	{
		ID: 3,
		Title: "Intro to Django",
		Description: "Django is awesome",
	},
}
