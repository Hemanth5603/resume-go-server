.PHONY: help build run test docker-build docker-run docker-compose-up docker-compose-down deploy-gcp deploy-cloud-run clean

# Variables
APP_NAME=resume-go-server
DOCKER_IMAGE=$(APP_NAME):latest
GCP_PROJECT_ID?=your-project-id
GCP_REGION?=us-central1
GCP_ZONE?=us-central1-a
INSTANCE_NAME?=resume-server

help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build the Go application
	go build -o bin/server cmd/server/main.go

run: ## Run the application locally
	go run cmd/server/main.go

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

docker-build: ## Build Docker image
	docker build -t $(DOCKER_IMAGE) .

docker-run: ## Run Docker container
	docker run -p 3000:3000 --env-file .env $(DOCKER_IMAGE)

docker-compose-up: ## Start services with docker-compose
	docker-compose up -d --build

docker-compose-down: ## Stop services with docker-compose
	docker-compose down

docker-compose-logs: ## View docker-compose logs
	docker-compose logs -f

docker-push-gcr: ## Push image to Google Container Registry
	docker tag $(DOCKER_IMAGE) gcr.io/$(GCP_PROJECT_ID)/$(APP_NAME):latest
	docker push gcr.io/$(GCP_PROJECT_ID)/$(APP_NAME):latest

deploy-gcp: ## Deploy to GCP VM
	@echo "Deploying to GCP VM..."
	chmod +x scripts/deploy.sh
	GCP_PROJECT_ID=$(GCP_PROJECT_ID) ZONE=$(GCP_ZONE) INSTANCE_NAME=$(INSTANCE_NAME) ./scripts/deploy.sh

deploy-cloud-run: ## Deploy to Google Cloud Run
	gcloud run deploy $(APP_NAME) \
		--source . \
		--platform managed \
		--region $(GCP_REGION) \
		--allow-unauthenticated \
		--project $(GCP_PROJECT_ID)

update-deployment: ## Update existing GCP deployment
	chmod +x scripts/update.sh
	INSTANCE_NAME=$(INSTANCE_NAME) ZONE=$(GCP_ZONE) ./scripts/update.sh

view-logs: ## View logs from GCP deployment
	chmod +x scripts/logs.sh
	INSTANCE_NAME=$(INSTANCE_NAME) ZONE=$(GCP_ZONE) ./scripts/logs.sh

ssh-vm: ## SSH into GCP VM
	gcloud compute ssh $(INSTANCE_NAME) --zone=$(GCP_ZONE) --project=$(GCP_PROJECT_ID)

clean: ## Clean build artifacts
	rm -rf bin/
	rm -f coverage.out
	docker-compose down -v
	docker rmi $(DOCKER_IMAGE) 2>/dev/null || true

deps: ## Download dependencies
	go mod download
	go mod tidy

fmt: ## Format code
	go fmt ./...

lint: ## Run linter
	golangci-lint run ./...

# Local development
dev: ## Run in development mode with hot reload (requires air)
	air

install-tools: ## Install development tools
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

