package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/navneetshukl/Golang_notes_API/routes"
)

func main() {
	//fmt.Println("Server running !!!")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mux := chi.NewRouter()
	mux.Post("/signup", routes.Signup)
	mux.Post("/login", routes.Login)
	mux.Get("/getnote", routes.GetNote)
	mux.Post("/createnote", routes.CreateNote)
	mux.Delete("/deletenote", routes.DeleteNote)
	http.ListenAndServe(":8080", mux)
}
