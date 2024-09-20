# Buffet POS Backend

## Prerequisites

- **[GoFiber](https://gofiber.io/)** - An Express-inspired web framework for building fast and scalable applications in Go.
- **[Gorm](https://gorm.io/)** - An ORM library for Go, making it easy to work with databases through object relational mapping.
- **[MariaDB](https://mariadb.org/)** - A popular open-source database used to store and manage application data.
- **[Docker](https://www.docker.com/)** - A platform for developing, shipping, and running applications in containers.

## Setup Instructions

##### Clone the repository

```bash
git clone https://github.com/cs471-buffetpos/buffet-pos-backend
```

##### Navigate to the project directory

```bash
cd buffet-pos-backend
```

##### Install dependencies

```bash
go mod tidy
```

##### Setup environment variables using .env

```bash
cp .env.example .env
```

##### Run database using Docker Compose

```bash
docker compose up -d
```

## Running the server

```bash
go run main.go
```

## Database Migration

```bash
make db-migrate
```
