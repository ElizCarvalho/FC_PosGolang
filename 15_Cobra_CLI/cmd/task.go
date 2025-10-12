package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// taskCmd represents the task command (subcomando de project)
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Gerenciar tarefas do projeto",
	Long: `Comandos para gerenciar tarefas dentro de um projeto.
	
Exemplos:
  course-cli project task add "Implementar login"
  course-cli project task list
  course-cli project task complete <id>`,
}

// taskAddCmd representa o comando para adicionar tarefa
var taskAddCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Adicionar nova tarefa",
	Long:  `Adiciona uma nova tarefa ao projeto atual.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("✅ Tarefa adicionada: %s\n", args[0])
	},
}

// taskListCmd representa o comando para listar tarefas
var taskListCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar todas as tarefas",
	Long:  `Lista todas as tarefas do projeto atual.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📋 Tarefas do projeto:")
		fmt.Println("1. Implementar autenticação")
		fmt.Println("2. Criar interface de usuário")
		fmt.Println("3. Configurar banco de dados")
	},
}

// taskCompleteCmd representa o comando para completar tarefa
var taskCompleteCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "Marcar tarefa como completa",
	Long:  `Marca uma tarefa como completa pelo ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("✅ Tarefa %s marcada como completa!\n", args[0])
	},
}

func init() {
	// Adicionar task como subcomando de project
	projectCmd.AddCommand(taskCmd)

	// Adicionar subcomandos de task
	taskCmd.AddCommand(taskAddCmd)
	taskCmd.AddCommand(taskListCmd)
	taskCmd.AddCommand(taskCompleteCmd)
}
