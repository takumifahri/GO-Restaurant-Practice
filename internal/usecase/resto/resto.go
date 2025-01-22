package resto

import (
	"github.com/takumifahri/GO-Restaurant-Practice/internal/model"
	"github.com/takumifahri/GO-Restaurant-Practice/internal/repository/menu"
)

type RestoUsecase struct {
	menuRepo menu.Repository
}

func GetUsecase(menuRepo menu.Repository) Usecase {
	return &RestoUsecase{
		menuRepo: menuRepo,
	}
}

func (r *RestoUsecase) GetMenu(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenu(menuType)
}