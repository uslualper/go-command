vet:
	go vet main.go

build:
	go build -o ./build/main.go 

run:
	go run main.go

prod-run:
	go run ./build/main.go --env=prod

docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-restart:
	docker-compose restart

docker-exec:  
	docker exec -it ${CONTAINER_NAME} bash

test:
	go test ./tests/...