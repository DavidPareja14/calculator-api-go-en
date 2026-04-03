# Specification (SDD) – Calculator REST API in Go (Clean Architecture + Chi)

## 1. Project Objective
Build a small **Golang** application that exposes a **REST API** to perform basic math operations:

- addition
- subtraction
- multiplication
- division

The main goal of this project is to serve as a baseline to compare token consumption between prompts written in Spanish vs English.  
Therefore, the development must be **clean**, **well-structured**, **reproducible**, and follow good practices.

---

## 2. Language Rules (MANDATORY)
⚠️ **All content must be written entirely in English.**

This includes:

- folder and file names (if applicable)
- code comments
- documentation (README)
- error messages
- API JSON responses
- all text generated in the final response

**Mixing English with any other language is forbidden.**

---

## 3. Functional Specification (Functional Spec)

### 3.1 Required Endpoints
The API must expose 4 independent endpoints:

- `GET /add?a={num}&b={num}`
- `GET /subtract?a={num}&b={num}`
- `GET /multiply?a={num}&b={num}`
- `GET /divide?a={num}&b={num}`

Where:

- `a` and `b` are real numbers (`float64`)
- both are mandatory

---

### 3.2 Successful Response
For successful responses, the API must return HTTP 200 with JSON like:

```json
{
  "operation": "add",
  "a": 10,
  "b": 5,
  "result": 15
}
```

---

### 3.3 Error Handling

#### 3.3.1 Missing or Invalid Parameters
If parameters are missing, invalid, or cannot be converted to a number:

- return HTTP 400
- return JSON with a clear message

Example:

```json
{
  "error": "The parameter 'a' is required and must be numeric."
}
```

---

#### 3.3.2 Division by Zero
Special case: division by zero

- return HTTP 400
- return JSON:

```json
{
  "error": "Division by zero is not allowed."
}
```

---

### 3.4 Healthcheck
There must be an additional endpoint:

- `GET /health`

HTTP 200 response:

```json
{
  "status": "ok"
}
```

---

## 4. Technical Specification (Technical Spec)

### 4.1 Required Architecture
The application must follow **Clean Architecture** principles, separating responsibilities into clear layers.

There must be explicit separation of:

- delivery layer (HTTP handlers)
- application layer (use case)
- domain layer (types, rules, errors)

---

### 4.2 Single Use Case
A single use case must be implemented, for example:

- `Calculate(operation, a, b)`

This use case must support:

- add
- subtract
- multiply
- divide

The operation must be represented as a controlled type, for example:

- constants in the domain (`OperationAdd`, `OperationSubtract`, etc.)
- or an enum-like type based on string

If an invalid operation is received, an error must be returned.

---

### 4.3 Use Case Interface
The handler must not depend on a concrete implementation.

An interface must exist such as:

```go
type CalculatorUseCase interface {
	Calculate(operation string, a float64, b float64) (float64, error)
}
```

And a concrete implementation must exist that fulfills this interface.

---

### 4.4 Handlers
There must be one handler per endpoint:

- addition handler
- subtraction handler
- multiplication handler
- division handler
- health handler

Each handler must:

1. read `a` and `b` from query params
2. validate that they exist
3. convert them to `float64`
4. call the use case with the corresponding operation
5. return a JSON response

---

### 4.5 Mandatory Router
The router must be:

- `github.com/go-chi/chi/v5`

Gin or other alternative frameworks must not be used.

---

### 4.6 Dependency Injection (manual)
No dependency injection frameworks must be used.

All composition must be done manually in `main.go` (composition root).

`main.go` must:

- create the use case implementation
- inject it into the handlers
- register the routes
- start the server

---

### 4.7 Standardized Error Handling
There must be a consistent way to return JSON errors.

Handlers must always return the following format:

```json
{
  "error": "error message"
}
```

---

## 5. Required Folder Structure
The project must be organized using a professional structure. Suggested example:

```text
/cmd/api/main.go
/internal/domain
/internal/usecases
/internal/handlers
/internal/server
```

Notes:

- `cmd/api/main.go` must be the entry point
- `internal/domain` must contain simple business rules (types, constants, errors)
- `internal/usecases` must contain the use case implementation
- `internal/handlers` must contain HTTP handlers
- `internal/server` may contain router/server configuration (optional)

---

## 6. Testing Specification (Testing Spec)

### 6.1 Mandatory Unit Tests
There must be unit tests for the `Calculate` use case that cover:

- addition
- subtraction
- multiplication
- division
- division by zero
- invalid operation

Allowed framework:

- Go standard `testing`

---

### 6.2 HTTP Tests (recommended)
Ideally include HTTP tests using:

- `net/http/httptest`

Validate at least:

- successful `/add` endpoint
- error due to invalid parameter
- division by zero error

---

## 7. Dockerization

### 7.1 Mandatory Dockerfile
A functional `Dockerfile` must be included using multi-stage build:

- build stage using the official Go image
- runtime stage using a lightweight image (for example alpine)

The container must:

- compile the application
- run the binary
- expose the configured port

---

### 7.2 docker-compose (optional)
If included, it must allow running the service easily with:

```bash
docker compose up --build
```

---

## 8. Runtime Configuration

### 8.1 Default Port
The application must run by default on port:

- `8080`

It must allow configuration through an environment variable:

- `PORT`

If not provided, it must use `8080`.

---

## 9. Quality Standards
The code must follow:

- idiomatic Go best practices
- proper error handling
- small and clear functions
- descriptive naming
- real separation of responsibilities
- avoid duplication of logic across handlers

---

## 10. Documentation (README)
A README written in English must be included with:

- project description
- general architecture overview (layers)
- how to run locally
- how to run using Docker
- request examples (curl)
- how to run tests

---

## 11. Mandatory Deliverables
The final result must include:

- full source code
- Clean Architecture structure
- working endpoints
- unit tests
- functional Dockerfile
- complete README

---

## 12. Suggested Task Plan (Task Breakdown)

1. Create base project structure
2. Implement domain (operation constants/types and errors)
3. Implement interface and `Calculate` use case
4. Implement HTTP handlers per endpoint
5. Configure router with Chi
6. Implement healthcheck
7. Implement unit tests for the use case
8. Implement basic HTTP tests (recommended)
9. Create Dockerfile
10. Create README with instructions and examples

---

## 13. Mandatory Final Validation
At the end, the response must include a section with:

- list of endpoints
- curl examples
- command to run tests
- command to build and run with Docker
