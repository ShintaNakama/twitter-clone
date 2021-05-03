package usecase

import (
	"context"

	"github.com/ShintaNakama/twitter-clone/backend/app/domain/entity"
	"github.com/ShintaNakama/twitter-clone/backend/app/domain/repository"
)

type UsersUsecase interface {
	CreateUser(ctx context.Context, input *UserInput) error
}

type usersUsecase struct {
	repo repository.UsersRepository
}

func NewUsersUsecase(r repository.UsersRepository) UsersUsecase {
	return &usersUsecase{
		repo: r,
	}
}

func (u *usersUsecase) CreateUser(ctx context.Context, input *UserInput) error {
	p := entity.NewUser(&entity.UserArgs{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
		Image: input.Image,
	})

	if err := u.repo.Insert(ctx, p); err != nil {
		return err
	}

	return nil
}

type UserInput struct {
	ID    string
	Email string
	Name  string
	Image string
}

type User struct {
	ID    string
	Email string
	Name  string
	Image string
}
