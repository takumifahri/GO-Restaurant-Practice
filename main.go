package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenuItem struct {
	Name string
	Price int
	IdMakanan string
}


func getMakananSaji(c echo.Context) error {
	// Menu Makanan Saji
	Menu := []MenuItem{
		{
			Name: "Nasi Lemak",
			Price: 10000,
			IdMakanan: "1",
		},
		{
			Name: "Nasi Goreng",
			Price: 12000,
			IdMakanan: "2",
		},
		{
			Name: "Nasi Kuning",
			Price: 15000,
			IdMakanan: "3",
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
			IdMakanan: "1",
		},
		{
			Name: "Esteh Tawar",
			Price: 2000,
			IdMakanan: "2",
		},
		{
			Name: "Kopi Hitam",
			Price: 26000,
			IdMakanan: "3",
		},
	}

	return t.JSON(http.StatusOK, map[string]interface{}{
		"data": MenusItemDrink,
	})
}


func main() {
	e := echo.New()
	// Untuk pemanggilannya layaknya endpoint pada javascript berupa url nya seperti localhost:3000/menu/food
	e.GET("/menu/food", getMakananSaji)
	e.GET("/menu/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":14045")) // untuk deklarasi localhost
}