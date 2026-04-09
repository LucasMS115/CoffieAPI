# Docker Run Notes — Issues & Fixes Log

---

## 1. Missing `go.sum` on `COPY`
**Problem:**
```
COPY go.mod go.sum ./
RUN go mod download

failed to calculate checksum of ref ...: "/go.sum": not found
```
The project only uses stdlib — no external dependencies — so `go.sum` was never generated.

**Fix:** Updated Dockerfile to copy `go.sum` only if it exists:
```dockerfile
COPY go.mod ./
RUN if [ -f go.sum ]; then go mod download; fi
```

---

## 2. Missing `.air.toml`
**Problem:** The `CMD` in the Dockerfile referenced `.air.toml` which didn't exist.

**Fix:** Created `.air.toml`:
```toml
[build]
cmd = "go build -o ./tmp/main ./cmd/server/main.go"
bin = "./tmp/main"
include_ext = ["go"]
exclude_dir = ["vendor", "tmp", "tests"]
```

---

## 3. `main.go` was an empty shell
**Problem:** The process exited immediately with code 0. `main.go` only contained comment placeholders — no actual server start logic.

**Fix:** Wrote real startup code that creates the server and calls `ListenAndServe()`.

---

## 4. Import name collision in `main.go`
**Problem:**
```
failed to build, error: exit status 1
```
Both `net/http` and `coffie/internal/http` use the package name `http`. Go can't distinguish `http.NewServer` from `http.ErrServerClosed`.

**Fix:** Aliased the internal package:
```go
import (
	"net/http"
	apphttp "coffie/internal/http"
)
```
Used `apphttp.NewServer(...)` and `http.ErrServerClosed` to disambiguate.

---

## Key Takeaways

- **`go.sum` won't exist** until external dependencies are added. The Dockerfile should accommodate this phase.
- **Shell files must have working minimal behavior** — even a skeleton `main.go` needs actual server startup code, not just comments.
- **Import collisions with stdlib** — our `internal/http` package shares the name `http` with `net/http`. Always use an alias when both are needed.
- **Air config deprecation** — `build.bin` is deprecated in air v1.65; should migrate to `build.entrypoint` later.
