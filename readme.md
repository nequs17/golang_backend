# Серверная часть написана Golang с использованием библиотеки Gorilla/Mux

_Пример работы api http://45.90.218.73:8080/openapi_



## Для сборки проекта

Собрка Golang
```
$ go mod tidy

$ go build -o my_app ./cmd

$ go run my_app

```

## Api
```
-User

/api/sockets/thermalmapdata - получение информации из БД по параметрам (см. в документации Swagger)

/api/sockets/thermalmapdataall - получение всех данных из БД

/api/user/auth - авторизация пользователя, cookie сохраняются

/api/user/logout - удаление cookie

/api/user/register - регистрация пользователя, присаивание группы user 

/api/user/verify - верификация пользователя путём подтверждения почты

-Admin

/api/admin/users - получение всех пользователей 

/api/admin/changerole - изменение роли пользователя (почта / уровень доступа)

```