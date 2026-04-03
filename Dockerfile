# Build stage — official Go image
FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /calculator-api ./cmd/api

# Runtime stage — lightweight Alpine image
FROM alpine:3.20

WORKDIR /app

COPY --from=build /calculator-api ./calculator-api

ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["./calculator-api"]
