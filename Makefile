
all: docker-build docker-up

docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-tests:
	docker-compose run backend go test -v ./...

docker-down:
	docker-compose down

docker-clean:
	docker-compose down --volumes --remove-orphans

