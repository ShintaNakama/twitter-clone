package env

import (
	"os"
)

// DSN DBのデータソース名を取得する
func DSN() string {
	var dsn = os.Getenv("DB_DSN")
	return dsn
}
