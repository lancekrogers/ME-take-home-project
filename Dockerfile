FROM golang:1.21-alpine3.18 as builder

COPY . /src

WORKDIR /src/cmd

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o ./challenge ./cmd

CMD ["./challenge"]
