# GoFiber Production Server

This repository contains a production-grade GoFiber server setup with a clean and scalable project structure.

## Project Structure

```
gofiber-server/
├── cmd/
│   └── server/
│       └── main.go
├── configs/
│   └── config.go
├── internal/
│   ├── api/
│   │   ├── middleware/
│   │   │   └── auth.go
│   │   ├── routes/
│   │   │   └── register.go
│   │   └── handlers/
│   │       └── user_handler.go
│   ├── service/
│   │   └── user_service.go
│   ├── repository/
│   │   └── user_repo.go
│   ├── model/
│   │   └── user.go
│   ├── database/
│   │   └── postgres.go
│   └── utils/
│       └── hash.go
├── pkg/
│   └── logger/
│       └── logger.go
├── test/
│   └── user_test.go
├── .env
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started) (optional)

### Installation

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/Hemanth5603/resume-go-server.git
    cd resume-go-server
    ```

2.  **Install dependencies:**
    ```sh
    go mod tidy
    ```

3.  **Run the server:**
    ```sh
    go run cmd/server/main.go
    ```

The server will start on `http://localhost:3000`.