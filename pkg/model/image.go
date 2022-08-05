package model

type Image struct{
	Id 				int 	`json:"-" db:"id"`
	Path 			string  `json:"path"`
	SeventyFivePath string 	`json:"-" db:"seventy_five_path"`
	HalfPath 		string 	`json:"-"db:"half_path"`
	TwentyFivePath 	string 	`json:"-"db:"twenty_five_path"`
}
