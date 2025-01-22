package database 

import (
	"github.com/takumifahri/GO-Restaurant-Practice/internal/model"
	"github.com/takumifahri/GO-Restaurant-Practice/internal/model/constant"
	// "fmt"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	Menu := []model.MenuItem{
		{
			Name: "Nasi Lemak",
			Price: 10000,
			KodeRahasia: "NasiLemakAduhai",
			Type: constant.MenuTypeFood,
		},
		{
			Name: "Nasi Goreng",
			Price: 12000,
			KodeRahasia: "NasiGorengJesGeses",
			Type: constant.MenuTypeFood,
		},
		{
			Name: "Nasi Kuning",
			Price: 15000,
			KodeRahasia: "NasiKuningPokAye",
			Type: constant.MenuTypeFood,
		},
	}



	MenusItemDrink := []model.MenuItem{
		{
			Name: "Esteh Manis",
			Price: 2000,
			KodeRahasia: "EstehManisMatsoleh",
			Type: constant.MenuTypeDrink,
		},
		{
			Name: "Esteh Tawar",
			Price: 2000,
			KodeRahasia: "TawarAcikiwir",
			Type: constant.MenuTypeDrink,
		},
		{
			Name: "Kopi Hitam",
			Price: 26000,
			KodeRahasia: "HitamLegamMenis",
			Type: constant.MenuTypeDrink,
		},
	}

	// kita validasi jika tidak ada data pada MenuItem di database kita create
	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&Menu)
		db.Create(&MenusItemDrink)
	}

}