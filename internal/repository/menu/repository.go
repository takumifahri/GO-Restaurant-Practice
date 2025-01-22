package menu

import (
	"github.com/takumifahri/GO-Restaurant-Practice/internal/model"
)

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}