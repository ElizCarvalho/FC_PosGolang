/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Retorna 'pong' para testar a conexão",
	Long: `Comando simples que retorna 'pong' quando chamado.
	
Útil para testar se a CLI está funcionando corretamente.
Use a flag --pong para retornar 'pong pong'.`,
	Run: func(cmd *cobra.Command, args []string) {
		pongFlag, _ := cmd.Flags().GetBool("pong")

		if pongFlag {
			fmt.Println("pong pong")
		} else {
			fmt.Println("pong")
		}
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)

	// Adicionar flag --pong
	pingCmd.Flags().Bool("pong", false, "Retorna 'pong pong' em vez de apenas 'pong'")
}
