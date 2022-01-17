package user

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
	util "github.com/timoteoBone/project-microservice/httpService/pkg/utils"
)

type Repository interface {
	CreateUser(ctx context.Context, rq entities.CreateUserRequest) (entities.CreateUserResponse, error)
	GetUser(ctx context.Context, rq entities.GetUserRequest) (entities.GetUserResponse, error)
}

type service struct {
	Repo   Repository
	Logger log.Logger
}

func NewService(repo Repository, logger log.Logger) *service {
	return &service{Repo: repo, Logger: logger}
}

func (s *service) CreateUser(ctx context.Context, rq entities.CreateUserRequest) (entities.CreateUserResponse, error) {

	logger := log.With(s.Logger, "create user request", "recevied")

	err := util.ValidateCreateUserRequest(rq)
	if err != nil {
		level.Error(logger).Log("error", err)
		return entities.CreateUserResponse{}, err
	}

	rq.Pass, err = util.HashPassword(rq.Pass)
	if err != nil {
		level.Error(logger).Log("error", err)
		return entities.CreateUserResponse{}, err
	}

	res, err := s.Repo.CreateUser(ctx, rq)

	if err != nil {
		level.Error(logger).Log("error", err)
		return entities.CreateUserResponse{}, err
	}

	return res, nil

}

func (s *service) GetUser(ctx context.Context, rq entities.GetUserRequest) (entities.GetUserResponse, error) {
	logger := log.With(s.Logger, "get user request", "recevied")

	if err := util.ValidateGetUserRequest(rq); err != nil {
		level.Error(logger).Log("error", err)
		return entities.GetUserResponse{}, err
	}

	res, err := s.Repo.GetUser(ctx, rq)
	if err != nil {
		level.Error(logger).Log("error", err)
		return entities.GetUserResponse{}, err
	}

	return res, nil
}
