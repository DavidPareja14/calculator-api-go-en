# Calculator REST API (Go)

REST API for basic arithmetic operations (add, subtract, multiply, divide) written in Go with Chi and Clean Architecture. It serves as a small baseline for comparing token usage between English and Spanish prompts.

All documentation in this file is in English.

---

## Architecture overview

The project follows Clean Architecture with explicit layers:

- **Domain** (`internal/domain`): operation types, parsing, domain errors.
- **Application** (`internal/usecases`): `Calculator` use case behind the `CalculatorUseCase` interface.
- **Delivery** (`internal/handlers`): HTTP handlers, query validation, JSON responses.
- **Composition** (`internal/server`, `cmd/api`): Chi router wiring and application entry point.

---

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/health` | Health check |
| `GET` | `/add?a={num}&b={num}` | Addition |
| `GET` | `/subtract?a={num}&b={num}` | Subtraction |
| `GET` | `/multiply?a={num}&b={num}` | Multiplication |
| `GET` | `/divide?a={num}&b={num}` | Division |

Parameters `a` and `b` are required numeric values. The server listens on port **8080** by default, configurable with the **`PORT`** environment variable.

---

## Run locally

Prerequisite: Go 1.22 or newer.

```bash
go run ./cmd/api
```

Optional port:

```bash
PORT=3000 go run ./cmd/api
```

---

## Run tests

```bash
go test ./...
```

---

## Docker

Build (from the repository root):

```bash
docker build -t calculator-api .
```

Run (publish container port 8080 on the host):

```bash
docker run --rm -p 8080:8080 calculator-api
```

Custom port inside the container:

```bash
docker run --rm -e PORT=3000 -p 3000:3000 calculator-api
```

The image uses a **multi-stage** build: compile with the official Go image, run the binary in a minimal **Alpine** image. Port **8080** is exposed; the app respects **`PORT`** at runtime.

---

## Request examples (`curl`)

Health:

```bash
curl -s "http://localhost:8080/health"
```

Successful operations:

```bash
curl -s "http://localhost:8080/add?a=10&b=5"
curl -s "http://localhost:8080/subtract?a=10&b=3"
curl -s "http://localhost:8080/multiply?a=6&b=7"
curl -s "http://localhost:8080/divide?a=20&b=4"
```

Invalid parameter (HTTP 400):

```bash
curl -s "http://localhost:8080/add?a=not-a-number&b=2"
```

Division by zero (HTTP 400):

```bash
curl -s "http://localhost:8080/divide?a=1&b=0"
```

---

## Quick reference (validation)

- **Endpoints:** `GET /health`, `GET /add`, `GET /subtract`, `GET /multiply`, `GET /divide` (with `a` and `b` query parameters where applicable).
- **Tests:** `go test ./...`
- **Docker build and run:** `docker build -t calculator-api .` then `docker run --rm -p 8080:8080 calculator-api`

## Prompt in english

-> The following prompt is the one I used to create the application

Implement the project by following **exactly** the specification defined in:

@docs/specification_calculator_api.md

Mandatory rules:
- Everything must be written in English (README, errors, JSON responses, comments).
- Use Clean Architecture.
- Use Chi (`github.com/go-chi/chi/v5`).
- Include tests.
- Include a functional Dockerfile.
- Include a README.

Implement using the following mandatory phases:
1. Base project structure.
2. Domain and `Calculate` use case.
3. HTTP handlers and routes.
4. Tests.
5. Dockerfile and README.

At the end of each phase:
- stop immediately,
- do not explain anything else,
- only respond: "Phase X completed. Type 'continue' to proceed."

Do not continue to the next phase until I say: "continue".