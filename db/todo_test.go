package db

import (
	"testing"
	"time"

	"github.com/DMV-Nicolas/todo/util"
	"github.com/stretchr/testify/require"
)

func randomTodo(t *testing.T) Todo {
	todo := Todo{
		Title: util.RandomString(10),
		Done:  util.RandomBool(),
	}

	createdTodo := testQueries.CreateTodo(todo)
	require.NotEmpty(t, createdTodo)
	require.NotZero(t, createdTodo.ID)

	require.Equal(t, todo.Title, createdTodo.Title)
	require.Equal(t, todo.Done, createdTodo.Done)

	require.WithinDuration(t, time.Now(), createdTodo.CreatedAt, time.Second)
	require.Zero(t, createdTodo.DeletedAt)

	obtainedTodo := testQueries.GetTodo(createdTodo.ID)
	require.NotEmpty(t, obtainedTodo)

	require.Equal(t, createdTodo.ID, obtainedTodo.ID)
	require.Equal(t, createdTodo.Title, obtainedTodo.Title)
	require.Equal(t, createdTodo.Done, obtainedTodo.Done)
	require.WithinDuration(t, createdTodo.CreatedAt, obtainedTodo.CreatedAt, time.Second)
	require.WithinDuration(t, createdTodo.DeletedAt.Time, obtainedTodo.DeletedAt.Time, time.Second)

	return obtainedTodo
}

func TestCreateTodo(t *testing.T) {
	randomTodo(t)
}

func TestGetTodo(t *testing.T) {
	todo := randomTodo(t)

	obtainedTodo := testQueries.GetTodo(todo.ID)
	require.NotEmpty(t, obtainedTodo)

	require.Equal(t, todo.ID, obtainedTodo.ID)
	require.Equal(t, todo.Title, obtainedTodo.Title)
	require.Equal(t, todo.Done, obtainedTodo.Done)
	require.WithinDuration(t, todo.CreatedAt, obtainedTodo.CreatedAt, time.Second)
	require.WithinDuration(t, todo.UpdatedAt, obtainedTodo.UpdatedAt, time.Second)
	require.WithinDuration(t, todo.DeletedAt.Time, obtainedTodo.DeletedAt.Time, time.Second)
}

func TestListTodos(t *testing.T) {
	n := 10
	for i := 0; i < n; i++ {
		randomTodo(t)
	}

	offset := n / 2
	limit := n / 2

	todos := testQueries.ListTodos(offset, limit)
	require.NotEmpty(t, todos)
	require.Len(t, todos, limit)

	for _, todo := range todos {
		require.NotEmpty(t, todo)
	}
}

func TestDeleteTodo(t *testing.T) {
	todo := randomTodo(t)

	testQueries.DeleteTodo(todo.ID)

	deletedTodo := testQueries.GetTodo(todo.ID)
	require.Empty(t, deletedTodo)
}

func TestUpdateTodoDone(t *testing.T) {
	todo := randomTodo(t)
	todo.Done = false

	testQueries.UpdateTodo(todo.ID, true)

	updatedTodo := testQueries.GetTodo(todo.ID)
	require.NotEmpty(t, updatedTodo)

	require.Equal(t, todo.Title, updatedTodo.Title)
	require.NotEqual(t, todo.Done, updatedTodo.Done)
	require.WithinDuration(t, todo.CreatedAt, updatedTodo.CreatedAt, time.Second)
	require.WithinDuration(t, time.Now(), updatedTodo.UpdatedAt, time.Second)
	require.WithinDuration(t, todo.DeletedAt.Time, todo.DeletedAt.Time, time.Second)
}
