package rabbitmq

import "github.com/streadway/amqp"

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery) error {
	msgs, err := ch.Consume(
		"minha-fila",  // queue
		"go-consumer", // consumer
		false,         // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range msgs {
		out <- msg
	}
	return nil
}
