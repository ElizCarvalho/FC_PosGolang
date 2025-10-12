package cmd

import (
	"testing"
)

func TestCategoryCmd(t *testing.T) {
	// Teste que o comando category existe e tem subcomandos
	cmd := categoryCmd

	if cmd == nil {
		t.Error("categoryCmd should not be nil")
	}

	if cmd.Name() != "category" {
		t.Errorf("Expected command name 'category', got %s", cmd.Name())
	}
}

func TestCategorySubcommands(t *testing.T) {
	// Teste que todos os subcomandos estão registrados
	cmd := categoryCmd

	expectedSubcommands := []string{"create", "list", "get", "update", "delete"}

	for _, subcmd := range expectedSubcommands {
		found := false
		for _, child := range cmd.Commands() {
			if child.Name() == subcmd {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected subcommand %s to be registered", subcmd)
		}
	}
}

func TestCreateCmdValidation(t *testing.T) {
	// Teste que o comando create tem validação de argumentos
	cmd := createCmd

	if cmd == nil {
		t.Error("createCmd should not be nil")
	}

	// Teste que o comando requer exatamente 2 argumentos
	if cmd.Args == nil {
		t.Error("createCmd should have argument validation")
	}

	// Teste que o comando tem a descrição correta
	if cmd.Short != "Criar uma nova categoria" {
		t.Errorf("Expected short description 'Criar uma nova categoria', got %s", cmd.Short)
	}
}

func TestGetCmdValidation(t *testing.T) {
	// Teste que o comando get tem validação de argumentos
	cmd := getCmd

	if cmd == nil {
		t.Error("getCmd should not be nil")
	}

	// Teste que o comando requer exatamente 1 argumento
	if cmd.Args == nil {
		t.Error("getCmd should have argument validation")
	}

	// Teste que o comando tem a descrição correta
	if cmd.Short != "Buscar categoria por ID" {
		t.Errorf("Expected short description 'Buscar categoria por ID', got %s", cmd.Short)
	}
}

func TestUpdateCmdValidation(t *testing.T) {
	// Teste que o comando update tem validação de argumentos
	cmd := updateCmd

	if cmd == nil {
		t.Error("updateCmd should not be nil")
	}

	// Teste que o comando requer exatamente 3 argumentos
	if cmd.Args == nil {
		t.Error("updateCmd should have argument validation")
	}

	// Teste que o comando tem a descrição correta
	if cmd.Short != "Atualizar uma categoria existente" {
		t.Errorf("Expected short description 'Atualizar uma categoria existente', got %s", cmd.Short)
	}
}

func TestDeleteCmdValidation(t *testing.T) {
	// Teste que o comando delete tem validação de argumentos
	cmd := deleteCmd

	if cmd == nil {
		t.Error("deleteCmd should not be nil")
	}

	// Teste que o comando requer exatamente 1 argumento
	if cmd.Args == nil {
		t.Error("deleteCmd should have argument validation")
	}

	// Teste que o comando tem a descrição correta
	if cmd.Short != "Deletar uma categoria" {
		t.Errorf("Expected short description 'Deletar uma categoria', got %s", cmd.Short)
	}
}

func TestListCmdNoArgs(t *testing.T) {
	// Teste que o comando list não precisa de argumentos
	cmd := listCmd

	if cmd == nil {
		t.Error("listCmd should not be nil")
	}

	// Teste que o comando não tem validação de argumentos (aceita 0)
	if cmd.Args != nil {
		t.Error("listCmd should not require arguments")
	}

	// Teste que o comando tem a descrição correta
	if cmd.Short != "Listar todas as categorias" {
		t.Errorf("Expected short description 'Listar todas as categorias', got %s", cmd.Short)
	}
}

func TestCommandStructure(t *testing.T) {
	// Teste que todos os comandos existem
	commands := []interface{}{
		categoryCmd, createCmd, listCmd, getCmd, updateCmd, deleteCmd,
	}

	for i, cmd := range commands {
		if cmd == nil {
			t.Errorf("Command %d should not be nil", i)
		}
	}
}
