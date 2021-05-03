package repository

import (
	"context"

	"github.com/ShintaNakama/twitter-clone/backend/app/domain/entity"
)

type UsersRepository interface {
	Insert(context.Context, *entity.User) error
}
