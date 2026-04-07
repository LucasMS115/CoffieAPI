# Coffee API — Service Structure

## High-Level Layout

```text
CoffieAPI/
├── cmd/
│   └── server/
│       └── main.go
├── docs/
├── internal/
│   ├── config/
│   ├── database/
│   ├── http/
│   │   ├── middleware/
│   │   └── response/
│   └── feature/
│       ├── user/
│       ├── coffee/
│       ├── recipe/
│       └── rating/
├── migrations/
├── tests/
│   ├── feature/
│   │   ├── coffee/
│   │   ├── rating/
│   │   ├── recipe/
│   │   └── user/
│   └── health/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── Makefile
└── README.md
```

## Current Bootstrap Flow

The application startup is currently split across focused packages:

1. `cmd/server/main.go` loads runtime configuration with `config.Load()`.
2. `internal/database/postgres.go` opens the PostgreSQL connection with `database.NewPostgresConn(...)`.
3. `internal/http/server.go` creates the `http.Server` and delegates route registration.
4. `internal/http/module_registration.go` wires infrastructure modules and feature modules.

This keeps `main.go` small and avoids leaking HTTP wiring or database setup details into feature code.

## File-by-File Breakdown

### `cmd/server/main.go`

Responsibility: application entry point.

What it does today:

- loads `DATABASE_URL` and `API_PORT` through `config.Load()`
- opens the database connection through `database.NewPostgresConn(...)`
- imports `coffie/docs` with a blank import so Swagger metadata is registered
- creates the HTTP server with `apphttp.NewServer(...)`
- starts `ListenAndServe()`

This file no longer contains custom database bootstrap logic.

### `docs/`

Responsibility: generated Swagger artifacts.

Contents:

- `docs.go`
- `swagger.json`
- `swagger.yaml`

These files are generated from annotations in `cmd/server/main.go` and handler files. They are build artifacts, not the source of truth.

### `internal/config/config.go`

Responsibility: central application configuration loading.

Current behavior:

- reads `DATABASE_URL`
- reads `API_PORT`
- applies local defaults when env vars are missing

This keeps environment access in one place instead of scattering `os.Getenv` across the codebase.

### `internal/database/postgres.go`

Responsibility: database connection creation.

Current behavior:

- opens the PostgreSQL connection pool with `sql.Open`
- verifies connectivity with `Ping`

This package is now the single infrastructure entry point for database startup.

### `internal/http/`

This package contains the shared HTTP infrastructure.

#### `internal/http/server.go`

Responsibility: create the `*http.Server` and the root `ServeMux`.

Current behavior:

- creates `http.NewServeMux()`
- registers Swagger routes
- registers health routes
- registers the user feature module

This file is intentionally thin. The actual route composition lives in helper functions.

#### `internal/http/module_registration.go`

Responsibility: modular route composition.

Current helpers:

- `registerSwaggerModule(...)`
- `registerHealthModule(...)`
- `registerUserModule(...)`

This is the current composition boundary between app infrastructure and feature wiring.

#### `internal/http/health.go`

Responsibility: health endpoint.

Current route:

- `GET /health`

Behavior:

- returns `200 OK`
- returns `{ "status": "ok" }`

#### `internal/http/middleware/`

Responsibility: shared middleware that is not feature-specific.

Current files:

- `auth.go`
- `logger.go`

#### `internal/http/response/json.go`

Responsibility: shared JSON response helpers.

Used for:

- normal JSON responses
- standardized error responses
- field-level validation errors

#### `internal/http/response/domain_error.go`

Responsibility: map known domain errors to HTTP responses.

Current mapping implemented:

- `userdomain.ErrUserAlreadyExists` -> `409 Conflict`
- unknown errors -> `500 Internal Server Error`

This is the start of a centralized error translation layer.

## Feature Layout

The project uses a feature-first structure under `internal/feature/`.

Current feature folders:

- `user/`
- `coffee/`
- `recipe/`
- `rating/`

The layout is not fully identical in every feature yet. The `user` feature is the most evolved one and already has an explicit `module.go`.

### Important Correction

The project does not currently use a `service/` directory inside each feature.

Business logic lives in `domain/*_service.go`, for example:

- `internal/feature/user/domain/user_service.go`
- `internal/feature/coffee/domain/coffee_service.go`
- `internal/feature/recipe/domain/recipe_service.go`
- `internal/feature/rating/domain/rating_service.go`

