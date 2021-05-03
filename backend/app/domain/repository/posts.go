package repository

import (
	"context"

	"github.com/ShintaNakama/twitter-clone/backend/app/domain/entity"
)

type PostsRepository interface {
	ListUserPosts(context.Context, string) ([]*entity.Post, error)
	Insert(context.Context, *entity.Post) error
}
