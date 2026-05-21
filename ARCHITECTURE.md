# DevArena Architecture

DevArena is an educational backend project designed to learn Go and backend development through practical implementation.

The project starts as a console battle simulator and gradually evolves into a production-like backend platform with REST API, PostgreSQL, Redis, MongoDB, gRPC, RabbitMQ, Kafka, Docker, Kubernetes, CI/CD, testing, security, observability and production-readiness practices.

The purpose of the architecture is not only to build a working application, but also to create a structured learning path from beginner level to confident intern, junior and middle-oriented backend development.

---

## Development Principle

The project is developed as a learning map.

Every topic must go through this cycle:

```text
1. Learn the concept
2. Apply it in DevArena
3. Explain why it is implemented this way
4. Commit the change
5. Mark the topic in LEARNING_CHECKLIST.md
6. Document important decisions in ARCHITECTURE.md
```

The goal is not to build fast.

The goal is to understand every layer.

---

## Project Concept

DevArena is a backend platform for a text-based arena game.

Users can:

```text
register
log in
create heroes
fight enemies
receive experience
level up
collect items
equip items
join tournaments
get rewards
appear on leaderboards
```

System components process:

```text
battle logic
rewards
notifications
ratings
audit logs
analytics
event streaming
background jobs
```

---

## Architecture Evolution

The project evolves step by step:

```text
Console Go application
→ Go packages
→ Error handling
→ Go concurrency
→ In-memory REST API
→ Layered backend architecture
→ PostgreSQL persistence
→ Docker Compose environment
→ Authentication and authorization
→ Redis mechanics
→ Automated tests
→ RabbitMQ workers
→ Kafka event streaming
→ MongoDB logs
→ gRPC internal services
→ Observability
→ Kubernetes deployment
→ CI/CD automation
→ Security hardening
→ Performance and production readiness
```

---

## Development Stages

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

## Initial Structure

```text
devarena/
  cmd/
    arena/
      main.go
  README.md
  ARCHITECTURE.md
  LEARNING_CHECKLIST.md
  .gitignore
  go.mod
```

---

## Final Target Structure

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
      app.go

    config/
      config.go

    handler/
      health_handler.go
      auth_handler.go
      hero_handler.go
      battle_handler.go
      inventory_handler.go
      rating_handler.go
      admin_handler.go

    service/
      auth_service.go
      hero_service.go
      battle_service.go
      inventory_service.go
      rating_service.go
      reward_service.go
      notification_service.go
      analytics_service.go

    repository/
      user_repository.go
      hero_repository.go
      battle_repository.go
      inventory_repository.go
      rating_repository.go
      refresh_token_repository.go
      battle_log_repository.go
      audit_log_repository.go

    model/
      user.go
      hero.go
      hero_class.go
      enemy.go
      battle.go
      battle_round.go
      item.go
      inventory.go
      reward.go
      rating.go
      event.go
      audit_log.go

    middleware/
      logger.go
      recovery.go
      auth.go
      rate_limiter.go
      request_id.go
      cors.go
      timeout.go

    response/
      response.go
      error.go

    validator/
      auth_validator.go
      hero_validator.go
      battle_validator.go
      inventory_validator.go

    infrastructure/
      postgres/
        connection.go
        user_repository.go
        hero_repository.go
        battle_repository.go
        inventory_repository.go
        rating_repository.go
        refresh_token_repository.go

      redis/
        client.go
        cache.go
        leaderboard.go
        rate_limiter.go
        session_store.go
        distributed_lock.go
        pubsub.go

      mongodb/
        client.go
        battle_log_repository.go
        audit_log_repository.go
        event_debug_log_repository.go

      rabbitmq/
        connection.go
        publisher.go
        consumer.go
        retry.go
        dlq.go

      kafka/
        producer.go
        consumer.go
        events.go
        outbox.go

      grpc/
        clients.go
        servers.go
        interceptors.go

      observability/
        logger.go
        metrics.go
        tracing.go

  proto/
    hero/
      hero.proto
    battle/
      battle.proto
    rating/
      rating.proto
    inventory/
      inventory.proto

  gen/
    hero/
    battle/
    rating/
    inventory/

  migrations/
    postgres/
      000001_create_users_table.up.sql
      000001_create_users_table.down.sql
      000002_create_heroes_table.up.sql
      000002_create_heroes_table.down.sql
      000003_create_enemies_table.up.sql
      000003_create_enemies_table.down.sql
      000004_create_battles_table.up.sql
      000004_create_battles_table.down.sql
      000005_create_battle_rounds_table.up.sql
      000005_create_battle_rounds_table.down.sql
      000006_create_items_table.up.sql
      000006_create_items_table.down.sql
      000007_create_hero_inventory_table.up.sql
      000007_create_hero_inventory_table.down.sql
      000008_create_ratings_table.up.sql
      000008_create_ratings_table.down.sql
      000009_create_refresh_tokens_table.up.sql
      000009_create_refresh_tokens_table.down.sql
      000010_create_outbox_events_table.up.sql
      000010_create_outbox_events_table.down.sql

  tests/
    integration/
      api_test.go
      postgres_test.go
      redis_test.go
      rabbitmq_test.go
      kafka_test.go

  deploy/
    k8s/
      namespace.yaml
      configmap.yaml
      secret.yaml
      api-deployment.yaml
      api-service.yaml
      notification-worker-deployment.yaml
      reward-worker-deployment.yaml
      analytics-worker-deployment.yaml
      postgres.yaml
      redis.yaml
      mongodb.yaml
      rabbitmq.yaml
      kafka.yaml
      ingress.yaml
      hpa.yaml

  .github/
    workflows/
      ci.yml
      docker.yml
      deploy.yml

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

