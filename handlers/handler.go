package handlers

import "github.com/DMV-Nicolas/todo/db"

type Handler struct {
	queries *db.Queries
}

func NewHandler(queries *db.Queries) *Handler {
	return &Handler{
		queries: queries,
	}
}