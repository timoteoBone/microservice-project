package user_test

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	logger "github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/timoteoBone/microservice-project/grpcService/pkg/entities"
	"github.com/timoteoBone/microservice-project/grpcService/pkg/errors"
	proto "github.com/timoteoBone/microservice-project/grpcService/pkg/pb"
	"github.com/timoteoBone/microservice-project/httpService/pkg/user"
	util "github.com/timoteoBone/microservice-project/httpService/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func dialer(mock *util.GrpcClientMock) func(context.Context, string) (net.Conn, error) {

	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	proto.RegisterUserServiceServer(server, mock)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestNewRepo(t *testing.T) {

	var log logger.Logger
	{
		log = logger.NewLogfmtLogger(os.Stderr)
		log = logger.NewSyncLogger(log)
		log = logger.With(log,
			"service", "grpcUserService",
			"time:", logger.DefaultTimestampUTC,
			"caller", logger.DefaultCaller,
		)
	}

	ctx := context.Background()

	grpcMock := new(util.GrpcClientMock)

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(grpcMock)))
	if err != nil {
		level.Error(log).Log(err)
	}
	defer conn.Close()

	repo := user.NewgRPClient(log, conn)
	assert.NotNil(t, repo)

}

func TestCreateCorrectUser(t *testing.T) {

	var log logger.Logger
	{
		log = logger.NewLogfmtLogger(os.Stderr)
		log = logger.NewSyncLogger(log)
		log = logger.With(log,
			"service", "grpcUserService",
			"time:", logger.DefaultTimestampUTC,
			"caller", logger.DefaultCaller,
		)
	}

	ctx := context.Background()

	grpcMock := util.NewgRPClientMock()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(grpcMock)))
	if err != nil {
		level.Error(log).Log(err)
	}
	defer conn.Close()

	grpClient := user.NewgRPClient(log, conn)

	t.Run("Create User Valid Case", func(t *testing.T) {

		grpcMock.On("CreateUser", mock.Anything, &proto.CreateUserRequest{
			Name:  "Timo",
			Pass:  "123",
			Age:   19,
			Email: "timoteo@globant.com",
		}).Return(&proto.CreateUserResponse{
			Status: &proto.Status{
				Message: "created successfully",
				Code:    0,
			},
		}, nil)

		resp, err := grpClient.CreateUser(ctx, entities.CreateUserRequest{
			Name:  "Timo",
			Pass:  "123",
			Age:   19,
			Email: "timoteo@globant.com",
		})
		assert.Equal(t, int32(0), resp.Status.Code)
		assert.Equal(t, "created successfully", resp.Status.Message)
		assert.Nil(t, err)

	})

}

func TestCreateUserAlreadyExists(t *testing.T) {

	var log logger.Logger
	{
		log = logger.NewLogfmtLogger(os.Stderr)
		log = logger.NewSyncLogger(log)
		log = logger.With(log,
			"service", "grpcUserService",
			"time:", logger.DefaultTimestampUTC,
			"caller", logger.DefaultCaller,
		)
	}

	ctx := context.Background()

	grpcMock := util.NewgRPClientMock()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(grpcMock)))
	if err != nil {
		level.Error(log).Log(err)
	}
	defer conn.Close()

	grpClient := user.NewgRPClient(log, conn)

	t.Run("Create User Invalid Case", func(t *testing.T) {

		grpcMock.On("CreateUser", mock.Anything, &proto.CreateUserRequest{
			Name:  "Timo",
			Pass:  "123",
			Age:   19,
			Email: "timoteo@globant.com",
		}).Return(&proto.CreateUserResponse{}, errors.NewUserAlreadyExists())

		resp, err := grpClient.CreateUser(ctx, entities.CreateUserRequest{
			Name:  "Timo",
			Pass:  "123",
			Age:   19,
			Email: "timoteo@globant.com",
		})
		assert.Empty(t, resp)
		assert.Error(t, err)

	})

}

