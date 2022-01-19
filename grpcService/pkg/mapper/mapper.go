package mapper

import "github.com/timoteoBone/microservice-project/grpcService/pkg/entities"

func CreateUserRequestToUser(userReq entities.CreateUserRequest) entities.User {

	user := entities.User{
		userReq.Name,
		userReq.Pass,
		userReq.Age,
		userReq.Email,
	}
	return user
}

/*
func UpdateUserToMap(user entities.UpdateUserRequest) map[string]interface{} {

}
*/
