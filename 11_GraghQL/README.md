# GraphQL

Usamos o https://gqlgen.com/ para gerar o codigo do graphql

## Como usar

```bash
printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
```

```bash
go mod tidy
```

## Inicializa o gqlgen config e gera os models

```bash
go run github.com/99designs/gqlgen init
```

## Inicia o servidor graphql

```bash
go run server.go
```

## Ajustar o projeto de acordo com o schema

```bash
go run github.com/99designs/gqlgen generate
```
