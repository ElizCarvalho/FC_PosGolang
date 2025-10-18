# 📊 Order System - Clean Architecture

> Sistema de pedidos implementado com Clean Architecture, gRPC, GraphQL e Event-Driven Architecture

## 📌 Sobre

Sistema de gerenciamento de pedidos que demonstra a aplicação de Clean Architecture em Go, oferecendo três interfaces diferentes para criação de pedidos:

- **REST API** (HTTP/JSON)
- **gRPC** (Protocol Buffers)
- **GraphQL** (Queries e Mutations)

Todos os pedidos criados são processados de forma assíncrona através de eventos publicados no RabbitMQ.

## 🚀 Quick Start

### 1. Setup Inicial

```bash
# Configurar ambiente e instalar dependências
make setup

# Subir infraestrutura (MySQL + RabbitMQ)
make docker-up

# Criar banco e tabelas
make db-create db-migrate
```

### 2. Executar a Aplicação

```bash
# Iniciar a aplicação
make run
```

A aplicação irá subir em três portas:

- **REST API**: <http://localhost:8080>
- **gRPC**: localhost:50051
- **GraphQL**: <http://localhost:8082>

## ✅ Validando as 3 Interfaces

### Teste REST API

```bash
make test-http
```

Ou manualmente:

```bash
curl -X POST http://localhost:8080/order \
  -H "Content-Type: application/json" \
  -d '{"id":"order-001","price":100.0,"tax":10.0}'
```

**Resposta esperada:**

```json
{
  "id": "order-001",
  "price": 100.0,
  "tax": 10.0,
  "final_price": 110.0
}
```

### Teste gRPC

```bash
make test-grpc
```

Ou com `grpcurl`:

```bash
grpcurl -plaintext -d '{"id":"order-002","price":200.0,"tax":20.0}' \
  localhost:50051 pb.OrderService/CreateOrder
```

**Resposta esperada:**

```json
{
  "id": "order-002",
  "price": 200,
  "tax": 20,
  "final_price": 220
}
```

### Teste GraphQL

```bash
make test-graphql
```

Ou acesse o playground em <http://localhost:8082> e execute:

```graphql
mutation {
  createOrder(input: {
    id: "order-003"
    Price: 300.0
    Tax: 30.0
  }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

**Resposta esperada:**

```json
{
  "data": {
    "createOrder": {
      "id": "order-003",
      "Price": 300,
      "Tax": 30,
      "FinalPrice": 330
    }
  }
}
```

## 🧪 Testes Automatizados

```bash
# Executar todos os testes
make test

# Testes com cobertura
make test-coverage

# Testes por camada
make test-entity      # Entidades de domínio
make test-usecase     # Casos de uso
make test-web         # Handlers HTTP
make test-events      # Sistema de eventos
```

## 🔧 Configuração

### Variáveis de Ambiente

Copie o arquivo de exemplo:

```bash
cp env.example cmd/ordersystem/.env
```

Principais configurações:

```env
# Database
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

# RabbitMQ
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
```

## 🏗️ Arquitetura

O projeto segue **Clean Architecture** com 3 camadas:

- **Domain** (`internal/entity`): Entidades e regras de negócio
- **Use Cases** (`internal/usecase`): Lógica de aplicação
- **Infrastructure** (`internal/infra`): REST, gRPC, GraphQL, Database

Eventos são processados de forma assíncrona via RabbitMQ (`events/`).

## 🛠️ Comandos Úteis

```bash
# Desenvolvimento
make run-dev          # Sobe infra + aplicação
make test             # Roda todos os testes
make test-coverage    # Testes com cobertura

# Infraestrutura
make docker-up        # Sobe MySQL + RabbitMQ
make docker-down      # Para containers
make db-reset         # Reseta banco de dados

# Validação
make test-http        # Testa REST API
make test-grpc        # Testa gRPC
make test-graphql     # Testa GraphQL

# Geração de código
make proto            # Gera código protobuf
make graphql          # Gera código GraphQL
make wire             # Gera injeção de dependências

# Limpeza
make clean            # Remove binários
make clean-all        # Remove tudo (binários + containers)
```

## 📝 Links Úteis

- **GraphQL Playground**: <http://localhost:8082>
- **RabbitMQ Management**: <http://localhost:15672> (guest/guest)
- **Makefile**: `make help` para ver todos os comandos

## 📄 Licença

Este projeto é parte do curso de Pós-Graduação em Go da Full Cycle.
