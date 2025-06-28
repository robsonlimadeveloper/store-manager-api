#!/bin/sh
set -e

# Diretório raiz do app
cd app

echo "Gerando mocks para módulo Store..."
mockgen -source=modules/store/repository.go -destination=modules/store/mocks/mock_store_repository.go -package=store
mockgen -source=modules/store/service.go -destination=modules/store/mocks/mock_store_service.go -package=store

echo "Gerando mocks para módulo Establishment..."
mockgen -source=modules/establishment/repository.go -destination=modules/establishment/mocks/mock_establishment_repository.go -package=establishment
mockgen -source=modules/establishment/service.go -destination=modules/establishment/mocks/mock_establishment_service.go -package=establishment