package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ElizCarvalho/fcutils/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// Declarar a fila
	_, err = ch.QueueDeclare(
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
	for i := 0; i < 10; i++ {
		body := "Hello World " + strconv.Itoa(i)
		err = rabbitmq.Publish(ch, body, "amq.direct")
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
		fmt.Printf("Mensagem enviada: %s\n", body)
	}
}
