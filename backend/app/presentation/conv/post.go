package conv

import (
	"github.com/ShintaNakama/twitter-clone/backend/app/usecase"
	"github.com/golang/protobuf/ptypes"
	"pb"
)

type PostConv struct{}

func (c *PostConv) ToPbs(ps []*usecase.Post) *pb.Posts {
	posts := make([]*pb.Post, len(ps))
	for i := range ps {
		posts[i] = c.ToPb(ps[i])
	}

	return &pb.Posts{
		Posts: posts,
	}
}

func (c *PostConv) ToPb(p *usecase.Post) *pb.Post {
	ts, _ := ptypes.TimestampProto(p.PostedAt)

	return &pb.Post{
		Id:       p.ID,
		UserId:   p.UserID,
		Body:     p.Body,
		PostedAt: ts,
	}
}
