package tasks

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
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

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}

	for i, t := range TaskList {
		if t.ID == id {
			TaskList = append(TaskList[:i], TaskList[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Задача не найдена", http.StatusNotFound)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}

	var updated Task
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	for i, t := range TaskList {
		if t.ID == id {
			TaskList[i].Title = updated.Title
			TaskList[i].Done = updated.Done

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(TaskList[i])
			return
		}
	}
	http.Error(w, "Задача не найдена", http.StatusNotFound)
}
