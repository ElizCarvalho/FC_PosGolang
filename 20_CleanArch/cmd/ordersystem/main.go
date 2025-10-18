package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/configs"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/events"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/internal/event/handler"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/internal/infra/graph"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/internal/infra/grpc/pb"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/internal/infra/grpc/service"
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
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Printf("Error closing database: %v\n", err)
		}
	}()

	rabbitMQChannel := getRabbitMQChannel(configs.RabbitMQURL)

	eventDispatcher := events.NewEventDispatcher()
	if err := eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}); err != nil {
		panic(err)
	}

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order", webOrderHandler.Create)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
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

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	if err := http.ListenAndServe(":"+configs.GraphQLServerPort, nil); err != nil {
		panic(err)
	}
}

func getRabbitMQChannel(rabbitMQURL string) *amqp.Channel {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
