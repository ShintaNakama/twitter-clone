package usecase

import (
	"context"
	"time"

	"github.com/ShintaNakama/twitter-clone/backend/app/domain/entity"
	"github.com/ShintaNakama/twitter-clone/backend/app/domain/repository"
)

type PostsUsecase interface {
	ListUserPosts(ctx context.Context, userID string) ([]*Post, error)
	GetPost(ctx context.Context, postID string) (*Post, error)
	CreatePost(ctx context.Context, input *PostInput) error
}

type postsUsecase struct {
	repo repository.PostsRepository
}

func NewPostsUsecase(r repository.PostsRepository) PostsUsecase {
	return &postsUsecase{
		repo: r,
	}
}

func (u *postsUsecase) ListUserPosts(ctx context.Context, userID string) ([]*Post, error) {
	posts, err := u.repo.ListUserPosts(ctx, userID)
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

func (u *postsUsecase) GetPost(ctx context.Context, userID string) (*Post, error) {
	return &Post{}, nil
}

func (u *postsUsecase) CreatePost(ctx context.Context, input *PostInput) error {
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
