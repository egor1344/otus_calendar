package main

import (
	"encoding/json"
	"log"

	"github.com/egor1344/otus_calendar/calendar/internal/domain/models"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"github.com/streadway/amqp"
)

// DbConfig - Конфигурация БД
type DbConfig struct {
	DbDNS string `envconfig:"DB_DSN" required:"true"`
}

// AmqpConfig - Конфигурация RabbitMQ
type AmqpConfig struct {
	AmqpDSN string `envconfig:"AMQP_DSN" required:"true"`
}

type config struct {
	DbConfig
	AmqpConfig
	ServerAddr string `envconfig:"SERVER_ADDR" required:"true"`
	QueueName  string `envconfig:"SERVER_ADDR" required:"true"`
}

func main() {
	var conf config
	err := envconfig.Process("reg_service", &conf)
	if err != nil {
		log.Fatal(err)
	}

	// PostgreSQL
	pgSQL, err := sqlx.Connect("pgx", conf.DbConfig.DbDNS)
	if err != nil {
		log.Fatal(err)
	}
	defer pgSQL.Close()

	// RabbitMQ
	conn, err := amqp.Dial(conf.AmqpConfig.AmqpDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	rmqCh, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer rmqCh.Close()

}

func startWorker(rmqCh *amqp.Channel, pgSQL *sqlx.DB, conf *config) {
	q, err := rmqCh.QueueDeclare(
		conf.QueueName, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
	event := models.Event{Date: ptypes.TimestampNow(), Title: "test", Description: "Description", UserId: 1}
	eventJSON, err := json.Marshal(event)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(eventJSON)
	err = rmqCh.Publish(
		q.Name,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        eventJSON,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
