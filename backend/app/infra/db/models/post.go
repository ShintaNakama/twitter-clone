package models

import (
	"time"
)

type Post struct {
	ID       string    `db:"id"`
	UserID   string    `db:"user_id"`
	Body     string    `db:"body"`
	PostedAt time.Time `db:"posted_at"`
}