## User

Represents a registered platform user.

```text
User:
- id
- email
- password_hash
- role
- created_at
- updated_at
```

Used for:

```text
authentication
authorization
JWT
password hashing
PostgreSQL constraints
indexes
refresh tokens
audit logs
```

---

## Hero

Represents a playable character owned by a user.

```text
Hero:
- id
- user_id
- name
- class
- level
- experience
- hp
- damage
- armor
- created_at
- updated_at
```

Used for:

```text
Go variables
structs
methods
pointers
REST CRUD
PostgreSQL relations
Redis cache
gRPC service methods
Kafka events
```

---

## Enemy

Represents an opponent in the arena.

```text
Enemy:
- id
- name
- level
- hp
- damage
- reward_exp
```

Used for:

```text
conditions
loops
arrays
slices
maps
battle logic
```

---

## Battle

Represents a fight between a hero and an enemy.

```text
Battle:
- id
- hero_id
- enemy_id
- status
- winner_type
- winner_id
- started_at
- finished_at
```

Used for:

```text
business logic
transactions
RabbitMQ jobs
Kafka events
MongoDB battle logs
gRPC streaming
```

---

## BattleRound

Represents one round inside a battle.

```text
BattleRound:
- id
- battle_id
- round_number
- attacker_type
- attacker_id
- defender_type
- defender_id
- damage
- created_at
```

Used for:

```text
loops
slices
ORDER BY
JOIN
LIMIT
battle replay
MongoDB embedded documents
```

---

## Item

Represents a game item.

```text
Item:
- id
- name
- type
- rarity
- damage_bonus
- armor_bonus
- created_at
```

Used for:

```text
interfaces
inventory
transactions
many-to-many relations
reward workers
```

---

## Inventory

Represents a hero inventory.

```text
InventoryItem:
- id
- hero_id
- item_id
- quantity
- equipped
```

Used for:

```text
slices
maps
JOIN
transactions
UNIQUE constraints
```

---

## Rating

Represents hero leaderboard data.

```text
Rating:
- hero_id
- wins
- losses
- score
```

Used for:

```text
GROUP BY
ORDER BY
LIMIT
Redis sorted sets
Kafka rating projector
```

---

## GameEvent

Represents a domain event.

```text
GameEvent:
- event_id
- event_type
- aggregate_id
- aggregate_type
- occurred_at
- payload
```

Used for:

```text
Kafka
RabbitMQ
outbox pattern
audit logs
analytics
event-driven architecture
```

