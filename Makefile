# Load environment variables from app.env file
include app.env

# Export environment variables
export $(shell sed 's/=.*//' app.env)

all: build

build:
	@echo "Building..."
	@go build -o main cmd/main.go

run:
	@go run cmd/main.go

docker-up:
	@echo "Starting Docker Compose services..."
	@docker-compose --env-file app.env -f ./docker/docker-compose.yml up -d

docker-down:
	@echo "Stopping Docker Compose services..."
	@docker-compose --env-file app.env -f ./docker/docker-compose.yml down -v --rmi all

migrate-run:
	migrate -path pkg/database/migrations -database "$(DB_URI)" -verbose up

migrate-revert:
	migrate -path pkg/database/migrations -database "$(DB_URI)" -verbose down 1

migrate-create:
	migrate create -ext sql -dir pkg/database/migrations -seq $(name)

swagger-update:
	@echo "Updating Swagger..."
	@swag init -g cmd/main.go