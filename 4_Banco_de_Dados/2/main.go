package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Flight struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Flight{}) //em ambiente de desenvolvimento

	//criar um novo voo
	//db.Create(&Flight{Name: "BSB-GIG", Price: 350})

	//criar em batch
	/*flighs := []Flight{
		{Name: "ABC-123", Price: 100},
		{Name: "DEF-456", Price: 200},
		{Name: "GHI-789", Price: 300},
	}
	db.Create(&flighs)*/

	//selecionar um voo
	var flight Flight
	fmt.Println("Selecionando voo com ID 2")
	db.Debug().First(&flight, 2)
	db.First(&flight, "id = ?", 2)
	fmt.Println(flight)

	//selecionar todos os voos
	var flights []Flight
	fmt.Println("Selecionando todos os voos")
	db.Debug().Find(&flights)
	fmt.Println(flights)
}
