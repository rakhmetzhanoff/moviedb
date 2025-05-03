FROM golang:1.23.5 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o moviedb ./cmd/main.go

FROM ubuntu:22.04
WORKDIR /app
COPY --from=builder /app/moviedb .
COPY --from=builder /app/internal/db/migrations ./internal/db/migrations
EXPOSE 8080
CMD ["./moviedb"]
