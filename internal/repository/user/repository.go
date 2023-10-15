package user

import (
	"context"

	"github.com/kelas.work/project-go-restful-api/internal/model"
)

type Repository interface {
	RegisterUser(ctx context.Context, Userdata model.User) (model.User, error)
	CheckRegistered(ctx context.Context, username string) (bool, error)
	GenerateUserHash(ctx context.Context, password string) (hash string, err error)
	VerifyLogin(ctx context.Context, username, password string, userData model.User) (bool, error)
	GetUserData(ctx context.Context, username string) (model.User, error)
	CreateUserSession(ctx context.Context, userID string) (model.UserSession, error)
	CheckSession(ctx context.Context, data model.UserSession) (userID string, err error)
}