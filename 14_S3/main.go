package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ==============================================================================
// CONFIGURAÇÕES
// ==============================================================================

const (
	endpoint        = "http://localhost:9000"
	region          = "us-east-1"
	accessKeyID     = "minioadmin"
	secretAccessKey = "minioadmin"
	bucketName      = "meu-bucket-teste"

	// Configurações de performance
	maxConcurrentUploads = 10              // Máximo de uploads simultâneos
	maxRetries           = 3               // Tentativas de retry
	chunkSize            = 5 * 1024 * 1024 // 5MB por chunk
)

// ==============================================================================
// TIPOS E STRUCTS
// ==============================================================================

// Resultado do upload
type UploadResult struct {
	Key     string
	Success bool
	Error   error
	Size    int64
	Time    time.Duration
}

// Worker para uploads concorrentes
type UploadWorker struct {
	client  *s3.S3
	bucket  string
	sem     chan struct{} // Semáforo para controlar concorrência
	results chan UploadResult
	wg      sync.WaitGroup
}

// ==============================================================================
// CONFIGURAÇÃO DO CLIENTE S3
// ==============================================================================

func setupS3Client() *s3.S3 {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(region),
		Credentials:      credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		S3ForcePathStyle: aws.Bool(true), // Necessário para MinIO
	}))
	return s3.New(sess)
}

func isMinIORunning() bool {
	client := setupS3Client()
	_, err := client.ListBuckets(&s3.ListBucketsInput{})
	return err == nil
}

// ==============================================================================
// OPERAÇÕES BÁSICAS DO S3
// ==============================================================================

func createBucket(client *s3.S3, bucket string) error {
	fmt.Printf("📦 Criando bucket '%s'...\n", bucket)
	_, err := client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}
	fmt.Println("✅ Bucket criado com sucesso!")
	return nil
}

func listBuckets(client *s3.S3) error {
	fmt.Println("📋 Listando buckets...")
	result, err := client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return err
	}

	for _, bucket := range result.Buckets {
		fmt.Printf("   - %s (criado em: %s)\n", *bucket.Name, bucket.CreationDate)
	}
	fmt.Println()
	return nil
}

func uploadFile(client *s3.S3, bucket, key, content string) error {
	fmt.Printf("⬆️  Fazendo upload do arquivo '%s'...\n", key)
	_, err := client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader([]byte(content)),
		ContentType: aws.String("text/plain"),
	})
	if err != nil {
		return err
	}
	fmt.Println("✅ Upload realizado com sucesso!")
	return nil
}

func listObjects(client *s3.S3, bucket string) error {
	fmt.Printf("📂 Listando objetos no bucket '%s'...\n", bucket)
	result, err := client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}

	if len(result.Contents) == 0 {
		fmt.Println("   (vazio)")
	}

	for _, obj := range result.Contents {
		fmt.Printf("   - %s (tamanho: %d bytes)\n", *obj.Key, *obj.Size)
	}
	fmt.Println()
	return nil
}

func downloadFile(client *s3.S3, bucket, key string) error {
	fmt.Printf("⬇️  Fazendo download do arquivo '%s'...\n", key)
	result, err := client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Download realizado! Conteúdo:\n   %s\n\n", string(body))
	return nil
}

func generatePresignedURL(client *s3.S3, bucket, key string) error {
	fmt.Printf("🔗 Gerando URL pré-assinada para '%s'...\n", key)

	req, _ := client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	url, err := req.Presign(15 * 60 * 1000000000) // 15 minutos
	if err != nil {
		return err
	}

	fmt.Printf("✅ URL gerada (válida por 15 minutos):\n   %s\n\n", url)
	return nil
}

func deleteObject(client *s3.S3, bucket, key string) error {
	fmt.Printf("🗑️  Deletando arquivo '%s'...\n", key)
	_, err := client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	fmt.Println("✅ Arquivo deletado com sucesso!")
	return nil
}

