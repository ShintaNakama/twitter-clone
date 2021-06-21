package resolvers

import (
	"context"
	"pb"

	"github.com/ShintaNakama/twitter-clone/bff/models"
)

type postResolver struct{ *Resolver }

func (r *postResolver) User(ctx context.Context, obj *models.Post) (*models.User, error) {
	u, err := r.backendClient.GetUserByPostID(ctx, &pb.GetUserByPostIDRequest{
		PostId: obj.ID,
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:    u.Id,
		Email: u.Email,
		Name:  u.Name,
		Image: u.Image,
	}, nil
}
