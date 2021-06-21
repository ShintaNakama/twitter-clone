package usecase

import (
	"context"

	"github.com/ShintaNakama/twitter-clone/backend/app/domain/entity"
	"github.com/ShintaNakama/twitter-clone/backend/app/domain/repository"
)

type UsersUsecase interface {
	CreateUser(ctx context.Context, input *UserInput) error
	Get(ctx context.Context, userID string) (*User, error)
	GetUserByPostID(ctx context.Context, postID string) (*User, error)
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

func (u *usersUsecase) Get(ctx context.Context, id string) (*User, error) {
	user, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:    user.GetID(),
		Email: user.GetEmail(),
		Name:  user.GetName(),
		Image: user.GetImage(),
	}, nil
}

func (u *usersUsecase) GetUserByPostID(ctx context.Context, postID string) (*User, error) {
	user, err := u.repo.GetUserByPostID(ctx, postID)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:    user.GetID(),
		Email: user.GetEmail(),
		Name:  user.GetName(),
		Image: user.GetImage(),
	}, nil
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