// ==============================================================================
// UPLOAD OTIMIZADO (CONCORRENTE)
// ==============================================================================

func NewUploadWorker(client *s3.S3, bucket string) *UploadWorker {
	return &UploadWorker{
		client:  client,
		bucket:  bucket,
		sem:     make(chan struct{}, maxConcurrentUploads),
		results: make(chan UploadResult, 100),
	}
}

func (w *UploadWorker) UploadMultipleFiles(filePaths []string) []UploadResult {
	fmt.Printf("🚀 Iniciando upload de %d arquivos com %d workers...\n",
		len(filePaths), maxConcurrentUploads)

	start := time.Now()

	// Inicia workers
	for _, filePath := range filePaths {
		w.wg.Add(1)
		go w.uploadFileAsync(filePath)
	}

	// Coleta resultados
	go func() {
		w.wg.Wait()
		close(w.results)
	}()

	var results []UploadResult
	for result := range w.results {
		results = append(results, result)
	}

	totalTime := time.Since(start)
	fmt.Printf("⏱️  Upload concluído em %v\n", totalTime)

	// Estatísticas
	successCount := 0
	totalSize := int64(0)
	for _, result := range results {
		if result.Success {
			successCount++
			totalSize += result.Size
		}
	}

	fmt.Printf("📊 Estatísticas:\n")
	fmt.Printf("   ✅ Sucessos: %d/%d\n", successCount, len(filePaths))
	fmt.Printf("   📦 Tamanho total: %.2f MB\n", float64(totalSize)/(1024*1024))
	fmt.Printf("   ⚡ Velocidade: %.2f MB/s\n",
		float64(totalSize)/(1024*1024)/totalTime.Seconds())

	return results
}

func (w *UploadWorker) uploadFileAsync(filePath string) {
	defer w.wg.Done()

	// Controla concorrência
	w.sem <- struct{}{}
	defer func() { <-w.sem }()

	start := time.Now()

	// Tenta upload com retry
	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			time.Sleep(time.Duration(attempt) * time.Second) // Backoff
		}

		size, err := w.uploadFileWithRetry(filePath)
		if err == nil {
			w.results <- UploadResult{
				Key:     filePath,
				Success: true,
				Size:    size,
				Time:    time.Since(start),
			}
			return
		}
		lastErr = err
	}

	// Falhou após todas as tentativas
	w.results <- UploadResult{
		Key:     filePath,
		Success: false,
		Error:   lastErr,
		Time:    time.Since(start),
	}
}

func (w *UploadWorker) uploadFileWithRetry(filePath string) (int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}

	// Para arquivos grandes, usa multipart upload
	if fileInfo.Size() > chunkSize {
		return w.multipartUpload(file, filePath, fileInfo.Size())
	}

	// Upload simples para arquivos pequenos
	_, err = w.client.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(w.bucket),
		Key:           aws.String(filePath),
		Body:          file,
		ContentLength: aws.Int64(fileInfo.Size()),
		ContentType:   aws.String(getContentType(filePath)),
	})

	return fileInfo.Size(), err
}

