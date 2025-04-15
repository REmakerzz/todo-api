package tasks

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"todo-api/db"
)

var lastID int

func getNextID() int {
	lastID++
	return lastID
}

func GetTasks(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	doneFilter := query.Get("done")

	var rows *sql.Rows
	var err error

	if doneFilter == "" {
		rows, err = db.DB.Query("SELECT id, title, done FROM tasks")
	} else {
		filterValue, err := strconv.ParseBool(doneFilter)
		if err != nil {
			http.Error(w, "Неверное значение параметра done. Используйте true или false", http.StatusBadRequest)
			return
		}

		rows, err = db.DB.Query("SELECT id, title, done FROM tasks WHERE done = ?", filterValue)
	}

	if err != nil {
		http.Error(w, "Ошибка при чтении из БД", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			http.Error(w, "Ошибка при чтении строки", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Неверный формат", http.StatusBadRequest)
		return
	}
	if t.Title == "" {
		http.Error(w, "Title не может быть пустым", http.StatusBadRequest)
		return
	}

	res, err := db.DB.Exec("INSERT INTO tasks (title, done) VALUES (?, ?)", t.Title, false)
	if err != nil {
		http.Error(w, "Ошибка при добавлении в БД", http.StatusInternalServerError)
		return
	}
	id, _ := res.LastInsertId()
	t.ID = int(id)
	t.Done = false

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
