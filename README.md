# Social App Backend

This is a backend system for a Twitter-like social media application, built with Go.

## tech Stack

- **Language:** [Go](https://go.dev/) (v1.24.0)
- **Web Framework:** [Gin](https://github.com/gin-gonic/gin) (v1.11.0)
- **ORM:** [GORM](https://gorm.io/) (v1.31.1)
- **Database:** (To be configured, currently using GORM with default/sqlite likely if not specified, but setup implies standard SQL structure)
- **Utilities:**
  - [Google UUID](https://github.com/google/uuid) for unique identifiers.

## Project Structure

```
backend/
├── cmd/
│   └── api/          # Application entry point
├── internal/
│   └── models/       # Database models (Base, User, Tweet, etc.)
├── go.mod            # Go module definition
└── go.sum            # Go module checksums
```

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.24 or higher.

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd social-app
   ```

2. Navigate to the backend directory:
   ```bash
   cd backend
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running the Application

To start the server, run:

```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:8080` (default Gin port).

## API Endpoints

### Health Check

- **GET /ping**
  - Returns `{"message": "pong"}` to verify the server is running.

## Development

- **Models:** defined in `internal/models`.
- **API Handlers:** (To be implemented/expanded)
