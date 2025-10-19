# 🐳 Docker Multi-Stage Build - Otimização de Imagens

> Guia didático sobre Dockerfile.prod, multi-stage builds e otimização de imagens Go

## 📌 Sobre

Este projeto demonstra como criar imagens Docker otimizadas para aplicações Go usando **multi-stage builds**. A técnica permite gerar imagens super leves e seguras para produção.

## 🔧 Dockerfile.prod - Análise Detalhada

### 📋 Estrutura do Dockerfile

```dockerfile
FROM golang:1.23.5 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server main.go

FROM scratch
COPY --from=builder /app/server .
CMD ["./server"]
```

### 🔍 Explicação Linha por Linha

#### **Estágio 1: Builder**

```dockerfile
FROM golang:1.23.5 AS builder
```

- **`FROM`**: Define a imagem base
- **`golang:1.23.5`**: Imagem oficial do Go (Debian-based)
- **`AS builder`**: Nomeia este estágio para referência posterior

```dockerfile
WORKDIR /app

```dockerfile
- **`WORKDIR`**: Define o diretório de trabalho dentro do container

```dockerfile
COPY . .
```

- **`COPY`**: Copia arquivos do contexto local para o container
- **`. .`**: Copia tudo do diretório atual para `/app`

```dockerfile
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server main.go
```

- **`CGO_ENABLED=0`**: Desabilita CGO (compilação estática)
- **`GOOS=linux`**: Define o sistema operacional de destino
- **`go build`**: Compila a aplicação
- **`-ldflags="-s -w"`**: Remove símbolos de debug e tabela de símbolos
- **`-o server`**: Nome do binário gerado
- **`main.go`**: Arquivo principal

#### **Estágio 2: Runtime**

```dockerfile
FROM scratch
```

- **`scratch`**: Imagem vazia (apenas o binário)

```dockerfile
COPY --from=builder /app/server .
```

- **`COPY --from=builder`**: Copia arquivo do estágio anterior
- **`/app/server`**: Caminho no estágio builder
- **`.`**: Destino no estágio atual

```dockerfile
CMD ["./server"]
```

- **`CMD`**: Comando executado quando o container inicia

## 🚀 Comandos para Criar Imagem

### 📦 Build Básico

```bash
# Build com nome e tag
docker build -f Dockerfile.prod -t goapp:prod .

# Build com tag latest
docker build -f Dockerfile.prod -t goapp:latest .

# Build com versão específica
docker build -f Dockerfile.prod -t goapp:v1.0.0 .
```

### 🔧 Build com Parâmetros Avançados

```bash
# Build com progresso detalhado
docker build -f Dockerfile.prod -t goapp:prod --progress=plain .

# Build sem cache
docker build -f Dockerfile.prod -t goapp:prod --no-cache .

# Build com argumentos de build
docker build -f Dockerfile.prod -t goapp:prod --build-arg VERSION=1.0.0 .
```

### 🏃‍♂️ Executar a Imagem

```bash
# Execução simples
docker run -p 8080:8080 goapp:prod

# Execução em background
docker run -d -p 8080:8080 --name goapp-prod goapp:prod

# Execução com variáveis de ambiente
docker run -p 8080:8080 -e PORT=8080 goapp:prod
```

## 📊 Verificar Tamanho da Imagem

### 🔍 Comandos para Análise

```bash
# Lista todas as imagens com tamanhos
docker images

# Lista imagens específicas
docker images goapp

# Mostra tamanho detalhado
docker images --format "table {{.Repository}}\t{{.Tag}}\t{{.Size}}" goapp

# Compara tamanhos
docker images --format "table {{.Repository}}\t{{.Tag}}\t{{.Size}}" | grep -E "(goapp|golang)"
```

### 📈 Análise de Camadas

```bash
# Histórico de camadas da imagem
docker history goapp:prod

# Histórico com tamanhos
docker history --human=true goapp:prod

# Análise detalhada
docker inspect goapp:prod
```

### 🎯 Comparação de Tamanhos

```dockerfile
# Imagem de desenvolvimento (com Go toolchain)
docker build -f Dockerfile -t goapp:dev .
docker images goapp

# Imagem de produção (otimizada)
docker build -f Dockerfile.prod -t goapp:prod .
docker images goapp
```

