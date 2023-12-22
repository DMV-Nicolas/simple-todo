package db

import (
	"log"
	"os"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	db, err := ConnectDatabase("../todo.db")
	if err != nil {
		log.Fatal("cannot connect database")
	}

	testQueries = NewQueries(db)

	os.Exit(m.Run())
}
