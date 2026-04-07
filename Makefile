.PHONY: up down logs test test-unit test-integration verify migrate build clean docs run

up:
	docker compose up --build -d

down:
	docker compose down

logs:
	docker compose logs -f api

test:
	docker compose exec api go test ./...

test-unit:
	go test ./tests/feature/... -v

test-integration:
	go test ./tests/health -v

verify:
	go test ./... -v

migrate:
	docker compose exec api go run cmd/migrate/main.go

build:
	docker compose build

run:
	go run ./cmd/server/main.go

docs:
	swag init -g cmd/server/main.go -o docs

clean:
	docker compose down -v
