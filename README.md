# DevArena

DevArena is a production-oriented educational backend project written in Go.

The project is designed as a long-term portfolio system for learning and demonstrating practical backend engineering: domain modeling, API design, persistence, caching, asynchronous processing, distributed communication, analytics, observability, testing, deployment and production-readiness practices.

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
ClickHouse for analytics and OLAP queries
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

DevArena is intended to evolve into a domain-first modular backend system.

The target architectural direction is:

```text
Domain-first design
Modular monolith first
Explicit application use cases
Ports and adapters
Infrastructure behind interfaces
Event-driven integrations
Production-oriented testing and observability
Analytics-oriented data processing
```

The project should grow around business capabilities and clear architectural boundaries:

```text
cmd/
  arena/
  api/
  reward-worker/
  notification-worker/
  analytics-worker/
  outbox-publisher/

internal/
  domain/
    user/
    auth/
    hero/
    enemy/
    battle/
    reward/
    inventory/
    rating/
    tournament/
    event/
    audit/

  app/
    user/
    auth/
    hero/
    battle/
    reward/
    inventory/
    rating/
    tournament/
    event/
    analytics/

  ports/
    repository/
    publisher/
    cache/
    lock/
    transaction/
    analytics/
    clock/
    idgen/
    password/
    token/

  adapters/
    http/
    grpc/
    postgres/
    redis/
    mongodb/
    clickhouse/
    rabbitmq/
    kafka/

  platform/
    config/
    logger/
    observability/
    shutdown/
    validation/
```

The intended architecture follows these principles:

```text
Domain logic is isolated from transport and infrastructure.
Application use cases coordinate domain logic through ports.
Ports define what the application needs from external systems.
Adapters implement ports using concrete technologies.
Infrastructure integrations are kept outside the domain layer.
Background workers reuse the same application layer as the API.
Application wiring is explicit and testable.
Analytics storage is separated from transactional business state.
```

Detailed architectural rules are documented in:

```text
ARCHITECTURE.md
```

---

## Domain Model

The core domain includes:

```text
User
Auth
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
AnalyticsEvent
```

The domain starts from a battle simulation model and grows into a backend system around progression, rewards, events, persistence and analytics.

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
Protocol Buffers
Middleware
Validation
JWT authentication
```

### Persistence and Storage

```text
PostgreSQL
Redis
MongoDB
ClickHouse
Database migrations
Repository pattern
Analytics storage
OLAP queries
```

### Messaging and Event Processing

```text
RabbitMQ
Kafka
Background workers
Event streaming
Outbox pattern
Analytics ingestion
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

## Planned Infrastructure Roles

### PostgreSQL

PostgreSQL is the main source of truth for transactional relational data.

Planned usage:

```text
users
refresh tokens
heroes
enemies
battles
battle rounds
items
hero inventory
ratings
tournaments
outbox events
```

PostgreSQL is responsible for consistency, transactions, constraints and relational queries.

---

### Redis

Redis is used for fast ephemeral and derived data.

Planned usage:

```text
hero cache
leaderboard cache
battle cooldowns
distributed locks
rate limiting
session or token metadata
```

Redis must not replace PostgreSQL as the source of truth for critical business data.

---

### MongoDB

MongoDB is used for flexible document-oriented data.

Planned usage:

```text
battle replay documents
audit documents
event debug logs
analytics-oriented snapshots
```

MongoDB is useful for flexible schemas and document-style reads.

---

### ClickHouse

ClickHouse is used for analytical and aggregated data.

Planned usage:

```text
battle analytics
hero progression analytics
reward distribution analytics
rating snapshots
event analytics
user activity analytics
```

ClickHouse is not a transactional source of truth.

It is intended for OLAP queries, large analytical scans, aggregated reporting and performance-oriented analytical reads.

ClickHouse data should be populated through event streams, background workers or explicit analytics pipelines.

---

### RabbitMQ

RabbitMQ is used for background jobs and task processing.

Planned usage:

```text
reward jobs
notification jobs
level-up jobs
dead letter queues
retry workflows
```

RabbitMQ is best suited for commands and jobs that should be processed by workers.

---

### Kafka

Kafka is used for event streaming and projections.

