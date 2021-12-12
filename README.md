# gprc-demo

## Деплой

Необходим [docker-compose V2](https://github.com/docker/compose/tree/v2).
```shell
make
```

## Примеры запросов

Все запросы в примерах делаются при помощи [gRPCurl](https://github.com/fullstorydev/grpcurl/).

### AddUser 
```shell
$ grpcurl -d '{"role":"USER", "name":"vasya"}' -plaintext localhost:8080 users.UsersService/AddUser
{
  "uid": {
    "id": 1
  },
  "role": "ADMIN",
  "name": "vasya",
  "lastUpdatedAt": "2021-12-12T18:28:49.609502Z"
}
```

### GetUser

```shell
$ grpcurl -d '{"id":1}' -plaintext localhost:8080 users.UsersService/GetUser
{
  "uid": {
    "id": 1
  },
  "role": "ADMIN",
  "name": "vasya",
  "lastUpdatedAt": "2021-12-12T18:28:49.609502Z"
}
```

### DeleteUser
```shell
$ grpcurl -d '{"id":1}' -plaintext localhost:8080 users.UsersService/DeleteUser
{
  "uid": {
    "id": 1
  },
  "role": "ADMIN",
  "name": "vasya",
  "lastUpdatedAt": "2021-12-12T18:51:37.161437Z"
}

$ grpcurl -d '{"id":1}' -plaintext localhost:8080 users.UsersService/GetUser   
ERROR:
  Code: Unknown
  Message: no rows in result set
```

### GetAllUsers
```shell
$ grpcurl -d '{"role":"ADMIN", "name":"petya"}' -plaintext localhost:8080 users.UsersService/AddUser           
...

$ grpcurl -d '{"role":"USER", "name":"sasha"}' -plaintext localhost:8080 users.UsersService/AddUser
...

$ grpcurl -d '{"role":"MODERATOR", "name":"grisha"}' -plaintext localhost:8080 users.UsersService/AddUser
...

$ grpcurl -plaintext localhost:8080 users.UsersService/GetAllUsers
{
  "users": [
    {
      "uid": {
        "id": 2
      },
      "role": "ADMIN",
      "name": "petya",
      "lastUpdatedAt": "2021-12-12T19:00:34.765087Z"
    },
    {
      "uid": {
        "id": 3
      },
      "name": "sasha",
      "lastUpdatedAt": "2021-12-12T19:00:57.107723Z"
    },
    {
      "uid": {
        "id": 4
      },
      "role": "MODERATOR",
      "name": "grisha",
      "lastUpdatedAt": "2021-12-12T19:01:06.455573Z"
    }
  ]
}
```

### ChangeUserRole
```shell
$ grpcurl -d '{"uid": {"id": 2}, "new_role":"USER"}' -plaintext localhost:8080 users.UsersService/ChangeUserRole
{
  "uid": {
    "id": 2
  },
  "name": "petya",
  "lastUpdatedAt": "2021-12-12T19:05:57.122279Z"
}
```
