# 📦 Estudos S3 com MinIO

> Ambiente local para estudar AWS S3 sem gastar dinheiro! Usando MinIO, um servidor de armazenamento de objetos compatível com a API do S3.

## 📌 Sobre

Este projeto demonstra como trabalhar com S3 localmente usando **MinIO**, que é 100% compatível com a API do AWS S3. Perfeito para:

- 🎓 Estudar sem medo de gastar dinheiro
- 🧪 Testar features do S3
- 🚀 Desenvolver aplicações que usam S3

## 🎯 O que é MinIO?

MinIO é um servidor de armazenamento de objetos de alta performance, compatível com a API do Amazon S3. Você pode usar o mesmo código que usaria com o S3 da AWS, mas rodando tudo localmente!

## 🔧 Configuração

### 1. Subir o MinIO

```bash
make minio-up
```

Isso vai iniciar:

- **API**: http://localhost:9000
- **Console Web**: http://localhost:9001
- **Credenciais**: minioadmin / minioadmin

### 2. Instalar dependências

```bash
make setup
```

### 3. Rodar a aplicação

```bash
# Exemplos básicos do S3
make run

# Demo de performance (comparação sequencial vs concorrente)
make run-demo
```

## 📚 Exemplos Implementados

O código demonstra as operações mais comuns do S3:

### 1. ✅ Criar Bucket

```go
client.CreateBucket(&s3.CreateBucketInput{
    Bucket: aws.String("meu-bucket"),
})
```

### 2. 📋 Listar Buckets

```go
client.ListBuckets(&s3.ListBucketsInput{})
```

### 3. ⬆️ Upload de Arquivo

```go
client.PutObject(&s3.PutObjectInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(filename),
    Body:   bytes.NewReader(content),
})
```

### 4. ⬇️ Download de Arquivo

```go
client.GetObject(&s3.GetObjectInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(filename),
})
```

### 5. 🔗 URL Pré-assinada

```go
req, _ := client.GetObjectRequest(&s3.GetObjectInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(filename),
})
url, _ := req.Presign(15 * time.Minute)
```

### 6. 🗑️ Deletar Objeto

```go
client.DeleteObject(&s3.DeleteObjectInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(filename),
})
```

## 🎮 Console Web

Acesse http://localhost:9001 para gerenciar visualmente:

- 📦 Buckets
- 📁 Arquivos
- ⚙️ Configurações
- 👥 Usuários

## 🧪 Testes

```bash
# Testes unitários
make test

# Benchmarks de performance
make benchmark
```

## ⚡ Demo de Performance

O projeto inclui um demo que compara upload sequencial vs concorrente:

```bash
# Demo interativo
make run-demo

# Ou diretamente
go run main.go demo
```

**Resultados esperados:**

- Upload sequencial: ~15-20 segundos (50 arquivos)
- Upload concorrente: ~3-5 segundos (50 arquivos)
- **Melhoria: 3-5x mais rápido!**

## ⚙️ Mecanismo de Paralelismo

O projeto implementa um **Worker Pool Pattern** com controle de concorrência para otimizar uploads:

### 🔧 **1. Worker Pool Pattern**

```go
type UploadWorker struct {
    client  *s3.S3
    bucket  string
    sem     chan struct{} // Semáforo para controlar concorrência
    results chan UploadResult
    wg      sync.WaitGroup
}
```

**Componentes:**

- **Workers**: goroutines que processam uploads
- **Semáforo**: controla quantos workers podem rodar simultaneamente
- **WaitGroup**: espera todos os workers terminarem
- **Channel de resultados**: coleta os resultados de cada worker

### ⚡ **2. Controle de Concorrência (Semáforo)**

```go
const maxConcurrentUploads = 10 // Máximo 10 uploads simultâneos

func (w *UploadWorker) uploadFileAsync(filePath string) {
    defer w.wg.Done()
    
    // Controla concorrência - só permite 10 workers simultâneos
    w.sem <- struct{}{}        // Pega um "slot"
    defer func() { <-w.sem }() // Libera o "slot" quando terminar
    
    // ... faz o upload ...
}
```

