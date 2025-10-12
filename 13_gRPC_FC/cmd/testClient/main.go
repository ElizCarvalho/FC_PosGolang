package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ElizCarvalho/FC_PosGolang/13_gRPC_FC/internal/pb"
)

func main() {
	// Conecta ao servidor gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Falha ao conectar: %v", err)
	}
	defer conn.Close()

	client := pb.NewCategoryServiceClient(conn)

	// Testa o streaming
	fmt.Println("🚀 Testando CreateCategoryStream...")
	fmt.Println("=====================================")

	req := &pb.CreateCategoryRequest{
		Name:        "Categoria Stream",
		Description: "Teste de streaming",
	}

	// Chama o método de streaming
	stream, err := client.CreateCategoryStream(context.Background(), req)
	if err != nil {
		log.Fatalf("Erro ao chamar CreateCategoryStream: %v", err)
	}

	// Processa as respostas do stream
	batchCount := 0
	totalCategories := 0

	for {
		// Recebe a próxima resposta do stream
		response, err := stream.Recv()
		if err == io.EOF {
			// Stream terminou
			break
		}
		if err != nil {
			log.Fatalf("Erro ao receber do stream: %v", err)
		}

		batchCount++
		categoriesInBatch := len(response.Categories)
		totalCategories += categoriesInBatch

		fmt.Printf("📦 Lote %d recebido: %d categorias\n", batchCount, categoriesInBatch)

		// Mostra as categorias do lote atual
		for i, category := range response.Categories {
			fmt.Printf("  %d. ID: %s | Nome: %s\n",
				i+1,
				category.Id,
				category.Name,
			)
		}

		fmt.Println("  ---")

		// Simula processamento do lote
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("=====================================")
	fmt.Printf("✅ Stream concluído!\n")
	fmt.Printf("📊 Total de lotes: %d\n", batchCount)
	fmt.Printf("📊 Total de categorias: %d\n", totalCategories)
}
