# Train Ticket Service with gRPC in Golang

This project demonstrates a train ticket booking system using gRPC in Go. It includes a gRPC server, a client, and a set of unit tests. The server and client communicate through gRPC, and Docker is used to containerize the application and run tests.

## Project Structure

- `cmd/server/` - Contains the main server code.
- `cmd/client/` - Contains the main client code.
- `protos/` - Contains the Protocol Buffers definition files and generated code.
- `server_test.go` - Unit tests for the server.
- `client_test.go` - Unit tests for the client.
- `Dockerfile.client` - Dockerfile to build the client application and run tests.
- `Dockerfile.server` - Dockerfile to build the server application and run tests.
- `Dockerfile.proto` - Dockerfile to build and compile proto and grpc.
- `docker-compose.yml` - compose file to run the docker images
- `Makefile` - makefile to build and run the containers locally


### Prerequisites

- Go 1.20 or higher
- Docker

## API Workflow

### 1. **PurchaseTicket**
- **Purpose:** Book a ticket.
- **Request:** Travel details and user info.
- **Response:** Receipt with seat allocation.

### 2. **GetReceipt**
- **Purpose:** Retrieve a user's ticket receipt.
- **Request:** User's email.
- **Response:** Receipt details.

### 3. **GetUsersInSection**
- **Purpose:** List users and seats in a specific section.
- **Request:** Train section.
- **Response:** List of users and their seats.

### 4. **RemoveUser**
- **Purpose:** Delete a user from the system.
- **Request:** User's receipt ID.
- **Response:** Success status.

### 5. **ModifySeat**
- **Purpose:** Update a user's seat.
- **Request:** Receipt ID and new seat number.
- **Response:** Success status.

## Testing
Basic test skeleton is added with GO inbuild test 
this should be improved by using third party framwork like goMock

Test report is already added in the same directory, or it can be run using
#### Run Tests Locally
go test -coverprofile=coverage.out
go tool cover -func=coverage.out

## TODO
### Auth : 
authentication and authorizaton is not implented.
### Persistance 
Persistance layer is just in memory for now, can be levarage to integrate with any datastore

### Middleware
Currently everything as a flat file in go , can be more modular with multiple packages.  



 