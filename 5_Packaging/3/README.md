# 🚀 Exemplo de Go Workspace

Este exemplo demonstra como usar o `go work` para gerenciar múltiplos módulos Go em um workspace.

## 📁 Estrutura do Projeto

```bash
3/
├── go.work                    # Workspace principal
├── math/                      # Módulo de matemática
│   ├── go.mod
│   └── math.go
├── utils/                     # Módulo de utilitários
│   ├── go.mod
│   └── formatter.go
└── calculator/                # Aplicação principal
    ├── go.mod
    └── main.go
```

## 🎯 O que este exemplo demonstra

### 1. **Múltiplos Módulos Independentes**

- `math`: Operações matemáticas básicas
- `utils`: Formatação e utilitários
- `calculator`: Aplicação que usa os outros módulos

### 2. **Compartilhamento de Dependências**

- Todos os módulos usam `github.com/google/uuid`
- O workspace resolve automaticamente as versões

### 3. **Desenvolvimento Local Integrado**

- Edite qualquer módulo e veja as mudanças imediatamente
- Não precisa publicar módulos para testar

## 🛠️ Como usar

### 1. **Verificar o workspace:**

```bash
go work edit -print
```

### 2. **Executar a aplicação:**

```bash
go run ./calculator/main.go
```

### 3. **Executar módulos individuais:**

```bash
# Testar apenas o módulo math
go test ./math

# Testar apenas o módulo utils  
go test ./utils
```

### 4. **Adicionar novo módulo:**

```bash
go work use ./novo-modulo
```

### 5. **Remover módulo:**

```bash
go work edit -dropuse ./modulo
```

## 🔍 Comandos úteis

```bash
# Listar todos os módulos no workspace
go list -m all

# Ver dependências de um módulo específico
go list -m all ./math

# Limpar cache de módulos
go clean -modcache

# Verificar se tudo está funcionando
go mod verify
```

## 💡 Vantagens do Workspace

- ✅ **Desenvolvimento mais rápido** - edições locais são refletidas imediatamente
- ✅ **Dependências compartilhadas** - evita duplicação de código
- ✅ **Resolução automática** - Go resolve conflitos de versão
- ✅ **Organização** - mantém módulos relacionados juntos

## ⚠️ Importante

- **Não commite o `go.work`** - é específico do ambiente local
- **Use apenas em desenvolvimento** - não para produção
- **Cada módulo deve ter seu próprio `go.mod`**

## 🎉 Resultado

O exemplo mostra como três módulos independentes podem trabalhar juntos em um workspace, compartilhando dependências e permitindo desenvolvimento integrado!
