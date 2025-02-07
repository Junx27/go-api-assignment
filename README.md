## Simple to use this program

```bash
go run main.go
```

## End point

### 1. Add user

POST `http://localhost:8080/add-user`

```json
{
  "id": 1,
  "name": "John Doe",
  "age": 30
}
```

### 2. Get User

GET `http://localhost:8080/get-user?id=1`

```json
{
  "id": 1,
  "name": "John Doe",
  "age": 30
}
```
