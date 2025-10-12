# 🔄 gRPC Bidirectional Streaming

## 🎯 O que foi implementado

Foi implementado o método `CreateCategoryStreamBidirectional` que demonstra **Bidirectional Streaming** no gRPC.

## 🧠 Conceito: Bidirectional Streaming

### O que é?

- **Padrão**: Cliente e servidor enviam MÚLTIPLAS mensagens simultaneamente
- **Fluxo**: Bidirecional (cliente ↔ servidor)
- **Independência**: Cliente e servidor podem enviar/receber independentemente
- **Assíncrono**: Não precisa esperar resposta para enviar próxima mensagem

### Quando usar?

- Chat em tempo real
- Upload de arquivos com feedback de progresso
- Processamento de stream de dados com confirmações
- Jogos multiplayer
- Colaboração em tempo real (Google Docs style)

### Diferença dos outros tipos de streaming

| Tipo | Quem Envia | Quem Recebe | Exemplo |
|------|-----------|-------------|---------|
| **Unary** | Cliente (1x) | Servidor (1x) | Login simples |
| **Server Streaming** | Cliente (1x) | Servidor (Nx) | Baixar lista de produtos |
| **Client Streaming** | Cliente (Nx) | Servidor (1x) | Upload de arquivo |
| **Bidirectional** | Ambos (Nx) | Ambos (Nx) | Chat em tempo real |

## 🏗️ Arquitetura da Implementação

### 1. **Definição no Proto** (`proto/course_category.proto`)

```protobuf
service CategoryService {
    // Unary RPC
    rpc CreateCategory(CreateCategoryRequest) returns (CategoryResponse) {}
    
    // Server Streaming RPC
    rpc CreateCategoryStream(CreateCategoryRequest) returns (stream CategoryList) {}
    
    // Bidirectional Streaming RPC
    rpc CreateCategoryStreamBidirectional(stream CreateCategoryRequest) returns (stream Category) {}
}
```

**Observação**:

- `stream` antes do **parâmetro** = Cliente envia múltiplas mensagens
- `stream` antes do **retorno** = Servidor envia múltiplas mensagens
- `stream` nos **dois** = Bidirectional!

### 2. **Service Layer** (`internal/service/category.go`)

```go
func (c *CategoryService) CreateCategoryStreamBidirectional(
    stream grpc.BidiStreamingServer[pb.CreateCategoryRequest, pb.Category],
) error {
    for {
        // 1. Recebe a próxima requisição do cliente
        req, err := stream.Recv()
        if err == io.EOF {
            // Cliente terminou de enviar
            return nil
        }
        if err != nil {
            return err
        }

        // 2. Processa a requisição
        category, err := c.CategoryDB.Create(req.Name, req.Description)
        if err != nil {
            return err
        }

        // 3. Envia a resposta de volta ao cliente IMEDIATAMENTE
        err = stream.Send(&pb.Category{
            Id:          category.ID,
            Name:        category.Name,
            Description: category.Description,
        })
        if err != nil {
            return err
        }
    }
}
```

**Pontos-chave:**

- `grpc.BidiStreamingServer[Request, Response]`: Interface para streaming bidirecional
- `stream.Recv()`: Recebe a próxima mensagem do cliente (bloqueia até receber ou EOF)
- `stream.Send()`: Envia mensagem para o cliente
- `io.EOF`: Cliente chamou `CloseSend()`, indicando fim do envio
- **Loop infinito**: Continua recebendo até o cliente fechar o stream

### 3. **Cliente** (`cmd/testBidiClient/main.go`)

```go
// 1. Abre o stream bidirectional
stream, err := client.CreateCategoryStreamBidirectional(context.Background())

// 2. Goroutine para RECEBER respostas (independente)
go func() {
    for {
        category, err := stream.Recv()
        if err == io.EOF {
            // Servidor terminou de enviar
            done <- true
            return
        }
        if err != nil {
            log.Fatalf("Erro: %v", err)
        }
        
        // Processa a categoria recebida
        fmt.Printf("Recebido: %s\n", category.Name)
    }
}()

// 3. ENVIA múltiplas requisições (main goroutine)
for _, cat := range categories {
    err := stream.Send(&pb.CreateCategoryRequest{
        Name:        cat.name,
        Description: cat.description,
    })
    if err != nil {
        log.Fatalf("Erro: %v", err)
    }
}

// 4. Fecha o stream de envio
stream.CloseSend()

// 5. Aguarda todas as respostas
<-done
```

**Pontos-chave:**

- **Goroutine para receber**: Permite processar respostas enquanto envia novas requisições
- `stream.Send()`: Envia requisição ao servidor
- `stream.Recv()`: Recebe resposta do servidor
- `stream.CloseSend()`: Informa ao servidor que cliente terminou de enviar
- **Canal `done`**: Sincroniza o fim do recebimento

## 📊 Fluxo de Execução

```mermaid
Cliente                                 Servidor
  |                                        |
  |----(1) Abre stream bidirectional----->|
  |                                        |
  |----(2) Envia: "Backend"-------------->|
  |                                        | [Cria categoria]
  |<---(3) Retorna: "Backend" (ID: 123)---|
  |                                        |
  |----(4) Envia: "Frontend"------------->|
  |                                        | [Cria categoria]
  |<---(5) Retorna: "Frontend" (ID: 456)--|
  |                                        |
  |----(6) Envia: "DevOps"--------------->|
  |                                        | [Cria categoria]
  |<---(7) Retorna: "DevOps" (ID: 789)----|
  |                                        |
  |----(8) CloseSend() [EOF]------------->|
  |                                        | [Retorna nil]
  |<---(9) EOF----------------------------|
  |                                        |
```

