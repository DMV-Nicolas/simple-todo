package db

import "gorm.io/gorm"

func NewQueries(db *gorm.DB) *Queries {
	return &Queries{
		db: db,
	}
}

type Queries struct {
	db *gorm.DB
}
