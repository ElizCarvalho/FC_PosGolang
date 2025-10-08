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

	// Declarar a fila
	q, err := ch.QueueDeclare(
		"minha-fila", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		panic(err)
	}

	// Publicar mensagem
	body := "Hello World"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Mensagem enviada: %s\n", body)
}
