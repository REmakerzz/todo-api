package main

import (
	"net/http"
	"todo-api/tasks"
)

func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tasks.GetTasks(w, r)
		} else if r.Method == http.MethodPost {
			tasks.AddTask(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	http.ListenAndServe(":8080", nil)
}
