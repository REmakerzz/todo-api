package tasks

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	query := r.URL.Query()
	doneFilter := query.Get("done")

	var rows *sql.Rows
	var err error

	if doneFilter == "" {
		rows, err = db.Query("SELECT id, title, done FROM tasks")
	} else {
		filterValue, err := strconv.ParseBool(doneFilter)
		if err != nil {
			http.Error(w, "Неверное значение параметра done. Используйте true или false", http.StatusBadRequest)
			return
		}

		rows, err = db.Query("SELECT id, title, done FROM tasks WHERE done = ?", filterValue)
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

func AddTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Неверный формат", http.StatusBadRequest)
		return
	}
	if t.Title == "" {
		http.Error(w, "Title не может быть пустым", http.StatusBadRequest)
		return
	}

	res, err := db.Exec("INSERT INTO tasks (title, done) VALUES (?, ?)", t.Title, false)
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

func DeleteTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Ошибка при удалении задачи", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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

	_, err = db.Exec("UPDATE tasks SET title = ?, done = ? WHERE id = ?", updated.Title, updated.Done, id)
	if err != nil {
		http.Error(w, "Ошибка при обновлении задачи", http.StatusInternalServerError)
		return
	}

	updated.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)

}
