# Ethereum Transaction Parser

## Description
This is a simple Ethereum blockchain parser implemented in Go. The parser allows querying transactions for subscribed Ethereum addresses. It provides an HTTP API to interact with the parser and retrieve information such as the latest block, subscribe to addresses, and get transactions for a given address.

## Features
- **Get Current Block**: Retrieve the latest parsed block.
- **Subscribe to an Address**: Subscribe to notifications for incoming and outgoing transactions for a specific Ethereum address.
- **Get Transactions**: Retrieve all inbound and outbound transactions for a subscribed address.
- **Swagger Documentation**: View and interact with the API documentation via Swagger UI.

## Project Structure
- `/internal/parser`: Contains the core parsing logic and in-memory storage.
- `/internal/rpcclient`: Handles JSON-RPC communication with the Ethereum blockchain.
- `/api`: Exposes the HTTP API to interact with the parser.
- `/docs`: Contains the generated Swagger documentation.

## Prerequisites
- Go (latest version recommended)

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/ethereum-tx-parser.git
   cd ethereum-tx-parser
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   go mod vendor
   ```

3. Install tools:
   ```bash
   make tools
   ```

## Usage

### 1. Build the Project
   ```bash
   make build
   ```
   - **Description**: Builds the project, generating an executable named `tx-parser`.

### 2. Run All Tests
   ```bash
   make test
   ```
   - **Description**: Runs all tests in the `internal` directory to ensure everything is working correctly.

### 3. Start the Application
   ```bash
   make run
   ```
   - **Description**: Executes the built `tx-parser` executable, starting the HTTP server.

### 4. Generate Swagger Documentation
   ```bash
   make swagger
   ```
   - **Description**: Generates the Swagger documentation using `swag init`. Make sure you've added all the necessary comments in your code for Swagger.

### 5. Install Tools
   ```bash
   make tools
   ```
   - **Description**: Installs the necessary tools, such as the `swag` CLI, for generating documentation.

## API Endpoints
- **GET** `/block/current`: Get the latest parsed block.
- **POST** `/subscribe/{address}`: Subscribe to an Ethereum address using the `address` path parameter.
- **GET** `/transactions/{address}`: Get transactions for an Ethereum address using the `address` path parameter.

### Example Requests
1. **Get Current Block**
   ```bash
   curl -X GET "http://localhost:8080/block/current"
   ```

2. **Subscribe to an Address**
   ```bash
   curl -X POST "http://localhost:8080/subscribe/0xYourEthereumAddress"
   ```

3. **Get Transactions for an Address**
   ```bash
   curl -X GET "http://localhost:8080/transactions/0xYourEthereumAddress"
   ```

## Swagger Documentation
- Once the server is running, access the Swagger UI at:
  ```
  http://localhost:8080/swagger/index.html
  ```