The project also does not currently use nested `request/` and `response/` folders under each feature HTTP package. Request and response types live in files such as `user_request.go` and `user_response.go`.

### `internal/feature/user/`

Current structure:

```text
user/
├── domain/
├── http/
├── store/
└── module.go
```

This is the current reference feature for module boundaries.

#### `internal/feature/user/module.go`

Responsibility: own the user feature composition.

Current behavior:

- creates the store with `userstore.NewUserStore(...)`
- creates the service with `userdomain.NewService(...)`
- creates the handler with `userhttp.NewHandler(...)`
- registers feature routes through `RegisterRoutes(...)`

This prevents `internal/http/module_registration.go` from having to know each internal dependency of the feature.

#### `internal/feature/user/domain/`

Responsibility: user entities, business logic, contracts, and domain errors.

Current files include:

- `user.go`
- `user_service.go`
- `user_store.go`
- `errors.go`

Notable current state:

- the user flow is creation-focused
- `UserStats` was removed
- the service exposes user registration behavior
- the store contract is minimal and focused on creation
- duplicate user creation is represented as `ErrUserAlreadyExists`

#### `internal/feature/user/store/`

Responsibility: persistence implementation for the user feature.

Current behavior:

- inserts users into PostgreSQL
- translates PostgreSQL unique constraint violations into `domain.ErrUserAlreadyExists`

#### `internal/feature/user/http/`

Responsibility: transport layer for the user feature.

Current files include:

- `user_handler.go`
- `user_request.go`
- `user_response.go`
- `user_adapters.go`
- `user_service.go`

Current HTTP behavior:

- route: `POST /api/users`
- request validation lives on `RegisterUser.Validate()`
- handler depends on a local `UserService` interface instead of a concrete service type
- domain errors are translated through `response.DomainError(...)`

### `internal/feature/coffee/`, `recipe/`, and `rating/`

These features currently follow a lighter structure:

```text
<feature>/
├── domain/
├── http/
├── store/
└── <feature>_adapters.go
```

They already use the same broad layering idea, but they do not yet have the same explicit module boundary as `user/module.go`.

## Migrations

The `migrations/` directory holds schema changes as versioned SQL files.

Current purpose:

- create the `users` table
- create the `coffees` table
- create the `recipes` table
- create the `ratings` table

These files remain the source of truth for database structure.

## Tests

The test layout has changed and should be documented as follows.

### `tests/feature/`

Feature-focused tests live here.

Current subfolders:

- `tests/feature/user/`
- `tests/feature/coffee/`
- `tests/feature/recipe/`
- `tests/feature/rating/`

The `user` feature has the most active coverage right now, including:

- service tests
- store tests
- HTTP handler tests
- request validation tests
- shared test helpers

### `tests/health/`

Health endpoint tests live here.

This replaces the older documentation that mentioned a `tests/integration/` layout.

## Routes Currently Wired by the Server

As of the current implementation, the server registers these modules:

- Swagger UI
- health
- user

That means these routes are currently active in the main server wiring:

- `GET /swagger/`
- `GET /health`
- `POST /api/users`

Other feature folders already exist in the repository, but they are not described here as fully wired modules unless they are actually registered by `internal/http/module_registration.go`.

## Swagger Workflow

Swagger generation works like this:

1. API metadata lives in `cmd/server/main.go`.
2. Endpoint annotations live in handler files.
3. `swag init -g cmd/server/main.go -o docs` regenerates the artifacts in `docs/`.
4. `make docs` wraps the same command.

The Swagger UI is served at:

- `http://localhost:8080/swagger/index.html`

## Summary

The current architecture is moving toward clearer module boundaries, but it is not yet fully uniform across all features.

The most important truths about the current structure are:

- startup is split between `config`, `database`, and `http`
- route wiring is centralized in `internal/http/module_registration.go`
- `user` is the feature with the clearest module boundary today
- business logic currently lives under `domain/*_service.go`, not under `service/`
- tests are organized under `tests/feature/` and `tests/health/`
- Swagger output is generated into `docs/`
| Path | Purpose |
|------|---------|
| `tests/integration/` | Full HTTP flow tests (request → response) with real handler wiring |
| `tests/integration/helpers/` | Shared test server setup, fixture data, DB seed functions |

**Testing strategy per base-instructions:**
- **Unit tests** — target the `service/` layer, mock the store, format: Given/When/Then
- **Integration tests** — use `httptest` to exercise the full HTTP handler chain
