package model

type Image struct{
	Id 		int 	`json:"-" db:"id"`
	Path 	string  `json:"path"`
}
