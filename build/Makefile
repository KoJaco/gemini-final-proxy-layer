# Variables
APP_NAME := companion
DOCKER_IMAGE := kojaco/$(APP_NAME)
DOCKER_TAG := latest

# Build the Docker image
build:
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) -f Dockerfile .

# Run the Docker container
run:
	docker run -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)

# Push the Docker image to a repository
push:
	docker push $(DOCKER_IMAGE):$(DOCKER_TAG)

# Stop all running containers
stop:
	docker-compose -f deployments/docker-compose.yaml down

# Start services with Docker Compose
start:
	docker-compose -f deployments/docker-compose.yaml up -d

# Rebuild and start
rebuild: stop build start

# Clean up
clean:
	docker system prune -f