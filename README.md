# DevArena

DevArena is an educational backend project for learning Go and modern backend development from scratch.

The project starts as a simple console battle simulator and gradually evolves into a production-like backend platform with REST API, PostgreSQL, Redis, MongoDB, gRPC, RabbitMQ, Kafka, Docker, Kubernetes, CI/CD, observability, security, testing and production-readiness practices.

The goal is not to build fast.  
The goal is to learn every important backend topic through implementation, commit history and documentation.

---

## Project Goal

DevArena is built as a long-term learning project for becoming a strong Go backend developer.

The project is designed to grow through levels:

```text
Beginner
→ Confident Intern
→ Junior Backend Developer
→ Strong Junior / Junior+
→ Middle-oriented Backend Developer
```

Every technology and concept is introduced through a real project feature.

The learning cycle for every topic:

```text
1. Learn the concept
2. Apply it in DevArena
3. Understand why this solution is used
4. Commit the change
5. Mark the topic in LEARNING_CHECKLIST.md
6. Document important decisions in ARCHITECTURE.md
```

---

## Project Concept

DevArena is a backend platform for a text-based arena game.

Users create heroes, fight enemies, earn experience, receive items, join tournaments, trade equipment, get ranked on leaderboards and generate game events processed by background workers and analytics services.

The project starts with a console application:

```text
Hero appears in the arena.
Hero has HP, damage, level and class.
Enemy appears.
Battle runs round by round.
Winner gets experience.
```

Later it evolves into a backend system:

```text
REST API
PostgreSQL persistence
Redis cache and cooldowns
MongoDB battle logs
RabbitMQ workers
Kafka event streaming
gRPC internal services
Docker Compose local environment
Kubernetes deployment
GitHub Actions CI/CD
Observability and production-readiness
```

---

## Current Stage

```text
Stage 1: Go Core
```

Current focus:

```text
Go packages and modules
```

---

## Planned Tech Stack

### Language and Runtime

```text
Go
Go modules
Go standard library
```

### Backend

```text
REST API
gRPC
HTTP
JSON
JWT authentication
Middleware
Validation
```

### Databases and Storage

```text
PostgreSQL
Redis
MongoDB
```

### Messaging and Event-Driven Architecture

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
CI/CD
```

### Quality and Production Readiness

```text
Testing
Linting
Structured logging
Metrics
Tracing
Health checks
Graceful shutdown
Security basics
Performance basics
```

---

## Learning Areas

DevArena covers the following areas:

```text
Go Core
Go Error Handling
Go Concurrency
Testing
HTTP and REST API
API Design
gRPC
PostgreSQL and SQL
Redis
MongoDB
RabbitMQ
Kafka
Docker
Kubernetes
CI/CD
Git and GitHub
Linux and Terminal
Networking Basics
Security
Observability
Backend Architecture
Data Modeling
Background Jobs
Performance Basics
Production Readiness
```

---

## Development Principle

This project is not developed as a quick demo.

It is developed as a learning map.

Every feature must answer:

```text
What concept am I learning?
Where is it used in DevArena?
Why is it implemented this way?
What problem does it solve?
How can I explain this on an interview?
```

---

## Main Domain Entities

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

---

## Project Evolution

```text
Stage 0: Repository Initialization
Stage 1: Console Arena and Go Basics
Stage 2: Go Error Handling
Stage 3: Go Concurrency
Stage 4: Domain Packages and Modules
Stage 5: In-memory REST API
Stage 6: Backend Architecture
Stage 7: PostgreSQL Persistence
Stage 8: Docker Compose Environment
Stage 9: Authentication and Authorization
Stage 10: Redis Mechanics
Stage 11: Testing
Stage 12: RabbitMQ Workers
Stage 13: Kafka Event Streaming
Stage 14: MongoDB Logs
Stage 15: gRPC Internal Services
Stage 16: Observability
Stage 17: Kubernetes Deployment
Stage 18: CI/CD
Stage 19: Security Hardening
Stage 20: Performance and Production Readiness
```

---

## Current Project Structure

```text
devarena/
  cmd/
    arena/
      main.go

  internal/
    battle/
      battle.go

    enemy/
      enemy.go

    hero/
      hero.go
      weapon.go

  README.md
  ARCHITECTURE.md
  LEARNING_CHECKLIST.md
  .gitignore
  go.mod
