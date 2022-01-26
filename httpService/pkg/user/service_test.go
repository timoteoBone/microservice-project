package user_test

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/timoteoBone/microservice-project/grpcService/pkg/entities"
	errors "github.com/timoteoBone/microservice-project/grpcService/pkg/errors"
	"github.com/timoteoBone/microservice-project/httpService/pkg/user"
	util "github.com/timoteoBone/microservice-project/httpService/pkg/utils"
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

	repo := util.NewRepositoryMock()

	srvc := user.NewService(&repo, logger)

	assert.NotNil(t, srvc)

}

func TestCreateValidUser(t *testing.T) {
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

	repo := util.NewRepositoryMock()

	srvc := user.NewService(&repo, logger)

	var (
		correctCreateRequest entities.CreateUserRequest = entities.CreateUserRequest{
			Name:  "Timo",
			Age:   19,
			Pass:  "123",
			Email: "timoteo@globant.com",
		}

		correctCreateResponse entities.CreateUserResponse = entities.CreateUserResponse{
			Status: entities.Status{Message: "created successfully"},
		}
	)

	t.Run("Create User Valid Case", func(t *testing.T) {
		ctx := context.Background()
		repo.On("CreateUser", ctx, correctCreateRequest).Return(correctCreateResponse, nil)

		res, err := srvc.CreateUser(ctx, correctCreateRequest)
		assert.ErrorIs(t, err, nil)
		assert.Equal(t, correctCreateResponse, res)
	})

}

func TestCreateUserEmptyFields(t *testing.T) {
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

	repo := util.NewRepositoryMock()

	srvc := user.NewService(&repo, logger)

	var (
		correctCreateRequest entities.CreateUserRequest = entities.CreateUserRequest{
			Name: "Timo",
		}
	)

	t.Run("Create User Valid Case", func(t *testing.T) {
		ctx := context.Background()

		res, err := srvc.CreateUser(ctx, correctCreateRequest)
		assert.Equal(t, err.Error(), errors.NewBadRequest().Error())
		assert.Empty(t, res)
	})

}

func TestGetExistingUser(t *testing.T) {
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

	repo := util.NewRepositoryMock()

	srvc := user.NewService(&repo, logger)

	var (
		correctGetUserRequest entities.GetUserRequest = entities.GetUserRequest{
			UserID: "2abc-323kol",
		}

		correctGetUserResponse entities.GetUserResponse = entities.GetUserResponse{
			Name:  "Timo",
			Id:    "2abc-323kol",
			Age:   19,
			Email: "timoteo@globant.com",
		}
	)

	t.Run("Get Existing User", func(t *testing.T) {
		ctx := context.Background()
		repo.On("GetUser", ctx, correctGetUserRequest).Return(correctGetUserResponse, nil)

		res, err := srvc.GetUser(ctx, correctGetUserRequest)
		assert.Equal(t, correctGetUserResponse, res)
		assert.ErrorIs(t, err, nil)

	})

}

func TestGetNonExistingUser(t *testing.T) {
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

	repo := util.NewRepositoryMock()

	srvc := user.NewService(&repo, logger)

	var (
		correctGetUserRequest entities.GetUserRequest = entities.GetUserRequest{
			UserID: "2abc-323kol",
		}
	)

	newErr := errors.NewUserNotFound()

	t.Run("Get Existing User", func(t *testing.T) {
		ctx := context.Background()
		repo.On("GetUser", ctx, correctGetUserRequest).Return(entities.GetUserResponse{}, newErr)

		res, err := srvc.GetUser(ctx, correctGetUserRequest)
		assert.Empty(t, res)
		assert.ErrorIs(t, err, newErr)

	})

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

	var (
		correctDeleteUserReq = entities.DeleteUserRequest{
			UserId: "1234567abcd",
		}

		correctDeleteUserResponse = entities.DeleteUserResponse{
			Status: entities.Status{
				Message: "User deleted succesfully",
				Code:    0,
			},
		}
	)

	repo := util.NewRepositoryMock()

	srvc := user.NewService(&repo, logger)

	t.Run("Delete Existing User", func(t *testing.T) {
		ctx := context.Background()
		repo.On("DeleteUser", ctx, correctDeleteUserReq).Return(correctDeleteUserResponse, nil)

		res, err := srvc.DeleteUser(ctx, correctDeleteUserReq)
		assert.Equal(t, correctDeleteUserResponse, res)
		assert.Nil(t, err)

	})

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

	var (
		correctDeleteUserReq = entities.DeleteUserRequest{
			UserId: "1234567abcd",
		}
	)

	newErr := errors.NewUserNotFound()

	repo := util.NewRepositoryMock()

	srvc := user.NewService(&repo, logger)

	t.Run("Delete Existing User", func(t *testing.T) {
		ctx := context.Background()
		repo.On("DeleteUser", ctx, correctDeleteUserReq).Return(entities.DeleteUserResponse{}, newErr)

		res, err := srvc.DeleteUser(ctx, correctDeleteUserReq)
		assert.Empty(t, res)
		assert.Equal(t, newErr, err)

	})

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

	var (
		correctUpdateRequest entities.UpdateUserRequest = entities.UpdateUserRequest{
			Name:  "Timo",
			Age:   19,
			Pass:  "123",
			Email: "timoteo@globant.com",
		}

		correctUpdateResponse entities.UpdateUserResponse = entities.UpdateUserResponse{
			Status: entities.Status{Message: "User updated succesfully", Code: 0},
		}
	)

	newErr := errors.NewUserNotFound()

	repo := util.NewRepositoryMock()

	srvc := user.NewService(&repo, logger)

	t.Run("Update Existing User", func(t *testing.T) {
		ctx := context.Background()
		repo.On("UpdateUser", ctx, correctUpdateRequest).Return(correctUpdateResponse, nil)

		res, err := srvc.UpdateUser(ctx, correctUpdateRequest)
		assert.Empty(t, res)
		assert.Equal(t, newErr, err)

	})

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

	var (
		correctUpdateRequest entities.UpdateUserRequest = entities.UpdateUserRequest{
			Name:  "Timo",
			Age:   19,
			Pass:  "123",
			Email: "timoteo@globant.com",
		}
	)

	newErr := errors.NewUserNotFound()

	repo := util.NewRepositoryMock()

	srvc := user.NewService(&repo, logger)

	t.Run("Update Non Existing User", func(t *testing.T) {
		ctx := context.Background()
		repo.On("UpdateUser", ctx, correctUpdateRequest).Return(entities.UpdateUserResponse{}, nil)

		res, err := srvc.UpdateUser(ctx, correctUpdateRequest)
		assert.Empty(t, res)
		assert.Equal(t, newErr, err)

	})

}
