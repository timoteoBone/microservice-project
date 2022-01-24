package user_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/timoteoBone/microservice-project/grpcService/pkg/entities"
	myErr "github.com/timoteoBone/microservice-project/grpcService/pkg/errors"
	service "github.com/timoteoBone/microservice-project/grpcService/pkg/user"
	"github.com/timoteoBone/microservice-project/grpcService/pkg/utils"
)

func TestNewService(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	repo := utils.NewRepoMock(logger, mock.Mock{})

	srvc := service.NewService(logger, &repo)

	assert.False(t, srvc == nil)
}

func TestServiceCreateUser(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	user := entities.User{
		Name:  "Timo",
		Pass:  "123",
		Age:   19,
		Email: "timoteo@globant.com",
	}

	userId := utils.GenerateId()

	correctCreateUserRequest := entities.CreateUserRequest{
		Name:  user.Name,
		Pass:  user.Pass,
		Age:   user.Age,
		Email: user.Email,
	}

	Succesfullresponse := entities.CreateUserResponse{
		Status: entities.Status{
			Message: "created successfully",
		}, UserId: userId,
	}

	repo := utils.NewRepoMock(logger, mock.Mock{})
	srvc := service.NewService(logger, &repo)

	t.Run("Create User Valid case", func(t *testing.T) {
		ctx := context.Background()
		repo.On("CreateUser", ctx, user).Return(userId, nil)

		res, err := srvc.CreateUser(ctx, correctCreateUserRequest)
		assert.ErrorIs(t, err, nil)
		assert.Equal(t, Succesfullresponse, res)

	})

}

func TestServiceCreateExistingUser(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	user := entities.User{
		Name: "Timo",
		Pass: "123",
		Age:  19,
	}

	userId := utils.GenerateId()

	correctCreateUserRequest := entities.CreateUserRequest{
		Name: user.Name,
		Pass: user.Pass,
		Age:  user.Age,
	}

	repo := utils.NewRepoMock(logger, mock.Mock{})
	srvc := service.NewService(logger, &repo)

	t.Run("Create User Valid case", func(t *testing.T) {
		ctx := context.Background()
		repo.On("CreateUser", ctx, user).Return(userId, myErr.NewDataBaseError())

		res, err := srvc.CreateUser(ctx, correctCreateUserRequest)
		assert.Equal(t, err.Error(), myErr.NewDataBaseError().Error())
		assert.Empty(t, res)

	})

}

func TestServiceGetExistingUser(t *testing.T) {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	user := entities.User{
		Name: "Timo",
		Pass: "123",
		Age:  19,
	}

	userId := utils.GenerateId()

	correctGetUserRequest := entities.GetUserRequest{
		UserID: userId,
	}

	correctGetUserResponse := entities.GetUserResponse{
		Name: user.Name,
		Id:   userId,
		Age:  user.Age,
	}

	repo := new(utils.RepoSitoryMock)
	srvc := service.NewService(logger, repo)

	ctx := context.Background()

	repo.Mock.On("GetUser", ctx, userId).Return(user, nil)

	res, err := srvc.GetUser(ctx, correctGetUserRequest)
	assert.Equal(t, correctGetUserResponse, res)
	assert.ErrorIs(t, err, nil)

}

func TestServiceGetNonExistingUser(t *testing.T) {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	userId := utils.GenerateId()

	correctGetUserRequest := entities.GetUserRequest{
		UserID: userId,
	}

	expectedErr := myErr.NewUserNotFound()

	repo := new(utils.RepoSitoryMock)
	srvc := service.NewService(logger, repo)
	ctx := context.Background()

	repo.Mock.On("GetUser", ctx, userId).Return(entities.User{}, sql.ErrNoRows)

	res, err := srvc.GetUser(ctx, correctGetUserRequest)
	assert.Equal(t, entities.GetUserResponse{}, res)
	assert.Equal(t, err.Error(), expectedErr.Error())

}

func TestDeleteExistingUser(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	userId := utils.GenerateId()

	correctDeleteUserRequest := entities.DeleteUserRequest{
		UserId: userId,
	}

	succesfullDeleteUserRequest := entities.DeleteUserResponse{
		Status: entities.Status{
			Message: "User deleted succesfully",
			Code:    0,
		},
	}

	repo := new(utils.RepoSitoryMock)
	srvc := service.NewService(logger, repo)

	ctx := context.Background()
	repo.Mock.On("DeleteUser", ctx, userId).Return(nil)

	res, err := srvc.DeleteUser(ctx, correctDeleteUserRequest)
	assert.Equal(t, succesfullDeleteUserRequest, res)
	assert.NoError(t, err)

}

func TestDeleteNonExistingUser(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	userId := utils.GenerateId()

	correctDeleteUserRequest := entities.DeleteUserRequest{
		UserId: userId,
	}

	repo := new(utils.RepoSitoryMock)
	srvc := service.NewService(logger, repo)

	ctx := context.Background()
	repo.Mock.On("DeleteUser", ctx, userId).Return(sql.ErrNoRows)

	res, err := srvc.DeleteUser(ctx, correctDeleteUserRequest)
	assert.Empty(t, res)
	assert.Equal(t, err.Error(), myErr.NewUserNotFound().Error())

}

func TestUpdateExistingUser(t *testing.T) {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	userId := utils.GenerateId()

	userMock := entities.User{
		Name:  "Timo",
		Pass:  "123",
		Age:   20,
		Email: "timoteo@globant.com",
	}

	updateUserRequest := entities.UpdateUserRequest{
		Name:  "Timo",
		Pass:  "123",
		Age:   20,
		Email: "timoteo@globant.com",
		Id:    userId,
	}

	repo := new(utils.RepoSitoryMock)
	srvc := service.NewService(logger, repo)

	ctx := context.Background()
	repo.Mock.On("UpdateUser", ctx, userMock, userId).Return(nil)

	res, err := srvc.UpdateUser(ctx, updateUserRequest)
	assert.Equal(t, "User updated succesfully", res.Status.Message)
	assert.Equal(t, int32(0), res.Status.Code)
	assert.Nil(t, err)

}

func TestUpdateNonExistingUser(t *testing.T) {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	userId := utils.GenerateId()

	userMock := entities.User{
		Name:  "Timo",
		Pass:  "123",
		Age:   20,
		Email: "timoteo@globant.com",
	}

	updateUserRequest := entities.UpdateUserRequest{
		Name:  "Timo",
		Pass:  "123",
		Age:   20,
		Email: "timoteo@globant.com",
		Id:    userId,
	}

	repo := new(utils.RepoSitoryMock)
	srvc := service.NewService(logger, repo)

	ctx := context.Background()
	repo.Mock.On("UpdateUser", ctx, userMock, userId).Return(sql.ErrNoRows)

	res, err := srvc.UpdateUser(ctx, updateUserRequest)
	assert.Equal(t, entities.UpdateUserResponse{}, res)
	assert.NotNil(t, err)

}
