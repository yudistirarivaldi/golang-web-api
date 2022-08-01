package book

import "encoding/json"

type BookInput struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
	//Subtitle string if you want call subtitle with _ example sub_title you can use `json:"sub_title"`
}