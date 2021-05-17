package interactor

import (
	"github.com/ShintaNakama/twitter-clone/backend/app/domain/repository"
	"github.com/ShintaNakama/twitter-clone/backend/app/infra/db"
	"github.com/ShintaNakama/twitter-clone/backend/app/presentation"
	"github.com/ShintaNakama/twitter-clone/backend/app/usecase"
	"github.com/go-gorp/gorp"
)

type Interactor interface {
	NewPostsRepository() repository.PostsRepository
	NewPostsUsecase() usecase.PostsUsecase
	NewTwitterCloneServer() presentation.TwitterCloneServer
	NewApi() *api
}

type interactor struct {
	conn *gorp.DbMap
}

func NewInteractor(conn *gorp.DbMap) Interactor {
	return &interactor{conn: conn}
}

type api struct {
	// grpc以外のpresentationが増えたらここも増える感じ
	presentation.TwitterCloneServer
}

func (i *interactor) NewApi() *api {
	api := &api{}
	api.TwitterCloneServer = i.NewTwitterCloneServer()

	return api
}

func (i *interactor) NewUsersRepository() repository.UsersRepository {
	return db.NewUsersRepository(i.conn)
}

func (i *interactor) NewPostsRepository() repository.PostsRepository {
	return db.NewPostsRepository(i.conn)
}

func (i *interactor) NewUsersUsecase() usecase.UsersUsecase {
	return usecase.NewUsersUsecase(i.NewUsersRepository())
}

func (i *interactor) NewPostsUsecase() usecase.PostsUsecase {
	return usecase.NewPostsUsecase(i.NewPostsRepository())
}

func (i *interactor) NewTwitterCloneServer() presentation.TwitterCloneServer {
	return presentation.NewTwitterCloneServer(i.NewUsersUsecase(), i.NewPostsUsecase())
}
