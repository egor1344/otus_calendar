package cmd

import (
	"encoding/json"
	//"github.com/egor1344/otus_calendar/calendar/proto/event"

	//"github.com/egor1344/otus_calendar/calendar/internal/domain/models"
	"github.com/egor1344/otus_calendar/calendar/internal/domain/models"
	"github.com/egor1344/otus_calendar/calendar/pkg/logger"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
	"strings"
	"sync"
	"time"
)

// construc - создание подключения к необходимым приложениям
func construc(db_dsn, amqp_dsn, amqp_queue_name string) {
	// PostgreSQL
	zapLog.Info(db_dsn)
	pgSQL, err := sqlx.Connect("pgx", db_dsn)
	if err != nil {
		zapLog.Fatal(err)
	}
	defer pgSQL.Close()

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
	ticker := time.NewTicker(time.Duration(periodScanDb) * time.Second)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ticker.C:
				zapLog.Info("Scan db")
				RunScheduler(rmqCh, pgSQL, amqp_queue_name)
			}
		}
	}()
	wg.Wait()
}

// RunScheduler - запуск планировщика
func RunScheduler(rmqCh *amqp.Channel, pgSQL *sqlx.DB, queue_name string) {
	eventList := getEventDB(pgSQL)
	for _, e := range eventList {
		zapLog.Info(e)
		pullInQueue(&e, rmqCh, queue_name)
	}
	markEventSend(eventList, pgSQL)
	deleteOldEvent(pgSQL)
}

// deleteOldEvent - удаление старых событий
func deleteOldEvent(pgSQL *sqlx.DB) {
	_, err := pgSQL.Exec("DELETE FROM events where (now()- make_interval(years := 1)) >= date_time")
	if err != nil {
		zapLog.Fatal(err)
	}
}

// markEventSend - пометка события завершенным
func markEventSend(eventList []models.Event, pgSQL *sqlx.DB) {
	countEvent := len(eventList)
	if countEvent > 0 {
		var IDList strings.Builder
		IDList.WriteString("update events set send=true where id::uuid in (")

		for i, e := range eventList {
			IDList.WriteString("'")
			IDList.WriteString(e.UUID)
			if i+1 != countEvent {
				IDList.WriteString("',")
			} else {
				IDList.WriteString("');")
			}
		}

		zapLog.Info(IDList.String())
		_, err := pgSQL.Exec(IDList.String())
		if err != nil {
			zapLog.Fatal(err)
		}
	}

}

// pullInQueue Отправка сообщений в очередь
func pullInQueue(event *models.Event, rmqCh *amqp.Channel, queue_name string) {
	eventJSON, err := json.Marshal(event)
	if err != nil {
		zapLog.Fatal(err)
	}
	q, err := rmqCh.QueueDeclare(
		queue_name, // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		zapLog.Fatal(err)
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
		zapLog.Fatal(err)
	}
}

// getEventDB - получение событий из БД
func getEventDB(pgSQL *sqlx.DB) (eventList []models.Event) {
	err := pgSQL.Select(&eventList, "SELECT id, title, date_time, duration, description, owner, before_time_pull"+
		" FROM events WHERE  (send=false and (now() >= date_time))")
	if err != nil {
		zapLog.Fatal(err)
	}
	zapLog.Info(eventList)
	return
}

var dbDsn string
var amqpDsn string
var amqpQueueName string
var periodScanDb int
var zapLog *zap.SugaredLogger

// GrpcServerCmd cobra run server
var SchedulerCmd = &cobra.Command{
	Use:   "scheduler",
	Short: "run scheduler",
	Run: func(cmd *cobra.Command, args []string) {
		construc(dbDsn, amqpDsn, amqpQueueName)
	},
}

func init() {
	//SchedulerCmd.Flags().StringVar(&dbDsn, "db_dsn", "host=db port=5432 user=postgres dbname=postgres sslmode=disable", "databse connection string")
	//SchedulerCmd.Flags().StringVar(&amqpDsn, "amqp_dsn", "amqp://guest:guest@rabbitmq:5672/", "rabbit connection string")
	//SchedulerCmd.Flags().StringVar(&amqpQueueName, "amqp_queue_name", "event_queue", "queue name")
	//SchedulerCmd.Flags().IntVar(&periodScanDb, "perion_scan_db", 1, "queue name")
	l, err := logger.GetLogger()
	zapLog = l
	if err != nil {
		log.Fatal("Error init logger", err)
	}
	err = viper.BindEnv("PERIOD_CLEAR_SECOND")
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
	dbDsn = viper.GetString("db_dsn")
	amqpDsn = viper.GetString("amqp_dsn")
	amqpQueueName = viper.GetString("queue_name")
	periodScanDb = viper.GetInt("period_clear_second")
	zapLog.Info(periodScanDb)
}
