package main

import (
	"net/http"

	"github.com/ElizCarvalho/FC_PosGolang/7_APIS/configs"
	"github.com/ElizCarvalho/FC_PosGolang/7_APIS/infra/database"
	"github.com/ElizCarvalho/FC_PosGolang/7_APIS/internal/entity"
	"github.com/ElizCarvalho/FC_PosGolang/7_APIS/internal/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	http.ListenAndServe(":"+config.WebServerPort, r)
}
