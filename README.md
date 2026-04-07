# Coffie API

REST API for managing coffee recipes, built with Go.

Swagger UI: `http://localhost:8080/swagger/index.html`

Run tests in `CoffieAPI` with `go test ./... -v`.

This is a learning project.
I'm using it learn a bit about go and claude code!

---

## Prerequisites

- **Go 1.26+** (for running without Docker)
- **Docker & Docker Compose** (for running with Docker)
- **PostgreSQL 17** (only for running without Docker — Docker Compose provides it automatically)

---

## Running with Docker (recommended)

The easiest way to run the project. Starts both the API and PostgreSQL automatically.

```bash
# Start all services (API + PostgreSQL) in the background
docker compose up -d

# Follow the API logs
docker compose logs -f api

# Stop everything
docker compose down
```

The API runs on `http://localhost:8080` with hot-reload via [air](https://github.com/air-verse/air) — edit a `.go` file and the server restarts automatically.

**Available `make` commands** (works on Linux/macOS; on Windows use `docker compose` directly):

| Command         | Description                        |
|-----------------|------------------------------------|
| `make up`       | Start all services                 |
| `make down`     | Stop all services                  |
| `make logs`     | Follow API logs                    |
| `make test`     | Run tests inside container         |
| `make build`    | Rebuild the API image              |
| `make docs`     | Regenerate Swagger docs            |
| `make clean`    | Stop everything and remove volumes |

---

## Running without Docker

You need PostgreSQL and Go installed locally.

### 1. Start PostgreSQL

Make sure PostgreSQL is running and the database exists:

```bash
createdb coffie_dev  # or use an existing database
```

### 2. Set environment variables

```bash
export DATABASE_URL="postgres://coffie:coffie_pass@localhost:5432/coffie_dev?sslmode=disable"
export API_PORT="8080"
```

Or create a `.env` file and load it with `set -a; . .env; set +a`.

### 3. Run migrations

```bash
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate -path migrations -database "$DATABASE_URL" up
```

### 4. Start the API

```bash
go run ./cmd/server/main.go
```

The server starts on `http://localhost:8080`.

### Swagger docs

Swagger is generated from:

- API metadata annotations in `cmd/server/main.go`
- endpoint annotations in the HTTP handlers under `internal/feature/**/http`

The generated artifacts live in `docs/` and should not be edited manually.

To regenerate them:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
make docs
```

Or run the command directly:

```bash
swag init -g cmd/server/main.go -o docs
```

---

## Testing the API

### Health check

```bash
curl http://localhost:8080/health
```

**Response:**
```json
{"status":"ok"}
```

---


See `service-structure.md` in the project root for a detailed breakdown of each component.
