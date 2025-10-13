package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/ElizCarvalho/FC_PosGolang/16_SQLC/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil) // Inicia uma transação com isolamento padrao (ACID)
	if err != nil {
		return err
	}
	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return fmt.Errorf("rollback error: %w, original error: %w", errRollback, err)
		}
		return err
	}

	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCategory.ID,
			Price:       argsCourse.Price,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	courseArgs := CourseParams{
		ID:          uuid.New().String(),
		Name:        "Curso de Backend C#",
		Description: sql.NullString{String: "Curso de Backend C#", Valid: true},
		Price:       100.00,
	}
	categoryArgs := CategoryParams{
		ID:          uuid.New().String(),
		Name:        "Backend C#",
		Description: sql.NullString{String: "Curso de Backend C#", Valid: true},
	}

	courseDB := NewCourseDB(conn)
	err = courseDB.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)
	if err != nil {
		panic(err)
	}

	queries := db.New(conn)
	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}
	for _, course := range courses {
		fmt.Printf("ID: %s, Name: %s, Description: %s, Price: %f, CategoryName: %s\n", course.ID, course.Name, course.Description.String, course.Price, course.CategoryName)
	}

}
