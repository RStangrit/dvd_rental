package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"main/config"
	emailclient "main/pkg/email_client"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Notification struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
}

func connect() (*amqp.Connection, *amqp.Channel) {
	params := config.LoadConfig()

	// conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@rabbitmq:5672/", params.RABBITMQ_USER, params.RABBITMQ_PASSWORD))
	conn, err := connectWithRetry(params)

	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	_, err = ch.QueueDeclare(
		"user_registered",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare queue")

	return conn, ch
}

func PublishRegistrationNotification(sender, receiver, subject, body string) {
	conn, ch := connect()
	defer conn.Close()
	defer ch.Close()

	notif := Notification{
		Sender:   sender,
		Receiver: receiver,
		Subject:  subject,
		Body:     body,
	}
	serializedBody, err := json.Marshal(notif)
	failOnError(err, "Failed to serialize message")

	err = ch.Publish(
		"",
		"user_registered",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        serializedBody,
		},
	)
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent notification for %s\n", receiver)
}

func StartConsumer() {
	conn, ch := connect()
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		"user_registered",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register consumer")

	log.Println(" [*] Waiting for messages...")
	for d := range msgs {
		var notif Notification
		if err := json.Unmarshal(d.Body, &notif); err != nil {
			log.Println("Invalid message:", err)
			continue
		}
		SendEmailNotification(notif.Sender, notif.Receiver, notif.Subject, notif.Body)
	}

}

func SendEmailNotification(sender, receiver, subject, body string) {
	emailclient.SendEmail(sender, receiver, subject, body)
}

func failOnError(err error, msg string) {
	if err != nil {
		// log.Fatalf("%s: %s", msg, err)
		fmt.Println("%s: %s", msg, err)
	}
}

func connectWithRetry(params *config.Config) (*amqp.Connection, error) {
	var conn *amqp.Connection
	var err error

	rabbitmqURL := fmt.Sprintf("amqp://%s:%s@rabbitmq:5672/", params.RABBITMQ_USER, params.RABBITMQ_PASSWORD)

	for i := 0; i < 10; i++ {
		conn, err = amqp.Dial(rabbitmqURL)
		if err == nil {
			log.Println("Successfully connected to RabbitMQ.")
			return conn, nil
		}
		log.Printf("RabbitMQ not ready (attempt %d/10), retrying in 3s...\n", i+1)
		time.Sleep(3 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to RabbitMQ after retries: %w", err)
}
