**Задание:**

Разработать сервис для управления задачами.

**Запуск приложения:**

1. Запуск docker-compose:

`docker-compose up`

2. Применение миграций:

`migrate -path ./migrations -database postgres://postgres:mypass@localhost:5428/test?sslmode=disable up`

**Документация приложения:**

http://localhost:8000/swagger/index.html
