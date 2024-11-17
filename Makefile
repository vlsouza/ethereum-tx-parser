# Name of the generated executable
BINARY_NAME=tx-parser

# Directory containing the internal project files
INTERNAL_DIR=./internal

# Command to generate Swagger documentation (replace with your specific Swagger command)
SWAGGER_CMD=swag init

# Rule to build the project
build:
	go build -o $(BINARY_NAME) main.go

# Rule to run all tests in the internal directory
test:
	go test $(INTERNAL_DIR)/...

# Rule to run the application
run:
	./$(BINARY_NAME)

# Rule to generate Swagger documentation
swagger:
	$(SWAGGER_CMD)

# Rule to install required tools
tools:
	go install github.com/swaggo/swag/cmd/swag@latest

# Rule to run all steps in sequence (optional)
all: build test run
