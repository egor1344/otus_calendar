FROM golang:1.13
#RUN apt update && apt upgrade && \
#    apt install --no-cache bash git openssh make
RUN mkdir -p /opt/integration_tests
WORKDIR /opt/integration_tests

COPY go.mod go.sum ./
RUN go mod download

COPY . .
CMD ["go", "test"]