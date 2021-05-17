package presentation

import (
	"context"

	"github.com/ShintaNakama/twitter-clone/backend/app/presentation/conv"
	"github.com/ShintaNakama/twitter-clone/backend/app/usecase"
	"github.com/golang/protobuf/ptypes/empty"
	"pb"
)

type TwitterCloneServer interface {
	ListPost(context.Context, *pb.ListPostRequest) (*pb.Posts, error)
	GetPost(context.Context, *pb.GetPostRequest) (*pb.Post, error)
	CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*empty.Empty, error)
	CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*empty.Empty, error)
}

type twitterCloneServer struct {
	usersUsecase usecase.UsersUsecase
	postsUsecase usecase.PostsUsecase
	postConv     conv.PostConv
}

func NewTwitterCloneServer(u usecase.UsersUsecase, p usecase.PostsUsecase) TwitterCloneServer {
	return &twitterCloneServer{
		usersUsecase: u,
		postsUsecase: p,
		postConv:     conv.PostConv{},
	}
}

func (s *twitterCloneServer) ListPost(ctx context.Context, req *pb.ListPostRequest) (*pb.Posts, error) {
	posts, err := s.postsUsecase.List(ctx)
	if err != nil {
		return nil, err
	}

	return s.postConv.ToPbs(posts), nil
}

func (s *twitterCloneServer) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	post, err := s.postsUsecase.Get(ctx, req.GetPostId())
	if err != nil {
		return nil, err
	}

	return s.postConv.ToPb(post), nil
}

func (s *twitterCloneServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*empty.Empty, error) {
	in := &usecase.UserInput{
		ID:    req.GetId(),
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	if err := s.usersUsecase.CreateUser(ctx, in); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *twitterCloneServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*empty.Empty, error) {
	in := &usecase.PostInput{
		ID:     req.GetId(),
		UserID: req.GetUserId(),
		Body:   req.GetBody(),
	}

	if err := s.postsUsecase.Create(ctx, in); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
