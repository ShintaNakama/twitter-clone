package conv

import (
	"pb"

	"github.com/ShintaNakama/twitter-clone/backend/app/usecase"
)

type UserConv struct{}

func (c *UserConv) ToPb(u *usecase.User) *pb.User {
	return &pb.User{
		Id:    u.ID,
		Email: u.Email,
		Name:  u.Name,
		Image: u.Image,
	}
}
