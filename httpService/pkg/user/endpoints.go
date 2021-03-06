package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/timoteoBone/microservice-project/grpcService/pkg/entities"
	errs "github.com/timoteoBone/microservice-project/grpcService/pkg/errors"
)

type Service interface {
	CreateUser(ctx context.Context, rq entities.CreateUserRequest) (entities.CreateUserResponse, error)
	GetUser(ctx context.Context, rq entities.GetUserRequest) (entities.GetUserResponse, error)
	DeleteUser(ctx context.Context, rq entities.DeleteUserRequest) (entities.DeleteUserResponse, error)
}

type Endpoints struct {
	CreateUs endpoint.Endpoint
	GetUs    endpoint.Endpoint
	DeleteUs endpoint.Endpoint
}

func MakeEndpoints(s Service) *Endpoints {

	return &Endpoints{
		CreateUs: MakeCreateUserEndpoint(s),
		GetUs:    MakeGetUserEndpoint(s),
		DeleteUs: MakeDeleteUserEndpoint(s),
	}
}

func MakeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, rq interface{}) (interface{}, error) {
		request, valid := rq.(entities.CreateUserRequest)

		if !valid {
			return nil, errs.NewFieldsMissing()
		}

		res, err := s.CreateUser(ctx, request)
		if err != nil {
			return nil, err
		}

		return res, nil

	}
}

func MakeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, rq interface{}) (interface{}, error) {
		request, valid := rq.(entities.GetUserRequest)
		if !valid {
			return nil, errs.NewFieldsMissing()
		}

		res, err := s.GetUser(ctx, request)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func MakeDeleteUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, rq interface{}) (interface{}, error) {
		request, valid := rq.(entities.DeleteUserRequest)
		if !valid {
			return nil, errs.NewFieldsMissing()
		}

		res, err := s.DeleteUser(ctx, request)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}
