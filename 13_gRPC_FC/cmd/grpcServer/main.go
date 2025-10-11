package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/ElizCarvalho/FC_PosGolang/13_gRPC_FC/internal/database"
	"github.com/ElizCarvalho/FC_PosGolang/13_gRPC_FC/internal/pb"
	"github.com/ElizCarvalho/FC_PosGolang/13_gRPC_FC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file:./db.dbgrpc")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer) // permite que o cliente descubra os serviços disponíveis

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	log.Printf("Server is running on port %s", listener.Addr().String())
	log.Fatal(grpcServer.Serve(listener))
}
