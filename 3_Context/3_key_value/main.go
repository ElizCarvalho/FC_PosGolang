package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", "42")
	bookHotel(ctx, "Avenida Brasil")
}

func bookHotel(ctx context.Context, name string) {
	//por convenção, o primeiro parâmetro de uma função que usa context é o próprio context
	requestID := ctx.Value("requestID")
	fmt.Println("requestID:", requestID)
}
