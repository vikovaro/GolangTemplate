## GolangTemplate

Шаблон REST API на Go (Gin) с авторизацией через JWT, ролями и базовыми операциями над пользователями.

## Возможности

- Регистрация пользователя: `POST /api/auth/register`
- Логин и выдача JWT: `POST /api/auth/login`
- Защищенные операции с пользователями:
  - `GET /api/users/:id`
  - `PUT /api/users/:id`
  - `DELETE /api/users/:id`
- JWT (HS256) содержит `user_id` и `role` (`user`/`admin`)
- Пароли хранятся как хеши bcrypt
- Swagger UI доступен по пути `/swagger/*`

## Технологии

- Web: `github.com/gin-gonic/gin`
- ORM/БД: `gorm.io/gorm` + `gorm.io/driver/postgres`
- Авторизация: `github.com/golang-jwt/jwt/v5` (HS256)
- Хеширование паролей: `golang.org/x/crypto/bcrypt`
- Swagger: `github.com/swaggo/gin-swagger` (`/swagger/*any`)
- Логи: `github.com/sirupsen/logrus`

## Требования

- Go (в `go.mod` указан `go 1.25.2`)
- PostgreSQL

## Быстрый старт

1. Создайте переменные окружения (удобно использовать `.env` в корне проекта — его подхватывает `godotenv`).
2. Запустите сервер:

```bash
go run ./cmd/api
```

При старте выполняется `AutoMigrate(&model.User{})`.

Swagger UI:

- `http://localhost:<PORT>/swagger/index.html`

## Конфигурация (env)

Используются следующие переменные:

- `PORT` (строка, default: `8080`)
- `DATABASE_URL` (строка, DSN для PostgreSQL, обязателен)
- `JWT_SECRET` (строка, секрет для подписи JWT, обязателен)

Пример `.env`:

```env
PORT=8080
DATABASE_URL=postgres://user:password@localhost:5432/golangtemplate?sslmode=disable
JWT_SECRET=change-me
```

## JWT и роли

- Авторизация в запросах выполняется заголовком:

`Authorization: Bearer <token>`

- Токен содержит:
  - `user_id` (id пользователя)
  - `role` (`user` или `admin`)
  - `exp` (истечение через 24 часа)

- Ограничения по `PUT /api/users/:id`:
  - редактировать можно только свою запись (`user_id` из токена == `:id`)
  - либо если роль `admin`

- Важно:
  - `DELETE /api/users/:id` требует только аутентификацию (проверка владения/роли в текущей реализации не выполняется).

## API

### Регистрация

`POST /api/auth/register`

Тело запроса (`application/json`):

```json
{
  "username": "string",
  "password": "string (min 6)",
  "phone": "string",
  "email": "string (email)"
}
```

Ответ:

- `201` — созданный пользователь (поле `password` в ответе отсутствует)
- `400` — ошибка валидации/дубликаты (username/email)

Пример:

```bash
curl -X POST http://localhost:8080/api/auth/register ^
  -H "Content-Type: application/json" ^
  -d "{\"username\":\"alice\",\"password\":\"secret123\",\"phone\":\"+123456789\",\"email\":\"alice@example.com\"}"
```

### Логин

`POST /api/auth/login`

Тело запроса:

```json
{
  "username": "string",
  "password": "string"
}
```

Ответ:

- `200`:
```json
{
  "token": "jwt-token"
}
```
- `401` — неверные учетные данные

Пример:

```bash
curl -X POST http://localhost:8080/api/auth/login ^
  -H "Content-Type: application/json" ^
  -d "{\"username\":\"alice\",\"password\":\"secret123\"}"
```

### Получение пользователя

`GET /api/users/:id`

Требуется `Authorization: Bearer <token>`.

Ответ:

- `200` — пользователь
- `404` — пользователь не найден

### Обновление пользователя

`PUT /api/users/:id`

Требуется `Authorization: Bearer <token>`.

Тело запроса: все поля опциональны (поля задаются как `string`):

```json
{
  "email": "optional",
  "username": "optional",
  "phone": "optional",
  "password": "optional (min 6)"
}
```

Ответ:

- `200` — обновленный пользователь
- `400` — ошибка запроса (например, конфликт username)
- `401` — не авторизован
- `403` — запрещено (нет прав на редактирование)
- `404` — пользователь не найден

### Удаление пользователя

`DELETE /api/users/:id`

Требуется `Authorization: Bearer <token>`.

Ответ:

- `204` — удалено
- `500` — ошибка при удалении

## Примечания по структуре проекта

- Точка входа: `cmd/api/main.go`
- Конфигурация: `internal/config/config.go` (поддержка `.env`)
- БД: `internal/database/*` (Postgres + миграции)
- Модули:
  - `internal/modules/auth` (register/login)
  - `internal/modules/user` (CRUD пользователей)
- Middleware:
  - `internal/middleware/auth.go` (проверка JWT)
  - `internal/middleware/role.go` (утилита для require-role; в текущих роутерах напрямую не подключена)
  - `internal/middleware/logger.go`, `internal/middleware/recovery.go`
    