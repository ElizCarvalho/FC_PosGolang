# ==============================================================================
# Variáveis
# ==============================================================================
APP_NAME=grpc-calculator
VERSION?=latest
PORT?=50051
PROTO_DIR=proto
SERVER_DIR=server
CLIENT_DIR=client

# Cores
BLUE=\033[0;34m
GREEN=\033[0;32m
RED=\033[0;31m
YELLOW=\033[0;33m
NC=\033[0m

# ==============================================================================
# Comandos de Desenvolvimento
# ==============================================================================
.PHONY: setup run-server run-client test lint proto clean help

setup: ## Configura o ambiente
	@echo "$(BLUE)🔧 Configurando ambiente gRPC...$(NC)"
	@go mod download
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "$(GREEN)✅ Ambiente configurado!$(NC)"

proto: ## Gera código Go a partir dos arquivos .proto
	@echo "$(BLUE)📝 Gerando código Go a partir dos arquivos .proto...$(NC)"
	@protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto
	@echo "$(GREEN)✅ Código gerado com sucesso!$(NC)"

run-server: ## Executa o servidor gRPC
	@echo "$(BLUE)🚀 Iniciando servidor gRPC na porta $(PORT)...$(NC)"
	@go run $(SERVER_DIR)/main.go

run-client: ## Executa o cliente gRPC
	@echo "$(BLUE)📱 Iniciando cliente gRPC...$(NC)"
	@go run $(CLIENT_DIR)/main.go

test: ## Executa os testes
	@echo "$(BLUE)🧪 Executando testes...$(NC)"
	@go test -v ./...

lint: ## Executa o linter
	@echo "$(BLUE)🔍 Verificando código...$(NC)"
	@go vet ./...
	@go fmt ./...

clean: ## Limpa arquivos gerados
	@echo "$(BLUE)🧹 Limpando arquivos gerados...$(NC)"
	@rm -f $(PROTO_DIR)/*.pb.go
	@rm -f $(PROTO_DIR)/*_grpc.pb.go
	@go clean
	@echo "$(GREEN)✅ Limpeza concluída!$(NC)"

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
# Ajuda
# ==============================================================================
help: ## Mostra essa ajuda
	@echo "$(BLUE)Comandos disponíveis para gRPC:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-20s$(NC) %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
