package sendmsg
import (
	// "fmt"
	"os"
	"log"
	"errors"
	"github.com/streadway/amqp"
)


var (
	ErrConnect = errors.New("failed to connect to RabbitMQ")
	ErrChannel = errors.New("failed to open a channel")
	ErrQueue = errors.New("failed to declare a queue")
	ErrPublish = errors.New("failed to publish a message")
)


func SendMessage(msg string) error{
	conn, err := amqp.Dial(os.Getenv("RABBIT_PATH"))

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
		"postUser", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
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

	log.Printf(" [x] Sent %s", body)
  return nil
}
