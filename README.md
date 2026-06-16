# Go Fiber CRUD API

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/postgresql-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

A robust, production-ready RESTful CRUD API built with Go and the Fiber framework. This project manages user Date of Birth (DOB) records and demonstrates best practices in Go backend development, including database interactions with SQLC, structured project layout, containerization, and data validation.

## ✨ Features

- **Blazing Fast Router**: Built on top of [Fiber](https://gofiber.io/), an Express-inspired web framework written in Go.
- **Type-Safe Database Queries**: Uses [SQLC](https://sqlc.dev/) to generate type-safe Go code from SQL, ensuring query safety at compile-time.
- **Robust Database**: Uses **PostgreSQL** for reliable data persistence.
- **Containerization**: Fully Dockerized for seamless local development and easy deployment.
- **Data Validation**: Comprehensive input validation to ensure data integrity and security.
- **Middleware Integration**: Includes essential middlewares for logging, error handling, and request lifecycle management.
- **Pagination**: Efficiently handles large datasets with built-in pagination support for API endpoints.

## 🚀 Setup Instructions

Follow these steps to get the project up and running on your local machine.

### 1. Clone the repository
```bash
git clone <your-repo-url>
cd go_postgresSQL
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Environment Variables
Create a `.env` file in the root directory based on a `.env.example` (or use the following template):

```env
PORT=8080
DATABASE_URL=postgres://postgres:password@localhost:5432/postgres?sslmode=disable
APP_ENV=development
```

### 4. Database Migrations
Make sure your PostgreSQL instance is running. Then, run the database migrations:

```bash
make migrateup
```

### 5. Run the Project
Start the Go Fiber application:

```bash
make run
```

## 🐳 Docker Usage

To spin up the entire application stack (API and Database) using Docker Compose, simply run:

```bash
make dockerup
```
This command builds the necessary images and starts the containers as defined in the `docker-compose.yml`.

## 🔌 API Examples

Here are some examples of the available endpoints for user management.

### Create a User

**Endpoint:** `POST /users`

**Request Body:**
```json
{
  "name": "John Doe",
  "dob": "1990-01-15T00:00:00Z"
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "ID": 1,
    "Name": "John Doe",
    "Dob": "1990-01-15",
    "CreatedAt": "2023-10-27T10:00:00Z",
    "UpdatedAt": "2023-10-27T10:00:00Z"
  }
}
```

### Get Users (with Pagination)

**Endpoint:** `GET /users?page=1&limit=10`

**Response (200 OK):**
```json
{
  "success": true,
  "data": [
    {
      "ID": 1,
      "Name": "John Doe",
      "Dob": "1990-01-15",
      "CreatedAt": "2023-10-27T10:00:00Z",
      "UpdatedAt": "2023-10-27T10:00:00Z"
    }
  ]
}
```

## 📈 Why README Matters

A high-quality `README.md` is the face of your project. Recruiters and engineers heavily judge projects by their README quality. It serves as the primary entry point for anyone discovering the repository, providing crucial context, setup instructions, and demonstrating the developer's ability to document and communicate effectively.
