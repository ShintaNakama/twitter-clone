package db_test

import (
	"testing"

	"github.com/ShintaNakama/twitter-clone/backend/db"
	"github.com/go-gorp/gorp"
)

func TestGetDbMap(t *testing.T) {
	dbmap, err := db.GetDbMap()
	if err != nil {
		t.Fatal("fatal GetDbMap")
	}

	_, ok := dbmap.Dialect.(gorp.MySQLDialect)
	if !ok {
		t.Errorf("error diarect. want: gorp.MySQLDialect, got: %T", dbmap.Dialect)
	}
}
