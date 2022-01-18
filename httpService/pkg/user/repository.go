package user

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
	util "github.com/timoteoBone/project-microservice/httpService/pkg/utils"

	proto "github.com/timoteoBone/project-microservice/grpcService/pkg/pb"
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

	protoReq := util.CreateToProto(rq)

	resp, err := client.CreateUser(ctx, protoReq)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return entities.CreateUserResponse{}, err
	}

	res := util.CreateFromProto(resp)

	return res, nil

}

func (repo *grpcClient) GetUser(ctx context.Context, rq entities.GetUserRequest) (entities.GetUserResponse, error) {
	logger := log.With(repo.logger, "get user request", "received")

	client := proto.NewUserServiceClient(repo.server)

	protoReq := util.GetToProto(rq)

	protoRes, err := client.GetUser(ctx, protoReq)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return entities.GetUserResponse{}, err
	}

	resp := util.GetFromProto(protoRes)

	return resp, nil
}
