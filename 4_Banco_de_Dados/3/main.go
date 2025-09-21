package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SerialNumber struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	FlightID int
}

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Flight struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int          //foreign key
	Category     Category     //relacionamento com a categoria
	SerialNumber SerialNumber //relacionamento com o serial number
	gorm.Model                //ID, CreatedAt, UpdatedAt, DeletedAt gerenciados pelo GORM
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//configurar o timezone
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&Flight{}, &Category{}, &SerialNumber{})

	//criar categoria Comercial, Economico, Executivo
	/*db.Create(&Category{Name: "Comercial"})
	db.Create(&Category{Name: "Economico"})
	db.Create(&Category{Name: "Executivo"})*/

	//criar voo Comercial, Economico, Executivo
	/*db.Create(&Flight{Name: "GIG-BSB", Price: 100, CategoryID: 1})
	db.Create(&Flight{Name: "SDU-CON", Price: 200, CategoryID: 2})
	db.Create(&Flight{Name: "BRZ-PTG", Price: 300, CategoryID: 3})*/

	//criar serial number
	db.Create(&SerialNumber{Name: "123456", FlightID: 1})
	db.Create(&SerialNumber{Name: "789012", FlightID: 2})
	db.Create(&SerialNumber{Name: "345678", FlightID: 3})

	//buscar todos os voos
	var flights []Flight
	db.Debug().Preload("Category").Preload("SerialNumber").Find(&flights) //preload é para buscar a categoria relacionada ao voo
	for _, flight := range flights {
		fmt.Println(flight.Name, flight.Price, flight.Category.Name, flight.SerialNumber.Name)
	}

}
