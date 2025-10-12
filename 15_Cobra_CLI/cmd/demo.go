package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// demoCmd represents the demo command
var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "Demonstração de manipulação de flags",
	Long: `Comando de demonstração que mostra como trabalhar com diferentes tipos de flags.
	
Exemplos:
  course-cli demo --name "João" --age 25 --active --tags "go,cli,cobra"
  course-cli demo --config-file config.json --timeout 30s --verbose
  course-cli demo --help`,
	Run: func(cmd *cobra.Command, args []string) {
		// ==============================================================================
		// MANIPULAÇÃO DE FLAGS - DIFERENTES TIPOS
		// ==============================================================================

		// 1. STRING - Flag de texto
		name, _ := cmd.Flags().GetString("name")
		configFile, _ := cmd.Flags().GetString("config-file")

		// 2. INT - Flag numérica inteira
		age, _ := cmd.Flags().GetInt("age")
		port, _ := cmd.Flags().GetInt("port")

		// 3. BOOL - Flag booleana
		active, _ := cmd.Flags().GetBool("active")
		verbose, _ := cmd.Flags().GetBool("verbose")
		debug, _ := cmd.Flags().GetBool("debug")

		// 4. FLOAT64 - Flag numérica decimal
		price, _ := cmd.Flags().GetFloat64("price")

		// 5. DURATION - Flag de duração de tempo
		timeout, _ := cmd.Flags().GetDuration("timeout")

		// 6. STRING SLICE - Flag que aceita múltiplos valores
		tags, _ := cmd.Flags().GetStringSlice("tags")

		// 7. INT SLICE - Flag que aceita múltiplos números
		ports, _ := cmd.Flags().GetIntSlice("ports")

		// 8. BOOL SLICE - Flag que aceita múltiplos booleanos
		features, _ := cmd.Flags().GetBoolSlice("features")

		// ==============================================================================
		// EXIBIÇÃO DOS VALORES DAS FLAGS
		// ==============================================================================

		fmt.Println("🎯 Demonstração de Manipulação de Flags")
		fmt.Println("=====================================")

		// String flags
		if name != "" {
			fmt.Printf("👤 Nome: %s\n", name)
		}
		if configFile != "" {
			fmt.Printf("📁 Arquivo de Config: %s\n", configFile)
		}

		// Int flags
		if age > 0 {
			fmt.Printf("🎂 Idade: %d anos\n", age)
		}
		if port > 0 {
			fmt.Printf("🔌 Porta: %d\n", port)
		}

		// Bool flags
		if active {
			fmt.Println("✅ Status: Ativo")
		} else {
			fmt.Println("❌ Status: Inativo")
		}

		if verbose {
			fmt.Println("🔍 Modo Verboso: Ativado")
		}

		if debug {
			fmt.Println("🐛 Modo Debug: Ativado")
		}

		// Float flags
		if price > 0 {
			fmt.Printf("💰 Preço: R$ %.2f\n", price)
		}

		// Duration flags
		if timeout > 0 {
			fmt.Printf("⏱️  Timeout: %v\n", timeout)
		}

		// String slice flags
		if len(tags) > 0 {
			fmt.Printf("🏷️  Tags: %s\n", strings.Join(tags, ", "))
		}

		// Int slice flags
		if len(ports) > 0 {
			fmt.Printf("🔌 Portas: %v\n", ports)
		}

		// Bool slice flags
		if len(features) > 0 {
			fmt.Printf("⚙️  Features: %v\n", features)
		}

		// ==============================================================================
		// VERIFICAÇÃO DE FLAGS OBRIGATÓRIAS
		// ==============================================================================

		requiredFlags := []string{"name"}
		missingFlags := []string{}

		for _, flag := range requiredFlags {
			if !cmd.Flags().Changed(flag) {
				missingFlags = append(missingFlags, flag)
			}
		}

		if len(missingFlags) > 0 {
			fmt.Printf("\n❌ Flags obrigatórias não fornecidas: %s\n", strings.Join(missingFlags, ", "))
			fmt.Println("Use --help para ver todas as opções disponíveis.")
			return
		}

		// ==============================================================================
		// VALIDAÇÃO DE VALORES
		// ==============================================================================

		if age < 0 || age > 150 {
			fmt.Println("❌ Idade deve estar entre 0 e 150 anos")
			return
		}

		if port < 1 || port > 65535 {
			fmt.Println("❌ Porta deve estar entre 1 e 65535")
			return
		}

		if price < 0 {
			fmt.Println("❌ Preço não pode ser negativo")
			return
		}

		fmt.Println("\n✅ Todas as flags foram processadas com sucesso!")
	},
}

func init() {
	rootCmd.AddCommand(demoCmd)

	// ==============================================================================
	// DEFINIÇÃO DE FLAGS - DIFERENTES TIPOS
	// ==============================================================================

	// 1. STRING - Flag de texto
	demoCmd.Flags().String("name", "", "Nome da pessoa (obrigatório)")
	demoCmd.Flags().String("config-file", "config.json", "Arquivo de configuração")

	// 2. INT - Flag numérica inteira
	demoCmd.Flags().Int("age", 0, "Idade da pessoa")
	demoCmd.Flags().Int("port", 8080, "Porta do servidor")

	// 3. BOOL - Flag booleana
	demoCmd.Flags().Bool("active", false, "Status ativo/inativo")
	demoCmd.Flags().Bool("verbose", false, "Modo verboso")
	demoCmd.Flags().Bool("debug", false, "Modo debug")

	// 4. FLOAT64 - Flag numérica decimal
	demoCmd.Flags().Float64("price", 0.0, "Preço do produto")

	// 5. DURATION - Flag de duração de tempo
	demoCmd.Flags().Duration("timeout", 30*time.Second, "Timeout para operações")

	// 6. STRING SLICE - Flag que aceita múltiplos valores
	demoCmd.Flags().StringSlice("tags", []string{}, "Tags para categorização")

	// 7. INT SLICE - Flag que aceita múltiplos números
	demoCmd.Flags().IntSlice("ports", []int{}, "Lista de portas")

	// 8. BOOL SLICE - Flag que aceita múltiplos booleanos
	demoCmd.Flags().BoolSlice("features", []bool{}, "Lista de features ativadas")

	// ==============================================================================
	// FLAGS OBRIGATÓRIAS
	// ==============================================================================

	demoCmd.MarkFlagRequired("name")

	// ==============================================================================
	// FLAGS COM SHORTCUTS (versões curtas)
	// ==============================================================================

	demoCmd.Flags().StringP("output", "o", "", "Arquivo de saída")
	demoCmd.Flags().BoolP("force", "f", false, "Forçar operação")
	demoCmd.Flags().IntP("count", "c", 1, "Número de repetições")
}
