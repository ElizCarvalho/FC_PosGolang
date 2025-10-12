# 🚀 Projeto gRPC - Course Category Service

> Um serviço gRPC para gerenciamento de categorias de cursos, implementado em Go

## 📌 Sobre

Este projeto demonstra a implementação de um serviço gRPC em Go para gerenciar categorias de cursos. O gRPC é um framework de comunicação RPC (Remote Procedure Call) de alta performance desenvolvido pelo Google, que utiliza Protocol Buffers como formato de serialização.

### 🌟 Features Implementadas

- ✅ **Unary RPC**: Requisição e resposta simples (CreateCategory, GetCategory, ListCategories)
- ✅ **Server-Side Streaming**: Servidor envia múltiplas respostas (CreateCategoryStream)
- ✅ **Bidirectional Streaming**: Cliente e servidor enviam múltiplas mensagens (CreateCategoryStreamBidirectional)
- ✅ Persistência com SQLite
- ✅ Reflection habilitado para introspecção
- ✅ Clientes de teste para demonstração

## 🔧 Configuração do Ambiente

### Pré-requisitos

- **Go 1.24+** instalado
- **Protocol Buffers Compiler (protoc)** instalado
- **Plugins do Go** para protoc instalados

### 1. Instalação do Protocol Buffers Compiler

#### macOS (usando Homebrew)

```bash
brew install protobuf
```

### 2. Instalação dos Plugins do Go

```bash
# Plugin para gerar código Go a partir de .proto
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Plugin para gerar código gRPC Go a partir de .proto
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 3. Verificação da Instalação

```bash
# Verificar se protoc está instalado
protoc --version

# Verificar se os plugins estão no PATH
protoc-gen-go --version
protoc-gen-go-grpc --version
```

## 📁 Estrutura do Projeto

```bash
13_gRPC_FC/
├── proto/
│   └── course_category.proto    # Definição do serviço gRPC
├── internal/
│   ├── pb/                      # Código gerado pelo protoc
│   │   ├── course_category.pb.go
│   │   └── course_category_grpc.pb.go
│   └── database/
│       ├── category.go          # Implementação do banco de dados
│       └── course.go
├── go.mod
├── go.sum
└── README.md
```

## 📋 Arquivo course_category.proto

O arquivo `proto/course_category.proto` é o **coração** do projeto gRPC. Ele define:

### 1. **Estruturas de Dados (Messages)**

```protobuf
message Category {
  string id = 1;
  string name = 2;
  string description = 3;
}
```

### 2. **Requisições e Respostas**

```protobuf
message CreateCategoryRequest {
    string name = 1;
    string description = 2;
}

message CategoryResponse {
    Category category = 1;
}
```

### 3. **Definição do Serviço**

```protobuf
service CategoryService {
    rpc CreateCategory(CreateCategoryRequest) returns (CategoryResponse) {}
}
```

### 📝 Explicação dos Números nos Campos

Os números (1, 2, 3) são **identificadores únicos** para cada campo na mensagem. Eles são usados para:

- **Serialização/Deserialização** eficiente
- **Compatibilidade** entre versões diferentes do .proto
- **Otimização** do tamanho da mensagem

⚠️ **IMPORTANTE**: Nunca altere esses números em campos existentes, pois isso quebra a compatibilidade!

## 🔨 Comando de Geração de Código

### Comando Principal

```bash
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

### Explicação dos Parâmetros

- `protoc`: Compilador do Protocol Buffers
- `--go_out=.`: Gera código Go para as mensagens no diretório atual
- `--go-grpc_out=.`: Gera código Go para os serviços gRPC no diretório atual
- `proto/course_category.proto`: Arquivo de definição

### O que o Comando Gera

1. **`course_category.pb.go`**: Contém as estruturas Go equivalentes às mensagens do .proto
2. **`course_category_grpc.pb.go`**: Contém interfaces e código para implementar o serviço gRPC

### Exemplo de Uso no Terminal

