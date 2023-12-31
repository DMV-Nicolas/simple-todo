package db

import (
	"time"

	"gorm.io/gorm"
)

func (q *Queries) CreateTodo(title string) Todo {
	todo := Todo{
		Title: title,
		Done:  false,
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	q.db.Create(&todo)

	return todo
}

func (q *Queries) GetTodo(id uint) Todo {
	todo := new(Todo)
	q.db.First(todo, id)
	return *todo
}

func (q *Queries) ListTodos(offset, limit int) []Todo {
	todos := make([]Todo, 0)
	q.db.Offset(offset).Limit(limit).Find(&todos)
	return todos
}

func (q *Queries) UpdateTodo(id uint, done bool) {
	todo := new(Todo)
	todo.ID = id

	q.db.Model(todo).Update("done", done)
}

func (q *Queries) DeleteTodo(id uint) {
	q.db.Delete(&Todo{}, id)
}