// Multipart upload para arquivos grandes
func (w *UploadWorker) multipartUpload(file *os.File, key string, size int64) (int64, error) {
	// Inicia multipart upload
	createResp, err := w.client.CreateMultipartUpload(&s3.CreateMultipartUploadInput{
		Bucket:      aws.String(w.bucket),
		Key:         aws.String(key),
		ContentType: aws.String(getContentType(key)),
	})
	if err != nil {
		return 0, err
	}

	var parts []*s3.CompletedPart
	partNumber := int64(1)
	buffer := make([]byte, chunkSize)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			// Aborta upload em caso de erro
			w.client.AbortMultipartUpload(&s3.AbortMultipartUploadInput{
				Bucket:   aws.String(w.bucket),
				Key:      aws.String(key),
				UploadId: createResp.UploadId,
			})
			return 0, err
		}

		// Upload da parte
		uploadResp, err := w.client.UploadPart(&s3.UploadPartInput{
			Bucket:     aws.String(w.bucket),
			Key:        aws.String(key),
			PartNumber: aws.Int64(partNumber),
			UploadId:   createResp.UploadId,
			Body:       bytes.NewReader(buffer[:n]),
		})
		if err != nil {
			w.client.AbortMultipartUpload(&s3.AbortMultipartUploadInput{
				Bucket:   aws.String(w.bucket),
				Key:      aws.String(key),
				UploadId: createResp.UploadId,
			})
			return 0, err
		}

		parts = append(parts, &s3.CompletedPart{
			ETag:       uploadResp.ETag,
			PartNumber: aws.Int64(partNumber),
		})

		partNumber++
	}

	// Completa o upload
	_, err = w.client.CompleteMultipartUpload(&s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(w.bucket),
		Key:      aws.String(key),
		UploadId: createResp.UploadId,
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: parts,
		},
	})

	return size, err
}

func getContentType(filename string) string {
	ext := filepath.Ext(filename)
	switch ext {
	case ".txt":
		return "text/plain"
	case ".json":
		return "application/json"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".pdf":
		return "application/pdf"
	default:
		return "application/octet-stream"
	}
}

// ==============================================================================
// UTILITÁRIOS
// ==============================================================================

func createTestFiles(count int) []string {
	var files []string

	for i := 0; i < count; i++ {
		filename := fmt.Sprintf("test-file-%d.txt", i)
		content := fmt.Sprintf("Conteúdo do arquivo de teste %d - %s", i, time.Now().Format(time.RFC3339))

		if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
			log.Printf("Erro ao criar arquivo %s: %v", filename, err)
			continue
		}

		files = append(files, filename)
	}

	return files
}

func cleanupTestFiles(files []string) {
	for _, file := range files {
		os.Remove(file)
	}
}

func createLocalFile(filename string) error {
	content := "Este arquivo foi criado localmente e será enviado para o S3!"
	return os.WriteFile(filename, []byte(content), 0644)
}

func uploadLocalFile(client *s3.S3, bucket, filePath string) error {
	fmt.Printf("⬆️  Fazendo upload do arquivo local '%s'...\n", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	_, err = client.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(filePath),
		Body:          file,
		ContentLength: aws.Int64(fileInfo.Size()),
		ContentType:   aws.String("text/plain"),
	})
	if err != nil {
		return err
	}

	fmt.Println("✅ Upload do arquivo local realizado com sucesso!")
	return nil
}

// ==============================================================================
// DEMOS E TESTES DE PERFORMANCE
// ==============================================================================

func testSequentialUpload(files []string) time.Duration {
	client := setupS3Client()
	bucket := "demo-sequential"

	// Cria bucket
	client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})

	start := time.Now()

	for _, filePath := range files {
		file, _ := os.Open(filePath)
		fileInfo, _ := file.Stat()

		client.PutObject(&s3.PutObjectInput{
			Bucket:        aws.String(bucket),
			Key:           aws.String(filePath),
			Body:          file,
			ContentLength: aws.Int64(fileInfo.Size()),
		})
		file.Close()
	}

	return time.Since(start)
}

func testConcurrentUpload(files []string) time.Duration {
	client := setupS3Client()
	bucket := "demo-concurrent"

	// Cria bucket
	client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})

	start := time.Now()

	worker := NewUploadWorker(client, bucket)
	worker.UploadMultipleFiles(files)

	return time.Since(start)
}

