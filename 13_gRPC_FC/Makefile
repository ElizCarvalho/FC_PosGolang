# ==============================================================================
# Variáveis
# ==============================================================================
APP_NAME=grpc-server
VERSION?=latest
PORT?=50051

# Cores
BLUE=\033[0;34m
GREEN=\033[0;32m
RED=\033[0;31m
YELLOW=\033[0;33m
NC=\033[0m

# ==============================================================================
# Comandos de Desenvolvimento
# ==============================================================================
.PHONY: setup run init-db test-client clean

setup: ## Configura o ambiente
	@echo "$(BLUE)🔧 Configurando ambiente...$(NC)"
	@go mod download
	@go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

init-db: ## Inicializa o banco de dados
	@echo "$(BLUE)🗄️ Criando tabelas...$(NC)"
	@go run cmd/initdb/main.go

run: init-db ## Roda o servidor gRPC
	@echo "$(BLUE)🚀 Iniciando servidor gRPC na porta $(PORT)...$(NC)"
	@go run cmd/grpcServer/main.go

test-client: ## Testa o cliente gRPC
	@echo "$(BLUE)🧪 Testando cliente gRPC...$(NC)"
	@grpcurl -plaintext localhost:$(PORT) list

test-create: ## Testa criação de categoria
	@echo "$(BLUE)🧪 Testando criação de categoria...$(NC)"
	@grpcurl -plaintext -d '{"name": "Teste", "description": "Categoria de teste"}' localhost:$(PORT) pb.CategoryService.CreateCategory

test-list: ## Testa listagem de categorias
	@echo "$(BLUE)🧪 Testando listagem de categorias...$(NC)"
	@grpcurl -plaintext localhost:$(PORT) pb.CategoryService.ListCategories

test-get: ## Testa busca de categoria por ID
	@echo "$(BLUE)🧪 Testando busca de categoria por ID...$(NC)"
	@echo "Exemplo: grpcurl -plaintext -d '{\"id\": \"780c5d43-eccd-4267-a03b-3bc728db4c09\"}' localhost:$(PORT) pb.CategoryService.GetCategory"
	@echo "Substitua CATEGORY_ID pelo ID de uma categoria existente"

test-stream: ## Testa o streaming de categorias
	@echo "$(BLUE)🧪 Testando CreateCategoryStream...$(NC)"
	@go run cmd/testClient/main.go

test-services: ## Lista todos os serviços disponíveis
	@echo "$(BLUE)🔍 Listando serviços disponíveis...$(NC)"
	@grpcurl -plaintext localhost:$(PORT) list

test-category-methods: ## Lista métodos do CategoryService
	@echo "$(BLUE)🔍 Listando métodos do CategoryService...$(NC)"
	@grpcurl -plaintext localhost:$(PORT) list pb.CategoryService

evans: ## Abre o Evans para interação
	@echo "$(BLUE)🔧 Abrindo Evans...$(NC)"
	@evans -r repl --host localhost --port $(PORT) --proto proto/course_category.proto

test-grpc: ## Testa o gRPC com grpcurl
	@echo "$(BLUE)🧪 Testando gRPC...$(NC)"
	@grpcurl -plaintext localhost:$(PORT) list

interactive-test: ## Teste interativo com grpcurl
	@echo "$(BLUE)🔧 Modo interativo com grpcurl...$(NC)"
	@echo "Comandos úteis:"
	@echo "  grpcurl -plaintext localhost:$(PORT) list"
	@echo "  grpcurl -plaintext localhost:$(PORT) list pb.CategoryService"
	@echo "  grpcurl -plaintext localhost:$(PORT) pb.CategoryService.ListCategories"
	@echo "  grpcurl -plaintext -d '{\"name\": \"Teste\", \"description\": \"Teste\"}' localhost:$(PORT) pb.CategoryService.CreateCategory"
	@echo "  grpcurl -plaintext -d '{\"id\": \"CATEGORY_ID\"}' localhost:$(PORT) pb.CategoryService.GetCategory"
	@bash

clean: ## Limpa arquivos temporários
	@echo "$(BLUE)🧹 Limpando arquivos...$(NC)"
	@rm -f db.dbgrpc

# ==============================================================================
# Ajuda
# ==============================================================================
help: ## Mostra essa ajuda
	@echo "$(BLUE)Comandos disponíveis:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-20s$(NC) %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
