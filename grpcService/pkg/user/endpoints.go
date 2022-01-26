package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/timoteoBone/microservice-project/grpcService/pkg/entities"
	"github.com/timoteoBone/microservice-project/grpcService/pkg/errors"
)

type Service interface {
	GetUser(ctx context.Context, userReq entities.GetUserRequest) (entities.GetUserResponse, error)
	CreateUser(ctx context.Context, userReq entities.CreateUserRequest) (entities.CreateUserResponse, error)
	DeleteUser(ctx context.Context, userReq entities.DeleteUserRequest) (entities.DeleteUserResponse, error)
	UpdateUser(ctx context.Context, userReq entities.UpdateUserRequest) (entities.UpdateUserResponse, error)
}

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
	DeleteUser endpoint.Endpoint
	UpdateUser endpoint.Endpoint
}

func MakeEndpoint(s Service) Endpoints {
	return Endpoints{
		CreateUser: MakeCreateUserEndpoint(s),
		GetUser:    MakeGetUserEndpoint(s),
		DeleteUser: MakeDeleteUserEndpoint(s),
		UpdateUser: MakeUpdateUserEndpoint(s),
	}
}

func MakeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		req, valid := request.(entities.CreateUserRequest)
		if !valid {
			return nil, errors.NewBadRequest()
		}

		c, err := s.CreateUser(ctx, req)
		if err != nil {
			return nil, err
		}

		return c, nil

	}
}

func MakeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, valid := request.(entities.GetUserRequest)
		if !valid {
			return nil, errors.NewBadRequest()
		}

		c, err := s.GetUser(ctx, req)
		if err != nil {
			return nil, err
		}

		return c, nil

	}
}

func MakeDeleteUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, valid := request.(entities.DeleteUserRequest)
		if !valid {
			return nil, errors.NewBadRequest()
		}

		c, err := s.DeleteUser(ctx, req)
		if err != nil {

			return nil, err
		}

		return c, nil

	}
}

func MakeUpdateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, valid := request.(entities.UpdateUserRequest)
		if !valid {
			return nil, errors.NewBadRequest()
		}
		c, err := s.UpdateUser(ctx, req)
		if err != nil {
			return nil, err
		}

		return c, nil

	}
}
