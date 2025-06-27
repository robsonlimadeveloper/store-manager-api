FROM golang:1.21-alpine

# Instala dependências necessárias para o swag
RUN apk add --no-cache git

# Instala o swag CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy
RUN go mod download

COPY . .

# Entra no diretório do app
WORKDIR /app/app

# Gera a documentação Swagger
RUN swag init --parseDependency --parseInternal

RUN go build -o main .

EXPOSE 8080
CMD ["./main"]
