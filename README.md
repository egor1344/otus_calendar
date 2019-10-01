# otus_calendar

```
protoc proto/server/server.proto --go_out=plugins=grpc:.
protoc proto/event/event.proto --go_out=plugins=grpc:.
protoc proto/calendar/calendar.proto --go_out=plugins=grpc:.
```