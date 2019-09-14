FROM golang:1.13-alpine
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh make
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8080