/*
Copyright © 2025 ElizCarvalho
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ElizCarvalho/FC_PosGolang/15_Cobra_CLI/internal/config"
	"github.com/ElizCarvalho/FC_PosGolang/15_Cobra_CLI/internal/database"
	"github.com/spf13/cobra"
)

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
	Run: func(cmd *cobra.Command, args []string) {
		db := config.GetDB()
		defer db.Close()

		categoryRepo := database.NewCategory(db)
		category, err := categoryRepo.Create(args[0], args[1])
		if err != nil {
			log.Fatalf("Erro ao criar categoria: %v", err)
		}

		fmt.Printf("✅ Categoria criada com sucesso!\n")
		fmt.Printf("ID: %s\n", category.ID)
		fmt.Printf("Nome: %s\n", category.Name)
		fmt.Printf("Descrição: %s\n", category.Description)
	},
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar todas as categorias",
	Long: `Lista todas as categorias cadastradas no banco de dados.
	
Exemplo:
  course-cli category list`,
	Run: func(cmd *cobra.Command, args []string) {
		db := config.GetDB()
		defer db.Close()

		categoryRepo := database.NewCategory(db)
		categories, err := categoryRepo.List()
		if err != nil {
			log.Fatalf("Erro ao listar categorias: %v", err)
		}

		if len(categories) == 0 {
			fmt.Println("📝 Nenhuma categoria encontrada.")
			return
		}

		fmt.Printf("📋 Categorias encontradas (%d):\n\n", len(categories))
		for i, category := range categories {
			fmt.Printf("%d. ID: %s\n", i+1, category.ID)
			fmt.Printf("   Nome: %s\n", category.Name)
			fmt.Printf("   Descrição: %s\n\n", category.Description)
		}
	},
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Buscar categoria por ID",
	Long: `Busca uma categoria específica pelo ID fornecido.
	
Exemplo:
  course-cli category get <category-id>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db := config.GetDB()
		defer db.Close()

		categoryRepo := database.NewCategory(db)
		category, err := categoryRepo.GetByID(args[0])
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("❌ Categoria com ID '%s' não encontrada.\n", args[0])
			} else {
				log.Fatalf("Erro ao buscar categoria: %v", err)
			}
			return
		}

		fmt.Printf("📋 Categoria encontrada:\n\n")
		fmt.Printf("ID: %s\n", category.ID)
		fmt.Printf("Nome: %s\n", category.Name)
		fmt.Printf("Descrição: %s\n", category.Description)
	},
}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [id] [name] [description]",
	Short: "Atualizar uma categoria existente",
	Long: `Atualiza uma categoria existente com novo nome e descrição.
	
Exemplo:
  course-cli category update <id> "Novo Nome" "Nova Descrição"`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		db := config.GetDB()
		defer db.Close()

		categoryRepo := database.NewCategory(db)
		err := categoryRepo.Update(args[0], args[1], args[2])
		if err != nil {
			log.Fatalf("Erro ao atualizar categoria: %v", err)
		}

		fmt.Printf("✅ Categoria atualizada com sucesso!\n")
		fmt.Printf("ID: %s\n", args[0])
		fmt.Printf("Novo Nome: %s\n", args[1])
		fmt.Printf("Nova Descrição: %s\n", args[2])
	},
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Deletar uma categoria",
	Long: `Remove uma categoria do banco de dados.
	
Exemplo:
  course-cli category delete <id>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db := config.GetDB()
		defer db.Close()

		categoryRepo := database.NewCategory(db)
		err := categoryRepo.Delete(args[0])
		if err != nil {
			log.Fatalf("Erro ao deletar categoria: %v", err)
		}

		fmt.Printf("✅ Categoria com ID '%s' deletada com sucesso!\n", args[0])
	},
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
