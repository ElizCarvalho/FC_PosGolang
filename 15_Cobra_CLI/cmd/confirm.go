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

// validateFlag valida se um valor está em uma lista de valores válidos
func validateFlag(value, flagName string, validValues []string, caseSensitive bool) error {
	for _, valid := range validValues {
		if caseSensitive && value == valid {
			return nil
		}
		if !caseSensitive && strings.EqualFold(value, valid) {
			return nil
		}
	}
	return fmt.Errorf("valor inválido para --%s: '%s'. Use: %s", flagName, value, strings.Join(validValues, "/"))
}

// validateConfirmFlags valida todas as flags do comando confirm
func validateConfirmFlags(cmd *cobra.Command) error {
	yesFlag, _ := cmd.Flags().GetString("yes")
	if err := validateFlag(yesFlag, "yes", []string{"y", "n", "yes", "no"}, false); err != nil {
		return err
	}

	mode, _ := cmd.Flags().GetString("mode")
	if err := validateFlag(mode, "mode", []string{"interactive", "batch"}, true); err != nil {
		return err
	}

	priority, _ := cmd.Flags().GetString("priority")
	if err := validateFlag(priority, "priority", []string{"low", "medium", "high"}, true); err != nil {
		return err
	}

	environment, _ := cmd.Flags().GetString("environment")
	if err := validateFlag(environment, "environment", []string{"dev", "staging", "prod"}, true); err != nil {
		return err
	}

	format, _ := cmd.Flags().GetString("format")
	if err := validateFlag(format, "format", []string{"json", "xml", "yaml"}, true); err != nil {
		return err
	}

	return nil
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
		return validateConfirmFlags(cmd)
	}
}
