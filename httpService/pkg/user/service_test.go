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

func TestCreateUser(t *testing.T) {
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
		assert.Equal(t, err.Error(), errors.NewFieldsMissing().Error())
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
