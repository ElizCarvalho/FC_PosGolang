/*
Copyright © 2025 ElizCarvalho

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RunEFunc é um tipo personalizado para funções que retornam erro
type RunEFunc func(cmd *cobra.Command, args []string) error

// RunEWithErrorHandling executa uma função RunE com tratamento elegante de erro
func RunEWithErrorHandling(fn RunEFunc) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := fn(cmd, args); err != nil {
			fmt.Printf("❌ Erro: %v\n", err)
			os.Exit(1)
		}
	}
}

// HandlerFunc é um tipo para funções que lidam com a lógica de negócio
type HandlerFunc func(args []string) error

// CreateHandler cria um handler para comandos que não precisam do cmd
func CreateHandler(handler HandlerFunc) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		return handler(args)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "course-cli",
	Short: "CLI para gerenciar categorias",
	Long: `Uma CLI completa para gerenciar categorias.
	
Permite criar, listar e buscar categorias através de comandos simples.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.15_Cobra_CLI.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
