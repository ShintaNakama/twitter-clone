package resolvers

import (
	"context"
	"pb"

	"github.com/ShintaNakama/twitter-clone/bff/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) PostList(ctx context.Context) ([]*models.Post, error) {
	posts, err := r.backendClient.ListPost(ctx, &pb.ListPostRequest{
		UserId: "",
	})
	if err != nil {
		return nil, err
	}

	res := make([]*models.Post, len(posts.GetPosts()))
	for i, p := range posts.GetPosts() {
		res[i] = &models.Post{
			ID:   p.GetId(),
			Body: p.GetBody(),
		}
	}

	return res, nil
}

func (r *queryResolver) Post(ctx context.Context, id string) (*models.Post, error) {
	post, err := r.backendClient.GetPost(ctx, &pb.GetPostRequest{
		PostId: id,
	})
	if err != nil {
		return nil, err
	}

	return &models.Post{
		ID:   post.GetId(),
		Body: post.GetBody(),
	}, nil
}
