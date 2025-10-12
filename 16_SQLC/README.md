# 📊 SQLC - Geração de Código Go a partir de SQL

> Projeto demonstrando o uso do SQLC para gerar código Go type-safe a partir de queries SQL

## 📌 Sobre

Este projeto demonstra como usar o **SQLC** para gerar código Go automaticamente a partir de arquivos SQL, proporcionando:

- ✅ **Type Safety**: Código Go com tipos seguros baseados nas queries SQL
- ✅ **Performance**: Código otimizado sem reflection
- ✅ **Produtividade**: Elimina boilerplate de acesso a banco de dados
- ✅ **Manutenibilidade**: Queries centralizadas e versionadas

## 🏗️ Arquitetura

```bash
16_SQLC/
├── data/                    # Dados do MySQL (volume Docker)
├── internal/
│   └── db/                 # Código gerado pelo SQLC
│       ├── db.go          # Interface e struct Queries
│       ├── models.go      # Structs das tabelas
│       └── query.sql.go   # Funções geradas das queries
├── sql/
│   ├── migrations/         # Migrações do banco de dados
│   │   ├── 000001_init.up.sql
│   │   └── 000001_init.down.sql
│   └── queries/           # Queries SQL para SQLC
│       └── query.sql
├── docker-compose.yaml     # Configuração do MySQL
├── go.mod                  # Dependências Go
├── sqlc.yaml              # Configuração do SQLC
├── Makefile               # Comandos de desenvolvimento
└── README.md              # Este arquivo
```

## 🗄️ Schema do Banco

### Tabela: `categories`

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | varchar(36) | Chave primária (UUID) |
| `name` | text | Nome da categoria |
| `description` | text | Descrição da categoria |

### Tabela: `courses`

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | varchar(36) | Chave primária (UUID) |
| `category_id` | varchar(36) | FK para categories |
| `name` | text | Nome do curso |
| `description` | text | Descrição do curso |
| `price` | decimal(10,2) | Preço do curso |

## 🔧 Configuração

### 1. Instalar Dependências

```bash
# Instalar golang-migrate (MySQL)
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Ou via Homebrew (Mac)
brew install golang-migrate

# Instalar SQLC
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Ou via Homebrew (Mac)
brew install sqlc
```

### 2. Configuração do SQLC

**Arquivo: `sqlc.yaml`**

```yaml
version: "2"
sql:
  - schema: "sql/migrations"
    queries: "sql/queries"
    engine: "mysql"
    gen:
      go:
        package: "db"
        out: "internal/db"
```

### 3. Configurar Banco de Dados

```bash
# Subir o MySQL via Docker
docker compose up -d

# Verificar se está rodando
docker compose ps
```

### 4. Executar Migrações

```bash
# Aplicar migrações (criar tabelas)
migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" -verbose up

# Reverter migrações (remover tabelas)
migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" -verbose down

# Verificar status das migrações
migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" version
```

### 5. Gerar Código com SQLC

```bash
# Gerar código Go a partir das queries SQL
sqlc generate

# Verificar se as queries estão corretas
sqlc compile

# Gerar documentação das queries
sqlc doc
```

## 🚀 Comandos Úteis

### Docker & MySQL

```bash
# Iniciar MySQL
docker compose up -d

# Parar MySQL
docker compose down

# Acessar MySQL via CLI
docker compose exec mysql bash
mysql -uroot -p courses

# Ver logs do MySQL
docker compose logs mysql

# Resetar dados (cuidado!)
docker compose down -v
```

### Migrações

```bash
# Aplicar todas as migrações
migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" up

# Reverter última migração
migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" down 1

# Forçar versão específica
migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" force 1

# Verificar status
migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" version
```

### SQLC (quando configurado)

```bash
# Gerar código Go
sqlc generate

# Verificar queries
sqlc compile

# Gerar documentação
sqlc doc

# Verificar configuração
sqlc config
```

## ⚡ SQLC - Código Gerado

### 📁 Estrutura dos Arquivos Gerados

#### `internal/db/models.go`

```go
type Category struct {
    ID          string
    Name        string
    Description sql.NullString
}

type Course struct {
    ID          string
    CategoryID  string
    Name        string
    Description sql.NullString
    Price       string
}
```

#### `internal/db/db.go`

```go
type DBTX interface {
    ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
    PrepareContext(context.Context, string) (*sql.Stmt, error)
    QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
    QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Queries struct {
    db DBTX
}

func New(db DBTX) *Queries {
    return &Queries{db: db}
}
```

#### `internal/db/query.sql.go`

```go
func (q *Queries) ListCategories(ctx context.Context) ([]Category, error) {
    rows, err := q.db.QueryContext(ctx, listCategories)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var items []Category
    for rows.Next() {
        var i Category
        if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
            return nil, err
        }
        items = append(items, i)
    }
    return items, nil
}
```

