package cmd

import (
	"github.com/egor1344/otus_calendar/calendar/pkg/logger"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
)

// construc - создание подключения к необходимым приложениям
func constructMailing(amqp_dsn, amqp_queue_name string) {

	// RabbitMQ
	zapLog.Info(amqp_dsn)
	conn, err := amqp.Dial(amqp_dsn)
	if err != nil {
		zapLog.Fatal(err)
	}
	defer conn.Close()
	rmqCh, err := conn.Channel()
	if err != nil {
		zapLog.Fatal(err)
	}
	defer rmqCh.Close()

	// Подключение к очереди
	queue, err := rmqCh.QueueDeclare(
		amqp_queue_name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		zapLog.Fatal(err)
	}

	runMailing(rmqCh, queue)

}

// runMailing - запуск расыльщика
func runMailing(rmqCh *amqp.Channel, queue amqp.Queue) {
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
		zapLog.Fatal(err)
	}
	forever := make(chan bool)
	// Получение сообщений из очереди
	go func() {
		for d := range messages {
			// Эмулирование работы отправки сообщения
			zapLog.Info("Received a message: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

var mailingAmqpDsn string
var mailingAmqpQueueName string

// Mailing cobra run server
var MailingCmd = &cobra.Command{
	Use:   "mailing",
	Short: "run mailing",
	Run: func(cmd *cobra.Command, args []string) {
		constructMailing(mailingAmqpDsn, mailingAmqpQueueName)
	},
}

func init() {
	//MailingCmd.Flags().StringVar(&mailingAmqpDsn, "amqp_dsn", "amqp://guest:guest@rabbitmq:5672/", "rabbit connection string")
	//MailingCmd.Flags().StringVar(&mailingAmqpQueueName, "amqp_queue_name", "event_queue", "queue name")
	l, err := logger.GetLogger()
	zapLog = l
	if err != nil {
		log.Fatal("Error init logger", err)
	}
	err = viper.BindEnv("PERIOD_CLEAR_MINUTE")
	if err != nil {
		zapLog.Fatal(err)
	}
	err = viper.BindEnv("QUEUE_NAME")
	if err != nil {
		zapLog.Fatal(err)
	}
	err = viper.BindEnv("AMQP_DSN")
	if err != nil {
		zapLog.Fatal(err)
	}
	err = viper.BindEnv("DB_DSN")
	if err != nil {
		zapLog.Fatal(err)
	}
	viper.AutomaticEnv()
	zapLog.Info(viper.AllSettings())
	mailingAmqpDsn = viper.GetString("amqp_dsn")
	mailingAmqpQueueName = viper.GetString("queue_name")
}
