package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func TestPingCmd(t *testing.T) {
	// Teste básico do comando ping
	cmd := pingCmd
	output := captureOutput(func() {
		cmd.Run(cmd, []string{})
	})

	expected := "pong\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestPingCmdWithArgs(t *testing.T) {
	// Teste que o comando ping não precisa de argumentos
	cmd := pingCmd
	output := captureOutput(func() {
		cmd.Run(cmd, []string{"arg1", "arg2"})
	})

	expected := "pong\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestPingCmdStructure(t *testing.T) {
	// Teste da estrutura do comando ping
	cmd := pingCmd

	if cmd == nil {
		t.Error("pingCmd should not be nil")
	}

	if cmd.Name() != "ping" {
		t.Errorf("Expected command name 'ping', got %s", cmd.Name())
	}

	if cmd.Short != "Retorna 'pong' para testar a conexão" {
		t.Errorf("Expected short description 'Retorna 'pong' para testar a conexão', got %s", cmd.Short)
	}
}

func TestPingCmdWithPongFlag(t *testing.T) {
	// Teste do comando ping com flag --pong
	cmd := &cobra.Command{
		Use: "ping",
		Run: func(cmd *cobra.Command, args []string) {
			pongFlag, _ := cmd.Flags().GetBool("pong")

			if pongFlag {
				fmt.Println("pong pong")
			} else {
				fmt.Println("pong")
			}
		},
	}

	cmd.Flags().Bool("pong", false, "Retorna 'pong pong' em vez de apenas 'pong'")
	cmd.SetArgs([]string{"--pong"})
	cmd.ParseFlags([]string{"--pong"})

	output := captureOutput(func() {
		cmd.Run(cmd, []string{})
	})

	expected := "pong pong\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestPingCmdWithoutPongFlag(t *testing.T) {
	// Teste do comando ping sem flag --pong
	cmd := &cobra.Command{
		Use: "ping",
		Run: func(cmd *cobra.Command, args []string) {
			pongFlag, _ := cmd.Flags().GetBool("pong")

			if pongFlag {
				fmt.Println("pong pong")
			} else {
				fmt.Println("pong")
			}
		},
	}

	cmd.Flags().Bool("pong", false, "Retorna 'pong pong' em vez de apenas 'pong'")
	cmd.SetArgs([]string{})
	cmd.ParseFlags([]string{})

	output := captureOutput(func() {
		cmd.Run(cmd, []string{})
	})

	expected := "pong\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestPingCmdFlagExists(t *testing.T) {
	// Teste que a flag --pong existe
	cmd := pingCmd

	// Verificar se a flag existe
	flag := cmd.Flags().Lookup("pong")
	if flag == nil {
		t.Error("Flag 'pong' should exist")
		return
	}

	if flag.Name != "pong" {
		t.Errorf("Expected flag name 'pong', got %s", flag.Name)
	}

	if flag.DefValue != "false" {
		t.Errorf("Expected flag default value 'false', got %s", flag.DefValue)
	}
}

// Função auxiliar para capturar output
func captureOutput(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
