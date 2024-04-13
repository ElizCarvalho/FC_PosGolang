package main

import "gorm.io/gorm"
import "gorm.io/driver/mysql"

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
	db.AutoMigrate(&Flight{})

	//criar um novo voo
	db.Create(&Flight{Name: "BSB-GIG", Price: 350})

	flighs := []Flight{
		{Name: "ABC-123", Price: 100},
		{Name: "DEF-456", Price: 200},
		{Name: "GHI-789", Price: 300},
	}
	db.Create(&flighs)
}
