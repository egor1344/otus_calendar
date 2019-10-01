GOPATH=$(pwd)
PWD=$(pwd)

build:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build .

test:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -v  ./...

lint:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN)  golangci-lint run --enable-all

run:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run cmd/calendar/main.go

generate_proto:
	protoc proto/event/event.proto --go_out=.
	protoc proto/server/server.proto --go_out=plugins=grpc:.
	protoc proto/calendar/calendar.proto --go_out=plugins=grpc:.

run_grpc_server:
	go run main.go grpc_server

build_server:
	go build -o service main.go

run_grpc_client:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run grps/client/client.go

.PHONY: lint  test build run generate_proto run_grpc_server run_grpc_client