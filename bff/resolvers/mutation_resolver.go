package resolvers

import (
	"context"
	"pb"

	"github.com/ShintaNakama/twitter-clone/bff/models"
)

type mutationResolver struct{ *Resolver }

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
