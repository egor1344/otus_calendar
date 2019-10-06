GOPATH=$(pwd)
PWD=$(pwd)

generate_proto:
	protoc proto/event/event.proto --go_out=plugins=grpc:.
	protoc proto/server/server.proto --go_out=plugins=grpc:.
	protoc proto/calendar/calendar.proto --go_out=plugins=grpc:.

run_grpc_server:
	go run main.go grpc_server

build_server:
	go build -o service main.go

run_grpc_client:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run grps/client/client.go

.PHONY: lint  test build run generate_proto run_grpc_server run_grpc_client