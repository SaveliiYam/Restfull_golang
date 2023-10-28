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

## Принцип работы
У пользователя есть возможность создавать задачи (lists). Задачи сохраняются в базу данных под своим номерм. В дальнейшем задачу можно удалять или изменять.
Также пользователь может добавлять подзадачи в основные задачи (items).
Например:
```
Список покупок (list)
    * Хлеб (item)
    * Молоко (item)
```

## Список конечных точек
* http://localhost:8080/auth/sign-up для регистрации
* http://localhost:8080/auth/sign-in для аутентификации (токен действителен 12 часов)
* http://localhost:8080/api/lists для создания и получения списков TODO.
* http://localhost:8080/api/lists/:id для получения, редактирования и удаления опреденного списка задания по переданного id.
* http://localhost:8080/api/lists/:id/items для создания и получения определенной подзадачи в задании
* http://localhost:8080/api/items/:id для получения, изменения и удаления определенной подзадачи по ее id

## Обращение к конечным точкам API

### POST http://localhost:8080/auth/sign-up
#### Пример ввода:
```
{
    "name": "name",
    "username": "username",
    "password": "password"
}
```
#### Пример вывода:
```
{
    "id": 1
}
```
### POST http://localhost:8080/auth/sign-in
#### Пример ввода:
```
{
    "username": "username",
    "password": "password"
}
```
#### Пример вывода:
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTg1MDI5NTQsImlhdCI6MTY5ODQ1OTc1NCwidXNlcl9pZCI6NX0.jiuI6_lwZiG-hh0TzeIesu8u2UlDUTIJ9M3UnaGvKB4"
}
```
### POST http://localhost:8080/api/lists
#### Пример ввода:
```
{
    "title": "Список покупок",
    "description": "Купить завтра"
}
```
#### Пример вывода:
```
{
    "id": 6
}
```
### GET http://localhost:8080/api/lists
#### Пример вывода:
```
{
    "data": [
        {
            "id": 6,
            "title": "Список покупок",
            "description": "Купить завтра"
        }
    ]
}
```
### GET http://localhost:8080/api/lists/6
#### Пример вывода:
```
{
    "id": 6,
    "title": "Список покупок",
    "description": "Купить завтра"
}
```
### PUT http://localhost:8080/api/lists/6
#### Пример ввода:
```
{
    "title": "Список покупок",
    "description": "Купить послезавтра"
}
```
#### Пример вывода:
```
{
    "status": "ok"
}
```
### POST http://localhost:8080/api/lists/6/items
#### Пример ввода:
```
{
    "title": "Хлеб",
    "description": "Белый"
}
```
#### Пример вывода:
```
{
    "id": 3
}
```
### GET http://localhost:8080/api/lists/6/items
#### Пример вывода:
```
[
    {
        "id": 3,
        "title": "Хлеб",
        "description": "Белый",
        "done": false
    },
    {
        "id": 4,
        "title": "Молоко",
        "description": "Деревенское",
        "done": false
    }
]
```
### GET http://localhost:8080/api/items/3
#### Пример вывода:
```
{
    "id": 3,
    "title": "Хлеб",
    "description": "Белый",
    "done": false
}
```