---

## Layered Architecture

The main backend request flow:

```text
HTTP request
→ middleware
→ handler
→ service
→ repository
→ database/storage
```

---

## Handler Layer

Responsible for:

```text
HTTP requests
JSON decoding
path parameters
query parameters
headers
basic request validation
calling services
mapping service errors to HTTP status codes
returning JSON responses
```

Handlers must not contain business logic.

---

## Service Layer

Responsible for:

```text
business logic
battle rules
hero leveling
reward calculation
cooldown checks
authorization rules
transaction orchestration
event publishing
calling repositories
calling infrastructure clients
```

Services contain the main application behavior.

---

## Repository Layer

Responsible for:

```text
data persistence
SQL queries
PostgreSQL CRUD
MongoDB CRUD
Redis operations
transactions
query methods
```

Repositories must not contain HTTP logic.

---

## Infrastructure Layer

Responsible for:

```text
PostgreSQL connection
Redis client
MongoDB client
Kafka producer
Kafka consumer
RabbitMQ publisher
RabbitMQ consumer
gRPC clients
gRPC servers
observability clients
```

---

## Planned REST API

### Health

```http
GET /health
GET /health/live
GET /health/ready
```

### Auth

```http
POST /api/v1/auth/register
POST /api/v1/auth/login
POST /api/v1/auth/refresh
POST /api/v1/auth/logout
GET  /api/v1/users/me
```

### Heroes

```http
POST   /api/v1/heroes
GET    /api/v1/heroes
GET    /api/v1/heroes/{id}
PUT    /api/v1/heroes/{id}
PATCH  /api/v1/heroes/{id}
DELETE /api/v1/heroes/{id}
```

### Battles

```http
POST /api/v1/battles
GET  /api/v1/battles/{id}
GET  /api/v1/heroes/{id}/battles
GET  /api/v1/battles/{id}/rounds
GET  /api/v1/battles/{id}/log
```

### Inventory

```http
GET    /api/v1/heroes/{id}/inventory
POST   /api/v1/heroes/{id}/inventory/items
PATCH  /api/v1/heroes/{id}/inventory/items/{item_id}/equip
DELETE /api/v1/heroes/{id}/inventory/items/{item_id}
```

### Rating

```http
GET /api/v1/leaderboard
GET /api/v1/heroes/{id}/rating
```

### Admin and Analytics

```http
GET /api/v1/admin/stats
GET /api/v1/admin/audit-logs
GET /api/v1/admin/events
```

---

## REST API Principles

The API must use:

```text
clear resource names
consistent status codes
JSON request and response bodies
unified error format
pagination
filtering
sorting
validation
request ID
authentication middleware
authorization rules
rate limiting
graceful shutdown
```

---

## Unified Error Format

All API errors should follow one structure:

```json
{
  "error": {
    "code": "hero_not_found",
    "message": "hero not found",
    "details": {
      "hero_id": 123
    }
  }
}
```

Common error codes:

```text
validation_error
unauthorized
forbidden
hero_not_found
battle_not_found
cooldown_active
rate_limit_exceeded
conflict
internal_error
```

---

## Planned PostgreSQL Tables

```text
users
heroes
enemies
battles
battle_rounds
items
hero_inventory
ratings
refresh_tokens
outbox_events
```

---

## PostgreSQL Responsibilities

PostgreSQL stores core relational data:

```text
users
heroes
battle metadata
battle rounds
items
inventory
ratings
refresh tokens
outbox events
```

---

## PostgreSQL Topics Used

```text
SELECT
INSERT
UPDATE
DELETE
WHERE
JOIN
LEFT JOIN
GROUP BY
HAVING
ORDER BY
LIMIT
OFFSET
COUNT
SUM
AVG
MIN
MAX
primary key
foreign key
unique constraint
not null constraint
check constraint
indexes
transactions
migrations
connection pool
pgx
context with DB queries
EXPLAIN
EXPLAIN ANALYZE
transaction isolation levels
SELECT FOR UPDATE
optimistic locking
upsert
CTE
window functions
N+1 problem
slow query analysis
```

