# Онлайн калькулятор

## Разворачиваем проект

```cmd
git clone https://github.com/nais2008/project_go_yandex
cd project_fo_yandex
go mod init ../..
```

## Запуск проекта

```cmd
go run ./cmd/main.go
```

## Использование

### Запрос

```cmd
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "выражение"
}'
```

## Варианты ответа

### Запрос 200 (OK)

```cmd
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "1+2"
}'
```

### Ответ 200

```json
{
  "result": "3"
}
```

### Запрос 422 (Unprocessable Entity)

```cmd
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "Слово или неполное выражение"
}'
```

### Ответ 422

```json
{
  "error": "Expression is not valid"
}
```

### Запрос 500 (Internal Server Error)

```cmd
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "Слово или неполное выражение"
}'
```

### Ответ 500

```json
{
  "error": "Internal server error"
}
```
