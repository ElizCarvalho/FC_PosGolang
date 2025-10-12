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
├── demo (demonstração de tipos de flags)
│   ├── --name [string] --age [int] --active [bool]
│   ├── --price [float] --timeout [duration]
│   ├── --tags [string-slice] --ports [int-slice]
│   └── --features [bool-slice]
└── confirm (flags com opções específicas)
    ├── --yes [y/n/yes/no] (padrão: y)
    ├── --mode [interactive/batch] (padrão: interactive)
    ├── --priority [low/medium/high] (padrão: medium)
    ├── --environment [dev/staging/prod] (padrão: dev)
    └── --format [json/xml/yaml] (padrão: json)
└── hooks (demonstração de hooks do Cobra)
    ├── --name [string] (exemplo de hook)
    └── subcommand --value [string] (herança de hooks)
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

## 🎛️ Manipulação de Flags

### Tipos de Flags Disponíveis

#### 1. **STRING** - Texto
```go
cmd.Flags().String("name", "", "Nome da pessoa")
```
```bash
--name "João"
```

#### 2. **INT** - Números Inteiros
```go
cmd.Flags().Int("age", 0, "Idade da pessoa")
```
```bash
--age 25
```

#### 3. **BOOL** - Valores Booleanos
```go
cmd.Flags().Bool("active", false, "Status ativo/inativo")
```
```bash
--active
```

#### 4. **FLOAT64** - Números Decimais
```go
cmd.Flags().Float64("price", 0.0, "Preço do produto")
```
```bash
--price 99.99
```

#### 5. **DURATION** - Duração de Tempo
```go
cmd.Flags().Duration("timeout", 30*time.Second, "Timeout para operações")
```
```bash
--timeout 1m
--timeout 30s
```

#### 6. **STRING SLICE** - Múltiplos Textos
```go
cmd.Flags().StringSlice("tags", []string{}, "Tags para categorização")
```
```bash
--tags "go,cli,demo"
```

#### 7. **INT SLICE** - Múltiplos Números
```go
cmd.Flags().IntSlice("ports", []int{}, "Lista de portas")
```
```bash
--ports 80,443,8080
```

#### 8. **BOOL SLICE** - Múltiplos Booleanos
```go
cmd.Flags().BoolSlice("features", []bool{}, "Lista de features ativadas")
```
```bash
--features true,false,true
```

### Técnicas de Manipulação

#### **Obter Valores das Flags**
```go
name, _ := cmd.Flags().GetString("name")
age, _ := cmd.Flags().GetInt("age")
active, _ := cmd.Flags().GetBool("active")
```

#### **Verificar se Flag foi Fornecida**
```go
if cmd.Flags().Changed("name") {
    // Flag foi fornecida
}
```

#### **Flags Obrigatórias**
```go
cmd.MarkFlagRequired("name")
```

#### **Flags com Shortcuts**
```go
cmd.Flags().StringP("output", "o", "", "Arquivo de saída")
cmd.Flags().BoolP("force", "f", false, "Forçar operação")
```

#### **Validação de Valores**
```go
if age < 0 || age > 150 {
    fmt.Println("❌ Idade deve estar entre 0 e 150 anos")
    return
}
```

### Flags com Opções Específicas

#### **Exemplo: Yes/No com Valor Padrão**
```go
cmd.Flags().String("yes", "y", "Confirmação (y/n ou yes/no) - padrão: y")
```

**Comportamento:**
- ✅ **Valor padrão**: `y` (sim)
- ✅ **Aceita**: `y`, `n`, `yes`, `no`
- ❌ **Rejeita**: `maybe`, `talvez`, etc.

#### **Validação Personalizada**
```go
cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
    yesFlag, _ := cmd.Flags().GetString("yes")
    validYes := []string{"y", "n", "yes", "no"}
    for _, valid := range validYes {
        if strings.ToLower(yesFlag) == valid {
            return nil
        }
    }
    return fmt.Errorf("valor inválido para --yes: '%s'. Use: y/n ou yes/no", yesFlag)
}
```

## 🪝 Hooks do Cobra

### O que são Hooks?

**Hooks** são funções que são executadas em **momentos específicos** do ciclo de vida dos comandos Cobra. Eles permitem executar código antes, durante e depois da execução dos comandos.

### Tipos de Hooks Disponíveis

#### **1. Hooks Básicos (Apenas do comando atual)**

