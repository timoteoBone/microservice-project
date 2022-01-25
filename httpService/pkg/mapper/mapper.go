package util

import (
	"github.com/timoteoBone/microservice-project/grpcService/pkg/entities"
	proto "github.com/timoteoBone/microservice-project/grpcService/pkg/pb"
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
	}
}

func DeleteToProto(req entities.DeleteUserRequest) *proto.DeleteUserRequest {
	return &proto.DeleteUserRequest{
		User_Id: req.UserId,
	}
}

func DeleteFromProto(resp *proto.DeleteUserResponse) entities.DeleteUserResponse {
	return entities.DeleteUserResponse{
		Status: entities.Status{
			Message: resp.Status.Message,
			Code:    resp.Status.Code,
		},
	}
}

func UpdateToProto(req entities.UpdateUserRequest) *proto.UpdateUserRequest {
	return &proto.UpdateUserRequest{
		Name:  req.Name,
		Pass:  req.Pass,
		Age:   req.Age,
		Email: req.Email,
		Id:    req.Id,
	}
}

func UpdateFromProto(resp *proto.UpdateUserResponse) entities.UpdateUserResponse {
	return entities.UpdateUserResponse{
		Status: entities.Status{
			Message: resp.Status.Message,
			Code:    resp.Status.Code,
		},
	}
}
