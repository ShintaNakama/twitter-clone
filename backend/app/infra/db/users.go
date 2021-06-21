package db

import (
	"context"
	"fmt"

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

func (r *usersImpl) Get(ctx context.Context, id string) (*entity.User, error) {
	var user *models.User
	q := fmt.Sprintf("SELECT %s FROM users WHERE id = ?", usersColumns)

	err := r.exec.SelectOne(&user, q, id)
	if err != nil {
		return nil, err
	}

	return entity.NewUser(&entity.UserArgs{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: user.Image,
	}), nil
}

func (r *usersImpl) GetUserByPostID(ctx context.Context, postID string) (*entity.User, error) {
	var user *models.User
	q := "SELECT users.id AS id, users.name AS Name, users.email AS Email, users.image AS Image  FROM users JOIN posts ON users.id = posts.user_id  WHERE posts.id = ?"

	err := r.exec.SelectOne(&user, q, postID)
	if err != nil {
		return nil, err
	}

	return entity.NewUser(&entity.UserArgs{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: user.Image,
	}), nil
}
