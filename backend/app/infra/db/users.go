package db

import (
	"context"

	"github.com/ShintaNakama/twitter-clone/backend/app/domain/entity"
	"github.com/ShintaNakama/twitter-clone/backend/app/domain/repository"
	"github.com/ShintaNakama/twitter-clone/backend/app/infra/db/models"
	"github.com/go-gorp/gorp"
)

const usersColumns = `
id,
name,
email,
image
`

type usersImpl struct {
	exec gorp.SqlExecutor
}

func NewUsersRepository(conn gorp.SqlExecutor) repository.UsersRepository {
	return &usersImpl{exec: conn}
}

func (r *usersImpl) Insert(ctx context.Context, user *entity.User) error {
	u := &models.User{
		ID:    user.GetID(),
		Name:  user.GetName(),
		Email: user.GetEmail(),
		Image: user.GetImage(),
	}

	if err := r.exec.Insert(u); err != nil {
		return err
	}

	return nil
}