**Resultado esperado:**

- `goapp:dev` ~800MB+ (com Go toolchain)
- `goapp:prod` ~10-20MB (apenas binário)

## 🪶 O que é o Scratch?

### 📚 Definição

**`scratch`** é uma imagem Docker especial que:

- **Não contém nada** - nem sistema operacional, nem shell, nem bibliotecas
- **É a base mais leve possível** - apenas o binário executável
- **É usada para aplicações estáticas** - que não dependem de bibliotecas externas

### ✅ Vantagens do Scratch

1. **Tamanho mínimo**: Imagem super leve
2. **Segurança máxima**: Sem dependências desnecessárias
3. **Performance**: Inicialização mais rápida
4. **Simplicidade**: Apenas o essencial

### ⚠️ Limitações do Scratch

1. **Sem shell**: Não pode executar `bash`, `sh`, etc.
2. **Sem ferramentas**: Sem `ls`, `cat`, `curl`, etc.
3. **Apenas binário**: Só executa o programa compilado
4. **Debugging limitado**: Difícil de debugar problemas

### 🔧 Quando Usar Scratch

- ✅ Aplicações Go com `CGO_ENABLED=0`
- ✅ Binários estáticos
- ✅ Microserviços simples
- ✅ Ambientes de produção

### 🚫 Quando NÃO Usar Scratch

- ❌ Aplicações que precisam de shell
- ❌ Binários que dependem de bibliotecas C
- ❌ Ambientes de desenvolvimento
- ❌ Aplicações que precisam de debugging

## 🛠️ Parâmetros de Build Go

### 🔧 CGO_ENABLED=0

```bash
CGO_ENABLED=0
```

- **O que faz**: Desabilita CGO (C bindings)
- **Resultado**: Binário estático, sem dependências externas
- **Necessário para**: Usar imagem `scratch`

### 🐧 GOOS=linux

```bash
GOOS=linux
```

- **O que faz**: Define o sistema operacional de destino
- **Resultado**: Binário compilado para Linux
- **Necessário para**: Containers Docker (que rodam Linux)

### ⚡ ldflags="-s -w"

```bash
-ldflags="-s -w"
```

- **`-s`**: Remove a tabela de símbolos
- **`-w`**: Remove informações de debug (DWARF)
- **Resultado**: Binário menor e mais rápido
- **Tamanho**: Reduz ~30-50% do tamanho

### 📊 Comparação de Tamanhos

```bash
# Sem otimização
go build -o server main.go
ls -lh server  # ~8MB

# Com otimização
CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server main.go
ls -lh server  # ~4MB
```

## 🧪 Testando a Aplicação

### 🚀 Execução Local

```bash
# Build e run
docker build -f Dockerfile.prod -t goapp:prod .
docker run -p 8080:8080 goapp:prod

# Teste
curl http://localhost:8080
# Resposta: "Aula de Deploy K8S"
```

### 🔍 Verificação de Funcionamento

```bash
# Status do container
docker ps

# Logs do container
docker logs goapp-prod

# Entrar no container (se necessário)
docker exec -it goapp-prod sh
```

## 📚 Recursos Adicionais

### 🔗 Links Úteis

- [Docker Multi-stage Builds](https://docs.docker.com/build/building/multi-stage/)
- [Go Docker Best Practices](https://docs.docker.com/language/golang/)
- [Scratch Image](https://hub.docker.com/_/scratch)

### 📖 Conceitos Relacionados

- **Multi-stage builds**: Técnica para otimizar imagens
- **Static linking**: Compilação sem dependências externas
- **Container security**: Minimizar superfície de ataque
- **Image optimization**: Reduzir tamanho e complexidade

## 🎯 Resumo

Este Dockerfile.prod demonstra as melhores práticas para:

- ✅ **Otimização de tamanho** (imagem mínima)
- ✅ **Segurança** (sem dependências desnecessárias)
- ✅ **Performance** (inicialização rápida)
- ✅ **Produção** (pronto para deploy)

A técnica de multi-stage build é essencial para aplicações Go em produção! 🚀
