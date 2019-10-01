FROM golang:1.13-alpine
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh make
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build_server
COPY service /usr/bin/
RUN chmod +x /usr/bin/service
EXPOSE 8080