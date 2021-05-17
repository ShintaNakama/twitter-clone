package repository

import (
	"context"

	"github.com/ShintaNakama/twitter-clone/backend/app/domain/entity"
)

type PostsRepository interface {
	List(context.Context) ([]*entity.Post, error)
	ListByUser(context.Context, string) ([]*entity.Post, error)
	Get(context.Context, string) (*entity.Post, error)
	Insert(context.Context, *entity.Post) error
}
