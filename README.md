# Hotel Booking Service

Сервис бронирования номеров в отеле с поддержкой управления квотами и валидацией заказов.

## Структура проекта

```plaintext
.
├── cmd
│   └── api
│       └── main.go
├── internal
│   ├── domain
│   │   ├── models
│   │   │   └── order.go
│   │   └── repository
│   │       └── repository.go
│   ├── infrastructure
│   │   └── memory
│   │       ├── order_repository.go
│   │       └── availability_repository.go
│   ├── application
│   │   └── booking
│   │       └── service.go
│   └── interfaces
│       └── http
│           └── handlers
│               └── booking.go
├── Makefile
├── go.mod
└── README.md
```

## Быстрый старт

1. Клонировать репозиторий
2. Выполнить команду `make run`
3. Сервис будет доступен на http://localhost:8080

## Доступные make команды

- `make run` - запуск сервиса
- `make test` - запуск тестов
- `make lint` - проверка кода линтером
- `make build` - сборка бинарного файла

## API Endpoints

### Создание заказа

`POST /orders`

Request body:

```json
{
    "hotel_id": "reddison",
    "room_id": "lux",
    "email": "guest@mail.ru",
    "from": "2024-01-02T00:00:00Z",
    "to": "2024-01-04T00:00:00Z"
}
```

Response (201 Created):

```json
{
    "hotel_id": "reddison",
    "room_id": "lux",
    "email": "guest@mail.ru",
    "from": "2024-01-02T00:00:00Z",
    "to": "2024-01-04T00:00:00Z"
}
```

## Обработка ошибок

Сервис возвращает следующие HTTP статусы:

- 201: Заказ успешно создан
- 400: Ошибка в запросе (неверный формат, некорректные даты)
- 405: Неверный HTTP метод
- 500: Внутренняя ошибка сервера

## Тестирование

В репозитории доступна Postman-коллекция для тестирования API. Импортируйте файл `Hotel_Booking_Service.postman_collection.json` в ваш Postman для начала тестирования.

## Дальнейшие улучшения

- Добавление email-уведомлений
- Реализация промокодов и скидок
- Программа лояльности
- Поддержка бронирования нескольких номеров