```bash
# Navegar para o diretório do projeto
cd 13_gRPC_FC

# Gerar código Go a partir do .proto
protoc --go_out=. --go-grpc_out=. proto/course_category.proto

# Verificar se os arquivos foram gerados
ls -la internal/pb/
```

## 🚀 Como Executar

### 1. Instalar Dependências

```bash
go mod tidy
```

### 2. Instalar Ferramentas de Teste

```bash
# Instalar grpcurl para testar o serviço gRPC
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# Instalar Evans (cliente gRPC interativo) - opcional
go install github.com/ktr0731/evans@latest
```

### 3. Gerar Código (se necessário)

```bash
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

### 4. Executar o Servidor

```bash
# Usando Makefile (recomendado)
make run

# Ou diretamente
go run cmd/grpcServer/main.go
```

## 📚 Conceitos Importantes

### Protocol Buffers (protobuf)

- **Formato de serialização** binário eficiente
- **Linguagem neutra** - funciona com várias linguagens
- **Schema definido** - garante compatibilidade entre versões
- **Performance superior** ao JSON

### gRPC

- **RPC framework** de alta performance
- **HTTP/2** como transporte
- **Streaming** bidirecional
- **Code generation** automático

### Vantagens do gRPC

- ✅ **Performance**: Mais rápido que REST/JSON
- ✅ **Type Safety**: Código gerado com tipos seguros
- ✅ **Streaming**: Suporte a streams bidirecionais
- ✅ **Multi-language**: Funciona com várias linguagens
- ✅ **Schema Evolution**: Compatibilidade entre versões

## 🧪 Testando o Serviço

### 1. Usando grpcurl (Recomendado)

O **grpcurl** é uma ferramenta de linha de comando para testar serviços gRPC, similar ao curl para APIs REST.

#### Listar Serviços Disponíveis

```bash
# Listar todos os serviços
grpcurl -plaintext localhost:50051 list

# Listar métodos de um serviço específico
grpcurl -plaintext localhost:50051 list pb.CategoryService
```

#### Criar uma Categoria

```bash
grpcurl -plaintext -d '{
  "name": "Backend",
  "description": "Cursos de desenvolvimento backend"
}' localhost:50051 pb.CategoryService.CreateCategory
```

#### Listar Categorias

```bash
# Listar todas as categorias
grpcurl -plaintext localhost:50051 pb.CategoryService.ListCategories
```

#### Buscar Categoria por ID

```bash
# Buscar categoria específica por ID
grpcurl -plaintext -d '{"id": "CATEGORY_ID"}' localhost:50051 pb.CategoryService.GetCategory
```

#### Exemplo de Resposta - CreateCategory

```json
{
  "category": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Backend",
    "description": "Cursos de desenvolvimento backend"
  }
}
```

#### Exemplo de Resposta - ListCategories

```json
{
  "categories": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "name": "Backend",
      "description": "Cursos de desenvolvimento backend"
    },
    {
      "id": "456e7890-e89b-12d3-a456-426614174001",
      "name": "Frontend",
      "description": "Cursos de desenvolvimento frontend"
    }
  ]
}
```

#### Exemplo de Resposta - GetCategory

```json
{
  "category": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Backend",
    "description": "Cursos de desenvolvimento backend"
  }
}
```

### 2. Usando Makefile (Mais Fácil)

```bash
# Listar serviços disponíveis
make test-services

# Listar métodos do CategoryService
make test-category-methods

# Listar categorias
make test-list

# Buscar categoria por ID
make test-get

# Criar categoria
make test-create

# Modo interativo com grpcurl
make interactive-test
```

### 3. Usando Evans (Cliente Interativo)

```bash
# Abrir Evans
make evans

# No Evans:
localhost:50051> show service
localhost:50051> service pb.CategoryService
localhost:50051> call CreateCategory
```

### 4. Cliente gRPC em Go (Exemplo)

```go
package main