### 🔧 Como Usar o Código Gerado

```go
package main

import (
    "context"
    "database/sql"
    "log"
    
    _ "github.com/go-sql-driver/mysql"
    "github.com/ElizCarvalho/FC_PosGolang/16_SQLC/internal/db"
)

func main() {
    // Conectar ao banco
    conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    
    // Criar instância do SQLC
    queries := db.New(conn)
    
    // Usar as funções geradas
    categories, err := queries.ListCategories(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    
    for _, category := range categories {
        log.Printf("Category: %s - %s", category.Name, category.Description.String)
    }
}
```

### 📋 Tipos de Queries SQLC

| Tipo | Descrição | Retorno | Exemplo |
|------|-----------|---------|---------|
| `:one` | Retorna um único registro | `(Model, error)` | `GetCategory` |
| `:many` | Retorna múltiplos registros | `([]Model, error)` | `ListCategories` |
| `:exec` | Executa sem retorno | `error` | `CreateCategory` |
| `:execrows` | Executa e retorna rows afetadas | `(sql.Result, error)` | `UpdateCategory` |

### 🎯 Vantagens do SQLC

- ✅ **Type Safety**: Detecta erros em tempo de compilação
- ✅ **Performance**: Código otimizado sem reflection
- ✅ **Produtividade**: Elimina boilerplate de acesso a dados
- ✅ **Manutenibilidade**: Queries centralizadas e versionadas
- ✅ **IntelliSense**: Autocomplete completo no IDE
- ✅ **Testabilidade**: Fácil de testar com mocks

## 📝 Exemplo de Uso

### 1. Criar Queries SQL

**Arquivo: `sql/queries/query.sql`**

```sql
-- name: ListCategories :many
SELECT * FROM categories;

-- name: GetCategory :one
SELECT * FROM categories WHERE id = ?;

-- name: CreateCategory :exec
INSERT INTO categories (id, name, description) VALUES (?, ?, ?);

-- name: UpdateCategory :exec
UPDATE categories SET name = ?, description = ? WHERE id = ?;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = ?;

-- name: ListCourses :many
SELECT * FROM courses;

-- name: GetCourse :one
SELECT * FROM courses WHERE id = ?;

-- name: CreateCourse :exec
INSERT INTO courses (id, category_id, name, description, price) VALUES (?, ?, ?, ?, ?);
```

### 2. Gerar Código Go

```bash
# Gerar código a partir das queries
sqlc generate

# Verificar se está tudo correto
sqlc compile
```

### 3. Criar Nova Migração

```bash
# Criar arquivo de migração
migrate create -ext sql -dir sql/migrations -seq add_users_table
```

### 4. Estrutura de Migração

**Arquivo: `000002_add_users_table.up.sql`**

```sql
CREATE TABLE users (
    id varchar(36) NOT NULL PRIMARY KEY,
    name text NOT NULL,
    email text UNIQUE NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP
);
```

**Arquivo: `000002_add_users_table.down.sql`**

```sql
DROP TABLE IF EXISTS users;
```

### 5. Aplicar Migração

```bash
migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" up
```

## 🔍 Verificações

### Verificar Tabelas no MySQL

```sql
-- Conectar ao MySQL
mysql -uroot -p courses

-- Listar tabelas
SHOW TABLES;

-- Descrever estrutura
DESCRIBE categories;
DESCRIBE courses;

-- Ver dados
SELECT * FROM categories;
SELECT * FROM courses;
```

### Verificar Migrações

```bash
# Status das migrações
migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" version

# Histórico de migrações
migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" version
```

## 🐛 Troubleshooting

### Problemas Comuns

1. **Erro de Conexão MySQL**

   ```bash
   # Verificar se container está rodando
   docker compose ps
   
   # Reiniciar container
   docker compose restart mysql
   ```

2. **Migração Falhou**

   ```bash
   # Verificar logs
   docker compose logs mysql
   
   # Forçar versão e tentar novamente
   migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" force 1
   ```

3. **Tabela já existe**

   ```bash
   # Reverter migrações
   migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" down
   
   # Aplicar novamente
   migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" up
   ```

## 📚 Próximos Passos

1. **Configurar SQLC** para gerar código Go
2. **Criar queries SQL** para operações CRUD
3. **Gerar código Go** com SQLC
4. **Implementar handlers** HTTP
5. **Criar testes** para as operações

## 🔗 Links Úteis

- [SQLC Documentation](https://docs.sqlc.dev/)
- [Golang Migrate](https://github.com/golang-migrate/migrate)
- [MySQL Docker Hub](https://hub.docker.com/_/mysql)
- [Go Database/SQL](https://pkg.go.dev/database/sql)

---

Desenvolvido com ❤️ para o curso Full Cycle Go
