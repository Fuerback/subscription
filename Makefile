img := subscription

docker-up: docker-image docker-run

docker-image:
	docker build -t $(img) .

docker-run:
	docker run -p 8080:8080 $(img)

docker-tests:
	docker run -p 8080:8080 $(img) go test -cover ./...

run:
	go run .

build:
	go build

run-tests:
	go test ./... -cover