**PreRun** - Antes da execução
```go
cmd.PreRun = func(cmd *cobra.Command, args []string) {
    fmt.Println("🚀 Preparando execução...")
}
```

**Run** - Execução principal
```go
cmd.Run = func(cmd *cobra.Command, args []string) {
    fmt.Println("🎯 Executando comando principal...")
}
```

**PostRun** - Após a execução
```go
cmd.PostRun = func(cmd *cobra.Command, args []string) {
    fmt.Println("🏁 Finalizando...")
}
```

#### **2. Hooks Persistentes (Herdados por subcomandos)**

**PersistentPreRun** - Antes de qualquer comando
```go
cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
    fmt.Println("🌐 Inicialização global...")
}
```

**PersistentPostRun** - Após qualquer comando
```go
cmd.PersistentPostRun = func(cmd *cobra.Command, args []string) {
    fmt.Println("🌐 Finalização global...")
}
```

#### **3. Hooks com Tratamento de Erro**

**PreRunE** - PreRun com erro
```go
cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
    if name == "erro" {
        return fmt.Errorf("❌ Nome 'erro' não é permitido")
    }
    return nil
}
```

**PostRunE** - PostRun com erro
```go
cmd.PostRunE = func(cmd *cobra.Command, args []string) error {
    // Lógica que pode falhar
    return nil
}
```

### Ordem de Execução dos Hooks

```
1. PersistentPreRun (global)
2. PreRunE (validação com erro)
3. PreRun (preparação)
4. Run (execução principal)
5. PostRunE (finalização com erro)
6. PostRun (finalização)
7. PersistentPostRun (global)
```

### Casos de Uso dos Hooks

#### **1. Inicialização de Recursos**
```go
cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
    // Conectar ao banco de dados
    // Carregar configurações
    // Inicializar logs
}
```

#### **2. Validação de Entrada**
```go
cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
    // Validar argumentos
    // Verificar permissões
    // Validar flags
    return nil
}
```

#### **3. Logging e Auditoria**
```go
cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
    log.Printf("Comando executado: %s", cmd.Name())
    log.Printf("Argumentos: %v", args)
}
```

#### **4. Limpeza de Recursos**
```go
cmd.PersistentPostRun = func(cmd *cobra.Command, args []string) {
    // Fechar conexões
    // Salvar logs
    // Limpar cache
}
```

### Benefícios dos Hooks

1. **Separação de responsabilidades** - Código organizado
2. **Reutilização** - Hooks persistentes são herdados
3. **Validação** - PreRunE para validações
4. **Logging** - Rastreamento de execução
5. **Inicialização** - Setup automático
6. **Limpeza** - Cleanup automático
7. **Tratamento de erro** - Controle de falhas

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

### Demonstração de Tipos de Flags

```bash
# Exemplo básico com diferentes tipos
./course-cli demo --name "Maria" --age 30 --active --price 99.99

# Exemplo com múltiplos valores
./course-cli demo --name "João" --tags "go,cli" --ports 80,443,8080

# Exemplo com duração e validação
./course-cli demo --name "Ana" --timeout 1m --age 25 --active

# Exemplo com flags obrigatórias (falha sem --name)
./course-cli demo --age 25  # ❌ Erro: flag obrigatória não fornecida
```

### Flags com Opções Específicas

```bash
# Usando valores padrão
./course-cli confirm
# Usa: yes=y, mode=interactive, priority=medium, environment=dev, format=json

# Personalizando valores
./course-cli confirm --yes no --mode batch --priority high --environment prod

# Validação de entrada
./course-cli confirm --yes maybe  # ❌ Erro: valor inválido

# Combinações válidas
./course-cli confirm --yes yes --mode interactive --priority low --format yaml
```

### Demonstração de Hooks

```bash
# Comando principal com hooks
./course-cli hooks --name "João"
# Executa: PersistentPreRun → PreRunE → PreRun → Run → PostRunE → PostRun → PersistentPostRun

# Subcomando que herda hooks persistentes
./course-cli hooks subcommand --value "Exemplo"
# Executa: PersistentPreRun → Run → PersistentPostRun

# Exemplo de validação com erro
./course-cli hooks --name "erro"
# ❌ Falha no PreRunE: Nome 'erro' não é permitido
```

## 🎯 Padrão RunEFunc - Tratamento Elegante de Erros

### O que é o RunEFunc?

O **RunEFunc** é um padrão elegante para tratar erros em comandos Cobra, separando a lógica de negócio dos comandos e proporcionando um tratamento de erro consistente.