---

## Planned Redis Usage

```text
hero:{id}
leaderboard:global
battle:cooldown:hero:{id}
rate_limit:ip:{ip}
refresh_token:{token_id}
lock:battle:hero:{id}
arena:events
```

Redis is used for:

```text
caching
leaderboard
rate limiting
session storage
distributed locking
battle cooldowns
pub/sub notifications
```

---

## Planned MongoDB Collections

```text
battle_logs
audit_logs
event_debug_logs
```

MongoDB is used for:

```text
battle replay documents
audit logs
debug event payloads
flexible schema data
embedded documents
aggregation pipeline
```

---

## Planned RabbitMQ Usage

RabbitMQ is used for background jobs.

```text
Exchanges:
- devarena.direct
- devarena.fanout
- devarena.topic
- devarena.dlx

Queues:
- notification.queue
- reward.queue
- levelup.queue
- dead_letter.queue

Routing keys:
- battle.finished
- hero.level_up
- reward.grant
- notification.send
```

RabbitMQ handles:

```text
reward jobs
notification jobs
level-up jobs
retry
dead letter queue
manual ack/nack
background processing
```

---

## Planned Kafka Usage

Kafka is used for event streaming and analytics.

```text
Topics:
- devarena.hero.events
- devarena.battle.events
- devarena.inventory.events
- devarena.rating.events
```

Kafka consumers:

```text
analytics-worker
audit-worker
rating-projector
```

Kafka event types:

```text
hero.created
hero.updated
battle.started
battle.finished
hero.level_up
item.dropped
item.equipped
rating.updated
```

Kafka design principles:

```text
partition by hero_id
preserve ordering inside partition
at-least-once delivery
idempotent consumers
consumer groups
offset management
dead letter topic
event versioning
```

---

## Planned gRPC Services

```text
HeroService
BattleService
RatingService
InventoryService
```

RPC examples:

```text
HeroService.GetHero
HeroService.ValidateHero
BattleService.StartBattle
BattleService.GetBattle
BattleService.StreamBattleRounds
BattleService.SendBattleActions
BattleService.LiveBattle
RatingService.GetHeroRating
RatingService.UpdateRating
```

RPC types to learn:

```text
unary RPC
server streaming
client streaming
bidirectional streaming
```

gRPC must use:

```text
Protocol Buffers
generated Go code
metadata
deadlines
status codes
interceptors
health checks
grpcurl
```

---

## Planned Docker Setup

Docker is used for local development and reproducible environments.

Services:

```text
api
postgres
redis
mongodb
rabbitmq
kafka
notification-worker
reward-worker
analytics-worker
```

Docker topics:

```text
image
container
Dockerfile
docker build
docker run
docker exec
docker logs
docker compose
volumes
networks
environment variables
multi-stage build
.dockerignore
non-root user
healthcheck
```

---

## Planned Kubernetes Resources

```text
namespace
deployment
replica set
pod
service
ingress
configmap
secret
readiness probe
liveness probe
resource requests
resource limits
horizontal pod autoscaler
persistent volume
persistent volume claim
job
cronjob
service account
RBAC
network policy
```

Kubernetes services:

```text
api-service
postgres-service
redis-service
mongodb-service
rabbitmq-service
kafka-service
```

Kubernetes deployments:

```text
api-deployment
notification-worker-deployment
reward-worker-deployment
analytics-worker-deployment
```

---

## Planned CI/CD

GitHub Actions workflows:

```text
ci.yml
docker.yml
deploy.yml
```

CI pipeline:

```text
checkout
setup Go
gofmt check
go vet
go test ./...
go test -race ./...
coverage report
golangci-lint
docker build
```

Docker pipeline:

```text
docker login
docker build
docker push
```

Deploy pipeline:

```text
load kubeconfig from secrets
kubectl apply
kubectl rollout status
```

CI/CD topics:

```text
pipeline
build
test
lint
docker build
docker push
deploy
GitHub Actions
secrets
artifacts
environments
branch protection
release tags
rollback
deployment strategies
```

---

## Go Concurrency Architecture

Concurrency will be introduced after core Go topics.

