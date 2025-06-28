FROM golang:1.23-alpine

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

# Instala o mockgen para gerar mocks
RUN go install go.uber.org/mock/mockgen@latest

# Gera os mocks TODO: Descomentar quando necessário
COPY scripts/generate_mocks.sh /scripts/generate_mocks.sh
RUN chmod +x /scripts/generate_mocks.sh && ./scripts/generate_mocks.sh

# Entra no diretório do app
WORKDIR /app/app

# Gera a documentação Swagger
RUN swag init --parseDependency --parseInternal

#Roda os testes
#RUN go test ./app/... -v

RUN go build -o main .

EXPOSE 8080
CMD ["./main"]
