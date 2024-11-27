
FROM golang:1.22.0 AS builder


WORKDIR /app
ENV GOARCH=amd64
ENV GOOS=linux

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/server


FROM ubuntu:latest

RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /app/main /app/main

COPY --from=builder /app/internal/database/migrations /app/internal/database/migrations

# Устанавливаем рабочую директорию
WORKDIR /app

# Даем права на выполнение для файла main
RUN chmod +x /app/main

EXPOSE 8080

# Выполняем миграции и затем запускаем приложение
CMD ["sh", "-c", "./main migrate && ./main"]
