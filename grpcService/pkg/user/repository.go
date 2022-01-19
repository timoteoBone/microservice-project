package user

import (
	"context"
	"database/sql"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	entities "github.com/timoteoBone/microservice-project/grpcService/pkg/entities"
	"github.com/timoteoBone/microservice-project/grpcService/pkg/utils"
)

type sqlRepo struct {
	DB     *sql.DB
	Logger log.Logger
}

func NewSQL(db *sql.DB, log log.Logger) *sqlRepo {
	return &sqlRepo{db, log}
}

func (repo *sqlRepo) CreateUser(ctx context.Context, user entities.User, newId string) (string, error) {

	repo.Logger.Log(repo.Logger, "Repository method", "Create user")

	stmt, err := repo.DB.PrepareContext(ctx, utils.CreateUserQuery)
	if err != nil {
		level.Error(repo.Logger).Log(err)
		return "", err
	}

	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, user.Name, newId, user.Pass, user.Age, user.Email)
	if err != nil {
		level.Error(repo.Logger).Log(err)
		return "", err
	}

	repo.Logger.Log(repo.Logger, res, "rows affected")

	return newId, nil
}

func (repo *sqlRepo) GetUser(ctx context.Context, userId string) (entities.User, error) {
	repo.Logger.Log(repo.Logger, "Repository method", "Get user")

	user := entities.User{}
	stmt, err := repo.DB.PrepareContext(ctx, utils.GetUserQuery)
	if err != nil {
		level.Error(repo.Logger).Log(err)
		return entities.User{}, err
	}

	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, userId).Scan(&user.Name, &user.Age, &user.Email)
	if err != nil {
		level.Error(repo.Logger).Log(err)
		return entities.User{}, err
	}

	return user, nil
}

func (repo *sqlRepo) DeleteUser(ctx context.Context, userId string) error {
	repo.Logger.Log(repo.Logger, "Repository method", "delete user")

	stmt, err := repo.DB.PrepareContext(ctx, utils.DeleteUserQuery)
	if err != nil {
		level.Error(repo.Logger).Log(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userId)
	if err != nil {
		level.Error(repo.Logger).Log(err)
		return err
	}

	return nil

}
