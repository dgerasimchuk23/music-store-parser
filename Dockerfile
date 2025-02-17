# Используем официальный образ Go для сборки бинарника
FROM golang:1.23 AS builder

WORKDIR /app

# Копируем файлы проекта
COPY go.mod go.sum ./
RUN GOPROXY=direct go mod download

COPY . .

# Собираем бинарный файл
RUN go build -o parser main.go

# Используем минимальный образ для продакшена
FROM alpine:latest

WORKDIR /root/

# Копируем скомпилированный бинарник
COPY --from=builder /app/parser .

# Копируем .env (если нужен)
COPY .env .

# Запускаем приложение
CMD ["./parser"]
