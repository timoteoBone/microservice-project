package util

import (
	"github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
	proto "github.com/timoteoBone/project-microservice/grpcService/pkg/pb"
)

func CreateToProto(req entities.CreateUserRequest) *proto.CreateUserRequest {
	return &proto.CreateUserRequest{
		Name:  req.Name,
		Pass:  req.Pass,
		Age:   req.Age,
		Email: req.Email,
	}
}

func GetToProto(req entities.GetUserRequest) *proto.GetUserRequest {
	return &proto.GetUserRequest{
		User_Id: req.UserID,
	}
}

func CreateFromProto(resp *proto.CreateUserResponse) entities.CreateUserResponse {
	return entities.CreateUserResponse{
		Status: entities.Status{
			Message: resp.Status.Message,
			Code:    resp.Status.Code,
		}, UserId: resp.User_Id,
	}
}

func GetFromProto(resp *proto.GetUserResponse) entities.GetUserResponse {
	return entities.GetUserResponse{
		Name:  resp.Name,
		Id:    resp.Id,
		Age:   resp.Age,
		Email: resp.Email,
		Status: entities.Status{
			Message: resp.Status.Message,
			Code:    resp.Status.Code,
		},
	}
}
