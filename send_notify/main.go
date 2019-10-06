package main

import (
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/kelseyhightower/envconfig"
	"github.com/streadway/amqp"
)

type AmqpConfig struct {
	AmqpDSN string `envconfig:"AMQP_DSN" required:"true"`
}

type config struct {
	AmqpConfig
	QueueName string `envconfig:"QUEUE_NAME" required:"true"`
}

func main() {
	// Подключение к rabbitmq
	var conf config
	err := envconfig.Process("send_notify", &conf)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := amqp.Dial(conf.AmqpConfig.AmqpDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Подключение к каналу
	rmqCh, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer rmqCh.Close()

	// Подключение к очереди
	queue, err := rmqCh.QueueDeclare(
		conf.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	messages, err := rmqCh.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	forever := make(chan bool)
	// Получение сообщений из очереди
	go func() {
		for d := range messages {
			// Эмулирование работы отправки сообщения
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
