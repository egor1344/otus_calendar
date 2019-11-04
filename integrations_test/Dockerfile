FROM golang:1.13
#RUN apt update && apt upgrade && \
#    apt install --no-cache bash git openssh make
RUN mkdir -p /opt/integration_tests
WORKDIR /opt/integration_tests

COPY go.mod ./
RUN go mod download

COPY . .