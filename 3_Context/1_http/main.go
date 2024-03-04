package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second) //cancela após timeout de 3 segundos
	defer cancel()                                         //cancela com a chamada da funcao cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		println("Reserva de hotel cancelada. Timeout excedido.")
		return
	case <-time.After(5 * time.Second):
		println("Hotel reservado.")
		return
	}
}
