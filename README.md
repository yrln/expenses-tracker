# Expense Tracker API

 Go + Fiber REST API for tracking personal expenses.

This project is a backend-only service where users can record expenses and query them by date, category, and other filters.
​

### Tech Stack
- Go (Golang)
- Fiber web framework
- MySQL
​
### Features
- CRUD operations for expenses (create, list, update, delete).
- Per-user expenses (each user has their own data).
- Filters by category.
- health check endpoint for uptime monitoring.
- JWT authentication.

### Project Structure


- `cmd/api/main.go`         # Application entrypoint, wiring routes and dependencies 
- `internal/config/`         # Load application configuration (env vars, ports, DB DSN) 
- `internal/database/`       # Database connection setup (MySQL driver)
- `internal/expenses/`       # Expense domain: handlers, models, repository (DB queries)
  - `handler.go`
  - `model.go`
  - `repository.go`
- `internal/auth/`           # Authentication logic (JWT, password hashing)
- `internal/playground/`     # Playground/scratch code, not part of core API 


## Getting Started
### Prerequisites
- Go (version 1.22)
- A running MySQL instance (local or AWS RDS).
- Git

### Environment variables

`APP_ADDR` : address for Fiber to listen to

example: `APP_ADDR=:3000`
  
`DB_DSN` : Database DSN to connect to mysql database 

example `DB_DSN="user:password@tcp(localhost:3306)/expense_tracker?charset=utf8mb4&parseTime=True&loc=Local"`
  
### Run locally

```
git clone https://github.com/yrln/expenses-tracker.git
cd expenses-tracker

go mod tidy
go run ./cmd/api
```
The API will start on http://localhost:3000 (as set in APP_ADDR).

### API Endpoints

#### Health
- GET `/health` – health check, returns service status.
​
#### Expenses
- GET `/expenses` – List expenses, optionally filtered by query params `category`
- GET `/expenses/:id` – Get single expense by ID.
- POST `/expenses` – Create a new expense. Payload example:
    ```
    {
        "amount": 24000,
        "category": "food",
        "note": "nasi goreng"
    }
    ```

#### Profile
- GET `/me` – get user account information.

#### Auth
- POST `/auth/register` – Register new user.
- POST `/auth/login` – Log in and receive a JWT.

Use header `Authorization: Bearer <token>` for access protected resources.

#### About
- GET `/about` – get backend about information.

#### Development notes
This is a learning project to practice Go backend development, REST API design, and working with relational databases.

The structure loosely follows common Go REST API layouts (separate config, database, and domain packages).

# To Be Added Soon
- Swagger API Documentation
- Redis integration
- Containerize App in Docker
- Add Test Function
- Improve CI with test
- Rate Limiting
