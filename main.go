package main

import (
	"net/http"
	"todo-api/storage"
	"todo-api/tasks"
)

func main() {

	db := storage.NewStorage()
	defer db.DB.Close()
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			tasks.GetTasks(w, r, db.DB)
		case http.MethodPost:
			tasks.AddTask(w, r, db.DB)
		default:
			http.NotFound(w, r)
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			tasks.DeleteTask(w, r, db.DB)
		case http.MethodPut:
			tasks.UpdateTask(w, r, db.DB)
		default:
			http.NotFound(w, r)
		}
	})

	http.ListenAndServe(":8080", nil)
}
