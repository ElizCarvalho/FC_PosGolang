package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	endpoint        = "http://localhost:9000"
	region          = "us-east-1"
	accessKeyID     = "minioadmin"
	secretAccessKey = "minioadmin"
	bucketName      = "meu-bucket-teste"
)

func main() {
	// Configura a sessão do S3 apontando para o MinIO
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(region),
		Credentials:      credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		S3ForcePathStyle: aws.Bool(true), // Necessário para MinIO
	}))

	client := s3.New(sess)

	fmt.Println("🚀 Iniciando exemplos com S3 (MinIO)")
	fmt.Println("=====================================")

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
}

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
