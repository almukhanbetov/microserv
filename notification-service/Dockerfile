FROM golang:1.24-alpine

WORKDIR /app

# Установка git
RUN apk update && apk add --no-cache git

# Копируем модули
COPY go.mod ./
COPY go.sum ./

# Используем прямой доступ к GitHub
ENV GOPROXY=direct

# Скачиваем зависимости
RUN go mod download

# Копируем остальной код
COPY . .

# Сборка
RUN go build -o main .

EXPOSE 9093

CMD ["./main"]
