package publisher

import (
	"errors"
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

var (
	ErrEnv        = errors.New("not found RABBIT_PATH  environment variable")
	ErrEnvChannel = errors.New("not found RABBIT_CHANNEL  environment variable")
	ErrConnect    = errors.New("failed to connect to RabbitMQ")
	ErrChannel    = errors.New("failed to open a channel")
	ErrQueue      = errors.New("failed to declare a queue")
	ErrPublish    = errors.New("failed to publish a message")
)

func SendMessage(msg string) error {
	rabbit_url := os.Getenv("RABBIT_PATH")
	rabbit_ch := os.Getenv("RABBIT_CHANNEL")
	if len(rabbit_url) == 0 {
		return ErrEnv
	}
	if len(rabbit_ch) == 0 {
		return ErrEnvChannel
	}

	conn, err := amqp.Dial(rabbit_url)

	if err != nil {
		defer conn.Close()
		return ErrConnect
	}

	ch, err := conn.Channel()
	if err != nil {
		defer conn.Close()
		return ErrChannel
	}

	q, err := ch.QueueDeclare(
		rabbit_ch, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		defer conn.Close()
		return ErrQueue
	}

	body := msg
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	if err != nil {
		defer conn.Close()
		return ErrPublish
	}

	// fmt.Println("*************** ", conn)
	fmt.Println(" [x] Sent ", body)
	return nil
}
