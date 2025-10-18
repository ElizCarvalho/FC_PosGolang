# 📊 Order System - Clean Architecture

Sistema de pedidos com Clean Architecture, REST, gRPC, GraphQL e Event-Driven Architecture.

## 🚀 Passos para Executar o Desafio

### 1. Instalar Dependências

```bash
make setup
```

### 2. Subir Infraestrutura (Docker)

```bash
make docker-up
```

> **💡 Nota:** O comando `docker compose up` sobe MySQL e RabbitMQ. As migrations SQL são executadas **automaticamente** na primeira vez que o MySQL inicializa.

### 3. Iniciar a Aplicação

```bash
make run
```

## 🌐 Portas dos Serviços

| Serviço | Porta | Endpoint |
|---------|-------|----------|
| **REST API** | **8080** | <http://localhost:8080> |
| **gRPC** | **50051** | localhost:50051 |
| **GraphQL** | **8082** | <http://localhost:8082> |
| GraphQL Playground | 8082 | <http://localhost:8082> |
| MySQL | 3306 | localhost:3306 |
| RabbitMQ | 5672 | localhost:5672 |
| RabbitMQ Management | 15672 | <http://localhost:15672> |

## ✅ Testar Endpoints

### Usando o arquivo api.http

O projeto inclui o arquivo **`api.http`** na raiz com todas as requisições prontas para testar:

- ✅ Criar e listar orders via **REST**
- ✅ Criar e listar orders via **gRPC** (comandos grpcurl)
- ✅ Criar e listar orders via **GraphQL**
- ✅ Health check

### Via Makefile

#### Criar Orders

```bash
make test-http      # REST: POST /order
make test-grpc      # gRPC: CreateOrder
make test-graphql   # GraphQL: createOrder mutation
```

#### Listar Orders

```bash
make test-http-list    # REST: GET /orders
make test-grpc-list    # gRPC: ListOrders
make test-graphql-list # GraphQL: listOrders query
make test-all-list     # Testa todas as interfaces
```

## 🧪 Testes

```bash
make test           # Todos os testes
make test-coverage  # Com cobertura
```

## 📝 Comandos Úteis

```bash
make help          # Ver todos os comandos
make docker-up     # Subir MySQL + RabbitMQ
make docker-down   # Parar containers
make proto         # Gerar código protobuf
make graphql       # Gerar código GraphQL
```

## 📦 Migrations

As migrations SQL estão em **`sql/migrations/`** e são executadas automaticamente quando o MySQL sobe pela primeira vez.

**Arquivos:**

- `001_create_orders_table.sql` - Cria a tabela `orders`

**Comandos úteis:**

```bash
make db-reset      # Reseta o banco (apaga e recria)
make db-migrate    # Executa migrations manualmente
```

---

## 🔗 Links Úteis

- **GraphQL Playground:** <http://localhost:8082>
- **RabbitMQ Management:** <http://localhost:15672> (guest/guest)
- **Health Check:** <http://localhost:8080/health>
