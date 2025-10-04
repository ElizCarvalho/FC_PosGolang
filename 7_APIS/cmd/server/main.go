package main

import (
	"net/http"

	"github.com/ElizCarvalho/FC_PosGolang/7_APIS/configs"
	"github.com/ElizCarvalho/FC_PosGolang/7_APIS/infra/database"
	"github.com/ElizCarvalho/FC_PosGolang/7_APIS/internal/entity"
	"github.com/ElizCarvalho/FC_PosGolang/7_APIS/internal/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open(config.DBName), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Product{})
	productDB := database.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":"+config.WebServerPort, nil)
}
