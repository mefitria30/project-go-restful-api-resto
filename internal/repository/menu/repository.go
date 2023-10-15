package menu

import (
	"context"

	"github.com/kelas.work/project-go-restful-api/internal/model"
)

type Repository interface {
	GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error)
	GetMenu(ctx context.Context, orderCode string) (model.MenuItem, error)
}