package resto

import (
	"context"

	"github.com/kelas.work/project-go-restful-api/internal/model"
)

type Usecase interface {
	GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error)
	Order(ctx context.Context, request model.OrderMenuRequest) (model.Order, error)
	GetOrderInfo(ctx context.Context, request model.GetOrderInfoRequest) (model.Order, error)
	RegisterUser(ctx context.Context, request model.RegisterRequest) (model.User, error)
	Login(ctx context.Context, request model.LoginRequest) (model.UserSession, error)
	CheckSession(ctx context.Context, data model.UserSession) (userID string, err error)
}