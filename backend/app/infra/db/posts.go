package db

import (
	"context"
	"fmt"

	"github.com/ShintaNakama/twitter-clone/backend/app/domain/entity"
	"github.com/ShintaNakama/twitter-clone/backend/app/domain/repository"
	"github.com/ShintaNakama/twitter-clone/backend/app/infra/db/models"
	"github.com/go-gorp/gorp"
)

const defaultListPostLimit int = 30

const postsColumns = `
id,
user_id,
body,
posted_at
`

type postsImpl struct {
	exec gorp.SqlExecutor
}

func NewPostsRepository(conn gorp.SqlExecutor) repository.PostsRepository {
	return &postsImpl{exec: conn}
}

func (r *postsImpl) List(ctx context.Context) ([]*entity.Post, error) {
	var posts []*models.Post
	q := fmt.Sprintf("SELECT %s FROM posts ORDER BY posted_at DESC LIMIT ?", postsColumns)
	_, err := r.exec.Select(&posts, q, defaultListPostLimit)
	if err != nil {
		return nil, err
	}

	res := make([]*entity.Post, len(posts))
	for i, p := range posts {
		res[i] = entity.NewPost(&entity.PostArgs{
			ID:       p.ID,
			UserID:   p.UserID,
			Body:     p.Body,
			PostedAt: p.PostedAt,
		})
	}
	return res, nil
}

func (r *postsImpl) ListByUser(ctx context.Context, userID string) ([]*entity.Post, error) {
	var posts []*models.Post
	q := "SELECT * FROM posts WHERE user_id = ? ORDER BY posted_at DESC LIMIT ?"

	_, err := r.exec.Select(&posts, q, userID, defaultListPostLimit)
	if err != nil {
		return nil, err
	}

	res := make([]*entity.Post, len(posts))
	for i, p := range posts {
		res[i] = entity.NewPost(&entity.PostArgs{
			ID:       p.ID,
			UserID:   p.UserID,
			Body:     p.Body,
			PostedAt: p.PostedAt,
		})
	}
	return res, nil
}

func (r *postsImpl) Get(ctx context.Context, id string) (*entity.Post, error) {
	var post *models.Post
	q := fmt.Sprintf("SELECT %s FROM posts WHERE id = ?", postsColumns)

	err := r.exec.SelectOne(&post, q, id)
	if err != nil {
		return nil, err
	}

	return entity.NewPost(&entity.PostArgs{
		ID:       post.ID,
		UserID:   post.UserID,
		Body:     post.Body,
		PostedAt: post.PostedAt,
	}), nil
}

func (r *postsImpl) Insert(ctx context.Context, post *entity.Post) error {
	p := &models.Post{
		ID:       post.GetID(),
		UserID:   post.GetUserID(),
		Body:     post.GetBody(),
		PostedAt: post.GetPostedAt(),
	}

	if err := r.exec.Insert(p); err != nil {
		return err
	}

	return nil
}
