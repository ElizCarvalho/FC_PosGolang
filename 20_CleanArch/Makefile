# ==============================================================================
# Variáveis
# ==============================================================================
APP_NAME=ordersystem
WEB_PORT?=8080
GRPC_PORT?=50051
GRAPHQL_PORT?=8082
DB_USER?=root
DB_PASSWORD?=root
DB_NAME?=orders

# Cores
BLUE=\033[0;34m
GREEN=\033[0;32m
YELLOW=\033[0;33m
NC=\033[0m

# ==============================================================================
# Setup e Desenvolvimento
# ==============================================================================
.PHONY: setup
setup: ## Configura ambiente
	@echo "$(BLUE)🔧 Configurando ambiente...$(NC)"
	@go mod download
	@go install github.com/99designs/gqlgen@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/google/wire/cmd/wire@latest
	@go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	@if [ ! -f cmd/ordersystem/.env ]; then cp env.example cmd/ordersystem/.env; fi
	@echo "$(GREEN)✅ Setup concluído!$(NC)"

.PHONY: run
run: ## Executa a aplicação
	@echo "$(BLUE)🚀 Iniciando aplicação...$(NC)"
	@echo "$(YELLOW)Web: http://localhost:$(WEB_PORT) | gRPC: $(GRPC_PORT) | GraphQL: http://localhost:$(GRAPHQL_PORT)$(NC)"
	@cd cmd/ordersystem && go run .

.PHONY: build
build: ## Compila o binário
	@echo "$(BLUE)🔨 Compilando...$(NC)"
	@go build -o $(APP_NAME) ./cmd/ordersystem
	@echo "$(GREEN)✅ Binário: ./$(APP_NAME)$(NC)"

# ==============================================================================
# Testes
# ==============================================================================
.PHONY: test
test: ## Executa todos os testes
	@echo "$(BLUE)🧪 Executando testes...$(NC)"
	@go test ./... -v

.PHONY: test-coverage
test-coverage: ## Testes com cobertura
	@echo "$(BLUE)🧪 Testes com cobertura...$(NC)"
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)✅ Relatório: coverage.html$(NC)"

# ==============================================================================
# Docker e Database
# ==============================================================================
.PHONY: docker-up
docker-up: ## Sobe MySQL + RabbitMQ
	@echo "$(BLUE)🐳 Subindo containers...$(NC)"
	@docker-compose up -d
	@echo "$(GREEN)✅ Containers iniciados!$(NC)"

.PHONY: docker-down
docker-down: ## Para containers
	@docker-compose down

.PHONY: db-create
db-create: ## Cria banco de dados
	@docker exec mysql mysql -u$(DB_USER) -p$(DB_PASSWORD) -e "CREATE DATABASE IF NOT EXISTS $(DB_NAME);"
	@echo "$(GREEN)✅ Banco criado!$(NC)"

.PHONY: db-migrate
db-migrate: ## Executa migrations
	@docker exec mysql mysql -u$(DB_USER) -p$(DB_PASSWORD) $(DB_NAME) -e "\
		CREATE TABLE IF NOT EXISTS orders (\
			id varchar(255) NOT NULL,\
			price float NOT NULL,\
			tax float NOT NULL,\
			final_price float NOT NULL,\
			PRIMARY KEY (id)\
		);"
	@echo "$(GREEN)✅ Migrations executadas!$(NC)"

.PHONY: db-reset
db-reset: ## Reseta banco
	@docker exec mysql mysql -u$(DB_USER) -p$(DB_PASSWORD) -e "DROP DATABASE IF EXISTS $(DB_NAME);"
	@$(MAKE) db-create db-migrate

# ==============================================================================
# Geração de Código
# ==============================================================================
.PHONY: proto
proto: ## Gera código protobuf
	@echo "$(BLUE)📝 Gerando protobuf...$(NC)"
	@protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
	@echo "$(GREEN)✅ Protobuf gerado!$(NC)"

.PHONY: graphql
graphql: ## Gera código GraphQL
	@echo "$(BLUE)📝 Gerando GraphQL...$(NC)"
	@go run github.com/99designs/gqlgen generate
	@echo "$(GREEN)✅ GraphQL gerado!$(NC)"

.PHONY: generate
generate: proto graphql ## Gera código protobuf e GraphQL

# ==============================================================================
# Testes Manuais
# ==============================================================================
.PHONY: test-http
test-http: ## Testa criação via REST
	@curl -X POST http://localhost:$(WEB_PORT)/order \
		-H "Content-Type: application/json" \
		-d '{"id":"test-'$$(date +%s)'","price":100.0,"tax":10.0}'

.PHONY: test-http-list
test-http-list: ## Testa listagem via REST
	@curl http://localhost:$(WEB_PORT)/orders -s | jq '.'

.PHONY: test-grpc
test-grpc: ## Testa criação via gRPC
	@grpcurl -plaintext -d '{"id":"test-'$$(date +%s)'","price":100.0,"tax":10.0}' \
		localhost:$(GRPC_PORT) pb.OrderService/CreateOrder

.PHONY: test-grpc-list
test-grpc-list: ## Testa listagem via gRPC
	@grpcurl -plaintext localhost:$(GRPC_PORT) pb.OrderService/ListOrders

.PHONY: test-graphql
test-graphql: ## Testa criação via GraphQL
	@curl -X POST http://localhost:$(GRAPHQL_PORT)/query \
		-H "Content-Type: application/json" \
		-d '{"query":"mutation { createOrder(input: {id: \"test-'$$(date +%s)'\", Price: 100.0, Tax: 10.0}) { id Price Tax FinalPrice } }"}' \
		-s | jq '.'

.PHONY: test-graphql-list
test-graphql-list: ## Testa listagem via GraphQL
	@curl -X POST http://localhost:$(GRAPHQL_PORT)/query \
		-H "Content-Type: application/json" \
		-d '{"query":"query { listOrders { id Price Tax FinalPrice } }"}' \
		-s | jq '.'

.PHONY: test-all-list
test-all-list: ## Testa listagem em todas interfaces
	@echo "$(BLUE)🧪 Testando listagem...$(NC)"
	@echo "\n$(YELLOW)REST:$(NC)" && $(MAKE) test-http-list
	@echo "\n$(YELLOW)gRPC:$(NC)" && $(MAKE) test-grpc-list
	@echo "\n$(YELLOW)GraphQL:$(NC)" && $(MAKE) test-graphql-list

# ==============================================================================
# Utilitários
# ==============================================================================
.PHONY: clean
clean: ## Remove arquivos gerados
	@rm -f $(APP_NAME) coverage.out coverage.html
	@echo "$(GREEN)✅ Limpeza concluída!$(NC)"

.PHONY: help
help: ## Mostra ajuda
	@echo "$(BLUE)╔════════════════════════════════════════╗$(NC)"
	@echo "$(BLUE)║   Order System - Clean Architecture   ║$(NC)"
	@echo "$(BLUE)╚════════════════════════════════════════╝$(NC)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-18s$(NC) %s\n", $$1, $$2}' | sort

.DEFAULT_GOAL := help
