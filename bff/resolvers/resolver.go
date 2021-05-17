package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"pb"

	"github.com/ShintaNakama/twitter-clone/bff/generated"
	"github.com/ShintaNakama/twitter-clone/bff/models"
)

type Resolver struct {
	backendClient pb.TwitterCloneClient
}

func NewResolver(bc pb.TwitterCloneClient) *Resolver {
	return &Resolver{
		backendClient: bc,
	}
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *models.UserInput) (bool, error) {
	_, err := r.backendClient.CreateUser(ctx, &pb.CreateUserRequest{
		Id:    input.ID,
		Email: input.Email,
		Name:  input.Name,
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, input *models.PostInput) (bool, error) {
	_, err := r.backendClient.CreatePost(ctx, &pb.CreatePostRequest{
		Id:     input.ID,
		UserId: input.UserID,
		Body:   input.Body,
		Images: &pb.Images{},
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}

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
			User: &models.User{},
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
		User: &models.User{},
		Body: post.GetBody(),
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
