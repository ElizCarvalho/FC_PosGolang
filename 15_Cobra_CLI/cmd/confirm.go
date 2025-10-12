package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// confirmCmd represents the confirm command
var confirmCmd = &cobra.Command{
	Use:   "confirm",
	Short: "Exemplo de flags com opções específicas (yes/no, y/n)",
	Long: `Demonstra como criar flags com opções específicas e valores padrão.
	
Exemplos:
  course-cli confirm --yes                    # Usa valor padrão 'y'
  course-cli confirm --yes no                # Força 'no'
  course-cli confirm --yes yes               # Força 'yes'
  course-cli confirm --mode interactive      # Modo interativo
  course-cli confirm --mode batch            # Modo batch
  course-cli confirm --mode interactive --yes no  # Combinação`,
	Run: func(cmd *cobra.Command, args []string) {
		// ==============================================================================
		// FLAGS COM OPÇÕES ESPECÍFICAS E VALORES PADRÃO
		// ==============================================================================

		// Flag com opções yes/no, y/n - valor padrão 'y'
		yesFlag, _ := cmd.Flags().GetString("yes")

		// Flag com opções específicas (interactive/batch)
		mode, _ := cmd.Flags().GetString("mode")

		// Flag com opções de prioridade (low/medium/high)
		priority, _ := cmd.Flags().GetString("priority")

		// Flag com opções de ambiente (dev/staging/prod)
		environment, _ := cmd.Flags().GetString("environment")

		// Flag com opções de formato (json/xml/yaml)
		format, _ := cmd.Flags().GetString("format")

		fmt.Println("🎯 Exemplo de Flags com Opções Específicas")
		fmt.Println("==========================================")

		// ==============================================================================
		// PROCESSAMENTO DA FLAG YES/NO
		// ==============================================================================

		fmt.Printf("📋 Confirmação: ")
		switch strings.ToLower(yesFlag) {
		case "y", "yes":
			fmt.Println("✅ SIM")
		case "n", "no":
			fmt.Println("❌ NÃO")
		default:
			fmt.Printf("❓ Valor inválido: '%s' (use: y/n ou yes/no)\n", yesFlag)
			return
		}

		// ==============================================================================
		// PROCESSAMENTO DAS OUTRAS FLAGS
		// ==============================================================================

		fmt.Printf("🔄 Modo: %s\n", mode)
		fmt.Printf("⚡ Prioridade: %s\n", priority)
		fmt.Printf("🌍 Ambiente: %s\n", environment)
		fmt.Printf("📄 Formato: %s\n", format)

		// ==============================================================================
		// VALIDAÇÃO DE COMBINAÇÕES
		// ==============================================================================

		if mode == "batch" && yesFlag == "n" {
			fmt.Println("⚠️  Aviso: Modo batch com confirmação 'não' pode causar problemas")
		}

		if environment == "prod" && priority == "low" {
			fmt.Println("⚠️  Aviso: Ambiente de produção com prioridade baixa")
		}

		// ==============================================================================
		// SIMULAÇÃO DE AÇÃO BASEADA NAS FLAGS
		// ==============================================================================

		if strings.ToLower(yesFlag) == "y" || strings.ToLower(yesFlag) == "yes" {
			fmt.Println("\n🚀 Executando ação...")

			switch mode {
			case "interactive":
				fmt.Println("   📱 Modo interativo ativado")
			case "batch":
				fmt.Println("   ⚙️  Modo batch ativado")
			}

			fmt.Printf("   📊 Prioridade: %s\n", priority)
			fmt.Printf("   🌍 Ambiente: %s\n", environment)
			fmt.Printf("   📄 Formato de saída: %s\n", format)

			fmt.Println("✅ Ação executada com sucesso!")
		} else {
			fmt.Println("\n⏸️  Ação cancelada pelo usuário")
		}
	},
}

func init() {
	rootCmd.AddCommand(confirmCmd)

	// ==============================================================================
	// FLAGS COM OPÇÕES ESPECÍFICAS E VALORES PADRÃO
	// ==============================================================================

	// Flag yes/no com valor padrão 'y' e validação
	confirmCmd.Flags().String("yes", "y", "Confirmação (y/n ou yes/no) - padrão: y")

	// Flag com opções específicas - modo de operação
	confirmCmd.Flags().String("mode", "interactive", "Modo de operação (interactive/batch)")

	// Flag com opções de prioridade
	confirmCmd.Flags().String("priority", "medium", "Prioridade da operação (low/medium/high)")

	// Flag com opções de ambiente
	confirmCmd.Flags().String("environment", "dev", "Ambiente de execução (dev/staging/prod)")

	// Flag com opções de formato
	confirmCmd.Flags().String("format", "json", "Formato de saída (json/xml/yaml)")

	// ==============================================================================
	// VALIDAÇÃO PERSONALIZADA
	// ==============================================================================

	// Adicionar validação customizada
	confirmCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		// Validar flag yes
		yesFlag, _ := cmd.Flags().GetString("yes")
		validYes := []string{"y", "n", "yes", "no"}
		for _, valid := range validYes {
			if strings.ToLower(yesFlag) == valid {
				return nil
			}
		}
		return fmt.Errorf("valor inválido para --yes: '%s'. Use: y/n ou yes/no", yesFlag)

		// Validar flag mode
		mode, _ := cmd.Flags().GetString("mode")
		validModes := []string{"interactive", "batch"}
		for _, valid := range validModes {
			if mode == valid {
				return nil
			}
		}
		return fmt.Errorf("valor inválido para --mode: '%s'. Use: interactive/batch", mode)

		// Validar flag priority
		priority, _ := cmd.Flags().GetString("priority")
		validPriorities := []string{"low", "medium", "high"}
		for _, valid := range validPriorities {
			if priority == valid {
				return nil
			}
		}
		return fmt.Errorf("valor inválido para --priority: '%s'. Use: low/medium/high", priority)

		// Validar flag environment
		environment, _ := cmd.Flags().GetString("environment")
		validEnvironments := []string{"dev", "staging", "prod"}
		for _, valid := range validEnvironments {
			if environment == valid {
				return nil
			}
		}
		return fmt.Errorf("valor inválido para --environment: '%s'. Use: dev/staging/prod", environment)

		// Validar flag format
		format, _ := cmd.Flags().GetString("format")
		validFormats := []string{"json", "xml", "yaml"}
		for _, valid := range validFormats {
			if format == valid {
				return nil
			}
		}
		return fmt.Errorf("valor inválido para --format: '%s'. Use: json/xml/yaml", format)
	}
}
