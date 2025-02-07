package main

import (
	"go-api-assignment/handler"
	"go-api-assignment/repository"
	"log"
	"net/http"
)

func main() {
	repo := repository.NewUserRepositoryImpl()
	handler := handler.NewUserHandler(repo)

	http.HandleFunc("/add-user", handler.AddUserHandler)
	http.HandleFunc("/get-user", handler.GetUserByIDHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
