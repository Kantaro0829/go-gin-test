FROM golang:1.16


COPY go.mod ./code/go.mod
COPY go.sum ./code/go.sum
WORKDIR /code
RUN ls
#RUN go mod download

COPY . /code

CMD go run main.go