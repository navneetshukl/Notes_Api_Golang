package models


// Creating Signup types

type Signup struct {
	Id     int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}