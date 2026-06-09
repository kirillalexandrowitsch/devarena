# DevArena

DevArena is a production-oriented educational backend project written in Go.

The project is designed as a long-term portfolio system for learning and demonstrating practical backend engineering: domain modeling, API design, persistence, caching, asynchronous processing, distributed communication, observability, testing, deployment and production-readiness practices.

DevArena starts from a simple arena game domain and evolves into a modular backend platform with real backend infrastructure and production-style engineering decisions.

---

## Purpose

DevArena is built to serve two goals:

1. Learn Go and backend engineering through implementation, not isolated examples.
2. Become a serious backend portfolio project that can be explained in interviews and reviewed through commit history.

The project is not intended to be a quick demo application.  
Each technical concept should eventually be represented by real project code, tests, documentation or infrastructure.

---

## Product Vision

DevArena is a backend platform for a text-based arena game.

Users create heroes, fight enemies, receive rewards, manage inventory, participate in tournaments, progress through ratings and generate game events that are processed by background workers and analytics services.

The final system is expected to support:

```text
Hero management
Enemy management
Battle simulation
Inventory and rewards
Experience and progression
Ratings and leaderboards
Tournaments
Battle history
Game events
Background workers
Analytics
Audit logging
```

---

## Target Backend Capabilities

The final version of DevArena should demonstrate the following backend capabilities:

```text
REST API for public application operations
gRPC for internal service communication
PostgreSQL for relational persistence
Redis for caching, locks and cooldowns
MongoDB for battle logs and analytics-oriented documents
RabbitMQ for background task processing
Kafka for event streaming
Docker Compose for local infrastructure
Kubernetes for deployment
GitHub Actions for CI/CD
Structured logging
Metrics
Tracing
Health checks
Graceful shutdown
Security basics
Testing culture
Performance awareness
```

---

## Target Architecture

DevArena is intended to evolve into a modular backend system with separated application layers.

Target architectural direction:

```text
cmd/
  api/
  arena/
  reward-worker/
  notification-worker/
  analytics-worker/

internal/
  app/
  config/
  domain/
  handler/
  service/
  repository/
  middleware/
  validator/
  response/
  infrastructure/

proto/
gen/
migrations/
deploy/
tests/
```

The intended architecture follows these principles:

```text
Domain logic is isolated from transport and infrastructure.
Handlers depend on services.
Services coordinate domain logic and repositories.
Repositories hide persistence details.
Infrastructure integrations are kept behind interfaces.
Background workers process asynchronous tasks and events.
Application wiring is explicit and testable.
```

---

## Domain Model

The core domain includes:

```text
User
Hero
HeroClass
Enemy
Battle
BattleRound
Item
Inventory
Reward
Rating
Tournament
BattleLog
AuditLog
GameEvent
OutboxEvent
```

The domain starts from a battle simulation model and grows into a backend system around progression, rewards, events and persistence.

---

## Technology Stack

### Language

```text
Go
Go modules
Go standard library
```

### API and Communication

```text
HTTP
REST
JSON
gRPC
Middleware
Validation
JWT authentication
```

### Persistence and Storage

```text
PostgreSQL
Redis
MongoDB
Database migrations
Repository pattern
```

### Messaging and Event Processing

```text
RabbitMQ
Kafka
Background workers
Event streaming
Outbox pattern
```

### Infrastructure

```text
Docker
Docker Compose
Kubernetes
GitHub Actions
CI/CD pipelines
```

### Quality and Production Readiness

```text
Unit tests
Integration tests
Table-driven tests
Race detection
Code formatting
Linting
Structured logging
Metrics
Tracing
Health checks
Graceful shutdown
Security basics
Performance profiling
```

---

## Learning Philosophy

DevArena is an educational project, but the code should be treated as production-oriented.

Every topic should be learned through one or more of the following:

```text
A real domain feature
A refactoring that improves the design
A test that proves behavior
A benchmark or profiling task
A documented architectural decision
A production-style infrastructure addition
```

Learning progress is tracked separately in:

```text
LEARNING_CHECKLIST.md
```

Architectural decisions and design evolution are documented in:

```text
ARCHITECTURE.md
```

---

## Development

Run the application:

```bash
go run ./cmd/arena
```

Run all tests:

```bash
go test ./...
```

Format code:

```bash
gofmt -w .
```

List packages:

```bash
go list ./...
```

Clean and update module dependencies:

```bash
go mod tidy
```

---

## Documentation

### README.md

Describes the project purpose, final vision, target architecture and development commands.

### ARCHITECTURE.md

Documents architectural decisions, target structure, package responsibilities and system evolution.

### LEARNING_CHECKLIST.md

Tracks learning progress and confirms which topics have been implemented in the project.

---

## Commit Style

The project uses clear commit messages that describe the purpose of each change.

Examples:

```text
init: create DevArena project
docs: update project readme
learn: add hero stat lookup
refactor: split domain packages
feat: add REST API health endpoint
test: add battle service tests
ci: add GitHub Actions workflow
deploy: add Kubernetes manifests
```

---

## Final Goal

The final version of DevArena should be explainable as:

```text
DevArena is a production-oriented educational backend platform written in Go.

It uses a text-based arena game domain to demonstrate practical backend engineering:
REST API, PostgreSQL, Redis, MongoDB, RabbitMQ, Kafka, gRPC, Docker, Kubernetes,
CI/CD, testing, observability, security and production-readiness practices.
```

The expected result is a serious backend portfolio project with understandable architecture, meaningful commit history, tested business logic and infrastructure close to real production systems.