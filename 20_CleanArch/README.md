# 📊 Order System - Clean Architecture

Sistema de pedidos com Clean Architecture, REST, gRPC, GraphQL e Event-Driven Architecture.

## 🚀 Executar o Desafio

**Comando único para executar tudo:**

```bash
make up
```

> ✅ **Tudo automático**: MySQL + RabbitMQ + Aplicação Go

## 🧪 Testar a Aplicação

### Teste Rápido (Todas as APIs)

```bash
make test-api       # Testa REST + Health Check
make test-all       # Testa REST + gRPC + GraphQL
```

### Testes Individuais

```bash
make test-rest      # Testa REST API
make test-grpc      # Testa gRPC
make test-graphql   # Testa GraphQL
curl http://localhost:8080/health  # Health check
```

### GraphQL Playground

Acesse: <http://localhost:8082>

## 🌐 Endpoints Disponíveis

| Serviço | Porta | URL |
|---------|-------|-----|
| **REST API** | 8080 | http://localhost:8080 |
| **GraphQL** | 8082 | http://localhost:8082 |
| **gRPC** | 50051 | localhost:50051 |

## 📋 Arquivo de Testes

O projeto inclui `api.http` com requisições prontas para testar todas as funcionalidades.

## 🛠️ Comandos Úteis

```bash
make up             # Executar tudo
make down           # Parar tudo
make dev            # Desenvolvimento local
make test-api       # Testar REST + Health
make test-all       # Testar REST + gRPC + GraphQL
make test-rest      # Testar apenas REST
make test-grpc      # Testar apenas gRPC
make test-graphql   # Testar apenas GraphQL
make test           # Executar testes unitários
make help           # Ver todos os comandos
```