func runPerformanceDemo() {
	fmt.Println("🚀 Demo de Performance: S3 Upload")
	fmt.Println("==================================\n")

	// Verifica se MinIO está rodando
	if !isMinIORunning() {
		fmt.Println("❌ MinIO não está rodando!")
		fmt.Println("   Execute: make minio-up")
		return
	}

	fmt.Println("📊 Comparando implementações...\n")

	// Teste com diferentes quantidades de arquivos
	testCases := []int{5, 20, 50}

	for _, fileCount := range testCases {
		fmt.Printf("🧪 Testando com %d arquivos:\n", fileCount)

		// Cria arquivos de teste
		files := createTestFiles(fileCount)
		defer cleanupTestFiles(files)

		// Teste sequencial
		fmt.Println("   📝 Implementação sequencial...")
		sequentialTime := testSequentialUpload(files)

		// Teste concorrente
		fmt.Println("   ⚡ Implementação concorrente...")
		concurrentTime := testConcurrentUpload(files)

		// Resultados
		improvement := float64(sequentialTime) / float64(concurrentTime)
		fmt.Printf("   📈 Resultado: %.1fx mais rápido!\n", improvement)
		fmt.Printf("   ⏱️  Sequencial: %v\n", sequentialTime)
		fmt.Printf("   ⚡ Concorrente: %v\n\n", concurrentTime)
	}

	fmt.Println("💡 Dicas de otimização:")
	fmt.Println("   • Use workers para uploads paralelos")
	fmt.Println("   • Implemente retry com backoff")
	fmt.Println("   • Use multipart upload para arquivos grandes")
	fmt.Println("   • Controle concorrência com semáforos")
	fmt.Println("   • Monitore uso de memória")
}

// ==============================================================================
// FUNÇÃO PRINCIPAL
// ==============================================================================

func main() {
	// Verifica argumentos da linha de comando
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "demo":
			runPerformanceDemo()
			return
		case "help":
			fmt.Println("🚀 S3 Study - Comandos disponíveis:")
			fmt.Println("   go run main.go        - Exemplos básicos do S3")
			fmt.Println("   go run main.go demo   - Demo de performance")
			fmt.Println("   go run main.go help   - Esta ajuda")
			return
		}
	}

	// Configura a sessão do S3 apontando para o MinIO
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(region),
		Credentials:      credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		S3ForcePathStyle: aws.Bool(true), // Necessário para MinIO
	}))

	client := s3.New(sess)

	fmt.Println("🚀 Iniciando exemplos com S3 (MinIO)")
	fmt.Println("=====================================\n")

	// 1. Criar bucket
	if err := createBucket(client, bucketName); err != nil {
		log.Printf("⚠️  Aviso ao criar bucket: %v\n", err)
	}

	// 2. Listar buckets
	if err := listBuckets(client); err != nil {
		log.Fatalf("❌ Erro ao listar buckets: %v", err)
	}

	// 3. Upload de arquivo (string)
	fileName := "teste.txt"
	fileContent := "Olá! Este é um arquivo de teste para estudar S3 com Go e MinIO! 🚀"
	if err := uploadFile(client, bucketName, fileName, fileContent); err != nil {
		log.Fatalf("❌ Erro ao fazer upload: %v", err)
	}

	// 4. Listar objetos no bucket
	if err := listObjects(client, bucketName); err != nil {
		log.Fatalf("❌ Erro ao listar objetos: %v", err)
	}

	// 5. Download de arquivo
	if err := downloadFile(client, bucketName, fileName); err != nil {
		log.Fatalf("❌ Erro ao fazer download: %v", err)
	}

	// 6. Upload de arquivo real (da máquina)
	localFile := "arquivo-local.txt"
	if err := createLocalFile(localFile); err != nil {
		log.Fatalf("❌ Erro ao criar arquivo local: %v", err)
	}
	if err := uploadLocalFile(client, bucketName, localFile); err != nil {
		log.Fatalf("❌ Erro ao fazer upload do arquivo local: %v", err)
	}

	// 7. Gerar URL pré-assinada (válida por 15 minutos)
	if err := generatePresignedURL(client, bucketName, fileName); err != nil {
		log.Fatalf("❌ Erro ao gerar URL: %v", err)
	}

	// 8. Deletar objeto
	if err := deleteObject(client, bucketName, fileName); err != nil {
		log.Fatalf("❌ Erro ao deletar objeto: %v", err)
	}

	fmt.Println("\n✅ Todos os exemplos executados com sucesso!")
	fmt.Println("\n💡 Acesse o console web: http://localhost:9001")
	fmt.Println("   Usuário: minioadmin")
	fmt.Println("   Senha: minioadmin")
	fmt.Println("\n🚀 Para ver o demo de performance:")
	fmt.Println("   go run main.go demo")
}

