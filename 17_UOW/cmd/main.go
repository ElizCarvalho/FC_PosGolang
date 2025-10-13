package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/ElizCarvalho/FC_PosGolang/17_UOW/internal/repository"
	"github.com/ElizCarvalho/FC_PosGolang/17_UOW/internal/usecase"
	"github.com/ElizCarvalho/FC_PosGolang/17_UOW/pkg/uow"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Configuração do banco de dados
	dsn := "root:root@tcp(localhost:3306)/courses?charset=utf8&parseTime=True&loc=Local"

	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco: %v", err)
	}
	defer dbConn.Close()

	// Teste de conexão
	if err := dbConn.Ping(); err != nil {
		log.Fatalf("Erro ao fazer ping no banco: %v", err)
	}

	fmt.Println("✅ Conectado ao banco de dados MySQL")

	ctx := context.Background()

	// Configuração do UOW
	uowInstance := uow.NewUow(ctx, dbConn)

	// Registra os repositórios no UOW
	uowInstance.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCategoryRepository(dbConn)
		repo.Queries = repository.NewCategoryRepository(dbConn).Queries
		return repo
	})

	uowInstance.Register("CourseRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCourseRepository(dbConn)
		repo.Queries = repository.NewCourseRepository(dbConn).Queries
		return repo
	})

	// Configuração dos repositórios
	categoryRepo := repository.NewCategoryRepository(dbConn)
	courseRepo := repository.NewCourseRepository(dbConn)

	// Configuração dos casos de uso
	addCourseUseCase := usecase.NewAddCourseUseCase(courseRepo, categoryRepo)
	addCourseUowUseCase := usecase.NewAddCourseUseCaseUow(uowInstance)

	// Dados de exemplo
	input := usecase.InputUseCase{
		CategoryName:     "Programação",
		CourseName:       "Go Expert",
		CourseCategoryID: 1,
	}

	fmt.Println("\n🚀 Testando adição de curso SEM UOW...")
	err = addCourseUseCase.Execute(ctx, input)
	if err != nil {
		fmt.Printf("❌ Erro sem UOW: %v\n", err)
	} else {
		fmt.Println("✅ Curso adicionado com sucesso (sem UOW)")
	}

	// Reset para o próximo teste
	input.CourseName = "Go Expert UOW"

	fmt.Println("\n🚀 Testando adição de curso COM UOW...")
	err = addCourseUowUseCase.Execute(ctx, input)
	if err != nil {
		fmt.Printf("❌ Erro com UOW: %v\n", err)
	} else {
		fmt.Println("✅ Curso adicionado com sucesso (com UOW)")
	}

	fmt.Println("\n🎉 Demonstração concluída!")
	fmt.Println("\n💡 Para executar os testes:")
	fmt.Println("   make test")
	fmt.Println("\n💡 Para ver a diferença entre UOW e sem UOW:")
	fmt.Println("   make test-uow")
	fmt.Println("   make test-without-uow")
}
