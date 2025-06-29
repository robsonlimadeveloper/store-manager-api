<h1 align="center">
  <br>
   Golang API Example (Store Manager)
  <br>
</h1>

<p align="justify">
RESTful API for managing establishments and stores, with JWT authentication, Swagger documentation and modular structure.
</p>

<p><strong>Develop with:</strong></p>

<p align="left">
	
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=go,postgres,docker,vscode" />
  </a>
</p>

## 📌 Technologies Used

- Golang

- Echo Framework

- PostgreSQL

- JWT (Autenticação simulada)

- Swagger (via Swag)

- Docker + Docker Compose

- GoMock + Testify (testes unitários)

## 📁 Directory Structure

```txt
store-manager-api/
├── app/
│   ├── core/
│   ├── config/
│   ├── docs/
│   ├── middlewares/
│   ├── modules/
│   │   ├── establishment/
│   │   └── store/
│   └── routes/
├── scripts/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Getting Started:

Follow the steps below to set up and run the database and backend application locally.

## Prerequisites

*_**[Docker Desktop](https://www.docker.com/products/docker-desktop/)**_ installed on the environment.* (`Windows`)


*_**[Docker Engine and Compose](https://docs.docker.com/engine/install/ubuntu/)**_ installed on the environment.* (`Linux`)

## Installation and Run

Steps:
```bash

# Clone the repository

git clone https://github.com/your-user/store-manager-api.git
cd store-manager-api

# Run the environment with PostgreSQL

docker-compose up --build
```

The API will be available at:
📍 http://localhost:8080

The Swagger documentation will be at:
📄 [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html#/)

### 🔐 (Simulated) Authentication
Use the login endpoint to get a JWT token:

POST /v1/api/login

```json
{
"username": "admin",
"password": "123"
}
```
Use the returned token as a Bearer Token in Swagger or in the headers of protected requests.

### Authenticate and using Swagger Page

1 - http://localhost:8080/v1/api/login

2 - Copy the token returned and use in Authorized Button option:

3 - Paste token with Bearer prefix. Example: _Bearer ue54da221dd..._ 


## 🧪 Run tests

```sh
# Enter the container

$ docker exec -it store-manager-api sh

# Run tests

@container go test ./... -v -cover

```
