package models

// Creating login type

type Login struct{
	Email string `json:"email"`
	Password string `json:"password"`
}