### Estrutura do Padrão

#### **1. Tipos Personalizados**
```go
// RunEFunc é um tipo personalizado para funções que retornam erro
type RunEFunc func(cmd *cobra.Command, args []string) error

// HandlerFunc é um tipo para funções que lidam com a lógica de negócio
type HandlerFunc func(args []string) error
```

#### **2. Funções Auxiliares**
```go
// RunEWithErrorHandling executa uma função RunE com tratamento elegante de erro
func RunEWithErrorHandling(fn RunEFunc) func(cmd *cobra.Command, args []string) {
    return func(cmd *cobra.Command, args []string) {
        if err := fn(cmd, args); err != nil {
            fmt.Printf("❌ Erro: %v\n", err)
            os.Exit(1)
        }
    }
}

// CreateHandler cria um handler para comandos que não precisam do cmd
func CreateHandler(handler HandlerFunc) RunEFunc {
    return func(cmd *cobra.Command, args []string) error {
        return handler(args)
    }
}
```

### Como Usar

#### **Antes (Deselegante)**
```go
var createCmd = &cobra.Command{
    Use: "create",
    Run: func(cmd *cobra.Command, args []string) {
        if categoryService == nil {
            log.Fatal("❌ Serviço não inicializado")
        }
        
        category, err := categoryService.Create(args[0], args[1])
        if err != nil {
            log.Fatalf("Erro ao criar: %v", err)
        }
        
        fmt.Printf("✅ Categoria criada: %s\n", category.Name)
    },
}
```

#### **Depois (Elegante)**
```go
// Handler separado (lógica de negócio)
func createCategoryHandler(args []string) error {
    if categoryService == nil {
        return fmt.Errorf("serviço de categoria não foi inicializado")
    }

    category, err := categoryService.Create(args[0], args[1])
    if err != nil {
        return fmt.Errorf("erro ao criar categoria: %w", err)
    }

    fmt.Printf("✅ Categoria criada com sucesso!\n")
    fmt.Printf("ID: %s\n", category.ID)
    fmt.Printf("Nome: %s\n", category.Name)
    return nil
}

// Comando limpo e focado
var createCmd = &cobra.Command{
    Use: "create",
    Run: RunEWithErrorHandling(CreateHandler(createCategoryHandler)),
}
```

### Benefícios do Padrão

1. **✅ Separação de Responsabilidades**
   - Comandos focam apenas na definição (Use, Short, Long, Args)
   - Handlers contêm toda a lógica de negócio
   - Tratamento de erro centralizado

2. **✅ Reutilização**
   - Handlers podem ser testados independentemente
   - Lógica de negócio pode ser reutilizada
   - Tratamento de erro consistente

3. **✅ Testabilidade**
   - Handlers são funções puras (fáceis de testar)
   - Mocking simplificado
   - Testes unitários mais focados

4. **✅ Manutenibilidade**
   - Código mais limpo e organizado
   - Fácil de entender e modificar
   - Padrão consistente em toda aplicação

5. **✅ Tratamento de Erro Elegante**
   - Sem `log.Fatal` espalhado pelo código
   - Mensagens de erro consistentes
   - Uso de `fmt.Errorf` com `%w` para wrapping

### Exemplo Completo

```go
// Handler para criação de categoria
func createCategoryHandler(args []string) error {
    if categoryService == nil {
        return fmt.Errorf("serviço de categoria não foi inicializado")
    }

    category, err := categoryService.Create(args[0], args[1])
    if err != nil {
        return fmt.Errorf("erro ao criar categoria: %w", err)
    }

    fmt.Printf("✅ Categoria criada com sucesso!\n")
    fmt.Printf("ID: %s\n", category.ID)
    fmt.Printf("Nome: %s\n", category.Name)
    fmt.Printf("Descrição: %s\n", category.Description)
    return nil
}

// Comando usando o padrão
var createCmd = &cobra.Command{
    Use:   "create [name] [description]",
    Short: "Criar uma nova categoria",
    Long:  `Cria uma nova categoria com nome e descrição fornecidos.`,
    Args:  cobra.ExactArgs(2),
    Run:   RunEWithErrorHandling(CreateHandler(createCategoryHandler)),
}
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
│   ├── config.go          # Comandos de configuração (flags locais/globais)
│   ├── confirm.go         # Flags com opções específicas (yes/no)
│   ├── demo.go            # Demonstração de tipos de flags
│   ├── hooks.go           # Demonstração de hooks do Cobra
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
