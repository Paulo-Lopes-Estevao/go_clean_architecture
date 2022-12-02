# Go lang Architecture

## Core Bussiness
👉🏾 [Domain](domain.md)

## Installation
  required
  - Golang

install package

`go mod tidy`

start up server

`go run main.go`

access the host `localhost:8000`

***Welcome***
```http
GET http://localhost:8000
Content-Type: application/json
```

***Add Park***
```http
POST http://localhost:8000/park
Content-Type: application/json

{
    "name": "Viana park",
    "vague": 15,
    "limit": 20

}
```

## Architecture

![Clean Architecture](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)