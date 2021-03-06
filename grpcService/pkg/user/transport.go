package user

import (
	"context"

	gr "github.com/go-kit/kit/transport/grpc"

	"github.com/timoteoBone/microservice-project/grpcService/pkg/entities"
	customErr "github.com/timoteoBone/microservice-project/grpcService/pkg/errors"
	proto "github.com/timoteoBone/microservice-project/grpcService/pkg/pb"
)

type gRPCSv struct {
	createUs gr.Handler
	getUs    gr.Handler
	deleteUs gr.Handler
	proto.UnimplementedUserServiceServer
}

func NewGrpcServer(end Endpoints) proto.UserServiceServer {

	return &gRPCSv{
		createUs: gr.NewServer(
			end.CreateUser,
			decodeCreateUserRequest,
			encodeCreateUserResponse,
		),

		getUs: gr.NewServer(
			end.GetUser,
			decodeGetUserRequest,
			encodeGetUserResponse,
		),

		deleteUs: gr.NewServer(
			end.DeleteUser,
			decodeDeleteUserRequest,
			encodeDeleteUserRequest,
		),
	}
}

func (g *gRPCSv) CreateUser(ctx context.Context, rq *proto.CreateUserRequest) (rs *proto.CreateUserResponse, err error) {
	_, resp, err := g.createUs.ServeGRPC(ctx, rq)

	if err != nil {

		return nil, err
	}

	return resp.(*proto.CreateUserResponse), nil
}

func (g *gRPCSv) GetUser(ctx context.Context, rq *proto.GetUserRequest) (rs *proto.GetUserResponse, err error) {
	_, resp, err := g.getUs.ServeGRPC(ctx, rq)

	if err != nil {
		return nil, err
	}

	return resp.(*proto.GetUserResponse), nil
}

func (g *gRPCSv) DeleteUser(ctx context.Context, rq *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	_, resp, err := g.deleteUs.ServeGRPC(ctx, rq)
	if err != nil {
		return nil, err
	}

	return resp.(*proto.DeleteUserResponse), nil
}

func decodeCreateUserRequest(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := request.(*proto.CreateUserRequest)

	if !err {
		return nil, customErr.NewGrpcError()
	}

	return entities.CreateUserRequest{
		Name:  res.Name,
		Age:   res.Age,
		Pass:  res.Pass,
		Email: res.Email,
	}, nil

}

func encodeCreateUserResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(entities.CreateUserResponse)
	status := &proto.Status{Message: res.Status.Message, Code: res.Status.Code}
	protoResp := &proto.CreateUserResponse{User_Id: res.UserId, Status: status}
	return protoResp, nil
}

func decodeGetUserRequest(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := request.(*proto.GetUserRequest)

	if !err {
		return nil, customErr.NewGrpcError()
	}

	return entities.GetUserRequest{
		UserID: res.User_Id,
	}, nil

}

func encodeGetUserResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res, valid := response.(entities.GetUserResponse)
	if !valid {
		return nil, customErr.NewGrpcError()
	}
	protoResp := &proto.GetUserResponse{Id: res.Id, Name: res.Name, Age: res.Age}
	return protoResp, nil
}

func decodeDeleteUserRequest(ctx context.Context, request interface{}) (interface{}, error) {
	res, valid := request.(*proto.DeleteUserRequest)
	if !valid {
		return nil, customErr.NewGrpcError()
	}

	return entities.DeleteUserRequest{
		UserId: res.User_Id,
	}, nil
}

func encodeDeleteUserRequest(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(entities.DeleteUserResponse)
	protoResp := &proto.DeleteUserResponse{Status: &proto.Status{Message: resp.Status.Message, Code: resp.Status.Code}}
	return protoResp, nil
}
