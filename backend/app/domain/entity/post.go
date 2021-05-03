package entity

import (
	"time"
)

type Post struct {
	id       string
	userID   string
	body     string
	postedAt time.Time
}

func (p *Post) GetID() string {
	return p.id
}

func (p *Post) GetUserID() string {
	return p.userID
}

func (p *Post) GetBody() string {
	return p.body
}

func (p *Post) GetPostedAt() time.Time {
	return p.postedAt
}

type PostArgs struct {
	ID       string
	UserID   string
	Body     string
	PostedAt time.Time
}

func NewPost(args *PostArgs) *Post {
	return &Post{
		id:       args.ID,
		userID:   args.UserID,
		body:     args.Body,
		postedAt: args.PostedAt,
	}
}
