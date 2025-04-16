# TODO API на Golang

Простой REST API для управления списком задач (TODO). Учебный проект для освоения разработки backend-приложений на Go.
Хранение данных реализовано через SQLite.

## Возможности

- Получение списка задач: `GET /tasks`
- Добавление новой задачи: `POST /tasks`
- Удаление задачи по ID: `DELETE /tasks/{id}`
- Обновление задачи по ID: `PUT /tasks/{id}`
- Фильтрация задач по статусу done в `GET /tasks?done=true|false`

---

##  Используемые технологии

- Язык: Go 1.19+
- База данных: SQLite
- Стандартная библиотека

---

## Установка и запуск

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

## Примеры запросов (Postman)
### Добавить задачу

**POST** `http://localhost:8080/tasks`
**Body** (JSON):

```json
{
  "title": "Сделать домашку по Go"
}
```

### Получить список задач

GET http://localhost:8080/tasks

### Удалить задачу

DELETE http://localhost:8080/tasks/{id}

### Обновить задачу

PUT http://localhost:8080/tasks/{id}

### Фильтрация задач по статусу done

GET http://localhost:8080/tasks?done=true|false

---

## Планы на доработку

- Добавить миграции БД

- Юнит-тесты

- Документация через Swagger

- Поддержка фронтенда

- Сделать docker-версию API
