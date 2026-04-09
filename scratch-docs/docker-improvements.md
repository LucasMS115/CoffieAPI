# Docker Improvements

## To-do

- Add a Windows-friendly alternative to the `Makefile` workflow, such as a PowerShell helper script or a short command reference that mirrors `make up`, `make smoke`, and `make up-db-exposed`.
- Add a dedicated smoke check for the standalone runtime image built from `Dockerfile`, so the non-Compose path is validated as explicitly as the dev stack.
- Decide whether the current working state should keep the DB-exposed stack running or return to the safer default stack without host port `5432`.

Why these are still TODOs:

- The core Docker improvements are already implemented and validated.
- These remaining items are workflow polish and developer-experience follow-ups, not structural Docker fixes.
- None of them block the current dev image, runtime image, or Compose-based startup path from working correctly.

---

## Done

### 1. Stronger Postgres readiness

- Added a Postgres `healthcheck` to `docker-compose.yml`.
- Changed the API dependency from `service_started` to `service_healthy`.
- Verified in a real startup that the database becomes healthy before the API starts.

Why this change:

- `service_started` only means the database container process has launched. It does not mean PostgreSQL is already ready to accept connections.
- The API now opens the database connection during startup, so a weak dependency gate can produce intermittent boot failures even when the configuration is correct.
- A healthcheck makes startup ordering match real service readiness instead of container start timing.
- This reduces flaky local startup behavior and makes the development stack more predictable.

### 2. Explicit dev image vs runtime image

- Created `Dockerfile.dev` for the Compose-based hot-reload workflow.
- Replaced the root `Dockerfile` with a standalone multi-stage runtime image.
- The root runtime image builds a binary and runs it without relying on a bind mount.

Why this change:

- The previous Docker setup mixed development assumptions with what looked like a generic runtime image.
- That made the image misleading because it only really worked when the source tree was mounted into the container.
- Splitting the images makes the workflow honest and easier to reason about: one image is for local iteration, the other is self-contained.
- A standalone runtime image is a cleaner base for future CI, deployment, and smoke-test workflows.

### 3. Pinned Air version

- Pinned Air in `Dockerfile.dev` with `ARG AIR_VERSION=v1.61.7`.
- Removed the previous floating `@latest` install pattern.

Why this change:

- `@latest` makes rebuilds non-deterministic because the tool can change without any repository change.
- Pinning the version improves reproducibility across teammates, CI machines, and future rebuilds.
- This reduces the risk of local regressions caused by a hot-reload tool upgrade rather than by application code.

### 4. Non-root container execution

- The dev image now creates and runs as `appuser`.
- The runtime image also runs as a non-root user.
- Adjusted writable cache directories so Go tooling can run without root.

Why this change:

- Running application containers as root is an unnecessary privilege escalation in most cases.
- A non-root default is a standard container hardening practice and avoids normalizing root execution as the default behavior.
- Applying it to both the dev and runtime images keeps the environments closer to each other and exposes permission issues earlier.

### 5. Smaller and cleaner Docker build context

- Extended `.dockerignore` to exclude local artifacts such as:
	- `tmp/`
	- `main.exe`
	- `*.exe`
	- `coverage.out`
	- `*.test`
- Left `scratch-docs/` untouched on purpose.

Why this change:

- Files sent in the Docker build context slow down builds even if the final image does not use them.
- Local binaries, temporary build outputs, and test artifacts do not belong in an image build context.
- Keeping them out reduces noise, speeds up builds, and lowers the chance of accidental leakage of local-only artifacts into image layers.
- `scratch-docs/` was intentionally preserved because the scope here was to avoid changing that part of the workspace policy.

### 6. Faster Docker dev loop

- Added named volumes for:
	- Go module cache
	- Go build cache
- Wired `GOMODCACHE` and `GOCACHE` explicitly in the API container.

Why this change:

- The dev container uses a bind mount, which is convenient but usually slower than container-local volumes for repetitive tool output.
- Go repeatedly reads modules and writes build cache data during hot reloads.
- Persisting those caches in named volumes reduces unnecessary rebuild work and usually improves Docker Desktop performance, especially on Windows.
- Explicit cache paths also make the container behavior easier to understand and troubleshoot.

### 7. Optional host exposure for Postgres

- Removed the Postgres host port from the default `docker-compose.yml`.
- Added `docker-compose.db-port.yml` as an opt-in override for `5432:5432`.
- Added `make up-db-exposed` for environments that have `make` installed.

Why this change:

- The API talks to Postgres over the internal Compose network, so the database does not need to be published to the host for normal development.
- Leaving the port closed by default follows the principle of least exposure and avoids unnecessary local port conflicts.
- The override preserves convenience for cases where a host SQL client is actually needed.
- This gives a safer default without removing the common debugging workflow.

### 8. Dev workflow smoke commands

- Added `smoke` and `smoke-db-exposed` targets to `Makefile`.
- Updated `README.md` to document:
	- dev vs runtime Docker workflow
	- optional DB port exposure
	- new smoke commands
	- standalone runtime image build

Why this change:

- Once the Docker workflows became more explicit, it became useful to have a very small validation step for them.
- A smoke command gives a fast answer to the most important question after startup: is the API reachable and responding?
- Recording the workflow in the README reduces ambiguity for future contributors and makes the intended Docker usage discoverable.

### 9. Air permission issue fixed

- While validating the non-root dev setup, Air failed when trying to write the rebuilt binary to `./tmp/main` on the bind mount.
- Updated `.air.toml` so Air builds to `/tmp/air/main` inside the container instead.
- Restarted the API container and confirmed the rebuild succeeds without the permission error.

Why this change:

- The move to non-root execution surfaced a real mismatch between the hot-reload tool and the host-mounted filesystem permissions.
- Writing the binary into a container-local temporary path removes that dependency on host mount write behavior.
- This keeps the security improvement from the non-root user while restoring reliable hot-reload builds.
- It also separates transient build artifacts from the repository working tree more cleanly.

---

## Validation Summary

The following were validated successfully:

- `docker build -t coffie-api .`
- `docker compose config`
- `docker compose -f docker-compose.yml -f docker-compose.db-port.yml config`
- `docker compose up --build -d`
- `docker compose -f docker-compose.yml -f docker-compose.db-port.yml up --build -d`
- `curl http://localhost:8080/health` returned `{"status":"ok"}`
- the DB-exposed workflow published `0.0.0.0:5432->5432/tcp`

## Current Caveat

- `make` is not installed in the current shell on this Windows machine, so the `Makefile` targets were documented and added, but validated through the equivalent `docker compose` commands instead.

Why this matters:

- The Docker workflows themselves are valid, but the developer experience differs by environment.
- On this machine, the Makefile acts more as documented command aliases than as the primary executable interface.
- That is acceptable for now, but it means Windows-friendly helper scripts would still improve usability.

## Current State

- The stack is currently running with the DB-exposed override enabled.
- Active ports:
	- API on `8080`
	- Postgres on `5432`