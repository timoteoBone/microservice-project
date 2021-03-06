package util

import (
	"github.com/timoteoBone/microservice-project/grpcService/pkg/errors"
	err "github.com/timoteoBone/microservice-project/grpcService/pkg/errors"

	"github.com/timoteoBone/microservice-project/grpcService/pkg/entities"
)

func ValidateCreateUserRequest(user entities.CreateUserRequest) error {
	if user.Age < 1 || len(user.Name) < 1 || len(user.Pass) < 1 || len(user.Email) < 1 {
		return err.NewFieldsMissing()
	}
	return nil
}

func ValidateGetUserRequest(id entities.GetUserRequest) error {
	if len(id.UserID) < 1 {
		return errors.NewFieldsMissing()
	}
	return nil
}
