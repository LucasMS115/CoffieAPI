.PHONY: up down logs test migrate build clean docs

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

docs:
	swag init -g cmd/server/main.go -o docs

clean:
	docker compose down -v
