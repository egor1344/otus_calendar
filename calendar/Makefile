build:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build .

test:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -v  ./...

lint:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN)  golangci-lint run --enable-all

run:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run cmd/calendar/main.go

generate_proto:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) protoc  models/event/event.proto --go_out=.
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) protoc  models/server/server.proto --go_out=plugins=grpc:.

run_grpc_server:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run grps/server/server.go

run_grpc_client:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run grps/client/client.go

.PHONY: lint  test build run generate_proto run_grpc_server run_grpc_client