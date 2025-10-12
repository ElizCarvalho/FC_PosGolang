# 🚀 CLI com Cobra - Comandos Encadeados

Uma CLI completa desenvolvida em Go usando a biblioteca Cobra, demonstrando comandos encadeados e validação de flags/argumentos.

## 📋 Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Estrutura de Comandos](#estrutura-de-comandos)
- [Comandos Encadeados](#comandos-encadeados)
- [Instalação e Uso](#instalação-e-uso)
- [Exemplos de Uso](#exemplos-de-uso)
- [Testes](#testes)
- [Makefile](#makefile)
- [Estrutura do Projeto](#estrutura-do-projeto)

## 🎯 Sobre o Projeto

Este projeto demonstra como criar uma CLI robusta em Go usando a biblioteca Cobra, com foco em:

- **Comandos encadeados** (subcomandos de subcomandos)
- **Validação de argumentos** e flags
- **Testes unitários** para comandos
- **Estrutura organizada** seguindo boas práticas

## 🔗 Estrutura de Comandos

```
course-cli
├── category (comandos de categoria)
│   ├── create [name] [description]
│   ├── list
│   ├── get [id]
│   ├── update [id] [name] [description]
│   └── delete [id]
├── ping (comando simples com flag)
│   └── --pong (flag para retornar "pong pong")
├── project (comando principal)
│   └── task (subcomando de project)
│       ├── add [description]
│       ├── list
│       └── complete [id]
├── config (comandos de configuração)
│   ├── set --key [key] --value [value] (flags locais)
│   ├── get --key [key] (flags locais)
│   ├── list --verbose (flag global)
│   └── reset --force --verbose (flags locais + global)
```

## 🏷️ Flags Locais vs Globais

### Flags Locais (Local Flags)

**Características:**
- Aplicam-se **apenas ao comando específico**
- Não são herdadas por subcomandos
- Cada comando tem suas próprias flags

**Exemplo:**
```go
// Flags locais para um comando específico
configSetCmd.Flags().String("key", "", "Chave da configuração")
configSetCmd.Flags().String("value", "", "Valor da configuração")
```

**Uso:**
```bash
# --key e --value são locais do comando 'set'
course-cli config set --key "debug_mode" --value "true"
```

### Flags Globais (Persistent Flags)

**Características:**
- Aplicam-se ao **comando e todos os seus subcomandos**
- São herdadas automaticamente
- Úteis para configurações gerais

**Exemplo:**
```go
// Flag global --verbose - disponível em TODOS os subcomandos
configCmd.PersistentFlags().Bool("verbose", false, "Modo verboso")
```

**Uso:**
```bash
# --verbose funciona em TODOS os subcomandos de config
course-cli config list --verbose
course-cli config set --key "test" --value "value" --verbose
```

### Comparação

| Tipo | Escopo | Herança | Exemplo |
|------|--------|---------|---------|
| **Local** | Apenas o comando | ❌ Não | `--key`, `--value`, `--force` |
| **Global** | Comando + subcomandos | ✅ Sim | `--verbose`, `--debug` |

### Quando Usar

**Use Flags Locais quando:**
- A flag é específica de um comando
- Não faz sentido em outros comandos
- Exemplo: `--force` só para reset, `--key` só para set/get

**Use Flags Globais quando:**
- A flag é útil em vários comandos
- É uma configuração geral
- Exemplo: `--verbose`, `--debug`, `--output-format`

## 🔗 Comandos Encadeados

### O que são Comandos Encadeados?

Comandos encadeados são uma hierarquia de comandos onde você tem **subcomandos de subcomandos**, criando uma estrutura em árvore.

### Exemplo Prático - Sistema de Projetos

```bash
# Comando de 1º nível
course-cli project

# Comando de 2º nível (encadeado)
course-cli project task

# Comandos de 3º nível (encadeados)
course-cli project task add "Implementar login"
course-cli project task list
course-cli project task complete 1
```

### Como Funciona a Hierarquia

1. **`project`** - Comando principal
2. **`task`** - Subcomando de `project`
3. **`add/list/complete`** - Subcomandos de `task`

### Implementação em Código

```go
// Comando principal
projectCmd := &cobra.Command{
    Use:   "project",
    Short: "Gerenciar projetos",
}

// Subcomando de project
taskCmd := &cobra.Command{
    Use:   "task",
    Short: "Gerenciar tarefas do projeto",
}

// Subcomandos de task
taskAddCmd := &cobra.Command{
    Use:   "add [description]",
    Short: "Adicionar nova tarefa",
    Args:  cobra.ExactArgs(1),
}

// Encadeamento
projectCmd.AddCommand(taskCmd)
taskCmd.AddCommand(taskAddCmd)
```

### Exemplos Reais de Comandos Encadeados

- **Git**: `git remote add origin <url>`
- **Docker**: `docker container run <image>`
- **Kubectl**: `kubectl get pods -n namespace`
- **AWS CLI**: `aws s3 cp file.txt s3://bucket/`

## 🚀 Instalação e Uso

### Pré-requisitos

- Go 1.23.5 ou superior
- Git

### Instalação

```bash
# Clone o repositório
git clone <url-do-repositorio>
cd 15_Cobra_CLI

# Instale as dependências
go mod download

# Compile a aplicação
go build -o course-cli .
```

### Uso Básico

```bash
# Executar a CLI
./course-cli

# Ver ajuda
./course-cli --help

# Ver ajuda de um comando específico
./course-cli category --help
./course-cli project task --help
```

## 📝 Exemplos de Uso

### Comandos de Categoria

```bash
# Criar categoria
./course-cli category create "Programação" "Cursos de programação"

# Listar categorias
./course-cli category list

# Buscar categoria por ID
./course-cli category get <id>

# Atualizar categoria
./course-cli category update <id> "Novo Nome" "Nova Descrição"

# Deletar categoria
./course-cli category delete <id>
```

### Comando Ping com Flag

```bash
# Ping normal
./course-cli ping
# Output: pong

# Ping com flag
./course-cli ping --pong
# Output: pong pong
```

### Comandos Encadeados - Projetos e Tarefas

```bash
# Ver ajuda do projeto
./course-cli project --help

# Ver ajuda das tarefas
./course-cli project task --help

# Adicionar tarefa
./course-cli project task add "Implementar autenticação"

# Listar tarefas
./course-cli project task list

# Completar tarefa
./course-cli project task complete 1
```

### Comandos de Configuração (Exemplo de Flags)

```bash
# Definir configuração (flags locais)
./course-cli config set --key "database_url" --value "sqlite://db.sqlite"

# Obter configuração (flags locais)
./course-cli config get --key "database_url"

# Listar configurações (flag global --verbose)
./course-cli config list --verbose

# Resetar configurações (flags locais + global)
./course-cli config reset --force --verbose
```

## 🧪 Testes

### Executar Testes

```bash
# Executar todos os testes
make test

# Executar apenas testes de comandos
make test-cmd

# Executar testes com cobertura
make test-coverage

# Executar testes com detecção de race conditions
make test-race
```

### Estrutura de Testes

```
cmd/
├── category_test.go    # Testes para comandos de categoria
├── ping_test.go        # Testes para comando ping
├── category.go         # Comandos implementados
├── ping.go
└── root.go
```

### Tipos de Testes

- **Validação de argumentos** para cada comando
- **Testes de estrutura** dos comandos
- **Testes de subcomandos** registrados
- **Testes de descrições** dos comandos
- **Testes de comportamento** básico
- **Testes de flags** (como --pong)

## 🔧 Makefile

### Comandos Disponíveis

```bash
# Configuração
make setup          # Configura o ambiente
make build          # Compila a aplicação
make run            # Executa a aplicação

# Testes
make test           # Executa todos os testes
make test-cmd       # Executa testes de comandos
make test-coverage  # Executa testes com cobertura
make test-race      # Executa testes com detecção de race conditions

# Demonstração
make demo-full      # Demonstração completa
make demo-categories # Demonstra comandos de categorias
make demo-ping      # Demonstra comando ping

# Limpeza
make clean          # Remove arquivos gerados

# Ajuda
make help           # Mostra todos os comandos disponíveis
```

## 📁 Estrutura do Projeto

```
15_Cobra_CLI/
├── cmd/                    # Comandos da CLI
│   ├── category.go         # Comandos de categoria (CRUD)
│   ├── category_test.go    # Testes de categoria
│   ├── ping.go            # Comando ping com flag
│   ├── ping_test.go       # Testes de ping
│   ├── project.go         # Comando principal de projeto
│   ├── task.go            # Subcomandos de tarefas
│   └── root.go            # Comando raiz
├── internal/              # Código interno da aplicação
│   ├── config/            # Configurações
│   │   └── database.go    # Configuração do banco
│   └── database/          # Camada de dados
│       └── category.go    # Operações de categoria
├── main.go               # Ponto de entrada
├── go.mod               # Dependências Go
├── go.sum               # Checksums das dependências
├── Makefile             # Automação de tarefas
└── README.md            # Este arquivo
```

## 🎯 Benefícios dos Comandos Encadeados

1. **Organização**: Comandos agrupados logicamente
2. **Hierarquia**: Estrutura clara e intuitiva
3. **Escalabilidade**: Fácil adicionar novos níveis
4. **UX**: Interface mais amigável
5. **Manutenção**: Código mais organizado

## 🔍 Resolução de Problemas

### Pasta Vermelha no VS Code

Se a pasta `cmd` aparecer vermelha no VS Code:

1. **Recarregue o VS Code**: `Cmd+Shift+P` → "Developer: Reload Window"
2. **Verifique arquivos não salvos**: `Ctrl+S` em todos os arquivos
3. **Reinicie o VS Code** completamente
4. **Verifique o status do Git**: `git status`

### Comandos não encontrados

```bash
# Recompile a aplicação
make clean
make build

# Verifique se o executável existe
ls -la course-cli
```

## 📚 Recursos Adicionais

- [Documentação do Cobra](https://cobra.dev/)
- [Go CLI Best Practices](https://github.com/spf13/cobra-cli)
- [Go Testing](https://golang.org/pkg/testing/)

## 🤝 Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

---

**Desenvolvido como parte do curso de Pós-Graduação em Golang** 🚀
