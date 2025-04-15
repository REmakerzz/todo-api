# 📝 TODO API на Golang

Простой REST API для управления списком задач. Учебный проект для освоения разработки backend-приложений на Go.
На данном этапе перевожу хранение данных с оперативной памяти (с слайсов) в БД(sqlite3).

## 🚀 Возможности

- Получение списка задач: `GET /tasks`
- Добавление новой задачи: `POST /tasks`
- Удаление задачи по ID: `DELETE /tasks/{id}`
- Обновление задачи по ID: `PUT /tasks/{id}`
- Фильтрация задач по статусу done в `GET /tasks`

---

## ⚙️  Используемые технологии

- Язык: Go 1.19+
- Пакеты: `net/http`, `encoding/json`
- Архитектура с разделением по модулям (handlers, модели)

---

## 📦 Установка и запуск

```bash
# Клонируем репозиторий
git clone https://github.com/REmakerzz/todo-api.git
cd todo-api

# Запускаем сервер
go run main.go
```

Сервер будет доступен по адресу:
http://localhost:8080

---

## 🔧 Примеры запросов
### ➕ Добавить задачу

**POST** `http://localhost:8080/tasks`
Body (JSON):

```json
{
  "title": "Сделать домашку по Go"
}
```

Ответ:

```json
{
  "id": 1,
  "title": "Сделать домашку по Go",
  "done": false
}
```

### 📋 Получить список задач

GET http://localhost:8080/tasks

### ❌ Удалить задачу

DELETE http://localhost:8080/tasks/{id}

### 🔃 Обновить задачу

PUT http://localhost:8080/tasks/{id}

### Фильтрация задач по статусу done

GET http://localhost:8080/tasks?done=true
GET http://localhost:8080/tasks?done=false

---

## 📌 Планы на доработку

- Хранение задач в файле или базе данных

- Юнит-тесты

- Документация через Swagger

- Поддержка фронтенда
