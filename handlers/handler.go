package handlers

import "github.com/DMV-Nicolas/simple-todo/db"

// NewHandler creates a new handler object
func NewHandler(queries *db.Queries) *Handler {
	return &Handler{
		queries: queries,
	}
}

// Handler is an struct for create handler functions with a query object
type Handler struct {
	queries *db.Queries
}
