tests:
	cd backend && go test -v ./...
	
all: docker-build docker-up

docker-build:
	docker-compose build --no-cache

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-clean:
	docker-compose down --volumes --remove-orphans
