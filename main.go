package main

import (
	"net/http"
	"todo-api/db"
	"todo-api/tasks"
)

func main() {
	db.Init()

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			tasks.GetTasks(w, r)
		case http.MethodPost:
			tasks.AddTask(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			tasks.DeleteTask(w, r)
		case http.MethodPut:
			tasks.UpdateTask(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	http.ListenAndServe(":8080", nil)
}
