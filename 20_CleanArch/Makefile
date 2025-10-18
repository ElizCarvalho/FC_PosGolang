# ==============================================================================
# Variáveis
# ==============================================================================
APP_NAME=ordersystem
VERSION?=latest
GO_VERSION=1.24

# Portas
WEB_PORT?=8080
GRPC_PORT?=50051
GRAPHQL_PORT?=8082

# Database
DB_HOST?=localhost
DB_PORT?=3306
DB_USER?=root
DB_PASSWORD?=root
DB_NAME?=orders

# Cores para output
BLUE=\033[0;34m
GREEN=\033[0;32m
RED=\033[0;31m
YELLOW=\033[0;33m
NC=\033[0m

# ==============================================================================
# Comandos de Setup
# ==============================================================================
.PHONY: setup
setup: ## Configura o ambiente de desenvolvimento
	@echo "$(BLUE)🔧 Configurando ambiente...$(NC)"
	@go mod download
	@echo "$(GREEN)✓ Dependências baixadas$(NC)"
	@go install github.com/99designs/gqlgen@latest
	@echo "$(GREEN)✓ gqlgen instalado$(NC)"
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "$(GREEN)✓ protoc plugins instalados$(NC)"
	@go install github.com/google/wire/cmd/wire@latest
	@echo "$(GREEN)✓ wire instalado$(NC)"
	@if [ ! -f .env ]; then cp env.example .env; echo "$(GREEN)✓ Arquivo .env criado$(NC)"; fi
	@echo "$(GREEN)✅ Setup concluído!$(NC)"

.PHONY: setup-tools
setup-tools: ## Instala ferramentas adicionais de desenvolvimento
	@echo "$(BLUE)🔧 Instalando ferramentas...$(NC)"
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	@echo "$(GREEN)✅ Ferramentas instaladas!$(NC)"

# ==============================================================================
# Comandos de Desenvolvimento
# ==============================================================================
.PHONY: run
run: ## Executa a aplicação
	@echo "$(BLUE)🚀 Iniciando aplicação...$(NC)"
	@echo "$(YELLOW)Web Server: http://localhost:$(WEB_PORT)$(NC)"
	@echo "$(YELLOW)gRPC Server: localhost:$(GRPC_PORT)$(NC)"
	@echo "$(YELLOW)GraphQL Playground: http://localhost:$(GRAPHQL_PORT)$(NC)"
	@go run ./cmd/ordersystem

.PHONY: run-dev
run-dev: docker-up ## Sobe infraestrutura e executa a aplicação
	@sleep 3
	@$(MAKE) run

.PHONY: build
build: ## Compila o binário da aplicação
	@echo "$(BLUE)🔨 Compilando aplicação...$(NC)"
	@go build -o $(APP_NAME) ./cmd/ordersystem
	@echo "$(GREEN)✅ Binário criado: ./$(APP_NAME)$(NC)"

.PHONY: build-linux
build-linux: ## Compila para Linux
	@echo "$(BLUE)🔨 Compilando para Linux...$(NC)"
	@GOOS=linux GOARCH=amd64 go build -o $(APP_NAME)-linux ./cmd/ordersystem
	@echo "$(GREEN)✅ Binário criado: ./$(APP_NAME)-linux$(NC)"

# ==============================================================================
# Comandos de Testes
# ==============================================================================
.PHONY: test
test: ## Executa todos os testes
	@echo "$(BLUE)🧪 Executando testes...$(NC)"
	@go test ./... -v

.PHONY: test-short
test-short: ## Executa testes rápidos
	@echo "$(BLUE)🧪 Executando testes rápidos...$(NC)"
	@go test ./... -short

.PHONY: test-coverage
test-coverage: ## Executa testes com cobertura
	@echo "$(BLUE)🧪 Executando testes com cobertura...$(NC)"
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)✅ Relatório gerado: coverage.html$(NC)"

.PHONY: test-race
test-race: ## Executa testes com detector de race conditions
	@echo "$(BLUE)🧪 Executando testes com race detector...$(NC)"
	@go test ./... -race

.PHONY: test-entity
test-entity: ## Executa testes da camada de entidade
	@echo "$(BLUE)🧪 Testando entidades...$(NC)"
	@go test ./internal/entity -v

.PHONY: test-usecase
test-usecase: ## Executa testes dos use cases
	@echo "$(BLUE)🧪 Testando use cases...$(NC)"
	@go test ./internal/usecase -v

.PHONY: test-web
test-web: ## Executa testes da camada web
	@echo "$(BLUE)🧪 Testando web handlers...$(NC)"
	@go test ./internal/infra/web -v

.PHONY: test-database
test-database: ## Executa testes do repositório
	@echo "$(BLUE)🧪 Testando database...$(NC)"
	@go test ./internal/infra/database -v

.PHONY: test-events
test-events: ## Executa testes de eventos
	@echo "$(BLUE)🧪 Testando eventos...$(NC)"
	@go test ./events -v

# ==============================================================================
# Comandos de Qualidade de Código
# ==============================================================================
.PHONY: lint
lint: ## Executa o linter
	@echo "$(BLUE)🔍 Verificando código...$(NC)"
	@golangci-lint run ./... || echo "$(YELLOW)⚠️  Instale golangci-lint: make setup-tools$(NC)"

.PHONY: fmt
fmt: ## Formata o código
	@echo "$(BLUE)✨ Formatando código...$(NC)"
	@go fmt ./...
	@goimports -w . || echo "$(YELLOW)⚠️  Instale goimports: make setup-tools$(NC)"
	@echo "$(GREEN)✅ Código formatado!$(NC)"

.PHONY: vet
vet: ## Executa go vet
	@echo "$(BLUE)🔍 Executando go vet...$(NC)"
	@go vet ./...

.PHONY: check
check: fmt vet lint test ## Executa todas as verificações

# ==============================================================================
# Comandos Docker
# ==============================================================================
.PHONY: docker-up
docker-up: ## Sobe os containers (MySQL e RabbitMQ)
	@echo "$(BLUE)🐳 Subindo containers...$(NC)"
	@docker-compose up -d
	@echo "$(GREEN)✅ Containers iniciados!$(NC)"
	@echo "$(YELLOW)MySQL: localhost:3306$(NC)"
	@echo "$(YELLOW)RabbitMQ: localhost:5672$(NC)"
	@echo "$(YELLOW)RabbitMQ Management: http://localhost:15672 (guest/guest)$(NC)"

.PHONY: docker-down
docker-down: ## Para os containers
	@echo "$(BLUE)🐳 Parando containers...$(NC)"
	@docker-compose down
	@echo "$(GREEN)✅ Containers parados!$(NC)"

.PHONY: docker-logs
docker-logs: ## Mostra logs dos containers
	@docker-compose logs -f

.PHONY: docker-clean
docker-clean: ## Remove containers e volumes
	@echo "$(BLUE)🐳 Limpando containers e volumes...$(NC)"
	@docker-compose down -v
	@echo "$(GREEN)✅ Limpeza concluída!$(NC)"

.PHONY: docker-restart
docker-restart: docker-down docker-up ## Reinicia os containers

# ==============================================================================
# Comandos de Geração de Código
# ==============================================================================
.PHONY: proto
proto: ## Gera código a partir dos arquivos .proto
	@echo "$(BLUE)📝 Gerando código protobuf...$(NC)"
	@protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
	@echo "$(GREEN)✅ Código protobuf gerado!$(NC)"

.PHONY: graphql
graphql: ## Gera código GraphQL
	@echo "$(BLUE)📝 Gerando código GraphQL...$(NC)"
	@go run github.com/99designs/gqlgen generate
	@echo "$(GREEN)✅ Código GraphQL gerado!$(NC)"

.PHONY: wire
wire: ## Gera código de injeção de dependências
	@echo "$(BLUE)📝 Gerando código Wire...$(NC)"
	@cd cmd/ordersystem && wire
	@echo "$(GREEN)✅ Código Wire gerado!$(NC)"

.PHONY: generate
generate: proto graphql wire ## Gera todo o código (proto, graphql, wire)
	@echo "$(GREEN)✅ Todo código gerado com sucesso!$(NC)"

# ==============================================================================
# Comandos de Database
# ==============================================================================
.PHONY: db-create
db-create: ## Cria o banco de dados
	@echo "$(BLUE)💾 Criando banco de dados...$(NC)"
	@docker exec -it mysql mysql -u$(DB_USER) -p$(DB_PASSWORD) -e "CREATE DATABASE IF NOT EXISTS $(DB_NAME);"
	@echo "$(GREEN)✅ Banco criado!$(NC)"

.PHONY: db-migrate
db-migrate: ## Executa as migrations
	@echo "$(BLUE)💾 Executando migrations...$(NC)"
	@docker exec -it mysql mysql -u$(DB_USER) -p$(DB_PASSWORD) $(DB_NAME) -e "\
		CREATE TABLE IF NOT EXISTS orders (\
			id varchar(255) NOT NULL,\
			price float NOT NULL,\
			tax float NOT NULL,\
			final_price float NOT NULL,\
			PRIMARY KEY (id)\
		);"
	@echo "$(GREEN)✅ Migrations executadas!$(NC)"

.PHONY: db-drop
db-drop: ## Remove o banco de dados
	@echo "$(RED)⚠️  Removendo banco de dados...$(NC)"
	@docker exec -it mysql mysql -u$(DB_USER) -p$(DB_PASSWORD) -e "DROP DATABASE IF EXISTS $(DB_NAME);"
	@echo "$(GREEN)✅ Banco removido!$(NC)"

.PHONY: db-reset
db-reset: db-drop db-create db-migrate ## Reseta o banco de dados

# ==============================================================================
# Comandos de Teste Manual
# ==============================================================================
.PHONY: test-http
test-http: ## Testa endpoint HTTP (requer httpie ou curl)
	@echo "$(BLUE)🧪 Testando endpoint HTTP...$(NC)"
	@curl -X POST http://localhost:$(WEB_PORT)/order \
		-H "Content-Type: application/json" \
		-d '{"id":"test-'$$(date +%s)'","price":100.0,"tax":10.0}' || \
		echo "$(RED)❌ Erro: Certifique-se que a aplicação está rodando$(NC)"

.PHONY: test-grpc
test-grpc: ## Testa endpoint gRPC (requer grpcurl)
	@echo "$(BLUE)🧪 Testando endpoint gRPC...$(NC)"
	@grpcurl -plaintext -d '{"id":"test-'$$(date +%s)'","price":100.0,"tax":10.0}' \
		localhost:$(GRPC_PORT) pb.OrderService/CreateOrder || \
		echo "$(YELLOW)⚠️  Instale grpcurl: make setup-tools$(NC)"

.PHONY: test-graphql
test-graphql: ## Abre o GraphQL Playground
	@echo "$(BLUE)🧪 Abrindo GraphQL Playground...$(NC)"
	@echo "$(YELLOW)Acesse: http://localhost:$(GRAPHQL_PORT)$(NC)"
	@open http://localhost:$(GRAPHQL_PORT) 2>/dev/null || \
		xdg-open http://localhost:$(GRAPHQL_PORT) 2>/dev/null || \
		echo "$(YELLOW)Abra manualmente: http://localhost:$(GRAPHQL_PORT)$(NC)"

# ==============================================================================
# Comandos de Limpeza
# ==============================================================================
.PHONY: clean
clean: ## Remove arquivos gerados
	@echo "$(BLUE)🧹 Limpando arquivos...$(NC)"
	@rm -f $(APP_NAME) $(APP_NAME)-linux
	@rm -f coverage.out coverage.html
	@rm -rf .docker
	@echo "$(GREEN)✅ Limpeza concluída!$(NC)"

.PHONY: clean-all
clean-all: clean docker-clean ## Remove tudo (arquivos e containers)
	@echo "$(GREEN)✅ Limpeza completa concluída!$(NC)"

# ==============================================================================
# Comandos de Informação
# ==============================================================================
.PHONY: info
info: ## Mostra informações do projeto
	@echo "$(BLUE)📊 Informações do Projeto$(NC)"
	@echo "$(YELLOW)Nome:$(NC) $(APP_NAME)"
	@echo "$(YELLOW)Versão Go:$(NC) $(GO_VERSION)"
	@echo "$(YELLOW)Porta Web:$(NC) $(WEB_PORT)"
	@echo "$(YELLOW)Porta gRPC:$(NC) $(GRPC_PORT)"
	@echo "$(YELLOW)Porta GraphQL:$(NC) $(GRAPHQL_PORT)"
	@echo ""
	@echo "$(BLUE)📦 Dependências:$(NC)"
	@go list -m all | head -10

.PHONY: deps
deps: ## Lista as dependências do projeto
	@echo "$(BLUE)📦 Dependências do projeto:$(NC)"
	@go list -m all

.PHONY: deps-update
deps-update: ## Atualiza as dependências
	@echo "$(BLUE)📦 Atualizando dependências...$(NC)"
	@go get -u ./...
	@go mod tidy
	@echo "$(GREEN)✅ Dependências atualizadas!$(NC)"

# ==============================================================================
# Comandos Compostos
# ==============================================================================
.PHONY: start
start: docker-up db-create db-migrate run ## Inicia todo o ambiente

.PHONY: restart
restart: docker-restart run ## Reinicia todo o ambiente

.PHONY: ci
ci: fmt vet lint test ## Executa pipeline de CI

# ==============================================================================
# Ajuda
# ==============================================================================
.PHONY: help
help: ## Mostra esta ajuda
	@echo "$(BLUE)╔════════════════════════════════════════════════════════════╗$(NC)"
	@echo "$(BLUE)║         Order System - Clean Architecture                 ║$(NC)"
	@echo "$(BLUE)╚════════════════════════════════════════════════════════════╝$(NC)"
	@echo ""
	@echo "$(YELLOW)Comandos disponíveis:$(NC)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-20s$(NC) %s\n", $$1, $$2}' | \
		sort
	@echo ""
	@echo "$(BLUE)Exemplos de uso:$(NC)"
	@echo "  $(YELLOW)make setup$(NC)          - Primeira vez no projeto"
	@echo "  $(YELLOW)make start$(NC)          - Inicia tudo (DB + App)"
	@echo "  $(YELLOW)make run-dev$(NC)        - Desenvolvimento rápido"
	@echo "  $(YELLOW)make test$(NC)           - Roda todos os testes"
	@echo "  $(YELLOW)make ci$(NC)             - Valida antes do commit"
	@echo ""

.DEFAULT_GOAL := help

