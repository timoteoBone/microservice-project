package user

import (
	"context"

	"database/sql"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/google/uuid"

	entities "github.com/timoteoBone/microservice-project/grpcService/pkg/entities"
	errors "github.com/timoteoBone/microservice-project/grpcService/pkg/errors"
	mapper "github.com/timoteoBone/microservice-project/grpcService/pkg/mapper"
)

type Repository interface {
	GetUser(ctx context.Context, userId string) (entities.User, error)
	CreateUser(ctx context.Context, user entities.User, newId string) (string, error)
	DeleteUser(ctx context.Context, userId string) error
}

type service struct {
	Repo   Repository
	Logger log.Logger
}

func NewService(l log.Logger, r Repository) *service {
	return &service{r, l}
}

func (s *service) CreateUser(ctx context.Context, userReq entities.CreateUserRequest) (entities.CreateUserResponse, error) {
	s.Logger.Log(s.Logger, "request", "create user", "received")

	response := entities.CreateUserResponse{}
	status := entities.Status{}

	user := mapper.CreateUserRequestToUser(userReq)
	newId := generateId()
	genId, err := s.Repo.CreateUser(ctx, user, newId)

	if err != nil {
		if err.Error() == "Error 1062: Duplicate entry '' for key 'USER.email'" {
			return entities.CreateUserResponse{}, errors.NewUserAlreadyExists()
		}
		level.Error(s.Logger).Log("error", err)
		return entities.CreateUserResponse{}, errors.NewDataBaseError()
	}

	status.Message = "created successfully"
	response.Status = status
	response.UserId = genId

	return response, nil
}

func (s *service) GetUser(ctx context.Context, user entities.GetUserRequest) (entities.GetUserResponse, error) {
	s.Logger.Log(s.Logger, "request", "get user", "received")

	res, err := s.Repo.GetUser(ctx, user.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			level.Error(s.Logger).Log("error", err)
			return entities.GetUserResponse{}, errors.NewUserNotFound()
		}
		level.Error(s.Logger).Log("error", err)
		return entities.GetUserResponse{}, errors.NewDataBaseError()
	}

	response := entities.GetUserResponse{
		Name: res.Name,
		Id:   user.UserID,
		Age:  res.Age,
	}

	return response, nil
}

func (s *service) DeleteUser(ctx context.Context, rq entities.DeleteUserRequest) (entities.DeleteUserResponse, error) {
	s.Logger.Log(s.Logger, "delete user", "recevied")

	userId := rq.UserId

	err := s.Repo.DeleteUser(ctx, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			level.Error(s.Logger).Log("error", err)
			return entities.DeleteUserResponse{}, errors.NewUserNotFound()
		}
		level.Error(s.Logger).Log("error", err)
		return entities.DeleteUserResponse{}, errors.NewDataBaseError()
	}

	return entities.DeleteUserResponse{
		Status: entities.Status{
			Message: "User deleted succesfully",
			Code:    0,
		},
	}, nil
}

func generateId() string {
	return uuid.NewString()
}
