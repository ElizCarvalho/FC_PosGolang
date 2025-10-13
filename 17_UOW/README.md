# 📊 Unit of Work (UOW) Pattern

> Implementação do padrão Unit of Work em Go para gerenciar transações de banco de dados

## 📌 Sobre

Este projeto demonstra a implementação do padrão **Unit of Work (UOW)** em Go, que é uma técnica para gerenciar transações de banco de dados de forma eficiente e consistente. O UOW permite agrupar múltiplas operações de banco de dados em uma única transação, garantindo atomicidade.

## 🔧 Configuração

### Pré-requisitos

- Go 1.19+
- Docker e Docker Compose
- MySQL 5.7+

### Instalação

1.**Clone o repositório**:

```bash
git clone https://github.com/ElizCarvalho/FC_PosGolang.git
cd FC_PosGolang/17_UOW
```

2.**Instale as dependências**:

```bash
go mod tidy
```

3.**Configure o banco de dados**:

```bash
# Suba o MySQL via Docker
docker-compose up -d

# Execute as migrações
mysql -h localhost -P 3306 -u root -proot courses < sql/schema.sql
```

4.**Gere o código SQLC** (se necessário):

```bash
sqlc generate
```

## 🚀 Executando

### Testes

```bash
# Execute todos os testes
go test ./...

# Execute testes específicos
go test ./internal/usecase/...

# Execute com verbose
go test -v ./internal/usecase/...
```

### Aplicação

```bash
# Execute a aplicação principal
go run cmd/main.go
```

## 📚 Estrutura do Projeto

```bash
17_UOW/
├── cmd/                    # Aplicação principal
├── internal/
│   ├── db/                # Configuração do banco
│   ├── entity/            # Entidades do domínio
│   ├── repository/        # Camada de repositório
│   └── usecase/           # Casos de uso
├── pkg/
│   └── uow/              # Implementação do UOW
├── sql/                  # Scripts SQL
│   ├── schema.sql        # Schema do banco
│   └── queries.sql       # Queries SQLC
└── docker-compose.yaml   # Configuração do MySQL
```

## 🧪 Testes

### Teste sem UOW

```bash
go test -v -run TestAddCourse
```

### Teste com UOW

```bash
go test -v -run TestAddCourseUow
```

## 📝 Padrão Unit of Work

O **Unit of Work** é um padrão que:

1. **Agrupa operações**: Múltiplas operações de banco em uma transação
2. **Garante atomicidade**: Todas as operações são commitadas ou rollbackadas juntas
3. **Melhora performance**: Reduz o número de round-trips ao banco
4. **Facilita manutenção**: Centraliza o gerenciamento de transações

### Exemplo de Uso

```go
// Sem UOW - operações separadas
func AddCourseWithoutUOW(course Course) error {
    // Inserir categoria
    categoryID, err := categoryRepo.Create(course.Category)
    if err != nil {
        return err
    }
    
    // Inserir curso
    course.CategoryID = categoryID
    return courseRepo.Create(course)
}

// Com UOW - operações em transação
func AddCourseWithUOW(course Course) error {
    return uow.Do(ctx, func(uow *uow.UowInterface) error {
        // Inserir categoria
        categoryID, err := uow.CategoryRepository.Create(course.Category)
        if err != nil {
            return err
        }
        
        // Inserir curso
        course.CategoryID = categoryID
        return uow.CourseRepository.Create(course)
    })
}
```

## 🔍 Funcionalidades

- ✅ Implementação do padrão Unit of Work
- ✅ Gerenciamento de transações
- ✅ Repositórios com interface
- ✅ Testes unitários e de integração
- ✅ SQLC para geração de código SQL
- ✅ Docker para ambiente de desenvolvimento

## 📚 Documentação

Mais detalhes sobre o padrão Unit of Work podem ser encontrados em:

- [Martin Fowler - Unit of Work](https://martinfowler.com/eaaCatalog/unitOfWork.html)
- [Go Design Patterns](https://github.com/tmrts/go-patterns)

## 🛠️ Tecnologias

- **Go 1.19+**
- **MySQL 5.7**
- **SQLC** - Geração de código SQL
- **Docker** - Containerização
- **Testify** - Framework de testes
