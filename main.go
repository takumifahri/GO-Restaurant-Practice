package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type MenuItem struct {
	Name string
	Price int
	id string
}


func getMakananSaji(c echo.Context) error {
	// Menu Makanan Saji
	Menu := []MenuItem{
		{
			Name: "Nasi Lemak",
			Price: 10000,
			id: "1",
		},
		{
			Name: "Nasi Goreng",
			Price: 12000,
			id: "2",
		},
		{
			Name: "Nasi Kuning",
			Price: 15000,
			id: "3",
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": Menu,
	}) //data null karena belum ada data
}

func getDrinkMenu(t echo.Context) error {
	// Menu minumuan
	MenusItemDrink := []MenuItem{
		{
			Name: "Esteh Manis",
			Price: 2000,
			id: "1",
		},
		{
			Name: "Esteh Tawar",
			Price: 2000,
			id: "2",
		},
		{
			Name: "Kopi Hitam",
			Price: 26000,
			id: "3",
		},
	}

	return t.JSON(http.StatusOK, map[string]interface{}{
		"data": MenusItemDrink,
	})
}

func GetMenu(c echo.Context) error {
	foodMenu := []MenuItem{}
	if err := getMakananSaji(c); err != nil {
		return err
	}

	drinkMenu := []MenuItem{}
	if err := getDrinkMenu(c); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"food":  foodMenu,
			"drink": drinkMenu,
		},
	})
}


func main() {
	e := echo.New()
	// Untuk pemanggilannya layaknya endpoint pada javascript berupa url nya seperti localhost:3000/menu/food
	e.GET("/menu", GetMenu)
	e.GET("/menu/food", getMakananSaji)
	e.GET("/menu/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":14045")) // untuk deklarasi localhost
}