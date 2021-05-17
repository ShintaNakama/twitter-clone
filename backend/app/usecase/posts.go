package usecase

import (
	"context"
	"time"

	"github.com/ShintaNakama/twitter-clone/backend/app/domain/entity"
	"github.com/ShintaNakama/twitter-clone/backend/app/domain/repository"
)

type PostsUsecase interface {
	List(ctx context.Context) ([]*Post, error)
	Get(ctx context.Context, postID string) (*Post, error)
	Create(ctx context.Context, input *PostInput) error
}

type postsUsecase struct {
	repo repository.PostsRepository
}

func NewPostsUsecase(r repository.PostsRepository) PostsUsecase {
	return &postsUsecase{
		repo: r,
	}
}

func (u *postsUsecase) List(ctx context.Context) ([]*Post, error) {
	posts, err := u.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*Post, len(posts))
	for i, p := range posts {
		res[i] = &Post{
			ID:     p.GetID(),
			UserID: p.GetUserID(),
			Body:   p.GetBody(),
		}
	}
	return res, nil
}

func (u *postsUsecase) ListUserPosts(ctx context.Context, userID string) ([]*Post, error) {
	posts, err := u.repo.ListByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]*Post, len(posts))
	for i, p := range posts {
		res[i] = &Post{
			ID:       p.GetID(),
			UserID:   p.GetUserID(),
			Body:     p.GetBody(),
			PostedAt: p.GetPostedAt(),
		}
	}
	return res, nil
}

func (u *postsUsecase) Get(ctx context.Context, id string) (*Post, error) {
	p, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &Post{
		ID:       p.GetID(),
		UserID:   p.GetUserID(),
		Body:     p.GetBody(),
		PostedAt: p.GetPostedAt(),
	}, nil
}

func (u *postsUsecase) Create(ctx context.Context, input *PostInput) error {
	p := entity.NewPost(&entity.PostArgs{
		ID:       input.ID,
		UserID:   input.UserID,
		Body:     input.Body,
		PostedAt: time.Now(),
	})

	if err := u.repo.Insert(ctx, p); err != nil {
		return err
	}

	return nil
}

type PostInput struct {
	ID     string
	UserID string
	Body   string
}

type Post struct {
	ID       string
	UserID   string
	Body     string
	PostedAt time.Time
}