Concurrency use cases:

```text
parallel battle simulation
battle event channels
reward worker pool
rating update synchronization
context cancellation
timeout-aware operations
graceful shutdown
background workers
Kafka consumers
RabbitMQ consumers
```

Concurrency topics:

```text
goroutines
channels
buffered channels
select
sync.WaitGroup
sync.Mutex
sync.RWMutex
sync.Once
atomic counters
race conditions
go test -race
context.Context
worker pool
fan-out/fan-in
backpressure
goroutine leaks
bounded concurrency
```

---

## Error Handling Architecture

Errors must be explicit and structured.

Error categories:

```text
validation errors
domain errors
repository errors
infrastructure errors
authentication errors
authorization errors
conflict errors
not found errors
retryable errors
non-retryable errors
```

Error handling topics:

```text
errors.New
fmt.Errorf
%w wrapping
errors.Is
errors.As
sentinel errors
custom error types
mapping errors to HTTP status codes
structured error responses
```

---

## Testing Architecture

Testing levels:

```text
unit tests
table-driven tests
handler tests
repository tests
integration tests
contract tests
benchmark tests
fuzz tests
```

Testing tools and concepts:

```text
testing package
httptest
mocks
fakes
stubs
test helpers
test fixtures
testcontainers
go test ./...
go test -race
coverage
golden files
deterministic tests
flaky test prevention
```

---

## Observability Architecture

Observability includes:

```text
structured logging
request ID
correlation ID
log levels
metrics
health checks
readiness
liveness
tracing
alerts
dashboards
```

Planned tools/concepts:

```text
Prometheus
Grafana
OpenTelemetry
trace ID
span
RED metrics
USE metrics
error rate
latency
throughput
```

---

## Security Architecture

Security topics:

```text
password hashing
bcrypt
JWT access tokens
refresh tokens
token revocation
authorization
roles
permissions
input validation
SQL injection prevention
rate limiting
secure headers
secret management
audit logs
least privilege
dependency vulnerability scanning
```

---

## Performance Architecture

Performance topics:

```text
pagination
indexes
connection pooling
caching
timeouts
N+1 query prevention
query optimization
EXPLAIN ANALYZE
pprof
CPU profiling
memory profiling
benchmarks
load testing
latency percentiles
throughput
backpressure
horizontal scaling
```

---

## Production Readiness Architecture

Production-readiness topics:

```text
configuration via environment variables
health checks
graceful shutdown
database migrations
safe migrations
Docker Compose
Kubernetes deployment
CI/CD checks
rollback strategy
feature flags
alerts
dashboards
runbooks
incident analysis
postmortem basics
```

---

## Core Production-Like Scenario

Final DevArena scenario:

```text
1. User registers
2. User logs in
3. User creates Hero
4. Hero starts Battle
5. API checks JWT
6. Middleware adds request ID
7. BattleService checks cooldown in Redis
8. BattleService uses Redis distributed lock
9. BattleService loads Hero and Enemy from PostgreSQL
10. BattleService runs battle algorithm
11. Battle result is saved in PostgreSQL transaction
12. Battle rounds are saved in PostgreSQL
13. Full battle log is saved in MongoDB
14. battle.finished event is written to outbox_events
15. Outbox publisher publishes event to Kafka
16. reward.grant job is published to RabbitMQ
17. RewardWorker grants experience or item
18. RatingProjector updates Redis leaderboard
19. AnalyticsWorker updates statistics
20. API returns battle result as JSON
21. Logs, metrics and traces are emitted
22. CI pipeline verifies tests and build
23. Docker image is built and pushed
24. Kubernetes deploy updates application
```

---

## Final Architecture Goal

At the end of the project, DevArena should demonstrate:

```text
strong Go fundamentals
clear error handling
practical concurrency
clean backend architecture
REST API design
PostgreSQL persistence
Redis caching and locking
MongoDB document modeling
RabbitMQ background jobs
Kafka event streaming
gRPC internal communication
Dockerized development environment
Kubernetes deployment
CI/CD automation
testing culture
security basics
observability
performance awareness
production-readiness thinking
```