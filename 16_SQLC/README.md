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
├── sql/
│   └── migrations/         # Migrações do banco de dados
│       ├── 000001_init.up.sql
│       └── 000001_init.down.sql
├── docker-compose.yaml     # Configuração do MySQL
├── go.mod                  # Dependências Go
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
# Instalar golang-migrate (SQLite)
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Ou via Homebrew (Mac)
brew install golang-migrate

# Instalar SQLC
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### 2. Configurar Banco de Dados

```bash
# Subir o MySQL via Docker
docker compose up -d

# Verificar se está rodando
docker compose ps
```

### 3. Executar Migrações

```bash
# Aplicar migrações (criar tabelas)
migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" -verbose up

# Reverter migrações (remover tabelas)
migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" -verbose down

# Verificar status das migrações
migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" version
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

## 📝 Exemplo de Uso

### 1. Criar Nova Migração

```bash
# Criar arquivo de migração
migrate create -ext sql -dir sql/migrations -seq add_users_table
```

### 2. Estrutura de Migração

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

### 3. Aplicar Migração

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
