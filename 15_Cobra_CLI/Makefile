# ==============================================================================
# Variáveis
# ==============================================================================
APP_NAME=course-cli
VERSION?=latest
DB_PATH?=./database.db

# Cores
BLUE=\033[0;34m
GREEN=\033[0;32m
RED=\033[0;31m
YELLOW=\033[0;33m
NC=\033[0m

# ==============================================================================
# Comandos de Desenvolvimento
# ==============================================================================
.PHONY: setup run build test clean help

setup: ## Configura o ambiente
	@echo "$(BLUE)🔧 Configurando ambiente...$(NC)"
	@go mod download
	@go mod tidy
	@go install github.com/spf13/cobra-cli@latest
	@echo "$(GREEN)✅ Ambiente configurado!$(NC)"

build: ## Compila a aplicação
	@echo "$(BLUE)🔨 Compilando aplicação...$(NC)"
	@go build -o $(APP_NAME) .
	@echo "$(GREEN)✅ Aplicação compilada como '$(APP_NAME)'!$(NC)"

run: ## Roda a aplicação
	@echo "$(BLUE)🚀 Executando aplicação...$(NC)"
	@go run main.go

test: ## Roda os testes
	@echo "$(BLUE)🧪 Executando testes...$(NC)"
	@go test -v ./...

test-cmd: ## Roda apenas os testes de comandos
	@echo "$(BLUE)🧪 Executando testes de comandos...$(NC)"
	@go test -v ./cmd/...

test-coverage: ## Roda testes com cobertura
	@echo "$(BLUE)🧪 Executando testes com cobertura...$(NC)"
	@go test -v -cover ./...

test-race: ## Roda testes com detecção de race conditions
	@echo "$(BLUE)🧪 Executando testes com detecção de race conditions...$(NC)"
	@go test -v -race ./...

clean: ## Remove arquivos gerados
	@echo "$(BLUE)🧹 Limpando arquivos...$(NC)"
	@rm -f $(APP_NAME)
	@rm -f database.db
	@go clean
	@echo "$(GREEN)✅ Limpeza concluída!$(NC)"

# ==============================================================================
# Comandos de Banco de Dados
# ==============================================================================
.PHONY: db-init db-reset

db-init: ## Inicializa o banco de dados
	@echo "$(BLUE)🗄️ Inicializando banco de dados...$(NC)"
	@DB_PATH=$(DB_PATH) go run main.go category list
	@echo "$(GREEN)✅ Banco de dados inicializado!$(NC)"

db-reset: ## Remove e recria o banco de dados
	@echo "$(BLUE)🗄️ Resetando banco de dados...$(NC)"
	@rm -f $(DB_PATH)
	@$(MAKE) db-init
	@echo "$(GREEN)✅ Banco de dados resetado!$(NC)"

# ==============================================================================
# Comandos de Exemplo
# ==============================================================================
.PHONY: demo-categories demo-ping demo-full

demo-categories: ## Demonstra comandos de categorias
	@echo "$(BLUE)📚 Demonstração - Categorias$(NC)"
	@echo "$(YELLOW)Criando categorias...$(NC)"
	@DB_PATH=$(DB_PATH) go run main.go category create "Programação" "Cursos de programação e desenvolvimento"
	@DB_PATH=$(DB_PATH) go run main.go category create "Design" "Cursos de design e UX/UI"
	@echo "$(YELLOW)Listando categorias...$(NC)"
	@DB_PATH=$(DB_PATH) go run main.go category list

demo-ping: ## Demonstra comando ping
	@echo "$(BLUE)🏓 Demonstração - Ping$(NC)"
	@go run main.go ping

demo-full: ## Demonstração completa
	@echo "$(BLUE)🎯 Demonstração Completa$(NC)"
	@$(MAKE) demo-ping
	@echo ""
	@$(MAKE) demo-categories

# ==============================================================================
# Comandos de Desenvolvimento Avançado
# ==============================================================================
.PHONY: dev-install dev-run dev-clean cobra-help

dev-install: ## Instala dependências de desenvolvimento
	@echo "$(BLUE)📦 Instalando dependências de desenvolvimento...$(NC)"
	@go install github.com/spf13/cobra-cli@latest
	@echo "$(GREEN)✅ Dependências instaladas!$(NC)"

dev-run: build ## Roda a aplicação compilada
	@echo "$(BLUE)🚀 Executando aplicação compilada...$(NC)"
	@./$(APP_NAME)

dev-clean: clean ## Limpeza completa incluindo dependências
	@echo "$(BLUE)🧹 Limpeza completa...$(NC)"
	@go clean -modcache
	@echo "$(GREEN)✅ Limpeza completa concluída!$(NC)"

cobra-help: ## Mostra exemplos de uso do cobra-cli
	@echo "$(BLUE)📚 Exemplos de uso do cobra-cli:$(NC)"
	@echo ""
	@echo "$(YELLOW)Adicionar novo comando:$(NC)"
	@echo "  cobra-cli add <nome-do-comando>"
	@echo ""
	@echo "$(YELLOW)Exemplos:$(NC)"
	@echo "  cobra-cli add student     # Adiciona comando 'student'"
	@echo "  cobra-cli add teacher     # Adiciona comando 'teacher'"
	@echo "  cobra-cli add report      # Adiciona comando 'report'"
	@echo ""
	@echo "$(YELLOW)Ver ajuda completa:$(NC)"
	@echo "  cobra-cli --help"

# ==============================================================================
# Ajuda
# ==============================================================================
help: ## Mostra essa ajuda
	@echo "$(BLUE)Comandos disponíveis:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-20s$(NC) %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
