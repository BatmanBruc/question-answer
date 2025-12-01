## Инструкция по использованию сервиса

**Запуск сервиса**

docker compose up

**Создание пользователя**

Получаем токен, для создание ответов
```
curl -X GET http://localhost:8080/user
```
**Создаем вопрос**
```
curl -X POST http://localhost:8080/questions/ \
-H "Content-Type: application/json" \
-d '{"text": "Главный вопрос жизни, вселенной и вообще"}'
```
**Создаем ответ**
```
curl -X POST http://localhost:8080/questions/{id}/answers/ \
-H "Authorization: Bearer token" \
-H "Content-Type: application/json" \
-d '{"text": "42"}'
```

## API Reference

**Questions**

GET /questions/ — список всех вопросов

POST /questions/ — создать новый вопрос

GET /questions/{id} — получить вопрос и все ответы на него

DELETE /questions/{id} — удалить вопрос (вместе с ответами)


**Answers**

POST /questions/{id}/answers/ — добавить ответ к вопросу

GET /answers/{id} — получить конкретный ответ

DELETE /answers/{id} — удалить ответ

**User**

GET /user — Создание пользователя

## Тесты

```
go test ./handlers
```