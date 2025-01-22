package menu

import (
	"github.com/takumifahri/GO-Restaurant-Practice/internal/model"
	// "github.com/takumifahri/GO-Restaurant-Practice/internal/model/constant"
	"gorm.io/gorm"
)

type MenuRepo struct {
	db *gorm.DB
}

func GetRepo(db *gorm.DB) Repository {
	return &MenuRepo{
		db: db,
	}
}


func (m *MenuRepo) GetMenu(menuType string) ([]model.MenuItem, error) {
	var MenuData []model.MenuItem 
	if err := m.db.Where(model.MenuItem{Type: model.MenuType(menuType)}).Find(&MenuData).Error; err != nil {
		return nil, err
	}

	return MenuData, nil
}
