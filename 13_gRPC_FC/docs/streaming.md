# 📡 gRPC Server-Side Streaming

## 🎯 O que foi implementado

Foi implementado o método `CreateCategoryStream` que demonstra **Server-Side Streaming** no gRPC.

## 🧠 Conceito: Server-Side Streaming

### O que é?

- **Padrão**: Cliente envia UMA requisição → Servidor envia MÚLTIPLAS respostas
- **Fluxo**: Unidirecional (servidor → cliente)
- **Benefício**: Cliente recebe dados progressivamente, sem esperar tudo terminar

### Quando usar?

- Processar grandes volumes de dados em lotes
- Enviar atualizações progressivas ao cliente
- Melhorar a percepção de performance
- Reduzir uso de memória no servidor

## 🏗️ Arquitetura da Implementação

### 1. **Definição no Proto** (`proto/course_category.proto`)

```protobuf
service CategoryService {
    // Unary RPC (1 request → 1 response)
    rpc CreateCategory(CreateCategoryRequest) returns (CategoryResponse) {}
    
    // Server Streaming RPC (1 request → N responses)
    rpc CreateCategoryStream(CreateCategoryRequest) returns (stream CategoryList) {}
}
```

**Observação**: A palavra-chave `stream` antes do tipo de retorno indica Server-Side Streaming.

### 2. **Database Layer** (`internal/database/category.go`)

```go
// CreateMultiple cria múltiplas categorias simuladas
func (c *Category) CreateMultiple(baseName, baseDescription string, count int) ([]Category, error) {
    var categories []Category
    
    for i := 0; i < count; i++ {
        category, err := c.Create(
            fmt.Sprintf("%s %d", baseName, i+1),
            fmt.Sprintf("%s - Categoria %d", baseDescription, i+1),
        )
        if err != nil {
            return nil, err
        }
        categories = append(categories, category)
    }
    
    return categories, nil
}
```

### 3. **Service Layer** (`internal/service/category.go`)

```go
func (c *CategoryService) CreateCategoryStream(
    req *pb.CreateCategoryRequest, 
    stream grpc.ServerStreamingServer[pb.CategoryList],
) error {
    // 1. Cria as categorias
    categories, err := c.CategoryDB.CreateMultiple(req.Name, req.Description, 10)
    if err != nil {
        return err
    }

    // 2. Processa em lotes
    batchSize := 3
    for i := 0; i < len(categories); i += batchSize {
        end := i + batchSize
        if end > len(categories) {
            end = len(categories)
        }

        // 3. Cria o lote atual
        var batch []*pb.Category
        for j := i; j < end; j++ {
            batch = append(batch, &pb.Category{
                Id:          categories[j].ID,
                Name:        categories[j].Name,
                Description: categories[j].Description,
            })
        }

        // 4. Envia o lote via stream
        err := stream.Send(&pb.CategoryList{Categories: batch})
        if err != nil {
            return err
        }
    }

    return nil
}
```

**Pontos-chave:**

- Parâmetro `stream grpc.ServerStreamingServer[pb.CategoryList]`: Interface para enviar múltiplas respostas
- Método `stream.Send()`: Envia cada lote para o cliente
- Retorno `nil`: Indica que o stream foi concluído com sucesso

### 4. **Cliente** (`cmd/testClient/main.go`)

```go
// 1. Chama o método de streaming
stream, err := client.CreateCategoryStream(context.Background(), req)
if err != nil {
    log.Fatalf("Erro ao chamar CreateCategoryStream: %v", err)
}

// 2. Processa as respostas do stream
for {
    response, err := stream.Recv()
    if err == io.EOF {
        // Stream terminou
        break
    }
    if err != nil {
        log.Fatalf("Erro ao receber do stream: %v", err)
    }

    // 3. Processa cada lote recebido
    fmt.Printf("📦 Lote recebido: %d categorias\n", len(response.Categories))
}
```

**Pontos-chave:**

- `stream.Recv()`: Bloqueia até receber próxima resposta ou `io.EOF`
- `io.EOF`: Indica que o servidor terminou o stream
- Loop contínuo até receber `EOF`

## 📊 Fluxo de Execução

```mermaid
Cliente                     Servidor
  |                            |
  |---(1) CreateCategoryStream--->|
  |                            |  [Cria 10 categorias]
  |<--(2) Lote 1 (3 items)-----|
  |                            |
  |<--(3) Lote 2 (3 items)-----|
  |                            |
  |<--(4) Lote 3 (3 items)-----|
  |                            |
  |<--(5) Lote 4 (1 item)------|
  |                            |
  |<--(6) EOF-------------------|
  |                            |
```

## 🚀 Como Testar

### 1. Iniciar o servidor

```bash
cd /Users/ecarvalho/Documents/GitHub/FC_PosGolang/13_gRPC_FC
go run cmd/grpcServer/main.go
```

### 2. Executar o cliente de teste

```bash
# Terminal 2
make test-stream
```

ou

```bash
go run cmd/testClient/main.go
```

### 3. Usar grpcurl (linha de comando)

```bash
grpcurl -plaintext \
  -d '{"name": "Teste", "description": "Streaming test"}' \
  localhost:50051 \
  pb.CategoryService.CreateCategoryStream
```

## 🐛 Troubleshooting

### Problema: Método não aparece na lista de serviços

**Causa**: Arquivos `.pb.go` não foram regenerados ou servidor está usando cache antigo.

**Solução**:

```bash
# 1. Limpar arquivos antigos
rm -f internal/pb/*.pb.go

# 2. Regenerar arquivos proto
protoc --go_out=. --go-grpc_out=. proto/course_category.proto

# 3. Limpar cache do Go
go clean -cache

# 4. Recompilar e executar
go build -o main cmd/grpcServer/main.go
./main
```

### Problema: `unknown method CreateCategoryStream`

**Causa**: Servidor foi iniciado antes dos arquivos `.pb.go` serem atualizados.

**Solução**: Reinicie o servidor após regenerar os arquivos proto.

## 📚 Referências

- [gRPC Streaming Concepts](https://grpc.io/docs/what-is-grpc/core-concepts/#server-streaming-rpc)
- [Protocol Buffers Guide](https://protobuf.dev/programming-guides/proto3/)

## 🎓 Aprendizados

1. **Server-Side Streaming** é ideal para enviar grandes volumes de dados progressivamente
2. A palavra-chave `stream` no proto define o tipo de streaming
3. O servidor controla quando enviar dados e quando finalizar o stream
4. O cliente recebe dados conforme são enviados, não precisa esperar tudo
5. `io.EOF` indica fim do stream no cliente
6. Sempre regenere os arquivos `.pb.go` após modificar o `.proto`