**Observações:**

- Cliente e servidor operam **independentemente**
- Servidor responde **imediatamente** após processar cada requisição
- Não precisa esperar todas as requisições para começar a responder
- Cliente pode continuar enviando enquanto processa respostas

## 🚀 Como Testar

### 1. Iniciar o servidor

```bash
cd /Users/ecarvalho/Documents/GitHub/FC_PosGolang/13_gRPC_FC
go run cmd/grpcServer/main.go
```

### 2. Executar o cliente de teste

```bash
# Terminal 2
make test-bidi
```

ou

```bash
go run cmd/testBidiClient/main.go
```

### 3. Resultado esperado

```bash
🔄 Testando CreateCategoryStreamBidirectional...
=================================================
📤 [1] Enviando: Backend
📥 [1] Recebido: Backend (ID: xxx)
📤 [2] Enviando: Frontend
📥 [2] Recebido: Frontend (ID: xxx)
...
🔒 Cliente terminou de enviar requisições
✅ Servidor terminou de enviar respostas
=================================================
✅ Bidirectional streaming concluído!
```

## 🔍 Comparação: Server vs Bidirectional Streaming

### Server Streaming (`CreateCategoryStream`)

```mermaid
Cliente → (1 requisição) → Servidor
        ← (N respostas)  ←
```

- Cliente envia 1 mensagem
- Servidor envia N mensagens
- **Uso**: Baixar dados, listagens paginadas

### Bidirectional Streaming (`CreateCategoryStreamBidirectional`)

```mermaid
Cliente ⇄ (N requisições/respostas) ⇄ Servidor
```

- Cliente envia N mensagens
- Servidor envia N mensagens
- **Uso**: Chat, colaboração em tempo real, upload com feedback

## 🐛 Troubleshooting

### Problema: Goroutine nunca termina

**Causa**: Cliente não chamou `CloseSend()`, servidor fica esperando mais mensagens.

**Solução**: Sempre chamar `stream.CloseSend()` após enviar todas as mensagens.

### Problema: Deadlock no cliente

**Causa**: Cliente tenta receber na mesma goroutine que envia, bloqueando o fluxo.

**Solução**: Use goroutines separadas para enviar e receber:

```go
// ✅ Correto: Goroutines separadas
go func() {
    // Recebe mensagens
    for { stream.Recv() }
}()

// Envia mensagens
for { stream.Send() }
```

```go
// ❌ Errado: Mesma goroutine
for {
    stream.Send()  // Envia
    stream.Recv()  // Recebe (pode bloquear!)
}
```

### Problema: `io.EOF` não é recebido

**Causa**:

- Cliente: Servidor retornou erro antes de fechar
- Servidor: Cliente não chamou `CloseSend()`

**Solução**: Sempre verificar erros antes de verificar `EOF`:

```go
msg, err := stream.Recv()
if err != nil {
    if err == io.EOF {
        // Stream fechado normalmente
    } else {
        // Erro real
        return err
    }
}
```

## 🎓 Conceitos Importantes

### 1. **Goroutines são essenciais**

No cliente, você **PRECISA** de goroutines separadas para enviar e receber, caso contrário pode causar deadlock.

### 2. **CloseSend() é obrigatório**

Sempre chame `CloseSend()` quando terminar de enviar. Isso envia `EOF` para o servidor.

### 3. **Ordem não é garantida**

Em cenários mais complexos, as respostas podem chegar fora de ordem. Use IDs para correlacionar requisições e respostas.

### 4. **Backpressure**

Se o servidor processa mais rápido que o cliente envia, ou vice-versa, o gRPC gerencia automaticamente com backpressure.

## 📚 Casos de Uso Reais

### 1. **Chat em Tempo Real**

```protobuf
rpc Chat(stream ChatMessage) returns (stream ChatMessage) {}
```

- Usuários enviam mensagens
- Servidor distribui para todos os participantes

### 2. **Upload com Progresso**

```protobuf
rpc UploadFile(stream FileChunk) returns (stream UploadProgress) {}
```

- Cliente envia chunks do arquivo
- Servidor retorna progresso (%, tempo restante)

### 3. **Game Server**

```protobuf
rpc GameSession(stream PlayerAction) returns (stream GameState) {}
```

- Cliente envia ações do jogador
- Servidor envia estado atualizado do jogo

## 🔗 Referências

- [gRPC Streaming Guide](https://grpc.io/docs/what-is-grpc/core-concepts/#bidirectional-streaming-rpc)
- [gRPC Go Tutorial](https://grpc.io/docs/languages/go/basics/)

## 🎯 Resumo

**Bidirectional Streaming** é o tipo mais poderoso e complexo de streaming no gRPC:

- ✅ Comunicação em tempo real
- ✅ Independência total entre cliente e servidor
- ✅ Performance otimizada (não espera requisição terminar)
- ⚠️ Requer gerenciamento cuidadoso de goroutines
- ⚠️ Sempre feche o stream com `CloseSend()`
