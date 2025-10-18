package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/configs"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/events"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/internal/event/handler"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/internal/infra/graph"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/internal/infra/grpc/pb"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/internal/infra/web/webserver"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}

	rabbitMQConn, rabbitMQChannel := getRabbitMQChannel(configs.RabbitMQURL)

	eventDispatcher := events.NewEventDispatcher()
	if err := eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}); err != nil {
		panic(err)
	}

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)

	// Web Server (REST)
	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandlerWithMethod("POST", "/order", webOrderHandler.Create)
	webserver.AddHandlerWithMethod("GET", "/orders", webOrderHandler.List)

	// Health Check
	healthHandler := NewHealthHandler(db, rabbitMQChannel)
	webserver.AddHandler("/health", healthHandler.Check)

	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	// gRPC Server
	grpcServer := grpc.NewServer()
	orderService := NewOrderService(db, eventDispatcher)
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			fmt.Printf("Error starting gRPC server: %v\n", err)
		}
	}()

	// GraphQL Server
	orderRepository := NewOrderRepository(db)
	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		OrderRepository:    orderRepository,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	graphqlServer := &http.Server{
		Addr:    ":" + configs.GraphQLServerPort,
		Handler: nil,
	}

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	go func() {
		if err := graphqlServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting GraphQL server: %v\n", err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\nShutting down servers gracefully...")

	// Timeout para shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown GraphQL server
	fmt.Println("Shutting down GraphQL server...")
	if err := graphqlServer.Shutdown(ctx); err != nil {
		fmt.Printf("GraphQL server forced to shutdown: %v\n", err)
	}

	// Shutdown gRPC server
	fmt.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()

	// Close RabbitMQ
	fmt.Println("Closing RabbitMQ connection...")
	if err := rabbitMQChannel.Close(); err != nil {
		fmt.Printf("Error closing RabbitMQ channel: %v\n", err)
	}
	if err := rabbitMQConn.Close(); err != nil {
		fmt.Printf("Error closing RabbitMQ connection: %v\n", err)
	}

	// Close database
	fmt.Println("Closing database connection...")
	if err := db.Close(); err != nil {
		fmt.Printf("Error closing database: %v\n", err)
	}

	fmt.Println("Shutdown complete")
}

func getRabbitMQChannel(rabbitMQURL string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return conn, ch
}
