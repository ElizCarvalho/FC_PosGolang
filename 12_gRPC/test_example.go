package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ElizCarvalho/FC_PosGolang/12_gRPC/proto"
)

func main() {
	fmt.Println("🧪 Teste de Conectividade gRPC")
	fmt.Println("==============================")

	// Conecta ao servidor gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("❌ Falha ao conectar: %v", err)
	}
	defer conn.Close()

	// Cria o cliente
	client := pb.NewCalculatorServiceClient(conn)

	// Contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Teste simples de conectividade
	fmt.Println("🔗 Testando conectividade...")

	addResp, err := client.Add(ctx, &pb.AddRequest{A: 1, B: 1})
	if err != nil {
		log.Printf("❌ Erro no teste: %v", err)
		return
	}

	if addResp.Result == 2 {
		fmt.Println("✅ Servidor gRPC está funcionando corretamente!")
		fmt.Printf("📊 Resultado do teste: 1 + 1 = %.0f\n", addResp.Result)
	} else {
		fmt.Printf("⚠️  Resultado inesperado: %.0f\n", addResp.Result)
	}

	fmt.Println("\n🎉 Teste concluído com sucesso!")
}
