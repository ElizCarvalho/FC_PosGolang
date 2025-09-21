package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Labels struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	Flights []Flight `gorm:"many2many:flights_labels;"` //relacionamento many to many
}

type SerialNumber struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	FlightID int
}

type Category struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	Flights []Flight //relacionamento com o voo has many
}

type Flight struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int          //foreign key
	Category     Category     //relacionamento com a categoria has one
	SerialNumber SerialNumber //relacionamento com o serial number has one
	Labels       []Labels     `gorm:"many2many:flights_labels;"` //relacionamento many to many
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

	//criar label
	label1 := Labels{Name: "Promocao"}
	label2 := Labels{Name: "Black Friday"}
	label3 := Labels{Name: "Cyber Monday"}
	db.Create(&label1)
	db.Create(&label2)
	db.Create(&label3)

	//criar voo com label
	db.Create(&Flight{Name: "GIG-BSB", Price: 100, CategoryID: 1, Labels: []Labels{label1}})
	db.Create(&Flight{Name: "SDU-CON", Price: 200, CategoryID: 2, Labels: []Labels{label2}})
	db.Create(&Flight{Name: "BRZ-PTG", Price: 300, CategoryID: 3, Labels: []Labels{label3}})

	//buscar todos os voos
	var flights []Flight
	db.Debug().Preload("Category").Preload("SerialNumber").Find(&flights) //preload é para buscar a categoria relacionada ao voo
	for _, flight := range flights {
		fmt.Println(flight.Name, flight.Price, flight.Category.Name, flight.SerialNumber.Name)
	}

	//buscar todas as categorias com os voos relacionados
	var categories []Category
	err = db.Model(&Category{}).Preload("Flights").Preload("Flights.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, "->")
		for _, flight := range category.Flights {
			fmt.Println("- ", flight.Name, flight.SerialNumber.Name)
		}
	}
	//buscar todas as labels com os voos relacionados
	var labels []Labels
	err = db.Model(&Labels{}).Preload("Flights").Find(&labels).Error
	if err != nil {
		panic(err)
	}
	for _, label := range labels {
		fmt.Println(label.Name, "->")
		for _, flight := range label.Flights {
			fmt.Println("- ", flight.Name)
		}
	}

	//lock pessimista com select for update em labels
	//usado quando tem muita concorrencia
	//lock no banco de dados, bloqueia o banco de dados para outras transações
	//as transacoes que estao na fila bloquedas esperam o unlock para continuar
	//é mais seguro, mas tem mais overhead ... usado para ACID
	//select * from labels where id = 1 for update
	//GORM sempre vai trabalhar com transaction por padrao
	tx := db.Begin()
	var l Labels
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).Find(&l, 1).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	l.Name = "Promocao 3"
	err = tx.Debug().Save(&l).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
	fmt.Println("Label atualizada com sucesso")

}
