# Scaffoldy 🚀

> [!WARNING]
> **Project Status: Early Development (WIP)**
>
> This project is currently in an early development phase and is **not yet finished**. Please note:
> - **Future Direction**: Scaffoldy is planned to transition into a standalone **library**.
> - **Upcoming Changes**: The current authentication (Auth) system and migration schemes will be removed or refactored soon as they are currently not optimal.
> - **Scaffolding Capability**: The scaffolding process is currently stable and works optimally for **small-scale projects**. 
> - **Fixed Structure**: The project structure is not yet dynamic; you must follow the existing folder and file structure convention for the scaffolding to work correctly.

Scaffoldy is a powerful Go-based scaffolding tool and API boilerplate designed to accelerate backend development. It automates the generation of boilerplate code based on your domain entities, allowing you to focus on business logic rather than repetitive tasks.

## ✨ Features

- **Dynamic Code Generation**: Automatically generate Handler, Repository, Request, and Service layers based on a Go struct.
- **Auto-Injection**: The scaffolding tool automatically registers new domains into the main API router.
- **RESTful API Boilerplate**: Built on the high-performance [Gin](https://github.com/gin-gonic/gin) web framework.
- **Authentication**: Pre-configured JWT (JSON Web Token) authentication middleware.
- **Database Abstraction**: Easy-to-use repository pattern with MariaDB/MySQL support.
- **Database Migrations**: Integrated migration tool using `golang-migrate`.
- **Environment Driven**: Configuration managed via `.env` files.
- **Health Checks**: Built-in `/health` endpoint for monitoring.

## 🛠️ Tech Stack

- **Language**: [Go](https://golang.org/) (1.25.3+)
- **Web Framework**: [Gin Gonic](https://github.com/gin-gonic/gin)
- **Database**: MariaDB / MySQL
- **Authentication**: JWT
- **Migration**: [golang-migrate](https://github.com/golang-migrate/migrate)
- **Utilities**: `godotenv`, `google/uuid`, `crypto`

## 🚀 Getting Started

### 1. Prerequisites
- Go installed on your machine.
- MariaDB or MySQL server running.

### 2. Installation
Clone the repository:
```bash
git clone https://github.com/maulanaaghnii/scaffoldy.git
cd scaffoldy
```

### 3. Configuration
Create a `.env` file in the root directory (you can copy from the provided example below):
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=scaffoldy_db

SERVER_PORT=8080
ENVIRONMENT=development
ALLOWED_ORIGINS=*
```

### 4. Database Migrations
Run the migrations to set up your database schema:
```bash
go run cmd/migrate/main.go -up
```

### 5. Running the API
Start the development server:
```bash
go run cmd/api/main.go
```
The server will be available at `http://localhost:8080`.

## 🏗️ How to use the Scaffolding Tool

The primary power of Scaffoldy lies in its ability to generate code for you.

### Step 1: Define your Entity
Create a new directory in `internal/` (e.g., `internal/product`) and define your struct in `entity.go`:

```go
package product

type Product struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Price       int    `json:"price" db:"price"`
	Description string `json:"description" db:"description"`
}
```

### Step 2: Run Scaffoldy
Execute the scaffolding command:
```bash
go run cmd/scaffoldy/main.go --domain-name Product --table-name products
```

**What it does:**
1. Parses your `Product` struct.
2. Creates `handler.go`, `repository.go`, `request.go`, and `service.go` in `internal/product/`.
3. Injects the new domain registration into `cmd/api/main.go`.

## 📁 Project Structure

```text
├── cmd/
│   ├── api/            # Main API entry point
│   ├── migrate/        # Migration utility
│   └── scaffoldy/      # The Scaffolding tool CLI
├── internal/           # Domain-driven logic (Auth, User, etc.)
├── migrations/         # SQL migration files
├── pkg/                # Shared packages (config, middleware, utils)
├── scaffold_components/# Code generation templates
└── shared/             # Shared entities or types
```

## 📜 License
This project is licensed under the MIT License.
