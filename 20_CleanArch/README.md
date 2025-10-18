# 📊 Order System - Clean Architecture

Sistema de pedidos com Clean Architecture, REST, gRPC, GraphQL e Event-Driven Architecture.

## 🚀 Quick Start

```bash
make setup              # Instalar dependências
make docker-up          # Subir MySQL + RabbitMQ
make db-create db-migrate  # Criar banco
make run                # Iniciar aplicação
```

**Portas:** REST (8080) | gRPC (50051) | GraphQL (8082)

## ✅ Testar Endpoints

### Criar Orders

```bash
make test-http      # REST: POST /order
make test-grpc      # gRPC: CreateOrder
make test-graphql   # GraphQL: createOrder mutation
```

### Listar Orders

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

---

**GraphQL Playground:** <http://localhost:8082>
**RabbitMQ Management:** <http://localhost:15672> (guest/guest)
**Health Check:** <http://localhost:8080/health>
