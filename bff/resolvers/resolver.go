package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"pb"

	"github.com/ShintaNakama/twitter-clone/bff/generated"
)

type Resolver struct {
	backendClient pb.TwitterCloneClient
}

func NewResolver(bc pb.TwitterCloneClient) *Resolver {
	return &Resolver{
		backendClient: bc,
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }
