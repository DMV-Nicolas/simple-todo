package db

func (q *Queries) CreateTodo(todo Todo) Todo {
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

func (q *Queries) UpdateTodo(todo Todo) Todo {
	q.db.Model(&todo).Updates(Todo{
		Title:       todo.Title,
		Description: todo.Description,
		Done:        todo.Done,
	})
	return todo
}

func (q *Queries) DeleteTodo(id uint) {
	q.db.Delete(&Todo{}, id)
}

func (q *Queries) UpdateTodoDone(id uint, done bool) {
	todo := new(Todo)
	todo.ID = id

	q.db.Model(todo).Update("done", done)
}
