FROM golang:1.23-alpine AS builder

WORKDIR /build

# Copy only the service's go.mod and go.sum first
COPY go.mod go.sum ./
RUN go mod download

# Copy only the src/cmd directory for this service
COPY cmd/main.go ./
COPY internal/ ./internal/

WORKDIR /build/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main

FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
