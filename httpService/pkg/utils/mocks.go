package util

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/timoteoBone/microservice-project/grpcService/pkg/entities"
	proto "github.com/timoteoBone/microservice-project/grpcService/pkg/pb"
)

type RepositoryMock struct {
	mock.Mock
}

func NewRepositoryMock() RepositoryMock {
	return RepositoryMock{mock.Mock{}}
}

func (repo *RepositoryMock) CreateUser(ctx context.Context, rq entities.CreateUserRequest) (entities.CreateUserResponse, error) {

	args := repo.Mock.Called(ctx, rq)
	response := args[0]

	return response.(entities.CreateUserResponse), args.Error(1)
}

func (repo *RepositoryMock) GetUser(ctx context.Context, rq entities.GetUserRequest) (entities.GetUserResponse, error) {
	args := repo.Mock.Called(ctx, rq)
	response := args[0]

	return response.(entities.GetUserResponse), args.Error(1)

}

func (repo *RepositoryMock) DeleteUser(ctx context.Context, rq entities.DeleteUserRequest) (entities.DeleteUserResponse, error) {
	args := repo.Mock.Called(ctx, rq)
	response := args[0]

	return response.(entities.DeleteUserResponse), args.Error(1)
}

func (repo *RepositoryMock) UpdateUser(ctx context.Context, rq entities.UpdateUserRequest) (entities.UpdateUserResponse, error) {
	args := repo.Mock.Called(ctx, rq)

	return args[0].(entities.UpdateUserResponse), args.Error(1)
}

type GrpcClientMock struct {
	mock.Mock
	proto.UnimplementedUserServiceServer
}

func NewgRPClientMock() *GrpcClientMock {
	return &GrpcClientMock{Mock: mock.Mock{}}
}

func (cli *GrpcClientMock) CreateUser(ctx context.Context, rq *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	args := cli.Mock.Called(ctx, rq)

	return args[0].(*proto.CreateUserResponse), args.Error(1)

}

func (cli *GrpcClientMock) GetUser(ctx context.Context, rq *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	args := cli.Mock.Called(ctx, rq)

	return args[0].(*proto.GetUserResponse), args.Error(1)
}

func (cli *GrpcClientMock) DeleteUser(ctx context.Context, rq *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	args := cli.Mock.Called(ctx, rq)

	return args[0].(*proto.DeleteUserResponse), args.Error(1)
}

func (cli *GrpcClientMock) UpdateUser(ctx context.Context, rq *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	args := cli.Mock.Called(ctx, rq)

	return args[0].(*proto.UpdateUserResponse), args.Error(1)
}
