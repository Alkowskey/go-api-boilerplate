# Go User API

## Project Structure

```
.
├── config/         # Configuration files
├── internal/       # Internal application code
│   └── domain/     # Domain-specific code
│       └── user/   # User domain implementation
├── main.go         # Application entry point
├── go.mod          # Go module definition
├── go.sum          # Go dependencies checksum
├── .env            # Environment variables
└── docker-compose.yml # Docker configuration
```

## Prerequisites
- Docker and Docker Compose (optional, for containerized deployment)

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/Alkowskey/go-api-boilerplate.git
   cd go-user-api
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Configure environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

## Docker Deployment

To run the application using Docker:

```bash
docker-compose up -d
```

## API Endpoints

### User Management

- `POST /api/users/register` - Register a new user
  - Required fields: name, email, password
  - Password must be at least 6 characters long

## Development

### Running Tests

```bash
go test ./...
```

### Code Style

This project follows the standard Go formatting guidelines. To format your code:

```bash
go fmt ./...
```

## Acknowledgments

- Clean Architecture principles
- Domain-Driven Design patterns 