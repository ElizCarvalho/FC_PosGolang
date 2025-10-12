package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Gerenciar configurações",
	Long: `Comandos para gerenciar configurações do sistema.
	
Exemplos:
  course-cli config set --key "database_url" --value "sqlite://db.sqlite"
  course-cli config get --key "database_url"
  course-cli config list --verbose
  course-cli config reset --force`,
}

// configSetCmd representa o comando para definir configuração
var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Definir uma configuração",
	Long:  `Define um valor para uma chave de configuração.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		value, _ := cmd.Flags().GetString("value")
		verbose, _ := cmd.Flags().GetBool("verbose")

		if verbose {
			fmt.Printf("🔧 [VERBOSE] Definindo configuração: %s = %s\n", key, value)
		} else {
			fmt.Printf("✅ Configuração definida: %s = %s\n", key, value)
		}
	},
}

// configGetCmd representa o comando para obter configuração
var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Obter uma configuração",
	Long:  `Obtém o valor de uma chave de configuração.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		verbose, _ := cmd.Flags().GetBool("verbose")

		if verbose {
			fmt.Printf("🔍 [VERBOSE] Buscando configuração para chave: %s\n", key)
		}

		// Simular busca de configuração
		fmt.Printf("📋 Valor da configuração '%s': sqlite://database.db\n", key)
	},
}

// configListCmd representa o comando para listar configurações
var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar todas as configurações",
	Long:  `Lista todas as configurações do sistema.`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")

		if verbose {
			fmt.Println("📋 [VERBOSE] Listando todas as configurações:")
			fmt.Println("  - database_url: sqlite://database.db")
			fmt.Println("  - debug_mode: false")
			fmt.Println("  - log_level: info")
		} else {
			fmt.Println("📋 Configurações:")
			fmt.Println("  database_url: sqlite://database.db")
			fmt.Println("  debug_mode: false")
			fmt.Println("  log_level: info")
		}
	},
}

// configResetCmd representa o comando para resetar configurações
var configResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resetar configurações",
	Long:  `Reseta todas as configurações para os valores padrão.`,
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		verbose, _ := cmd.Flags().GetBool("verbose")

		if !force {
			fmt.Println("❌ Use --force para confirmar o reset das configurações")
			return
		}

		if verbose {
			fmt.Println("🔄 [VERBOSE] Resetando configurações para valores padrão...")
		}

		fmt.Println("✅ Configurações resetadas com sucesso!")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Adicionar subcomandos
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configResetCmd)

	// ==============================================================================
	// FLAGS GLOBAIS (Persistent Flags) - Aplicam-se a TODOS os subcomandos
	// ==============================================================================

	// Flag global --verbose - disponível em TODOS os subcomandos de config
	configCmd.PersistentFlags().Bool("verbose", false, "Modo verboso (aplicado a todos os subcomandos)")

	// ==============================================================================
	// FLAGS LOCAIS (Local Flags) - Aplicam-se APENAS ao comando específico
	// ==============================================================================

	// Flags locais para config set
	configSetCmd.Flags().String("key", "", "Chave da configuração (obrigatório)")
	configSetCmd.Flags().String("value", "", "Valor da configuração (obrigatório)")
	configSetCmd.MarkFlagRequired("key")
	configSetCmd.MarkFlagRequired("value")

	// Flags locais para config get
	configGetCmd.Flags().String("key", "", "Chave da configuração (obrigatório)")
	configGetCmd.MarkFlagRequired("key")

	// Flags locais para config reset
	configResetCmd.Flags().Bool("force", false, "Forçar reset das configurações")
}
