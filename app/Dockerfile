FROM golang:1.16


COPY go.mod ./app/go.mod
COPY go.sum ./app/go.sum
COPY main.go ./app/main.go
WORKDIR /app
RUN ls

CMD go run main.go