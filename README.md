# Music Store Parser

Этот проект представляет собой парсер цен с сайта **Doctorhead.ru** и сохраняет данные в базу PostgreSQL.

## Запуск с помощью Docker

### **1. Установка Docker и Docker Compose**
Перед запуском убедитесь, что у вас установлены **Docker** и **Docker Compose**.
- [Скачать Docker](https://www.docker.com/get-started)

### **2. Запуск контейнеров**
```sh
docker-compose up -d
```

### **3. Проверка логов**
```sh
docker logs -f music-store-parser
```

## Запуск без `docker-compose`
### **1. Запуск PostgreSQL отдельно**
```sh
docker run --name postgres_db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=parser_db -p 5432:5432 -d postgres:15
```

### **2. Запуск парсера**
```sh
docker run --rm --env-file .env --network host sprut13/music-store-parser:latest
```

## Образ в Docker Hub
Образ доступен по ссылке:  
**[sprut13/music-store-parser](https://hub.docker.com/repository/docker/sprut13/music-store-parser/general)**

## API эндпоинты
- `GET /products` — Получить список товаров
- `POST /products` — Добавить товар
- `PUT /products/:id` — Обновить цену товара
- `DELETE /products/:id` — Удалить товар

## Технологии
- **Go 1.23**
- **PostgreSQL 15**
- **Docker & Docker Compose**
- **Gin (для API)**
- **Goquery (для парсинга HTML)**

## Разработка
Для локального запуска без Docker:
```sh
go run main.go
