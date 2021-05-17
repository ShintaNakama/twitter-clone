package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ShintaNakama/twitter-clone/backend/env"

	_ "github.com/go-sql-driver/mysql"
)

const (
	driver  = "mysql"
	options = "?parseTime=true&loc=Asia%2FTokyo&interpolateParams=true"
)

// EstablishConnection DB接続を確立する
func EstablishConnection() (*sql.DB, error) {
	db, err := sql.Open(driver, dsn())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(10 * time.Second)

	return db, err
}

// CloseConnection DB接続を閉じる
func CloseConnection(conn *sql.DB) {
	if err := conn.Close(); err != nil {
		log.Fatalln("failed to db close")
	}
}

// DSN データソース名を取得する
func dsn() string {
	fmt.Println(env.DSN() + options)
	return env.DSN() + options
}
