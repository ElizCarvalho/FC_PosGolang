package cmd

import (
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Gerenciar projetos",
	Long: `Comandos para gerenciar projetos e suas tarefas.
	
Exemplos:
  course-cli project create "Meu Projeto"
  course-cli project list
  course-cli project task add "Nova Tarefa"
  course-cli project task list`,
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
