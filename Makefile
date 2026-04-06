.PHONY: up down logs test migrate build clean

up:
	docker compose up --build -d

down:
	docker compose down

logs:
	docker compose logs -f api

test:
	docker compose exec api go test ./...

migrate:
	docker compose exec api go run cmd/migrate/main.go

build:
	docker compose build

clean:
	docker compose down -v
