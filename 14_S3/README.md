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

