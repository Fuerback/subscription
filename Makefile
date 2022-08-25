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

tests:
	go test ./... -cover

mock-update:
	mockgen -source=./adapter/http/rest/subscriptionservice/new.go -destination=./core/domain/mocks/fakesubscriptionservice.go -package=mocks
	mockgen -source=./adapter/http/rest/productservice/new.go -destination=./core/domain/mocks/fakeproductservice.go -package=mocks
	mockgen -source=./core/usecase/productusecase/new.go -destination=./core/domain/mocks/fakeproductusecase.go -package=mocks
	mockgen -source=./core/usecase/subscriptionusecase/new.go -destination=./core/domain/mocks/fakesubscriptionusecase.go -package=mocks
	mockgen -source=./adapter/sqlite/productrepository/new.go -destination=./core/domain/mocks/fakeproductrepository.go -package=mocks
	mockgen -source=./adapter/sqlite/subscriptionrepository/new.go -destination=./core/domain/mocks/fakesubscriptionrepository.go -package=mocks

fmt:
	go fmt ./...