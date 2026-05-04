# Go-Fiber & React-TS Full-Stack Boilerplate

This repository is a high-performance, containerized starter kit combining the speed of Go (Fiber) with the type-safety of React (TypeScript). It follows a modular Layered MVC architecture to ensure the project remains maintainable as it grows.

## 📂 What's Inside

### 1. Server (`/server`)

A modular Go backend built for high-concurrency and speed.

- `cmd/api/`: The application entry point (`main.go`).
- `internal/handlers/`: Controllers that parse requests and return responses.
- `internal/services/`: The business logic layer.
- `internal/repository/`: Handles all database interactions (SQL/NoSQL).
- `internal/models/`: Data structures and schemas.

### 2. Client (`/frontend`)

A modern frontend optimized for developer experience and performance.

- **Vite**: Next-generation frontend tooling for instant HMR.
- **React + TypeScript**: Component-based UI with strict type definitions.
- **Tailwind CSS**: Utility-first styling.
- **Axios**: Pre-configured for API communication with the Go backend.

### 3. Infrastructure (`/docker`)

Complete containerized environment managed via `docker-compose.yaml`.

- **PostgreSQL**: Primary relational database.
- **MongoDB**: Secondary NoSQL database.
- **Management UIs**: Includes **pgAdmin** and **Mongo Express** for visual data management.

## ❓ Why This Stack?

- **Why Go (Fiber)?**: Fiber is inspired by Express.js but leverages Go's zero-memory allocation and high-concurrency capabilities, making it one of the fastest frameworks for microservices and APIs.
- **Why TypeScript?**: Using TypeScript on the frontend allows you to mirror your Go structs as TS interfaces. This ensures that the data your client sends matches exactly what your server expects, catching errors during development instead of at runtime.
- **Why Docker?**: It isolates the databases and the server environment, ensuring that every developer on the team is running the exact same versions of PostgreSQL and MongoDB without manual local installation.

## 🛠️ Quick Start

### Prerequisites

- **Node.js** & **npm**
- **Go**
- **Docker Desktop**

### Installation

1.  **Clone the repository**:

    ```bash
    git clone https://github.com/your-repo/go-fiber-react-ts-boilerplate.git
    cd go-fiber-react-ts-boilerplate
    ```

2.  **Install Frontend Dependencies**:

    ```bash
    cd frontend
    npm install
    cd ..
    ```

3.  **Install Backend Dependencies**:

    ```bash
    cd server
    go mod tidy
    cd ..
    ```

## 🚀 How to Run

### Method 1: Full Containerization (Recommended)

From the root directory, run the following to build the Go binary and spin up all services:

```powershell
docker-compose up --build
```

- **API**: `http://localhost:8080`
- **pgAdmin**: `http://localhost:5050`
- **Mongo Express**: `http://localhost:8081`

### Method 2: Local Development

If you want to run the server or frontend locally while keeping the databases in Docker:

#### Start Databases:

```powershell
docker-compose up db mongodb -d
```

#### Run Backend:

```powershell
cd server
go run cmd/api/main.go
```

#### Run Frontend:

```powershell
cd frontend
npm run dev
```

---

### 🔐 Default Credentials

| Service | Username | Password | Port |
| :--- | :--- | :--- | :--- |
| **PostgreSQL** | `boilerplate` | `pass_boilerplate` | `5432` |
| **MongoDB** | `boilerplate` | `pass_boilerplate` | `27017` |
| **pgAdmin** | `boilerplate@dev.com` | `pass_boilerplate` | `5050` |
