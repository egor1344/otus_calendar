package main

import (
	"encoding/json"
	"log"
	"time"

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
	QueueName string `envconfig:"QUEUE_NAME" required:"true"`
}

// Event - структура события
type Event struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Datetime    time.Time `json:"date_time" db:"date_time"`
	Duration    int64     `json:"duration" db:"duration"`
	Description string    `json:"description" db:"description"`
	UserID      int64     `json:"ower" db:"owner"`
}

func main() {
	var conf config
	err := envconfig.Process("publish", &conf)
	if err != nil {
		log.Fatal(err)
	}

	// PostgreSQL
	log.Println(conf.DbConfig.DbDNS)
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
	startWorker(rmqCh, pgSQL, &conf)
}

func startWorker(rmqCh *amqp.Channel, pgSQL *sqlx.DB, conf *config) {

	eventList := getEventDB(pgSQL)
	for _, e := range eventList {
		log.Println(e)
		pullInQueue(&e, rmqCh, conf)
	}
}

// pullInQueue Отправка сообщений в очередь
func pullInQueue(event *Event, rmqCh *amqp.Channel, conf *config) {
	eventJSON, err := json.Marshal(event)
	if err != nil {
		log.Fatal(err)
	}
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
	err = rmqCh.Publish(
		"",
		q.Name,
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

// getEventDB - получение событий из БД
func getEventDB(pgSQL *sqlx.DB) (eventList []Event) {
	list := []Event{}
	err := pgSQL.Select(&list, "SELECT * FROM events")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(list)
	return list
}
