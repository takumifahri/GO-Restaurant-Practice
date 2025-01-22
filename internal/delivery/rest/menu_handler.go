package rest

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func getMakananSaji(c echo.Context) error {
	// Kita buat endpointnya 
	db, err := gorm.Open(postgres.Open(AlamatDB))
	if err != nil {
		panic(err)
	}

	var MenuData []model.MenuItem

	db.Where(MenuItem{Type: menuTypeFood}).Find(&MenuData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": MenuData,
	}) //data null karena belum ada data
}

func getDrinkMenu(t echo.Context) error {
	// Kita buat endpointnya 
	db, err := gorm.Open(postgres.Open(AlamatDB))
	if err != nil {
		panic(err)
	}

	var MenuData []MenuItem

	db.Where(MenuItem{Type: menuTypeDrink}).Find(&MenuData)
	return t.JSON(http.StatusOK, map[string]interface{}{
		"data": MenuData,
	})
}

func GetMenu(c echo.Context) error {
	MenuType := c.FormValue("menu_type")
	
	db, err := gorm.Open(postgres.Open(AlamatDB))
	if err != nil {
		panic(err)
	}
	var MenuData []MenuItem

	db.Where(MenuItem{Type: menuType(MenuType)}).Find(&MenuData)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": MenuData,
	
	})
}