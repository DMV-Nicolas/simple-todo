package db

import "gorm.io/gorm"

// NewQueries creates a new queries object
func NewQueries(db *gorm.DB) *Queries {
	return &Queries{
		db: db,
	}
}

// Queries is an struct for to interact with the database
type Queries struct {
	db *gorm.DB
}
