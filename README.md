<h1 align="center">
  <br>
   Golang API Example (Store Manager)
  <br>
</h1>

<p align="justify">
A simple Golang API to manager stores and establishments.
</p>

<p><strong>Develop with:</strong></p>

<p align="left">
	
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=go,postgres,docker,vscode" />
  </a>
</p>

## Getting Started:

Follow the steps below to set up and run the database and backend application locally.

## Prerequisites

*_**[Docker Desktop](https://www.docker.com/products/docker-desktop/)**_ installed on the environment.* (`Windows`)


*_**[Docker Engine and Compose](https://docs.docker.com/engine/install/ubuntu/)**_ installed on the environment.* (`Linux`)

## Installation and Run

Build project

DEBUG_
```sh
$ docker-compose up --build
```
or

```sh
$ docker-compose up --build -d
```
## Swagger Documentation

[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html#/)


### Authenticate and using Swagger Page

1 - http://localhost:8080/v1/api/login

2 - Copy the token returned and use in Authorized Button option:

3 - Paste token with Bearer prefix. Example: _Bearer ue54da221dd..._ 


## Run tests

```sh
$ docker exec -it store-manager-api sh
```

```sh
@container go test ./... -v -cover
```
