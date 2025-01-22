package main

import (
	// "fmt"
	// "gorm.io/gorm"
	// "gorm.io/driver/postgres"
	"net/http"
	"github.com/labstack/echo/v4"
)

const(
	AlamatDB = "host=localhost user=postgres password=P0sTgreSqL@! port=5433 dbname=restaurant_db sslmode=disable"
)



type mahasiswa struct {
	Nama string
	NIM string
	Prodi string
	Angkatan int
}



func getMahasiswa(c echo.Context) error {
	// Data Mahasiswa 
	datas := []mahasiswa{
		{
			Nama: "Rizky",
			NIM: "123456",
			Prodi: "Teknik Informatika",
			Angkatan: 2019,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": datas,
	})
}

func main() {
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlconn)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// initDB()

	e := echo.New()
	// Untuk pemanggilannya layaknya endpoint pada javascript berupa url nya seperti localhost:3000/menu/food
	

	e.GET("/Mahasiswa", getMahasiswa)
	e.Logger.Fatal(e.Start(":14045")) // untuk deklarasi localhost
}