# ==============================================================================
# Variáveis
# ==============================================================================
APP_NAME=s3-study

# Cores
BLUE=\033[0;34m
GREEN=\033[0;32m
RED=\033[0;31m
YELLOW=\033[0;33m
NC=\033[0m

# ==============================================================================
# Comandos de Desenvolvimento
# ==============================================================================
.PHONY: setup minio-up minio-down run run-optimized test benchmark clean console

setup: ## Configura o ambiente
	@echo "$(BLUE)🔧 Configurando ambiente...$(NC)"
	@go mod download
	@echo "$(GREEN)✅ Ambiente configurado!$(NC)"

minio-up: ## Sobe o MinIO
	@echo "$(BLUE)🚀 Iniciando MinIO...$(NC)"
	@docker-compose up -d
	@echo "$(GREEN)✅ MinIO rodando!$(NC)"
	@echo "$(YELLOW)   API: http://localhost:9000$(NC)"
	@echo "$(YELLOW)   Console: http://localhost:9001$(NC)"
	@echo "$(YELLOW)   Usuário: minioadmin$(NC)"
	@echo "$(YELLOW)   Senha: minioadmin$(NC)"

minio-down: ## Para o MinIO
	@echo "$(BLUE)🛑 Parando MinIO...$(NC)"
	@docker-compose down
	@echo "$(GREEN)✅ MinIO parado!$(NC)"

run: ## Roda a aplicação básica
	@echo "$(BLUE)🚀 Executando aplicação básica...$(NC)"
	@go run main.go

run-demo: ## Roda o demo de performance
	@echo "$(BLUE)🚀 Executando demo de performance...$(NC)"
	@go run main.go demo

test: ## Roda os testes
	@echo "$(BLUE)🧪 Executando testes...$(NC)"
	@go test -v ./...

benchmark: ## Roda benchmarks de performance
	@echo "$(BLUE)⚡ Executando benchmarks...$(NC)"
	@go test -bench=. -benchmem ./...

console: ## Abre o console do MinIO
	@echo "$(BLUE)🌐 Abrindo console do MinIO...$(NC)"
	@open http://localhost:9001 || xdg-open http://localhost:9001

clean: ## Remove arquivos gerados
	@echo "$(BLUE)🧹 Limpando arquivos...$(NC)"
	@rm -f arquivo-local.txt test-file-*.txt large-file.txt
	@docker-compose down -v
	@echo "$(GREEN)✅ Limpeza concluída!$(NC)"

# ==============================================================================
# Ajuda
# ==============================================================================
help: ## Mostra essa ajuda
	@echo "$(BLUE)Comandos disponíveis:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-20s$(NC) %s\n", $$1, $$2}'

.DEFAULT_GOAL := help