**Por que usar semáforo?**

- **Evita sobrecarga**: não cria 1000 goroutines de uma vez
- **Controla recursos**: limita uso de memória e conexões
- **Performance otimizada**: 10 workers é o "sweet spot" para S3

### 🚀 **3. Goroutines + Channels**

```go
func (w *UploadWorker) UploadMultipleFiles(filePaths []string) []UploadResult {
    // 1. Inicia workers para cada arquivo
    for _, filePath := range filePaths {
        w.wg.Add(1)
        go w.uploadFileAsync(filePath) // Goroutine por arquivo
    }
    
    // 2. Coleta resultados em paralelo
    go func() {
        w.wg.Wait()      // Espera todos terminarem
        close(w.results) // Fecha o channel
    }()
    
    // 3. Coleta todos os resultados
    var results []UploadResult
    for result := range w.results {
        results = append(results, result)
    }
    
    return results
}
```

### 📊 **4. Fluxo Completo**

```mermaid
Arquivos: [file1.txt, file2.txt, ..., file50.txt]
    ↓
Cria 50 goroutines (mas só 10 rodam simultaneamente)
    ↓
Semáforo controla: 10 workers ativos por vez
    ↓
Cada worker:
  1. Pega um "slot" do semáforo
  2. Faz upload do arquivo
  3. Envia resultado para o channel
  4. Libera o "slot"
    ↓
WaitGroup espera todos terminarem
    ↓
Coleta todos os resultados
```

### 🔄 **5. Retry com Backoff**

```go
for attempt := 0; attempt < maxRetries; attempt++ {
    if attempt > 0 {
        time.Sleep(time.Duration(attempt) * time.Second) // 1s, 2s, 3s
    }
    
    if err := uploadFile(); err == nil {
        return // Sucesso!
    }
}
```

### 🎯 **6. Comparação de Abordagens**

#### **vs Upload Sequencial:**

```go
// ❌ Sequencial (lento)
for _, file := range files {
    uploadFile(file) // Bloqueia até terminar
}

// ✅ Concorrente (rápido)
for _, file := range files {
    go uploadFileAsync(file) // Roda em paralelo
}
```

#### **vs Sem Controle:**

```go
// ❌ Sem semáforo (pode estourar memória)
for _, file := range files {
    go uploadFile(file) // Cria 1000 goroutines!
}

// ✅ Com semáforo (controlado)
for _, file := range files {
    go uploadFileAsync(file) // Máximo 10 simultâneos
}
```

### 💡 **Resumo do Mecanismo:**

1. **Goroutines** - paralelismo real
2. **Semáforo** - controle de recursos
3. **WaitGroup** - sincronização
4. **Channels** - comunicação entre goroutines
5. **Retry** - resiliência a falhas

É um padrão clássico de **Worker Pool** com **controle de concorrência** - muito usado em sistemas de alta performance! 🚀

## 🧹 Limpeza

Para parar o MinIO:

```bash
make minio-down
```

Para remover tudo (incluindo volumes):

```bash
make clean
```

## 💡 Dicas

1. **Migrando para AWS Real**
   - Basta trocar o endpoint para o da AWS
   - Usar credenciais reais (AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY)
   - Remover `S3ForcePathStyle: true`

2. **Persistência**
   - Os dados ficam salvos no volume Docker
   - Use `make clean` para resetar tudo

3. **Performance**
   - MinIO é extremamente rápido
   - Perfeito para testes de integração

## 🔗 Links Úteis

- [MinIO Documentation](https://min.io/docs/minio/linux/index.html)
- [AWS SDK for Go](https://aws.amazon.com/sdk-for-go/)
- [AWS S3 API Reference](https://docs.aws.amazon.com/AmazonS3/latest/API/Welcome.html)

## 🚀 Próximos Passos

- [ ] Implementar upload de múltiplos arquivos
- [ ] Adicionar streaming de arquivos grandes
- [ ] Implementar versionamento de objetos
- [ ] Adicionar políticas de acesso (IAM)
- [ ] Implementar lifecycle policies
