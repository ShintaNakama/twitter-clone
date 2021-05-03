package db_test

import (
	"testing"

	"github.com/ShintaNakama/twitter-clone/backend/db"
)

var testDSN = "root:twitter@tcp(127.0.0.1:53306)/twitter_clone_db"

func TestEstablishConnection(t *testing.T) {
	d, err := db.EstablishConnection(testDSN)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	db.CloseConnection(d)
}

func TestCloseConnection(t *testing.T) {
	d, err := db.EstablishConnection(testDSN)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	db.CloseConnection(d)
}
