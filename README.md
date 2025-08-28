# GIFMA Backend

A Go Fiber web application with health check functionality.

## Features

- Health check endpoint at `/health`
- Root endpoint at `/`
- Built with Go Fiber framework
- Includes logging and recovery middleware

## Prerequisites

- Go 1.21 or higher

## Installation

1. Clone or navigate to this directory
2. Download dependencies:
   ```bash
   go mod tidy
   ```

## Running the Application

Start the server:
```bash
go run main.go
```

The server will start on port 3000.

## API Endpoints

### Health Check
- **GET** `/health`
- Returns service status and timestamp

### Root
- **GET** `/`
- Returns welcome message

## Example Usage

```bash
# Health check
curl http://localhost:3000/health

# Root endpoint
curl http://localhost:3000/
```

## Response Examples

Health Check Response:
```json
{
  "status": "ok",
  "timestamp": "2024-01-01T12:00:00Z",
  "service": "gifma-backend",
  "version": "1.0.0"
}
```

Root Response:
```json
{
  "message": "Welcome to GIFMA Backend API",
  "version": "1.0.0"
}
```
