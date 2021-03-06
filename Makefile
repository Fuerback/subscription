img := subscription

docker-up: build docker-image docker-run

docker-image:
	docker build -t $(img) .

docker-run:
	docker run -p 8080:8080 $(img)

docker-tests:
	docker run -p 8080:8080 $(img) go test ./...

run:
	go run .

build:
	go build

run-tests:
	go test ./... -cover