# 🚀 Quick Start Guide

## 📋 Pré-requisitos

```bash
# Verificar Go
go version

# Instalar protoc (se necessário)
brew install protobuf

# Instalar plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## 🏃 Execução Rápida

### 1️⃣ Iniciar o Servidor

```bash
cd /Users/ecarvalho/Documents/GitHub/FC_PosGolang/13_gRPC_FC
go run cmd/grpcServer/main.go
```

### 2️⃣ Testar os Endpoints

**Listar serviços disponíveis:**

```bash
grpcurl -plaintext localhost:50051 list pb.CategoryService
```

**Criar uma categoria (Unary):**

```bash
grpcurl -plaintext -d '{"name": "Backend", "description": "Desenvolvimento backend"}' \
  localhost:50051 pb.CategoryService.CreateCategory
```

**Server-Side Streaming:**

```bash
make test-stream
```

**Bidirectional Streaming:**

```bash
make test-bidi
```

## 🎯 Tipos de Streaming

### 1. Unary RPC (tradicional)

```mermaid
Cliente → (1 request) → Servidor
        ← (1 response) ←
```

**Métodos:** `CreateCategory`, `GetCategory`, `ListCategories`

### 2. Server-Side Streaming

```mermaid
Cliente → (1 request) → Servidor
        ← (N responses) ←
```

**Método:** `CreateCategoryStream`

- Cliente envia 1 requisição
- Servidor retorna múltiplas respostas em lotes
- Uso: Download de dados, listagens grandes

### 3. Bidirectional Streaming

```mermaid
Cliente ⇄ (N requests/responses) ⇄ Servidor
```

**Método:** `CreateCategoryStreamBidirectional`

- Cliente envia N requisições
- Servidor retorna N respostas
- Comunicação simultânea e independente
- Uso: Chat, colaboração em tempo real

## 📊 Exemplo de Uso

### Server-Side Streaming

```bash
# Terminal 1: Servidor
go run cmd/grpcServer/main.go

# Terminal 2: Cliente
make test-stream
```

**Saída:**

```bash
📦 Lote 1 recebido: 3 categorias
📦 Lote 2 recebido: 3 categorias
📦 Lote 3 recebido: 3 categorias
📦 Lote 4 recebido: 1 categorias
✅ Stream concluído!
```

### Bidirectional Streaming

```bash
# Terminal 1: Servidor
go run cmd/grpcServer/main.go

# Terminal 2: Cliente
make test-bidi
```

**Saída:**

```bash
📤 [1] Enviando: Backend
📥 [1] Recebido: Backend (ID: xxx)
📤 [2] Enviando: Frontend
📥 [2] Recebido: Frontend (ID: xxx)
...
✅ Bidirectional streaming concluído!
```

## 🔧 Comandos Úteis do Makefile

```bash
make help                   # Lista todos os comandos
make run                    # Inicia o servidor
make test-stream           # Testa server-side streaming
make test-bidi             # Testa bidirectional streaming
make test-create           # Testa criação de categoria
make test-list             # Testa listagem de categorias
make test-services         # Lista serviços disponíveis
make clean                 # Limpa arquivos temporários
```

## 🐛 Troubleshooting

### Problema: "method not implemented"

**Solução:** Regenere os arquivos proto e recompile

```bash
rm -f internal/pb/*.pb.go
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
go clean -cache
go build -o main cmd/grpcServer/main.go
./main
```

### Problema: "connection refused"

**Solução:** Verifique se o servidor está rodando

```bash
# Verificar processos
ps aux | grep grpcServer

# Iniciar servidor
go run cmd/grpcServer/main.go
```

### Problema: "table not found"

**Solução:** Inicialize o banco de dados

```bash
sqlite3 db.dbgrpc "CREATE TABLE IF NOT EXISTS categories (id TEXT PRIMARY KEY, name TEXT, description TEXT);"
```

## 📚 Próximos Passos

1. Ler a documentação detalhada:
   - [Server-Side Streaming](streaming.md)
   - [Bidirectional Streaming](bidirectional-streaming.md)

2. Explorar o código:
   - `proto/course_category.proto` - Definições dos serviços
   - `internal/service/category.go` - Implementação dos serviços
   - `cmd/testClient/main.go` - Cliente de teste (server streaming)
   - `cmd/testBidiClient/main.go` - Cliente de teste (bidirectional)

3. Experimentar:
   - Modificar o tamanho dos lotes no server streaming
   - Adicionar mais categorias no bidirectional streaming
   - Implementar client-side streaming

## 🎓 Conceitos Importantes

### EOF (End Of File)

- Indica que o stream terminou
- Cliente: `stream.CloseSend()` envia EOF para o servidor
- Servidor: `return nil` envia EOF para o cliente

### Goroutines no Streaming

- Bidirectional streaming **REQUER** goroutines separadas
- Uma para enviar, outra para receber
- Evita deadlocks

### Backpressure

- gRPC gerencia automaticamente velocidades diferentes
- Servidor rápido + cliente lento = servidor espera
- Cliente rápido + servidor lento = cliente espera

## 🔗 Links Úteis

- [Documentação gRPC](https://grpc.io/docs/)
- [Protocol Buffers](https://protobuf.dev/)
- [gRPC Go Tutorial](https://grpc.io/docs/languages/go/basics/)
