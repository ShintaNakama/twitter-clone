package presentation

import (
	"context"

	"github.com/ShintaNakama/twitter-clone/backend/app/presentation/conv"
	"github.com/ShintaNakama/twitter-clone/backend/app/usecase"
	"github.com/golang/protobuf/ptypes/empty"
	"pb"
)

type TwitterCloneServer interface {
	ListUserPosts(context.Context, *pb.ListUserPostsRequest) (*pb.Posts, error)
	GetPost(context.Context, *pb.GetPostRequest) (*pb.Post, error)
	CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*empty.Empty, error)
	CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*empty.Empty, error)
}

type twitterCloneServer struct {
	usecase  usecase.PostsUsecase
	postConv conv.PostConv
}

func NewTwitterCloneServer(usecase usecase.PostsUsecase) TwitterCloneServer {
	return &twitterCloneServer{
		usecase:  usecase,
		postConv: conv.PostConv{},
	}
}

func (s *twitterCloneServer) ListUserPosts(ctx context.Context, req *pb.ListUserPostsRequest) (*pb.Posts, error) {
	posts, err := s.usecase.ListUserPosts(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return s.postConv.ToPbs(posts), nil
}

func (s *twitterCloneServer) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	return &pb.Post{}, nil
}

func (s *twitterCloneServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *twitterCloneServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*empty.Empty, error) {
	in := &usecase.PostInput{
		ID:     req.GetId(),
		UserID: req.GetUserId(),
		Body:   req.GetBody(),
	}

	if err := s.usecase.CreatePost(ctx, in); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
