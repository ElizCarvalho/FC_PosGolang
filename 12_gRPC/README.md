# 🧮 Calculadora gRPC

> Exemplo de implementação de serviço gRPC em Go com operações matemáticas básicas

## 📌 Sobre

Este projeto demonstra como implementar um serviço gRPC em Go, incluindo:
- Definição de serviços usando Protocol Buffers
- Implementação de servidor gRPC
- Cliente gRPC para consumir os serviços
- Tratamento de erros e validações
- Estrutura de projeto organizada

## 🔧 Configuração

### Pré-requisitos
- Go 1.24+
- Protocol Buffers compiler (protoc)
- Plugins do Go para protobuf

### Instalação
```bash
# Configurar o ambiente
make setup

# Gerar código Go a partir dos arquivos .proto
make proto
```

## 🚀 Como Usar

### 1. Iniciar o Servidor
```bash
make run-server
```

### 2. Executar o Cliente (em outro terminal)
```bash
make run-client
```

## 📚 API

### Serviços Disponíveis

#### CalculatorService
- **Add**: Soma dois números
- **Subtract**: Subtrai dois números  
- **Multiply**: Multiplica dois números
- **Divide**: Divide dois números (com validação de divisão por zero)
- **SquareRoot**: Calcula a raiz quadrada (com validação de números negativos)

### Exemplo de Uso

```go
// Conectar ao servidor
conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
client := pb.NewCalculatorServiceClient(conn)

// Fazer uma operação
resp, err := client.Add(ctx, &pb.AddRequest{A: 10, B: 5})
fmt.Printf("Resultado: %.2f\n", resp.Result)
```

## 🧪 Testes

```bash
# Executar testes
make test

# Verificar código
make lint
```

## 📝 Comandos Disponíveis

```bash
make help              # Mostra todos os comandos
make setup             # Configura o ambiente
make proto             # Gera código Go a partir dos .proto
make run-server        # Executa o servidor
make run-client        # Executa o cliente
make test              # Executa os testes
make lint              # Verifica o código
make clean             # Limpa arquivos gerados
```

## 🏗️ Estrutura do Projeto

```
12_gRPC/
├── proto/                 # Arquivos .proto
│   ├── calculator.proto
│   ├── calculator.pb.go
│   └── calculator_grpc.pb.go
├── server/               # Implementação do servidor
│   └── main.go
├── client/               # Implementação do cliente
│   └── main.go
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 🔍 Protocol Buffers

O arquivo `calculator.proto` define:
- Serviço `CalculatorService` com 5 operações
- Mensagens de request e response para cada operação
- Tratamento de erros através de campos opcionais

## 🛠️ Desenvolvimento

### Adicionando Novas Operações

1. Edite o arquivo `proto/calculator.proto`
2. Execute `make proto` para gerar o código
3. Implemente a nova operação no servidor
4. Atualize o cliente para testar

### Exemplo de Nova Operação

```protobuf
// No calculator.proto
rpc Power(PowerRequest) returns (PowerResponse);

message PowerRequest {
  double base = 1;
  double exponent = 2;
}

message PowerResponse {
  double result = 1;
}
```

## 📖 Documentação

- [gRPC Go Quick Start](https://grpc.io/docs/languages/go/quickstart/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)
- [Go gRPC Documentation](https://pkg.go.dev/google.golang.org/grpc)
