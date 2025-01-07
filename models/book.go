package models

type Book struct {
	ID		 	uint		`json:"id"`
	Title	 	string		`json:"title"`
	Category 	string		`json:"category"`
	Author	 	string		`json:"author"`
	Synopsis 	string		`json:"synopsis"`
}