func TestGetCorrectUser(t *testing.T) {

	var log logger.Logger
	{
		log = logger.NewLogfmtLogger(os.Stderr)
		log = logger.NewSyncLogger(log)
		log = logger.With(log,
			"service", "grpcUserService",
			"time:", logger.DefaultTimestampUTC,
			"caller", logger.DefaultCaller,
		)
	}

	ctx := context.Background()

	grpcMock := util.NewgRPClientMock()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(grpcMock)))
	if err != nil {
		level.Error(log).Log(err)
	}
	defer conn.Close()

	grpClient := user.NewgRPClient(log, conn)

	t.Run("Get User Valid Case", func(t *testing.T) {

		grpcMock.On("GetUser", mock.Anything, &proto.GetUserRequest{
			User_Id: "389384340jfd0f03",
		}).Return(&proto.GetUserResponse{
			Name:  "Timo",
			Id:    "389384340jfd0f03",
			Email: "timoteo@globant.com",
			Age:   19,
		}, nil)

		resp, err := grpClient.GetUser(ctx, entities.GetUserRequest{
			UserID: "389384340jfd0f03",
		})
		assert.Exactly(t, entities.GetUserResponse{
			Name:  "Timo",
			Id:    "389384340jfd0f03",
			Email: "timoteo@globant.com",
			Age:   19,
		}, resp)
		assert.Nil(t, err)

	})

}

func TestGetUserThatDoesntExists(t *testing.T) {

	var log logger.Logger
	{
		log = logger.NewLogfmtLogger(os.Stderr)
		log = logger.NewSyncLogger(log)
		log = logger.With(log,
			"service", "grpcUserService",
			"time:", logger.DefaultTimestampUTC,
			"caller", logger.DefaultCaller,
		)
	}

	ctx := context.Background()

	grpcMock := util.NewgRPClientMock()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(grpcMock)))
	if err != nil {
		level.Error(log).Log(err)
	}
	defer conn.Close()

	grpClient := user.NewgRPClient(log, conn)

	t.Run("Get User Invalid Case", func(t *testing.T) {

		grpcMock.On("GetUser", mock.Anything, &proto.GetUserRequest{
			User_Id: "389384340jfd0f03",
		}).Return(&proto.GetUserResponse{}, errors.NewUserNotFound())

		resp, err := grpClient.GetUser(ctx, entities.GetUserRequest{
			UserID: "389384340jfd0f03",
		})
		assert.Exactly(t, entities.GetUserResponse{}, resp)
		assert.Error(t, err)

	})

}

func TestDeleteCorrectUser(t *testing.T) {

	var log logger.Logger
	{
		log = logger.NewLogfmtLogger(os.Stderr)
		log = logger.NewSyncLogger(log)
		log = logger.With(log,
			"service", "grpcUserService",
			"time:", logger.DefaultTimestampUTC,
			"caller", logger.DefaultCaller,
		)
	}

	ctx := context.Background()

	grpcMock := util.NewgRPClientMock()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(grpcMock)))
	if err != nil {
		level.Error(log).Log(err)
	}
	defer conn.Close()

	grpClient := user.NewgRPClient(log, conn)

	t.Run("Delete User Valid Case", func(t *testing.T) {

		grpcMock.On("DeleteUser", mock.Anything, &proto.DeleteUserRequest{
			User_Id: "389384340jfd0f03",
		}).Return(&proto.DeleteUserResponse{
			Status: &proto.Status{
				Message: "User deleted succesfully",
				Code:    0,
			},
		}, nil)

		resp, err := grpClient.DeleteUser(ctx, entities.DeleteUserRequest{
			UserId: "389384340jfd0f03",
		})
		assert.Exactly(t, entities.DeleteUserResponse{
			Status: entities.Status{
				Message: "User deleted succesfully",
				Code:    0,
			},
		}, resp)
		assert.Nil(t, err)

	})

}

