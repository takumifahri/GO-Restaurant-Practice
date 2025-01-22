package resto

import "github.com/takumifahri/GO-Restaurant-Practice/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}