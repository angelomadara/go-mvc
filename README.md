# Go MVC Boilerplate

A simple MVC (Model-View-Controller) boilerplate written in Go using the Gorilla Mux router with MySQL support.

## Project Structure

```
├── main.go              # Application entry point
├── go.mod              # Go module file
├── .env.example        # Environment variables example
├── config/             # Database configuration
│   └── database.go
├── models/             # Data models and repository interfaces
│   └── user.go
├── views/              # View layer for rendering responses
│   └── user_view.go
├── controllers/        # HTTP request handlers
│   └── user_controller.go
└── routes/             # Route definitions
    └── routes.go
```

## Getting Started

1. Install dependencies:
```bash
go mod tidy
```

2. Setup MySQL database:
   - Create a database named `mvc_app` (or use your preferred name)
   - Copy `.env.example` to `.env` and update with your MySQL credentials

3. Run the application:
```bash
go run main.go
```

The server will start on port 8080 and connect to your MySQL database. The users table will be created automatically if it doesn't exist.

## API Endpoints

- `GET /users` - Get all users
- `GET /users/{id}` - Get user by ID
- `POST /users` - Create a new user
- `PUT /users/{id}` - Update user by ID
- `DELETE /users/{id}` - Delete user by ID

## Example Usage

Create a user:
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'
```

Get all users:
```bash
curl http://localhost:8080/users
```