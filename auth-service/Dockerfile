FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /workspace

# Copy go.work and go.work.sum first for caching
COPY go.work go.work.sum ./

# Copy all module go.mod/go.sum files for dependency caching
COPY auth-service/go.mod auth-service/go.sum ./auth-service/
COPY pkg/go.mod pkg/go.sum ./pkg/

# Download dependencies using workspace mode
ENV GOWORK=/workspace/go.work
RUN go mod download -x 2>&1 || true
WORKDIR /workspace/auth-service
RUN go mod download
WORKDIR /workspace/pkg
RUN go mod download

# Copy the full source code
WORKDIR /workspace
COPY auth-service/ ./auth-service/
COPY pkg/ ./pkg/

# Build the auth-service binary
WORKDIR /workspace/auth-service
RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-service-bin ./cmd/main.go

# --- Final stage ---
FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=builder /auth-service-bin .

EXPOSE 8080

CMD ["./auth-service-bin"]
