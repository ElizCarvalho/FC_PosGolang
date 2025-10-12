# ==============================================================================
# Variáveis
# ==============================================================================
DB_NAME=courses
DB_USER=root
DB_PASSWORD=root
DB_HOST=localhost
DB_PORT=3306
MIGRATIONS_PATH=sql/migrations
DATABASE_URL=mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

# Cores
BLUE=\033[0;34m
GREEN=\033[0;32m
RED=\033[0;31m
YELLOW=\033[0;33m
NC=\033[0m

# ==============================================================================
# Comandos de Desenvolvimento
# ==============================================================================
.PHONY: setup run test lint docs help

setup: ## Configura o ambiente
	@echo "$(BLUE)🔧 Configurando ambiente...$(NC)"
	@go mod download
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@echo "$(GREEN)✅ Ambiente configurado!$(NC)"

run: ## Roda a aplicação
	@echo "$(BLUE)🚀 Iniciando aplicação...$(NC)"
	@echo "$(YELLOW)⚠️  Certifique-se de que o MySQL está rodando: make docker-up$(NC)"

test: ## Roda os testes
	@echo "$(BLUE)🧪 Executando testes...$(NC)"
	@go test -v ./...

lint: ## Executa o linter
	@echo "$(BLUE)🔍 Verificando código...$(NC)"
	@go vet ./...
	@go fmt ./...

docs: ## Gera documentação
	@echo "$(BLUE)📚 Gerando documentação...$(NC)"
	@echo "$(YELLOW)📖 Documentação disponível no README.md$(NC)"

# ==============================================================================
# Comandos Docker
# ==============================================================================
docker-up: ## Sobe o MySQL via Docker
	@echo "$(BLUE)🐳 Iniciando MySQL...$(NC)"
	@docker compose up -d
	@echo "$(GREEN)✅ MySQL rodando em localhost:3306$(NC)"

docker-down: ## Para o MySQL
	@echo "$(BLUE)🐳 Parando MySQL...$(NC)"
	@docker compose down
	@echo "$(GREEN)✅ MySQL parado$(NC)"

docker-logs: ## Mostra logs do MySQL
	@echo "$(BLUE)📋 Logs do MySQL:$(NC)"
	@docker compose logs mysql

docker-reset: ## Reseta dados do MySQL (CUIDADO!)
	@echo "$(RED)⚠️  ATENÇÃO: Isso vai apagar todos os dados!$(NC)"
	@docker compose down -v
	@docker compose up -d
	@echo "$(GREEN)✅ Dados resetados$(NC)"

# ==============================================================================
# Comandos de Migração
# ==============================================================================
migrate-create: ## Cria nova migração
	@echo "$(BLUE)📝 Criando nova migração...$(NC)"
	@migrate create -ext=sql -dir=$(MIGRATIONS_PATH) -seq $(name)
	@echo "$(GREEN)✅ Migração criada!$(NC)"

migrate-up: ## Aplica migrações
	@echo "$(BLUE)⬆️  Aplicando migrações...$(NC)"
	@migrate -path=$(MIGRATIONS_PATH) -database="$(DATABASE_URL)" -verbose up
	@echo "$(GREEN)✅ Migrações aplicadas!$(NC)"

migrate-down: ## Reverte migrações
	@echo "$(BLUE)⬇️  Revertendo migrações...$(NC)"
	@migrate -path=$(MIGRATIONS_PATH) -database="$(DATABASE_URL)" -verbose down
	@echo "$(GREEN)✅ Migrações revertidas!$(NC)"

migrate-status: ## Mostra status das migrações
	@echo "$(BLUE)📊 Status das migrações:$(NC)"
	@migrate -path=$(MIGRATIONS_PATH) -database="$(DATABASE_URL)" version

migrate-force: ## Força versão específica da migração
	@echo "$(BLUE)🔧 Forçando versão $(version)...$(NC)"
	@migrate -path=$(MIGRATIONS_PATH) -database="$(DATABASE_URL)" force $(version)

# ==============================================================================
# Comandos de Banco de Dados
# ==============================================================================
db-connect: ## Conecta ao MySQL
	@echo "$(BLUE)🔌 Conectando ao MySQL...$(NC)"
	@docker compose exec mysql mysql -uroot -p$(DB_PASSWORD) $(DB_NAME)

db-tables: ## Lista tabelas do banco
	@echo "$(BLUE)📋 Tabelas do banco:$(NC)"
	@docker compose exec mysql mysql -uroot -p$(DB_PASSWORD) $(DB_NAME) -e "SHOW TABLES;"

db-describe: ## Descreve estrutura das tabelas
	@echo "$(BLUE)📊 Estrutura das tabelas:$(NC)"
	@docker compose exec mysql mysql -uroot -p$(DB_PASSWORD) $(DB_NAME) -e "DESCRIBE categories; DESCRIBE courses;"

db-seed: ## Insere dados de exemplo
	@echo "$(BLUE)🌱 Inserindo dados de exemplo...$(NC)"
	@echo "$(YELLOW)⚠️  Implementar inserção de dados de exemplo$(NC)"

# ==============================================================================
# Comandos SQLC
# ==============================================================================
sqlc-generate: ## Gera código Go com SQLC
	@echo "$(BLUE)⚡ Gerando código Go...$(NC)"
	@sqlc generate
	@echo "$(GREEN)✅ Código gerado!$(NC)"

sqlc-compile: ## Verifica queries SQL
	@echo "$(BLUE)🔍 Verificando queries...$(NC)"
	@sqlc compile
	@echo "$(GREEN)✅ Queries válidas!$(NC)"

sqlc-doc: ## Gera documentação das queries
	@echo "$(BLUE)📚 Gerando documentação...$(NC)"
	@sqlc doc
	@echo "$(GREEN)✅ Documentação gerada!$(NC)"

# ==============================================================================
# Comandos de Desenvolvimento Completo
# ==============================================================================
dev-setup: docker-up migrate-up ## Setup completo para desenvolvimento
	@echo "$(GREEN)🎉 Ambiente de desenvolvimento pronto!$(NC)"

dev-reset: docker-reset migrate-up ## Reset completo do ambiente
	@echo "$(GREEN)🔄 Ambiente resetado e pronto!$(NC)"

# ==============================================================================
# Ajuda
# ==============================================================================
help: ## Mostra esta ajuda
	@echo "$(BLUE)Comandos disponíveis:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-20s$(NC) %s\n", $$1, $$2}'

.DEFAULT_GOAL := help