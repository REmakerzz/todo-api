package tasks

import (
	"encoding/json"
	"net/http"
)

var lastID int

func getNextID() int {
	lastID++
	return lastID
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TaskList)
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	var t Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Неверный формат", http.StatusBadRequest)
		return
	}
	if t.Title == "" {
		http.Error(w, "Title не может быть пустым", http.StatusBadRequest)
		return
	}
	t.Done = false
	t.ID = getNextID()
	TaskList = append(TaskList, t)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}