import (
    "context"
    "log"
    
    "google.golang.org/grpc"
    pb "github.com/ElizCarvalho/FC_PosGolang/13_gRPC_FC/internal/pb"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    
    client := pb.NewCategoryServiceClient(conn)
    
    resp, err := client.CreateCategory(context.Background(), &pb.CreateCategoryRequest{
        Name: "Backend",
        Description: "Cursos de desenvolvimento backend",
    })
    
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Categoria criada: %+v", resp.Category)
}
```

## 🛠️ Comandos Úteis

### Regenerar Código

#### Obs: apagar os arquivos gerados para garantir que o comando vai funcionar

```bash
# Limpar arquivos gerados
rm -rf internal/pb/*

# Regenerar
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

### Verificar Dependências

```bash
go mod verify
```

### Executar Testes

```bash
go test ./...
```

## 🔧 Troubleshooting

### Problema: Evans não lista serviços

Se o Evans não mostrar os serviços quando você executa `show service`, tente:

1. **Reiniciar o Evans**:

   ```bash
   # Fechar o Evans atual (Ctrl+C)
   # Reabrir com arquivo proto
   evans -r repl --host localhost --port 50051 --proto proto/course_category.proto
   ```

2. **Usar grpcurl** (mais confiável):

   ```bash
   grpcurl -plaintext localhost:50051 list
   ```

3. **Verificar se o servidor está rodando**:

   ```bash
   lsof -i :50051
   ```

### Problema: "no such table: categories"

Se aparecer esse erro, execute:

```bash
# Criar a tabela no banco
make init-db
```

### Problema: Servidor não inicia

1. **Verificar se a porta está livre**:

   ```bash
   lsof -i :50051
   ```

2. **Matar processo na porta**:

   ```bash
   kill -9 $(lsof -ti:50051)
   ```

3. **Reiniciar o servidor**:

   ```bash
   make run
   ```

## 📋 Comandos do Makefile

O projeto inclui um Makefile com comandos úteis para desenvolvimento:

```bash
# Ver todos os comandos disponíveis
make help

# Configurar ambiente
make setup

# Inicializar banco de dados
make init-db

# Executar servidor
make run

# Testar gRPC
make test-grpc

# Criar categoria de teste
make test-create

# Abrir Evans (cliente interativo)
make evans

# Modo interativo com grpcurl
make interactive-test

# Limpar arquivos temporários
make clean
```

## 📚 Documentação de Streaming

Este projeto implementa três tipos de comunicação gRPC:

### 1. Server-Side Streaming

📄 **[Documentação completa: Server-Side Streaming](docs/streaming.md)**

- Como funciona o streaming unidirecional (servidor → cliente)
- Implementação do `CreateCategoryStream`
- Casos de uso e melhores práticas
- Troubleshooting

**Teste rápido:**

```bash
make test-stream
```

### 2. Bidirectional Streaming

📄 **[Documentação completa: Bidirectional Streaming](docs/bidirectional-streaming.md)**

- Como funciona o streaming bidirecional (cliente ↔ servidor)
- Implementação do `CreateCategoryStreamBidirectional`
- Gerenciamento de goroutines
- Comparação com outros tipos de streaming
- Casos de uso reais (chat, upload com progresso, etc.)

**Teste rápido:**

```bash
make test-bidi
```

## 📖 Recursos para Estudo

- [Documentação oficial do gRPC](https://grpc.io/docs/)
- [Protocol Buffers Guide](https://developers.google.com/protocol-buffers/docs/overview)
- [gRPC Go Quick Start](https://grpc.io/docs/languages/go/quickstart/)
- [grpcurl - Ferramenta de teste gRPC](https://github.com/fullstorydev/grpcurl)
- [gRPC Streaming Concepts](https://grpc.io/docs/what-is-grpc/core-concepts/#server-streaming-rpc)

## 🤝 Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature
3. Commit suas mudanças
4. Push para a branch
5. Abra um Pull Request

## 📄 Licença

Este projeto é parte do curso de Pós-Graduação em Go da Full Cycle.

---

**💡 Dica**: Este README serve como material de estudo e referência. Mantenha-o sempre atualizado conforme o projeto evolui!
