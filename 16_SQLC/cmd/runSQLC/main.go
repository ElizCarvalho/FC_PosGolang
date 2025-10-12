package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/ElizCarvalho/FC_PosGolang/16_SQLC/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	queries := db.New(conn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Curso de Backend Golang", Valid: true},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description.String)
	// }

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "d6e860f2-8a81-48bc-bdec-e93159b811e8",
	// 	Name:        "Backend Updated",
	// 	Description: sql.NullString{String: "Curso de Backend Golang Updated", Valid: true},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	err = queries.DeleteCategory(ctx, "d6e860f2-8a81-48bc-bdec-e93159b811e8")
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

}
