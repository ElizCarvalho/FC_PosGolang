package main

import (
	"fmt"

	"github.com/ElizCarvalho/fcutils/rabbitmq"
	"github.com/streadway/amqp"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	err = rabbitmq.Consume(ch, msgs, "minha-fila")
	if err != nil {
		panic(err)
	}

	fmt.Println("🎧 Aguardando mensagens...")
	for msg := range msgs {
		fmt.Println("📨 Mensagem recebida:", string(msg.Body))
		msg.Ack(false) // requeue false para não reenviar a mensagem
	}
}
