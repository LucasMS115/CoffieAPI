.PHONY: up up-db-exposed down logs smoke smoke-db-exposed test test-unit test-integration verify migrate build build-runtime clean docs run

up:
	docker compose up --build -d

up-db-exposed:
	docker compose -f docker-compose.yml -f docker-compose.db-port.yml up --build -d

smoke:
	curl --silent --show-error http://localhost:8080/health

smoke-db-exposed:
	docker compose -f docker-compose.yml -f docker-compose.db-port.yml up --build -d
	curl --silent --show-error http://localhost:8080/health

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

build-runtime:
	docker build -t coffie-api .

run:
	go run ./cmd/server/main.go

docs:
	swag init -g cmd/server/main.go -o docs

clean:
	docker compose down -v
