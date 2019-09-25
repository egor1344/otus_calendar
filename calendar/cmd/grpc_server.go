package cmd

func construct(dsn string) (*api.CalendarServer, error) {
	eventStorage, err := maindb.NewPgEventStorage(dsn)
	if err != nil {
		return nil, err
	}
	eventService := &services.EventService{
		EventStorage: eventStorage,
	}
	server := &api.CalendarServer{
		EventService: eventService,
	}
	return server, nil
}

var host string
var port string
var dsn string

var GrpcServerCmd = &cobra.Command{
	Use: "grpc_server",
	Short: "run grpc server", 
	Run: func(cmd *cobra.Command, args []string) {
		server, err := construct()
		if err != nil {
			logger.Fatal(err)
		}
		address := host + ":" + port
		err = server.Serve(add)
	}
}

func init() {
	GrpcServerCmd.Flags().StringVar(&host, "host", "localhost", "Host server")
	GrpcServerCmd.Flags().StringVar(&port, "port", "8000", "Port server")
	GrpcServerCmd.Flags().StringVar(&dsn, "dsn", "host=127.0.0.1 user=event_user password=event_pwd dbname=event_db", "databse connection string")
}