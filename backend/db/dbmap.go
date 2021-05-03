package db

import (
	"log"
	"os"

	"github.com/ShintaNakama/twitter-clone/backend/app/infra/db/models"
	"github.com/ShintaNakama/twitter-clone/backend/env"
	"github.com/go-gorp/gorp"
)

func GetDbMap() (*gorp.DbMap, error) {
	db, err := EstablishConnection(env.DSN())
	if err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "utf8mb4"}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "twitter-clone:", log.Lmicroseconds))

	dbmap.AddTableWithName(models.User{}, "users").SetKeys(false, "id")
	dbmap.AddTableWithName(models.Post{}, "posts").SetKeys(false, "id")

	return dbmap, nil
}