Planned usage:

```text
hero events
battle events
inventory events
rating events
audit events
analytics events
```

Kafka is best suited for durable event streams, consumer groups, ordering within partitions and analytics pipelines.

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

A learning topic should not be marked as complete only because it was demonstrated through isolated terminal output.

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

Describes the project purpose, final vision, target architecture, technology stack and development commands.

### ARCHITECTURE.md

Documents architectural decisions, target structure, package responsibilities, dependency rules and system evolution.

### LEARNING_CHECKLIST.md

Tracks learning progress and confirms which topics have been implemented in the project.

---

## Repository Rules

Every meaningful topic or feature should follow this workflow:

```text
1. Understand the topic.
2. Implement it in the correct project layer.
3. Add or update tests when appropriate.
4. Run checks.
5. Commit the change.
6. Mark the topic in LEARNING_CHECKLIST.md only after real implementation.
```

Recommended checks before commit:

```bash
gofmt -w .
go test ./...
go run ./cmd/arena
```

---

## Architecture Rules

New code should respect the target dependency direction:

```text
cmd
→ internal/app
→ internal/domain
```

```text
internal/app
→ internal/ports
```

```text
internal/adapters
→ internal/ports
→ internal/domain
```

Forbidden shortcuts:

```text
Do not put business logic in main.go.
Do not put business logic in HTTP handlers.
Do not put database-specific code in domain packages.
Do not put Redis, MongoDB, ClickHouse, RabbitMQ or Kafka clients in domain packages.
Do not make application use cases depend on concrete infrastructure implementations.
Do not create a central generic model package for all entities.
Do not duplicate domain rules in workers.
```

---

## Commit Style

The project uses clear commit messages that describe the purpose of each change.

Examples:

```text
init: create DevArena project
docs: update project readme
docs: add ClickHouse to architecture and learning checklist
learn: add hero stat lookup
refactor: split domain packages
feat: add REST API health endpoint
test: add battle service tests
ci: add GitHub Actions workflow
deploy: add Kubernetes manifests
```

---

## Planned Production Scenario

The final production-like scenario should look like this:

```text
1. User registers.
2. User logs in.
3. User creates a hero.
4. Hero starts a battle.
5. API validates JWT.
6. Middleware adds request ID and correlation ID.
7. Battle use case checks cooldown through Redis.
8. Battle use case acquires a distributed lock.
9. Battle use case loads hero and enemy through repository ports.
10. Domain battle logic runs the battle.
11. Battle result is saved in PostgreSQL inside a transaction.
12. Battle rounds are saved.
13. Outbox event is written in the same transaction.
14. Transaction is committed.
15. Outbox publisher publishes battle.finished to Kafka.
16. Reward job is published to RabbitMQ.
17. Reward worker grants experience and items.
18. Rating projector updates leaderboard projection.
19. Analytics worker writes analytical data to ClickHouse.
20. Audit log is recorded.
21. API returns battle result as JSON.
22. Logs, metrics and traces are emitted.
```

---

## Final Goal

At the end of the project, DevArena should demonstrate:

```text
Strong Go fundamentals
Domain-first design
Clear application use cases
Ports and adapters
REST API design
gRPC internal communication
PostgreSQL persistence
Redis caching and locking
MongoDB document modeling
ClickHouse analytics and OLAP queries
RabbitMQ background jobs
Kafka event streaming
Outbox pattern
Dockerized local development
Kubernetes deployment
CI/CD automation
Structured testing
Security basics
Observability
Performance awareness
Production-readiness thinking
```

The final project should be explainable as:

```text
DevArena is a production-oriented educational backend platform written in Go.

It starts from a text-based arena game domain and evolves into a modular backend system.

The core business logic is kept in domain packages, application use cases coordinate behavior through ports, and infrastructure integrations are implemented as adapters.

The system demonstrates REST API, PostgreSQL, Redis, MongoDB, ClickHouse, RabbitMQ, Kafka, gRPC, Docker, Kubernetes, CI/CD, testing, observability, security, analytics and production-readiness practices.
```

The expected result is a serious backend portfolio project with understandable architecture, meaningful commit history, tested business logic and infrastructure close to real production systems.