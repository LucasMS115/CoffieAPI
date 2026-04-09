# Coffie API

REST API for managing coffee recipes, built with Go.

Swagger UI: `http://localhost:8080/swagger/index.html`

Run tests in `CoffieAPI` with `go test ./... -v` or `go test ./tests/feature/user -v` for a specific feature.

This is a learning project.
I'm using it learn a bit about go and claude code!

---

## Prerequisites

- **Go 1.26+** (for running without Docker)
- **Docker & Docker Compose** (for running with Docker)
- **PostgreSQL 17** (only for running without Docker — Docker Compose provides it automatically)

---

## Running with Docker (recommended)

```bash
# Start all services (API + PostgreSQL) in the background
docker compose up -d

# Follow the API logs
docker compose logs -f api

# Stop everything
docker compose down
```

The default Docker workflow is development-oriented.

- `docker compose up` uses `Dockerfile.dev`
- the API runs with hot-reload via [air](https://github.com/air-verse/air)
- the source tree is bind-mounted into the container
- Go module and build caches are persisted in named Docker volumes for faster rebuilds
- PostgreSQL is health-checked before the API starts

This starts both the API and PostgreSQL automatically.


The API runs on `http://localhost:8080` with hot-reload via [air](https://github.com/air-verse/air) — edit a `.go` file and the server restarts automatically.

PostgreSQL is no longer exposed to the host by default. If you want to connect to it from your host machine with a SQL client, use:

```bash
docker compose -f docker-compose.yml -f docker-compose.db-port.yml up -d
```

To verify the running stack quickly:

```bash
make smoke
```

**Available `make` commands** (works on Linux/macOS; on Windows use `docker compose` directly):

| Command | Description |
|-------------------|------------------------------------------|
| `make up` | Start the dev stack |
| `make up-db-exposed` | Start the dev stack with Postgres exposed on the host |
| `make smoke` | Call the health endpoint on the running stack |
| `make smoke-db-exposed` | Start the DB-exposed dev stack and call the health endpoint |
| `make down` | Stop all services |
| `make logs` | Follow API logs |
| `make test` | Run tests inside the dev container |
| `make test-unit` | Run unit/feature tests locally |
| `make test-integration` | Run health/integration tests locally |
| `make verify` | Run the full local Go test suite |
| `make build` | Build the dev image used by Compose |
| `make build-runtime` | Build the standalone runtime image |
| `make run` | Run the API locally without Docker |
| `make docs` | Regenerate Swagger docs |
| `make clean` | Stop everything and remove volumes |

### Building the standalone runtime image

The root `Dockerfile` now builds a self-contained runtime image:

- multi-stage build
- no source bind mount required
- runs as a non-root user

Build it with:

```bash
docker build -t coffie-api .
```

Or via Make:

```bash
make build-runtime
```

Run it with your own `DATABASE_URL` and `API_PORT` values:

```bash
docker run --rm -p 8080:8080 \
	-e DATABASE_URL="postgres://coffie:coffie_pass@host.docker.internal:5432/coffie_dev?sslmode=disable" \
	-e API_PORT="8080" \
	coffie-api
```

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
