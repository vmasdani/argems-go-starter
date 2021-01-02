package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Route(r *mux.Router, db *gorm.DB) {
	r.Use(BrotliCompressorMiddleware)

	// r.HandleFunc("/", http.FileServer("./"))\

	r.HandleFunc("/todos", AllTodos(db)).Methods("GET")
	r.HandleFunc("/todos/{id}", GetTodo(db)).Methods("GET")
	r.HandleFunc("/todos/{id}", DeleteTodo(db)).Methods("DELETE")
	r.HandleFunc("/todos", PostTodo(db)).Methods("POST")

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./frontend"))))
}
