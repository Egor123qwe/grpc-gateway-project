# Go parameters
GOCMD = go
GOTEST = $(GOCMD) test
GOCOVER = $(GOCMD) tool cover
SWAGGER = swagger
DOCKER_COMPOSE = docker-compose

# Directories
INTERNAL_DIR = ./internal/service
SWAGGER_FILE = ./cmd/docs/swagger.json

# Targets

# Run tests with code coverage and generate html-page
test:
	$(GOTEST) $(INTERNAL_DIR) -coverprofile=coverage.out /
	$(GOCOVER) -html=coverage.out

# Serve Swagger API documentation
swagger-serve:
	$(SWAGGER) serve $(SWAGGER_FILE)

# Build and run the application with Docker Compose
docker-compose-up:
	$(DOCKER_COMPOSE) up --build

grpc-generate:
	protoc --go_out=proto/api --go-grpc_out=proto/api proto/api/api.proto


# Default target
default: test
