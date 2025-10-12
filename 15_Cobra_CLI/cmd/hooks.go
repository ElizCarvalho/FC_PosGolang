package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

// hooksCmd represents the hooks command
var hooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: "Demonstração de hooks do Cobra",
	Long: `Demonstra como usar hooks (ganchos) do Cobra para executar código
em momentos específicos do ciclo de vida dos comandos.

Hooks disponíveis:
- PreRun: Antes da execução do comando
- Run: Execução principal do comando
- PostRun: Após a execução do comando
- PersistentPreRun: Antes da execução (herdado por subcomandos)
- PersistentPostRun: Após a execução (herdado por subcomandos)

Exemplos:
  course-cli hooks --name "Teste"
  course-cli hooks subcommand --value "Exemplo"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🎯 Executando comando principal (Run)")
		fmt.Println("=====================================")

		name, _ := cmd.Flags().GetString("name")
		if name != "" {
			fmt.Printf("👤 Nome: %s\n", name)
		}

		fmt.Println("📋 Este é o hook Run - execução principal do comando")
		fmt.Println("⏱️  Simulando processamento...")
		time.Sleep(1 * time.Second)
		fmt.Println("✅ Processamento concluído!")
	},
}

// hooksSubCmd representa um subcomando para demonstrar herança de hooks
var hooksSubCmd = &cobra.Command{
	Use:   "subcommand",
	Short: "Subcomando que herda hooks persistentes",
	Long:  `Subcomando que demonstra como hooks persistentes são herdados.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🎯 Executando subcomando (Run)")
		fmt.Println("==============================")

		value, _ := cmd.Flags().GetString("value")
		if value != "" {
			fmt.Printf("📝 Valor: %s\n", value)
		}

		fmt.Println("📋 Este é o hook Run do subcomando")
		fmt.Println("🔄 Subcomando processando...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("✅ Subcomando concluído!")
	},
}

func init() {
	rootCmd.AddCommand(hooksCmd)
	hooksCmd.AddCommand(hooksSubCmd)

	// Flags para o comando principal
	hooksCmd.Flags().String("name", "", "Nome para o exemplo")

	// Flags para o subcomando
	hooksSubCmd.Flags().String("value", "", "Valor para o exemplo")

	// ==============================================================================
	// HOOKS DO COMANDO PRINCIPAL
	// ==============================================================================

	// PreRun - Executado ANTES do comando principal
	hooksCmd.PreRun = func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 PreRun: Preparando execução do comando principal...")
		fmt.Println("   📊 Verificando permissões...")
		fmt.Println("   🔍 Validando argumentos...")
		fmt.Println("   ✅ Pré-requisitos verificados!")
		fmt.Println()
	}

	// PostRun - Executado APÓS o comando principal
	hooksCmd.PostRun = func(cmd *cobra.Command, args []string) {
		fmt.Println()
		fmt.Println("🏁 PostRun: Finalizando comando principal...")
		fmt.Println("   📝 Salvando logs...")
		fmt.Println("   🧹 Limpando recursos...")
		fmt.Println("   ✅ Comando principal finalizado!")
	}

	// ==============================================================================
	// HOOKS PERSISTENTES (Herdados por subcomandos)
	// ==============================================================================

	// PersistentPreRun - Executado ANTES de qualquer comando (herdado)
	hooksCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		fmt.Println("🌐 PersistentPreRun: Inicialização global...")
		fmt.Println("   🔧 Configurando ambiente...")
		fmt.Println("   📊 Carregando configurações...")
		fmt.Println("   🚀 Sistema inicializado!")
		fmt.Println()
	}

	// PersistentPostRun - Executado APÓS qualquer comando (herdado)
	hooksCmd.PersistentPostRun = func(cmd *cobra.Command, args []string) {
		fmt.Println()
		fmt.Println("🌐 PersistentPostRun: Finalização global...")
		fmt.Println("   📊 Gerando relatório de execução...")
		fmt.Println("   🔄 Atualizando estatísticas...")
		fmt.Println("   🏁 Sessão finalizada!")
	}

	// ==============================================================================
	// HOOKS DE ERRO
	// ==============================================================================

	// PreRunE - Versão com tratamento de erro do PreRun
	hooksCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		fmt.Println("🔍 PreRunE: Validação com tratamento de erro...")

		// Simular validação que pode falhar
		name, _ := cmd.Flags().GetString("name")
		if name == "erro" {
			return fmt.Errorf("❌ Nome 'erro' não é permitido")
		}

		fmt.Println("   ✅ Validação concluída com sucesso!")
		return nil
	}

	// PostRunE - Versão com tratamento de erro do PostRun
	hooksCmd.PostRunE = func(cmd *cobra.Command, args []string) error {
		fmt.Println("🔍 PostRunE: Finalização com tratamento de erro...")

		// Simular operação que pode falhar
		name, _ := cmd.Flags().GetString("name")
		if name == "falha" {
			return fmt.Errorf("❌ Falha ao finalizar para nome 'falha'")
		}

		fmt.Println("   ✅ Finalização concluída com sucesso!")
		return nil
	}

	// ==============================================================================
	// HOOKS DE VALIDAÇÃO
	// ==============================================================================

	// ValidArgs - Validação de argumentos
	hooksCmd.ValidArgs = []string{"arg1", "arg2", "arg3"}

	// ValidArgsFunction - Validação dinâmica de argumentos
	hooksCmd.ValidArgsFunction = func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// Simular sugestões de argumentos
		suggestions := []string{"sugestao1", "sugestao2", "sugestao3"}
		return suggestions, cobra.ShellCompDirectiveDefault
	}

	// ==============================================================================
	// HOOKS DE LOGGING
	// ==============================================================================

	// Adicionar logging personalizado
	hooksCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		// Hook original
		fmt.Println("🌐 PersistentPreRun: Inicialização global...")
		fmt.Println("   🔧 Configurando ambiente...")
		fmt.Println("   📊 Carregando configurações...")
		fmt.Println("   🚀 Sistema inicializado!")
		fmt.Println()

		// Logging adicional
		log.Printf("Comando executado: %s", cmd.Name())
		log.Printf("Argumentos: %v", args)
		log.Printf("Timestamp: %s", time.Now().Format("2006-01-02 15:04:05"))
	}
}
