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

	// Testa o bidirectional streaming
	fmt.Println("🔄 Testando CreateCategoryStreamBidirectional...")
	fmt.Println("=================================================")

	// Abre o stream bidirectional
	stream, err := client.CreateCategoryStreamBidirectional(context.Background())
	if err != nil {
		log.Fatalf("Erro ao chamar CreateCategoryStreamBidirectional: %v", err)
	}

	// Canal para sincronizar goroutines
	done := make(chan bool)

	// Goroutine para RECEBER respostas do servidor
	go func() {
		count := 0
		for {
			category, err := stream.Recv()
			if err == io.EOF {
				// Servidor terminou de enviar
				fmt.Println("\n✅ Servidor terminou de enviar respostas")
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("Erro ao receber do stream: %v", err)
			}

			count++
			fmt.Printf("📥 [%d] Recebido: %s (ID: %s)\n", count, category.Name, category.Id)
		}
	}()

	// ENVIA múltiplas requisições ao servidor
	categories := []struct {
		name        string
		description string
	}{
		{"Backend", "Desenvolvimento backend"},
		{"Frontend", "Desenvolvimento frontend"},
		{"DevOps", "Infraestrutura e deploy"},
		{"Mobile", "Desenvolvimento mobile"},
		{"Data Science", "Ciência de dados"},
	}

	for i, cat := range categories {
		req := &pb.CreateCategoryRequest{
			Name:        cat.name,
			Description: cat.description,
		}

		fmt.Printf("📤 [%d] Enviando: %s\n", i+1, cat.name)
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Erro ao enviar requisição: %v", err)
		}

		// Simula delay entre envios
		time.Sleep(500 * time.Millisecond)
	}

	// Fecha o stream de envio (cliente terminou de enviar)
	fmt.Println("\n🔒 Cliente terminou de enviar requisições")
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Erro ao fechar stream de envio: %v", err)
	}

	// Aguarda todas as respostas serem recebidas
	<-done

	fmt.Println("=================================================")
	fmt.Println("✅ Bidirectional streaming concluído!")
}
