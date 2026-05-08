# Day 29: Dockerizing Go Applications

This directory demonstrates how to containerize a Go web application using **Docker**, utilizing multi-stage builds for small, secure images, and **Docker Compose** for orchestrating the app alongside a PostgreSQL database.

## 📝 Key Concepts Covered

- **Multi-Stage Dockerfile**:
  - **Stage 1 (Builder)**: Uses the full `golang` Alpine image to download dependencies and compile the Go binary with optimizations (`-ldflags="-w -s"`).
  - **Stage 2 (Final Image)**: Uses a minimal `alpine` image (only ~15MB) to run the compiled binary, dropping all source code and Go tools to reduce attack surface and image size.
- **Security Best Practices**:
  - Creating a non-root user (`appuser`) inside the container to execute the application, preventing root access if the container is compromised.
- **Docker Compose Orchestration**:
  - Defining `app` and `db` (PostgreSQL) services in a single `docker-compose.yml` file.
  - Managing internal networking so the Go app can reach the database using the hostname `db`.
  - Passing environment variables directly to the container.
- **Healthchecks & Dependencies**: Using `depends_on` with `condition: service_healthy` to ensure the Go application waits until PostgreSQL is fully initialized and ready to accept connections before starting.

## 📂 Files

- [Dockerfile](Dockerfile): The multi-stage build instructions for the Go app.
- [docker-compose.yml](docker-compose.yml): The orchestration file defining the app, database, and volumes.
- [main.go](main.go) / [config/](config/) / etc.: The actual Go application codebase (a standard Gin/GORM API).

## 🚀 How to Run

1.  Make sure Docker and Docker Compose are installed and running on your system.
2.  From this directory, run:
    ```bash
    docker-compose up --build
    ```
    *(Use `-d` to run in detached mode).*
3.  The API will be available at `http://localhost:8080` and is connected to the containerized PostgreSQL database.
4.  To stop and remove containers, networks, and volumes:
    ```bash
    docker-compose down -v
    ```
