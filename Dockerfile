FROM golang:1.18

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build

EXPOSE 8080

CMD ["./subscription"]