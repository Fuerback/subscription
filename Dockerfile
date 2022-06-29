FROM golang:1.16 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o /subscription -a -ldflags '-linkmode external -extldflags "-static"' .

EXPOSE 8080

CMD [ "/subscription" ]