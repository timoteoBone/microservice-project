package mapper

import (
	"github.com/timoteoBone/microservice-project/grpcService/pkg/entities"
)

func CreateUserRequestToUser(userReq entities.CreateUserRequest) entities.User {

	user := entities.User{
		userReq.Name,
		userReq.Pass,
		userReq.Age,
		userReq.Email,
	}
	return user
}

func UpdateUserToUser(userData entities.UpdateUserRequest) entities.User {
	return entities.User{
		Name:  userData.Name,
		Pass:  userData.Pass,
		Age:   userData.Age,
		Email: userData.Email,
	}
}