```

---

## Future Project Structure

```text
devarena/
  cmd/
    arena/
      main.go
    api/
      main.go
    notification-worker/
      main.go
    reward-worker/
      main.go
    analytics-worker/
      main.go

  internal/
    app/
    config/
    handler/
    service/
    repository/
    model/
    middleware/
    response/
    validator/
    infrastructure/

  proto/
  gen/
  migrations/
  tests/
  deploy/
  .github/

  Dockerfile
  docker-compose.yml
  Makefile
  .env.example
  .dockerignore
  .gitignore
  README.md
  ARCHITECTURE.md
  LEARNING_CHECKLIST.md
  go.mod
  go.sum
```

---

## Go Module

This project uses Go modules.

Current module path:

```text
github.com/rudyakovk/devarena
```

Internal packages are imported using this module path.

Examples:

```go
import "github.com/rudyakovk/devarena/internal/hero"
import "github.com/rudyakovk/devarena/internal/enemy"
import "github.com/rudyakovk/devarena/internal/battle"
```

The `internal` directory is used to keep application packages private to this module.

This means external projects cannot import packages such as:

```text
github.com/rudyakovk/devarena/internal/hero
```

Only packages inside the same module can use these internal packages.

Useful Go module commands:

```bash
go mod tidy
go list ./...
go test ./...
```

---

## Current Internal Packages

### `internal/hero`

Contains hero-related domain logic:

```text
Hero
Weapon
Sword
Axe
Hero methods
Weapon interface
```

### `internal/enemy`

Contains enemy-related domain logic:

```text
Enemy
Enemy.TakeDamage
```

### `internal/battle`

Contains battle-related domain logic:

```text
Battle
Battle.Start
```

---

## Run

Current console version:

```bash
go run ./cmd/arena
```

---

## Test

Run all tests:

```bash
go test ./...
```

Later the project will include:

```bash
go test -race ./...
go test -cover ./...
```

---

## Useful Development Commands

Format code:

```bash
gofmt -w .
```

List all packages:

```bash
go list ./...
```

Clean and update module dependencies:

```bash
go mod tidy
```

Run the console arena:

```bash
go run ./cmd/arena
```

Run all tests:

```bash
go test ./...
```

---

## Git Commit Style

The project uses meaningful commit messages.

Examples:

```text
init: create DevArena project
docs: add learning roadmap
learn: add hero variables
learn: add basic hero data types
learn: add damage calculation function
learn: add battle condition
learn: add simple battle loop
learn: add hero constants
learn: add fixed attack array
learn: add hero inventory slice
learn: copy hero inventory slice
learn: add hero stats map
learn: iterate hero collections with range
learn: add Hero Enemy and Battle structs
learn: add Hero and Battle methods
learn: use pointers to update hero and battle state
learn: add Weapon interface
refactor: split arena domain into packages
learn: document Go module usage
feat: add REST API health endpoint
feat: add PostgreSQL hero repository
test: add battle service table-driven tests
ci: add GitHub Actions workflow
deploy: add Kubernetes manifests
```

---

## Documentation Files

### README.md

Explains what the project is, how to run it and what it is designed to teach.

### ARCHITECTURE.md

Describes the current and future architecture of the project.

### LEARNING_CHECKLIST.md

Tracks all topics that must be learned, implemented, committed and documented.

---

## Final Goal

At the end of development, DevArena should look like a serious backend portfolio project.

It should demonstrate:

```text
Strong Go fundamentals
Practical backend architecture
REST API design
PostgreSQL persistence
Redis caching and locking
MongoDB document modeling
RabbitMQ background workers
Kafka event streaming
gRPC service communication
Dockerized local environment
Kubernetes deployment
CI/CD automation
Testing culture
Security basics
Observability
Production-readiness thinking
```

The final project should be explainable in interviews as:

```text
DevArena is a production-like educational backend platform written in Go.
It started as a console battle simulator and evolved into a modular backend system with REST API, PostgreSQL, Redis, MongoDB, RabbitMQ, Kafka, gRPC, Docker, Kubernetes, CI/CD, testing and observability.
```