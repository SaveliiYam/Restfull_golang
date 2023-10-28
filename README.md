# REST API Для создания TODO списков на Golang

## В сервисе используется

* Дизайн REST API 
* Работа с фреймворком gin.
* Техника внедрения зваисимостей. 
* Используется БД PostgresSQL из Docker, работа с библиотекой sqlx.
* Конфигурация собирается с помощью spf13/viper + использование переменных окружения.
* Регистрация и аутентификация. Работа JWT токенами.


## Для запуска приложения 
```
make build && make run
```

Также при первом запуске необходимо применить миграцию баз данных

```
make migrate
```

## Список конечных точек
* localhost:8080/auth/sign-up для регистрации
* localhost:8080/auth/sign-in для аутентификации (токен действителен 12 часов)
* localhost:8080/api/lists для создания и получения списков TODO.
* localhost:8080/api/lists/:id для получения, редактирования и удаления опреденного списка задания по переданного id.
* localhost:8080/api/lists/:id/items для создания и получения определенной подзадачи в задании
* localhost:8080/api/items/:id для получения, изменения и удаления определенной подзадачи по ее id