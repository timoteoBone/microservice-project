package user_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/user"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/utils"
)

func TestNewRepo(t *testing.T) {
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

	db, _ := utils.NewMock(logger)

	repo := user.NewSQL(db, logger)

	assert.NotNil(t, repo)
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

	userMock := entities.User{
		Name:  "Timo",
		Age:   19,
		Pass:  "1234",
		Email: "timoteo@globant.com",
	}

	db, mock := utils.NewMock(logger)
	defer db.Close()

	userId := utils.GenerateId()

	repo := user.NewSQL(db, logger)

	testCases := []struct {
		Name           string
		User           entities.User
		UserID         string
		buildMock      func(mock sqlmock.Sqlmock, user entities.User)
		assertResponse func(t *testing.T, id string, err error)
	}{

		{
			Name:   "Create User Valid Case",
			User:   userMock,
			UserID: userId,
			buildMock: func(mock sqlmock.Sqlmock, user entities.User) {
				mock.ExpectPrepare(utils.CreateUserQuery)
				mock.ExpectExec(utils.CreateUserQuery).WithArgs(user.Name, userId, user.Pass, user.Age, user.Email).WillReturnResult(sqlmock.NewResult(int64(1), int64(1)))
			},
			assertResponse: func(t *testing.T, id string, err error) {
				assert.Equal(t, userId, id)
				assert.NoError(t, err)
			},
		},
		{
			Name: "Create Already Existing User",
			User: userMock,
			buildMock: func(mock sqlmock.Sqlmock, user entities.User) {
				mock.ExpectPrepare(utils.CreateUserQuery)
				mock.ExpectExec(utils.CreateUserQuery).WithArgs(user.Name, userId, user.Pass, user.Age, user.Email).WillReturnError(sqlmock.ErrCancelled)
			},
			assertResponse: func(t *testing.T, id string, err error) {
				assert.Equal(t, "", id)
				assert.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			ctx := context.Background()
			tc.buildMock(mock, tc.User)

			id, err := repo.CreateUser(ctx, tc.User, userId)
			tc.assertResponse(t, id, err)
		})
	}
}

func TestGetUser(t *testing.T) {
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

	userMock := entities.User{
		Name:  "Timoteo",
		Pass:  "sdsfodemf",
		Age:   19,
		Email: "timoteo@globant.com",
	}

	db, mock := utils.NewMock(logger)
	defer db.Close()

	userId := utils.GenerateId()

	repo := user.NewSQL(db, logger)

	testCases := []struct {
		Name           string
		UserID         string
		buildMock      func(mock sqlmock.Sqlmock, userId string)
		assertResponse func(t *testing.T, response entities.User, err error)
	}{
		{
			Name:   "Get Existing User",
			UserID: userId,
			buildMock: func(mock sqlmock.Sqlmock, userId string) {
				res := sqlmock.NewRows([]string{"first_name", "age", "email"}).AddRow(userMock.Name, userMock.Age, userMock.Email)
				mock.ExpectPrepare(utils.GetUserQuery)
				mock.ExpectQuery(utils.GetUserQuery).WithArgs(userId).WillReturnRows(res)
			},
			assertResponse: func(t *testing.T, resp entities.User, err error) {
				assert.Nil(t, err)
				assert.Equal(t, userMock.Email, resp.Email)
			},
		},

		{
			Name:   "Get non existing user",
			UserID: userId,
			buildMock: func(mock sqlmock.Sqlmock, userId string) {
				res := sqlmock.NewRows([]string{"first_name", "age", "email"})
				mock.ExpectPrepare(utils.GetUserQuery)
				mock.ExpectQuery(utils.GetUserQuery).WithArgs(userId).WillReturnRows(res)
			},
			assertResponse: func(t *testing.T, resp entities.User, err error) {
				assert.Error(t, err)
				assert.Empty(t, resp)
			},
		},
	}

	for _, tc := range testCases {

		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()
			tc.buildMock(mock, userId)

			res, err := repo.GetUser(ctx, tc.UserID)
			tc.assertResponse(t, res, err)
		})
	}

}

func TestDeleteUser(t *testing.T) {
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

	db, mock := utils.NewMock(logger)
	defer db.Close()

	userId := utils.GenerateId()

	repo := user.NewSQL(db, logger)

	testCases := []struct {
		Name           string
		UserID         string
		buildMock      func(mock sqlmock.Sqlmock, userId string)
		assertResponse func(t *testing.T, err error)
	}{
		{
			Name:   "Delete Existing User",
			UserID: userId,
			buildMock: func(mock sqlmock.Sqlmock, userId string) {
				mock.ExpectPrepare(utils.DeleteUserQuery)
				mock.ExpectExec(utils.DeleteUserQuery).WithArgs(userId).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			assertResponse: func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},

		{
			Name:   "Delete non existing user",
			UserID: userId,
			buildMock: func(mock sqlmock.Sqlmock, userId string) {
				mock.ExpectPrepare(utils.DeleteUserQuery)
				mock.ExpectExec(utils.DeleteUserQuery).WithArgs(userId).WillReturnError(sql.ErrNoRows)
			},
			assertResponse: func(t *testing.T, err error) {
				assert.ErrorIs(t, sql.ErrNoRows, err)
			},
		},
	}

	for _, tc := range testCases {

		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()
			tc.buildMock(mock, userId)

			err := repo.DeleteUser(ctx, tc.UserID)
			tc.assertResponse(t, err)
		})
	}

}
