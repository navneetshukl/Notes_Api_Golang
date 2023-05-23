package models

// Creating Notes type


type Notes struct{
	Id int  `json:"id"`
	Title string  `json:"title"`
	Email string `json:"email"`
}