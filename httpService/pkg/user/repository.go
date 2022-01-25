package user

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/timoteoBone/microservice-project/grpcService/pkg/entities"
	mapper "github.com/timoteoBone/microservice-project/httpService/pkg/mapper"

	proto "github.com/timoteoBone/microservice-project/grpcService/pkg/pb"
	"google.golang.org/grpc"
)

type grpcClient struct {
	server *grpc.ClientConn
	logger log.Logger
}

func NewgRPClient(log log.Logger, sv *grpc.ClientConn) *grpcClient {
	return &grpcClient{sv, log}
}

func (repo *grpcClient) CreateUser(ctx context.Context, rq entities.CreateUserRequest) (entities.CreateUserResponse, error) {
	logger := log.With(repo.logger, "create user", "recevied")

	client := proto.NewUserServiceClient(repo.server)

	protoReq := mapper.CreateToProto(rq)

	resp, err := client.CreateUser(ctx, protoReq)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return entities.CreateUserResponse{}, err
	}

	res := mapper.CreateFromProto(resp)

	return res, nil

}

func (repo *grpcClient) GetUser(ctx context.Context, rq entities.GetUserRequest) (entities.GetUserResponse, error) {
	logger := log.With(repo.logger, "get user request", "received")
	client := proto.NewUserServiceClient(repo.server)
	protoReq := mapper.GetToProto(rq)

	protoRes, err := client.GetUser(ctx, protoReq)

	if err != nil {

		level.Error(logger).Log(err)
		return entities.GetUserResponse{}, err
	}

	resp := mapper.GetFromProto(protoRes)

	return resp, nil
}

func (repo *grpcClient) DeleteUser(ctx context.Context, rq entities.DeleteUserRequest) (entities.DeleteUserResponse, error) {
	logger := log.With(repo.logger, "delete user request", "received")

	client := proto.NewUserServiceClient(repo.server)

	protoReq := mapper.DeleteToProto(rq)

	resp, err := client.DeleteUser(ctx, protoReq)
	if err != nil {
		level.Error(logger).Log(err)
		return entities.DeleteUserResponse{}, err
	}

	ret := mapper.DeleteFromProto(resp)

	return ret, nil

}

func (repo *grpcClient) UpdateUser(ctx context.Context, rq entities.UpdateUserRequest) (entities.UpdateUserResponse, error) {
	logger := log.With(repo.logger, "update user request", "received")

	client := proto.NewUserServiceClient(repo.server)

	protoReq := mapper.UpdateToProto(rq)

	protoResp, err := client.UpdateUser(ctx, protoReq)
	if err != nil {
		level.Error(logger).Log(err)
		return entities.UpdateUserResponse{}, err
	}

	resp := mapper.UpdateFromProto(protoResp)

	return resp, nil

}
