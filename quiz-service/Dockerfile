FROM golang:alpine AS builder

WORKDIR /app
# Copy the entire workspace to satisfy go.work
COPY . .

WORKDIR /app/quiz-service
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/quiz-service/main .

EXPOSE 8080
CMD ["./main"]
