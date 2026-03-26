package main

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todos []Todo

func getTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

// POST - yangi todo qo'shadi
func createTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	json.NewDecoder(r.Body).Decode(&newTodo)

	todos = append(todos, newTodo)
	json.NewEncoder(w).Encode(newTodo)
}

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getTodos(w, r)
		} else if r.Method == "POST" {
			createTodo(w, r)
		}
	})
	println("Server http://localhost:8080/todos manzilida ishga tushdi...")

	http.ListenAndServe(":8080", nil)
}
