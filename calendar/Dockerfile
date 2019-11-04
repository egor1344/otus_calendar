FROM golang:1.13-alpine
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh make
WORKDIR /app
COPY go.mod go.sum ./
COPY wait-for /usr/bin/wait-for
RUN chmod +x /usr/bin/wait-for
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
