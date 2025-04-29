# Currency Exchange Microservices

This project is a simple system of two Go-based microservices:

- `currency`: Fetches daily exchange rates and stores them in a PostgreSQL database.
- `gateway`: Provides a REST API for users, handling authentication and forwarding requests to `currency`.

## 🧱 Architecture

- **gateway**: REST API server that:
    - Handles user login via an external auth service
    - Validates JWT tokens
    - Forwards authorized requests to the currency service
- **currency**:
    - Worker fetches exchange rates once per day from a public API
    - Stores rates in a database
    - Serves gRPC or REST requests for current or historical exchange rates

## 📦 Technologies

- Golang
- PostgreSQL
- gRPC & REST
- Docker & Docker Compose
- Prometheus + Grafana (monitoring)
- Viper (config management)
- JWT (authorization)

