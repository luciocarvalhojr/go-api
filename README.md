# Go API - DevSecOps Project

A production-ready Go API following the [Standard Go Project Layout](https://github.com/golang-standards/project-layout) with a focus on DevSecOps and security automation.

## 🚀 Features

- **Standard Layout**: Separates entry point (`cmd/`) from business logic (`internal/`).
- **DevSecOps Pipeline**: Automated Linting, SAST, SCA, and Container Scanning via GitHub Actions.
- **Secure Docker Image**: Multi-stage build using a **Distroless** base image to minimize attack surface.
- **RESTful API**: Built with the [Gin Gonic](https://gin-gonic.com/) framework.

---

## 🛠️ Prerequisites

- **Go**: v1.26 or higher
- **Docker**: For containerization
- **(Optional) Local Security Tools**:
  - `golangci-lint` (Linting)
  - `gosec` (Security Scan)
  - `govulncheck` (Vulnerability Scan)

---

## 💻 Local Development

### Run the API
```bash
go run cmd/api/main.go
```
The server will start at `http://localhost:8080`.

### Run Tests
Run all unit tests with race detection and coverage:
```bash
go test -v -race -cover ./...
```

### Manage Dependencies
```bash
go mod tidy
```

---

## 🔒 Security & Quality Commands

Use these commands locally to catch issues before pushing to GitHub.

### 1. Linting
Requires [golangci-lint](https://golangci-lint.run/usage/install/).
```bash
golangci-lint run
```

### 2. SAST (Static Application Security Testing)
Scans code for security flaws like hardcoded secrets or SQL injection.
```bash
go install github.com/securego/gosec/v2/cmd/gosec@latest
gosec ./...
```

### 3. SCA (Software Composition Analysis)
Checks for known vulnerabilities in your dependencies.
```bash
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...
```

---

## 🐳 Docker

### Build the Image
```bash
docker build -t go-api .
```

### Run the Container
```bash
docker run -p 8080:8080 go-api
```

---

## 🤖 CI/CD (DevSecOps Pipeline)

The project includes a GitHub Actions workflow in `.github/workflows/devsecops.yml` that automatically runs on every push and pull request to `main`:

1.  **Test & Lint**: Runs `golangci-lint` and `go test`.
2.  **Security Scan**: Runs `gosec` (SAST) and `govulncheck` (SCA).
3.  **Docker Scan**: Builds the image and scans it with **Trivy** for vulnerabilities.

---

## 🛣️ API Endpoints

| Method | Endpoint      | Description           |
| :----- | :------------ | :-------------------- |
| GET    | `/albums`     | Get all albums        |
| GET    | `/albums/:id` | Get album by ID       |
| POST   | `/albums`     | Add a new album       |
