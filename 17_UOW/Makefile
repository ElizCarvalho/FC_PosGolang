# ==============================================================================
# Variáveis
# ==============================================================================
APP_NAME=uow-demo
VERSION?=latest
PORT?=3306
DB_HOST?=localhost
DB_PORT?=3306
DB_USER?=root
DB_PASS?=root
DB_NAME?=courses

# Cores
BLUE=\033[0;34m
GREEN=\033[0;32m
RED=\033[0;31m
YELLOW=\033[0;33m
NC=\033[0m

# ==============================================================================
# Comandos de Desenvolvimento
# ==============================================================================
.PHONY: setup run test lint docs clean

setup: ## Configura o ambiente
	@echo "$(BLUE)🔧 Configurando ambiente...$(NC)"
	@go mod download
	@go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

run: ## Roda a aplicação
	@echo "$(BLUE)🚀 Iniciando aplicação...$(NC)"
	@go run cmd/main.go

test: ## Roda os testes
	@echo "$(BLUE)🧪 Executando testes...$(NC)"
	@go test -v ./...

test-uow: ## Roda apenas os testes de UOW
	@echo "$(BLUE)🧪 Executando testes de UOW...$(NC)"
	@go test -v -run TestAddCourseUow ./internal/usecase/

test-without-uow: ## Roda apenas os testes sem UOW
	@echo "$(BLUE)🧪 Executando testes sem UOW...$(NC)"
	@go test -v -run TestAddCourse ./internal/usecase/

lint: ## Executa o linter
	@echo "$(BLUE)🔍 Verificando código...$(NC)"
	@golangci-lint run

docs: ## Gera documentação SQLC
	@echo "$(BLUE)📚 Gerando código SQLC...$(NC)"
	@sqlc generate

# ==============================================================================
# Comandos de Banco de Dados
# ==============================================================================
db-up: ## Sobe o banco de dados
	@echo "$(BLUE)🐳 Subindo MySQL...$(NC)"
	@docker-compose up -d

db-down: ## Para o banco de dados
	@echo "$(BLUE)🐳 Parando MySQL...$(NC)"
	@docker-compose down

db-reset: ## Reseta o banco de dados
	@echo "$(BLUE)🔄 Resetando banco de dados...$(NC)"
	@docker-compose down -v
	@docker-compose up -d
	@sleep 5
	@make db-migrate

db-migrate: ## Executa as migrações
	@echo "$(BLUE)📊 Executando migrações...$(NC)"
	@docker exec -i mysql mysql -u $(DB_USER) -p$(DB_PASS) $(DB_NAME) < sql/schema.sql

db-shell: ## Acessa o shell do MySQL
	@echo "$(BLUE)🐚 Acessando shell do MySQL...$(NC)"
	@docker exec -it mysql mysql -u $(DB_USER) -p$(DB_PASS) $(DB_NAME)

# ==============================================================================
# Comandos Docker
# ==============================================================================
docker-build: ## Build da imagem Docker
	@echo "$(BLUE)🐳 Construindo imagem Docker...$(NC)"
	@docker build -t $(APP_NAME):$(VERSION) .

docker-run: ## Roda o container
	@echo "$(BLUE)🐳 Iniciando container...$(NC)"
	@docker run -p $(PORT):$(PORT) $(APP_NAME):$(VERSION)

# ==============================================================================
# Comandos de Limpeza
# ==============================================================================
clean: ## Limpa arquivos temporários
	@echo "$(BLUE)🧹 Limpando arquivos temporários...$(NC)"
	@go clean
	@rm -f server

clean-db: ## Limpa dados do banco
	@echo "$(BLUE)🧹 Limpando dados do banco...$(NC)"
	@docker-compose down -v

# ==============================================================================
# Comandos de Desenvolvimento Completo
# ==============================================================================
dev-setup: setup db-up ## Configuração completa para desenvolvimento
	@echo "$(GREEN)✅ Ambiente de desenvolvimento configurado!$(NC)"
	@echo "$(YELLOW)💡 Execute 'make test' para testar$(NC)"

dev-test: db-up test ## Executa testes com banco configurado
	@echo "$(GREEN)✅ Testes executados!$(NC)"

# ==============================================================================
# Ajuda
# ==============================================================================
help: ## Mostra essa ajuda
	@echo "$(BLUE)Comandos disponíveis:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-20s$(NC) %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