func TestDeleteUserThatDoesntExists(t *testing.T) {

	var log logger.Logger
	{
		log = logger.NewLogfmtLogger(os.Stderr)
		log = logger.NewSyncLogger(log)
		log = logger.With(log,
			"service", "grpcUserService",
			"time:", logger.DefaultTimestampUTC,
			"caller", logger.DefaultCaller,
		)
	}

	ctx := context.Background()

	grpcMock := util.NewgRPClientMock()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(grpcMock)))
	if err != nil {
		level.Error(log).Log(err)
	}
	defer conn.Close()

	grpClient := user.NewgRPClient(log, conn)

	t.Run("Delete User Invalid Case", func(t *testing.T) {

		grpcMock.On("DeleteUser", mock.Anything, &proto.DeleteUserRequest{
			User_Id: "389384340jfd0f03",
		}).Return(&proto.DeleteUserResponse{}, errors.NewUserNotFound())

		resp, err := grpClient.DeleteUser(ctx, entities.DeleteUserRequest{
			UserId: "389384340jfd0f03",
		})
		assert.Exactly(t, entities.DeleteUserResponse{}, resp)
		assert.Error(t, err)

	})

}

func TestUpdateCorrectUser(t *testing.T) {

	var log logger.Logger
	{
		log = logger.NewLogfmtLogger(os.Stderr)
		log = logger.NewSyncLogger(log)
		log = logger.With(log,
			"service", "grpcUserService",
			"time:", logger.DefaultTimestampUTC,
			"caller", logger.DefaultCaller,
		)
	}

	ctx := context.Background()

	grpcMock := util.NewgRPClientMock()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(grpcMock)))
	if err != nil {
		level.Error(log).Log(err)
	}
	defer conn.Close()

	grpClient := user.NewgRPClient(log, conn)

	t.Run("Update User Valid Case", func(t *testing.T) {

		grpcMock.On("UpdateUser", mock.Anything, &proto.UpdateUserRequest{
			Name:  "Timo",
			Pass:  "123",
			Age:   19,
			Email: "timoteo@globant.com",
		}).Return(&proto.UpdateUserResponse{
			Status: &proto.Status{
				Message: "User updated succesfully",
				Code:    0,
			},
		}, nil)

		resp, err := grpClient.UpdateUser(ctx, entities.UpdateUserRequest{
			Name:  "Timo",
			Pass:  "123",
			Age:   19,
			Email: "timoteo@globant.com",
		})
		assert.Exactly(t, entities.UpdateUserResponse{
			Status: entities.Status{
				Message: "User updated succesfully",
				Code:    0,
			},
		}, resp)
		assert.Nil(t, err)

	})

}

func TestUpdateUserThatDoesntExists(t *testing.T) {

	var log logger.Logger
	{
		log = logger.NewLogfmtLogger(os.Stderr)
		log = logger.NewSyncLogger(log)
		log = logger.With(log,
			"service", "grpcUserService",
			"time:", logger.DefaultTimestampUTC,
			"caller", logger.DefaultCaller,
		)
	}

	ctx := context.Background()

	grpcMock := util.NewgRPClientMock()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(grpcMock)))
	if err != nil {
		level.Error(log).Log(err)
	}
	defer conn.Close()

	grpClient := user.NewgRPClient(log, conn)

	t.Run("Update User Invalid Case", func(t *testing.T) {

		grpcMock.On("UpdateUser", mock.Anything, &proto.UpdateUserRequest{
			Name:  "Timo",
			Pass:  "123",
			Age:   19,
			Email: "timoteo@globant.com",
		}).Return(&proto.UpdateUserResponse{}, errors.NewUserNotFound())

		resp, err := grpClient.UpdateUser(ctx, entities.UpdateUserRequest{
			Name:  "Timo",
			Pass:  "123",
			Age:   19,
			Email: "timoteo@globant.com",
		})
		assert.Exactly(t, entities.UpdateUserResponse{}, resp)
		assert.Error(t, err)

	})

}