// ==============================================================================
// TESTES (BENCHMARKS)
// ==============================================================================

// Benchmark do upload sequencial (implementação original)
func BenchmarkSequentialUpload(b *testing.B) {
	client := setupS3Client()
	bucket := "benchmark-bucket"

	// Cria bucket se não existir
	client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})

	// Cria arquivos de teste
	files := createTestFiles(10)
	defer cleanupTestFiles(files)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Upload sequencial (como na implementação original)
		for _, filePath := range files {
			file, _ := os.Open(filePath)
			fileInfo, _ := file.Stat()

			client.PutObject(&s3.PutObjectInput{
				Bucket:        aws.String(bucket),
				Key:           aws.String(filePath),
				Body:          file,
				ContentLength: aws.Int64(fileInfo.Size()),
			})
			file.Close()
		}
	}
}

// Benchmark do upload concorrente (implementação otimizada)
func BenchmarkConcurrentUpload(b *testing.B) {
	client := setupS3Client()
	bucket := "benchmark-bucket-concurrent"

	// Cria bucket se não existir
	client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})

	// Cria arquivos de teste
	files := createTestFiles(10)
	defer cleanupTestFiles(files)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		worker := NewUploadWorker(client, bucket)
		worker.UploadMultipleFiles(files)
	}
}

// Teste de performance com diferentes números de arquivos
func TestPerformanceComparison(t *testing.T) {
	client := setupS3Client()
	bucket := "performance-test"

	client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})

	testCases := []struct {
		name  string
		count int
	}{
		{"10 arquivos", 10},
		{"50 arquivos", 50},
		{"100 arquivos", 100},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			files := createTestFiles(tc.count)
			defer cleanupTestFiles(files)

			// Teste sequencial
			start := time.Now()
			for _, filePath := range files {
				file, _ := os.Open(filePath)
				fileInfo, _ := file.Stat()
				client.PutObject(&s3.PutObjectInput{
					Bucket:        aws.String(bucket),
					Key:           aws.String(filePath),
					Body:          file,
					ContentLength: aws.Int64(fileInfo.Size()),
				})
				file.Close()
			}
			sequentialTime := time.Since(start)

			// Teste concorrente
			start = time.Now()
			worker := NewUploadWorker(client, bucket)
			worker.UploadMultipleFiles(files)
			concurrentTime := time.Since(start)

			t.Logf("Sequencial: %v", sequentialTime)
			t.Logf("Concorrente: %v", concurrentTime)
			t.Logf("Melhoria: %.1fx", float64(sequentialTime)/float64(concurrentTime))
		})
	}
}

// Teste de memória com arquivos grandes
func TestMemoryUsage(t *testing.T) {
	client := setupS3Client()
	bucket := "memory-test"

	client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})

	// Cria arquivo grande (10MB)
	largeFile := "large-file.txt"
	content := make([]byte, 10*1024*1024) // 10MB
	for i := range content {
		content[i] = byte(i % 256)
	}

	os.WriteFile(largeFile, content, 0644)
	defer os.Remove(largeFile)

	// Testa upload com multipart
	worker := NewUploadWorker(client, bucket)
	results := worker.UploadMultipleFiles([]string{largeFile})

	if len(results) != 1 || !results[0].Success {
		t.Fatalf("Upload falhou: %v", results[0].Error)
	}

	t.Logf("Arquivo grande (%d MB) enviado com sucesso", len(content)/(1024*1024))
}
