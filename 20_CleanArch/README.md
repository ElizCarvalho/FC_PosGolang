# 📊 Order System - Clean Architecture

> Sistema de pedidos implementado com Clean Architecture, gRPC, GraphQL e Event-Driven Architecture

## 📌 Sobre

Este projeto implementa um sistema de pedidos seguindo os princípios da Clean Architecture, com múltiplas interfaces (REST API, gRPC, GraphQL) e sistema de eventos usando RabbitMQ.

## 🔧 Configuração

### Variáveis de Ambiente

Crie um arquivo `.env` baseado no `env.example`:

```bash
cp env.example .env
```

Configure as seguintes variáveis:

```env
# Database Configuration
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=orders

# Server Ports
WEB_SERVER_PORT=8080
GRPC_SERVER_PORT=50051
GRAPHQL_SERVER_PORT=8082

# RabbitMQ Configuration
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
```

### Dependências

```bash
# Instalar dependências
go mod tidy

# Instalar ferramentas
go install github.com/99designs/gqlgen@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/google/wire/cmd/wire@latest
```

### Banco de Dados

```sql
CREATE DATABASE orders;
CREATE TABLE orders (
    id varchar(255) NOT NULL,
    price float NOT NULL,
    tax float NOT NULL,
    final_price float NOT NULL,
    PRIMARY KEY (id)
);
```

### RabbitMQ

```bash
# Usando Docker
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

## 🚀 Execução

```bash
# Executar a aplicação
go run ./cmd/ordersystem

# Ou compilar e executar
go build ./cmd/ordersystem
./ordersystem
```

## 📚 APIs

### REST API

#### POST /order
Criar um novo pedido

```bash
curl -X POST http://localhost:8080/order \
  -H "Content-Type: application/json" \
  -d '{
    "id": "123",
    "price": 100.0,
    "tax": 10.0
  }'
```

### gRPC

#### CreateOrder
```protobuf
service OrderService {
  rpc CreateOrder(CreateOrderRequest) returns (CreateOrderResponse);
}
```

### GraphQL

Acesse o playground em: `http://localhost:8082`

#### Mutation
```graphql
mutation {
  createOrder(input: {
    id: "123"
    Price: 100.0
    Tax: 10.0
  }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

## 🧪 Testes

```bash
# Executar todos os testes
go test ./... -v

# Executar com cobertura
go test ./... -cover

# Executar testes específicos
go test ./internal/entity -v
go test ./internal/usecase -v
go test ./internal/infra/web -v
go test ./events -v
```

## 🏗️ Arquitetura

```
cmd/
├── ordersystem/          # Entry point da aplicação
internal/
├── entity/               # Entidades de domínio
├── usecase/              # Casos de uso
├── infra/                # Infraestrutura
│   ├── database/         # Repositórios
│   ├── grpc/            # Serviços gRPC
│   ├── graph/           # GraphQL resolvers
│   └── web/             # Handlers HTTP
└── event/               # Eventos de domínio
events/                  # Sistema de eventos
configs/                 # Configurações
```

## 🔄 Event-Driven Architecture

O sistema utiliza eventos para comunicação assíncrona:

- **OrderCreated**: Disparado quando um pedido é criado
- **RabbitMQ**: Broker de mensagens para processamento assíncrono

## 📝 Documentação

- **Swagger**: Disponível em `/swagger/` (quando implementado)
- **GraphQL Playground**: `http://localhost:8082`
- **RabbitMQ Management**: `http://localhost:15672` (guest/guest)

## 🛠️ Desenvolvimento

### Regenerar Código

```bash
# GraphQL
gqlgen generate

# gRPC
protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

# Wire (Dependency Injection)
wire ./cmd/ordersystem
```

### Estrutura de Commits

Seguindo Conventional Commits:

```
feat: add new feature
fix: fix bug
refactor: refactor code
docs: update documentation
test: add tests
chore: maintenance
```

## 📄 Licença

Este projeto é parte do curso de Pós-Graduação em Go da Full Cycle.
