# Build stage
FROM golang:1.24-alpine AS builder

# Instalar dependências necessárias
RUN apk add --no-cache git protobuf-dev

# Instalar ferramentas necessárias
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.1
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
RUN go install github.com/99designs/gqlgen@v0.17.76
RUN go install github.com/google/wire/cmd/wire@latest

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Gerar código protobuf e GraphQL
RUN protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
RUN go run github.com/99designs/gqlgen generate

# Compilar a aplicação
RUN go build -o ordersystem ./cmd/ordersystem

# Runtime stage
FROM alpine:3.22

# Instalar dependências de runtime e criar usuário não-root
RUN apk add --no-cache ca-certificates && \
    adduser -D -s /bin/sh appuser

# Definir diretório de trabalho
WORKDIR /app

# Copiar binário compilado
COPY --from=builder /app/ordersystem .

# Copiar arquivos de configuração
COPY --from=builder /app/env.example .env

# Mudar para usuário não-root
USER appuser

# Expor portas
EXPOSE 8080 50051 8082

# Comando de inicialização
CMD ["./ordersystem"]
