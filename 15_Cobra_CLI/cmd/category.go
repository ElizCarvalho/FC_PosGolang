/*
Copyright © 2025 ElizCarvalho
*/
package cmd

import (
	"database/sql"
	"fmt"

	"github.com/ElizCarvalho/FC_PosGolang/15_Cobra_CLI/internal/service"
	"github.com/spf13/cobra"
)

// categoryService é a instância do serviço (injetada via DI)
var categoryService service.CategoryService

// SetCategoryService define o serviço de categoria (DI)
func SetCategoryService(service service.CategoryService) {
	categoryService = service
}

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Comandos para gerenciar categorias",
	Long: `Comandos para criar, listar, buscar, atualizar e deletar categorias.
	
Exemplos de uso:
  course-cli category create "Programação" "Cursos de programação"
  course-cli category list
  course-cli category get <id>
  course-cli category update <id> "Novo Nome" "Nova Descrição"
  course-cli category delete <id>`,
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [name] [description]",
	Short: "Criar uma nova categoria",
	Long: `Cria uma nova categoria com nome e descrição fornecidos.
	
Exemplo:
  course-cli category create "Programação" "Cursos de programação e desenvolvimento"`,
	Args: cobra.ExactArgs(2),
	Run:  RunEWithErrorHandling(CreateHandler(createCategoryHandler)),
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar todas as categorias",
	Long: `Lista todas as categorias cadastradas no banco de dados.
	
Exemplo:
  course-cli category list`,
	Run: RunEWithErrorHandling(CreateHandler(listCategoriesHandler)),
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Buscar categoria por ID",
	Long: `Busca uma categoria específica pelo ID fornecido.
	
Exemplo:
  course-cli category get <category-id>`,
	Args: cobra.ExactArgs(1),
	Run:  RunEWithErrorHandling(CreateHandler(getCategoryHandler)),
}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [id] [name] [description]",
	Short: "Atualizar uma categoria existente",
	Long: `Atualiza uma categoria existente com novo nome e descrição.
	
Exemplo:
  course-cli category update <id> "Novo Nome" "Nova Descrição"`,
	Args: cobra.ExactArgs(3),
	Run:  RunEWithErrorHandling(CreateHandler(updateCategoryHandler)),
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Deletar uma categoria",
	Long: `Remove uma categoria do banco de dados.
	
Exemplo:
  course-cli category delete <id>`,
	Args: cobra.ExactArgs(1),
	Run:  RunEWithErrorHandling(CreateHandler(deleteCategoryHandler)),
}

func init() {
	rootCmd.AddCommand(categoryCmd)

	// Adicionar subcomandos ao categoryCmd
	categoryCmd.AddCommand(createCmd)
	categoryCmd.AddCommand(listCmd)
	categoryCmd.AddCommand(getCmd)
	categoryCmd.AddCommand(updateCmd)
	categoryCmd.AddCommand(deleteCmd)
}

// InitializeCategoryService inicializa o serviço de categoria com dependências
func InitializeCategoryService(service service.CategoryService) {
	SetCategoryService(service)
}

// Handlers para operações de categoria (lógica de negócio separada dos comandos)

// createCategoryHandler lida com a criação de categorias
func createCategoryHandler(args []string) error {
	if categoryService == nil {
		return fmt.Errorf("serviço de categoria não foi inicializado")
	}

	category, err := categoryService.Create(args[0], args[1])
	if err != nil {
		return fmt.Errorf("erro ao criar categoria: %w", err)
	}

	fmt.Printf("✅ Categoria criada com sucesso!\n")
	fmt.Printf("ID: %s\n", category.ID)
	fmt.Printf("Nome: %s\n", category.Name)
	fmt.Printf("Descrição: %s\n", category.Description)
	return nil
}

// listCategoriesHandler lida com a listagem de categorias
func listCategoriesHandler(args []string) error {
	if categoryService == nil {
		return fmt.Errorf("serviço de categoria não foi inicializado")
	}

	categories, err := categoryService.List()
	if err != nil {
		return fmt.Errorf("erro ao listar categorias: %w", err)
	}

	if len(categories) == 0 {
		fmt.Println("📝 Nenhuma categoria encontrada.")
		return nil
	}

	fmt.Printf("📋 Categorias encontradas (%d):\n\n", len(categories))
	for i, category := range categories {
		fmt.Printf("%d. ID: %s\n", i+1, category.ID)
		fmt.Printf("   Nome: %s\n", category.Name)
		fmt.Printf("   Descrição: %s\n\n", category.Description)
	}
	return nil
}

// getCategoryHandler lida com a busca de categoria por ID
func getCategoryHandler(args []string) error {
	if categoryService == nil {
		return fmt.Errorf("serviço de categoria não foi inicializado")
	}

	category, err := categoryService.GetByID(args[0])
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("❌ Categoria com ID '%s' não encontrada.\n", args[0])
			return nil // Não é um erro fatal, apenas não encontrou
		}
		return fmt.Errorf("erro ao buscar categoria: %w", err)
	}

	fmt.Printf("📋 Categoria encontrada:\n\n")
	fmt.Printf("ID: %s\n", category.ID)
	fmt.Printf("Nome: %s\n", category.Name)
	fmt.Printf("Descrição: %s\n", category.Description)
	return nil
}

// updateCategoryHandler lida com a atualização de categorias
func updateCategoryHandler(args []string) error {
	if categoryService == nil {
		return fmt.Errorf("serviço de categoria não foi inicializado")
	}

	err := categoryService.Update(args[0], args[1], args[2])
	if err != nil {
		return fmt.Errorf("erro ao atualizar categoria: %w", err)
	}

	fmt.Printf("✅ Categoria atualizada com sucesso!\n")
	fmt.Printf("ID: %s\n", args[0])
	fmt.Printf("Novo Nome: %s\n", args[1])
	fmt.Printf("Nova Descrição: %s\n", args[2])
	return nil
}

// deleteCategoryHandler lida com a deleção de categorias
func deleteCategoryHandler(args []string) error {
	if categoryService == nil {
		return fmt.Errorf("serviço de categoria não foi inicializado")
	}

	err := categoryService.Delete(args[0])
	if err != nil {
		return fmt.Errorf("erro ao deletar categoria: %w", err)
	}

	fmt.Printf("✅ Categoria com ID '%s' deletada com sucesso!\n", args[0])
	return nil
}
