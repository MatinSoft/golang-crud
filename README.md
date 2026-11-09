# Customer CRUD Application with Go, Gin, DDD, Clean Architecture & CQRS

## Project Structure
- `domain/`: Domain models, Value Objects (`PhoneNumber` validated with Google libphonenumber, `Email`, `BankAccountNumber`), and Domain Events.
- `application/`: CQRS command & query handlers, CustomerService business logic.
- `ports/`: Interfaces for Repositories and Event Store.
- `infrastructure/`: In-Memory / GORM Repositories and Event Store implementation.
- `presentation/`: Gin REST API handlers, routes, and Swagger UI integration.

## API Endpoints
- `POST /api/v1/customers` - Create customer with validations
- `GET /api/v1/customers` - List all customers
- `GET /api/v1/customers/:id` - Get customer by ID
- `DELETE /api/v1/customers/:id` - Delete customer
- `GET /swagger/*any` - Swagger UI documentation

## Running Tests
```bash
go test -v ./...
```

## Running with Docker Compose
```bash
docker-compose up --build
```
