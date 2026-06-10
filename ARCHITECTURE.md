# DevArena Architecture

DevArena is a production-oriented educational backend platform written in Go.

The project uses a text-based arena game domain to learn and demonstrate practical backend engineering: domain modeling, application use cases, persistence, caching, messaging, distributed communication, analytics, observability, testing, deployment and production-readiness practices.

The architecture is intentionally designed as a **domain-first modular backend system**.

---

## Architectural Direction

DevArena follows this architectural direction:

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

The project should not be organized around generic technical folders only.

Instead of growing into a structure like this:

```text
internal/
  model/
  service/
  repository/
  handler/
```

DevArena should grow around business capabilities and architectural boundaries:

```text
internal/
  domain/
    hero/
    enemy/
    battle/
    reward/
    inventory/
    rating/
    event/

  app/
    hero/
    battle/
    reward/
    inventory/
    rating/
    event/

  ports/
  adapters/
  platform/
```

Technical layers still exist, but they are organized around clear ownership and dependency direction.

---

## Core Architecture Rule

The most important rule:

```text
Domain logic must not depend on infrastructure.
```

Domain packages must not import:

```text
HTTP frameworks
database drivers
Redis clients
Kafka clients
RabbitMQ clients
MongoDB clients
ClickHouse clients
gRPC transport code
environment configuration
logging frameworks
observability frameworks
```

Domain packages may use:

```text
Go standard library
other domain packages when the dependency is meaningful
small domain-specific value objects and interfaces
```

Application packages coordinate domain logic and depend on ports.

Adapters implement ports using real technologies.

---

## Target System Overview

DevArena is a backend platform for a text-based arena game.

The final system supports:

```text
User registration and authentication
Hero creation and management
Enemy management
Battle simulation
Battle rounds
Inventory management
Reward calculation and granting
Experience and progression
Ratings and leaderboards
Tournaments
Battle history
Domain events
Outbox events
Background workers
Analytics
Audit logging
```

---

## Target Runtime Components

The final runtime topology includes:

```text
API application
Console arena application
Reward worker
Notification worker
Analytics worker
Outbox publisher
PostgreSQL
Redis
MongoDB
ClickHouse
RabbitMQ
Kafka
gRPC internal services
```

The project may start as a console application, but the final architecture must support multiple entrypoints.

---

## Target Repository Structure

The target repository structure is:

```text
devarena/
  cmd/
    arena/
      main.go
    api/
      main.go
    reward-worker/
      main.go
    notification-worker/
      main.go
    analytics-worker/
      main.go
    outbox-publisher/
      main.go

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
        handler/
        middleware/
        request/
        response/
        router/

      grpc/
        server/
        client/
        interceptor/

      postgres/
        repository/
        transaction/
        migration/

      redis/
        cache/
        lock/
        leaderboard/
        ratelimit/

      mongodb/
        repository/

      clickhouse/
        analytics/
        migration/

      rabbitmq/
        publisher/
        consumer/

      kafka/
        producer/
        consumer/

    platform/
      config/
      logger/
      observability/
      shutdown/
      validation/

  proto/
    hero/
    battle/
    rating/
    inventory/

  gen/
    hero/
    battle/
    rating/
    inventory/

  migrations/
    postgres/
    clickhouse/

  tests/
    integration/
    e2e/

  deploy/
    docker/
    k8s/

  .github/
    workflows/

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

## Dependency Direction

Allowed dependency flow:

```text
cmd
→ internal/app
→ internal/domain
→ Go standard library
```

```text
cmd
→ internal/adapters
→ internal/ports
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

```text
internal/platform
→ infrastructure libraries
```

Forbidden dependency flow:

```text
domain → adapters
domain → platform
domain → database driver
domain → HTTP transport
domain → Kafka/RabbitMQ/Redis/MongoDB/ClickHouse clients
app    → concrete database implementation
app    → concrete analytics storage implementation
app    → concrete message broker implementation
```

The intended dependency graph:

```text
cmd
 ├── app
 │    ├── domain
 │    └── ports
 ├── adapters
 │    ├── ports
 │    └── domain
 └── platform
```

---

## Layer Responsibilities

### `cmd/`

Contains executable entrypoints.

Examples:

```text
cmd/arena
cmd/api
cmd/reward-worker
cmd/notification-worker
cmd/analytics-worker
cmd/outbox-publisher
```

Responsibilities:

```text
Load configuration
Create dependencies
Wire application services
Start the process
Handle shutdown
Return process exit status
```

`cmd` packages should not contain business logic.

---

### `internal/domain/`

Contains business concepts and business rules.

Examples:

```text
domain/user
domain/auth
domain/hero
domain/enemy
domain/battle
domain/reward
domain/inventory
domain/rating
domain/tournament
domain/event
domain/audit
```

Responsibilities:

```text
Entities
Value objects
Domain services
Domain events
Business invariants
Pure calculations
State transitions
Domain-specific interfaces when needed
```

Domain code should be easy to test without databases, HTTP servers or external systems.

Domain code must not know about PostgreSQL, Redis, MongoDB, ClickHouse, RabbitMQ, Kafka, HTTP, gRPC or Kubernetes.

---

### `internal/app/`

Contains application use cases.

Examples:

```text
app/hero/CreateHero
app/battle/StartBattle
app/reward/GrantReward
app/rating/GetLeaderboard
app/event/PublishOutboxEvents
app/analytics/RecordBattleAnalytics
```

Responsibilities:

```text
Coordinate domain objects
Call repositories through ports
Run transactions through transaction ports
Check authorization decisions when needed
Publish events through publisher ports
Coordinate cache and lock ports
Coordinate analytics ports
Return application-level results
```

Application layer should not know which exact database, broker, cache or analytics implementation is used.

---

### `internal/ports/`

Contains interfaces required by application services.

Examples:

```text
HeroRepository
BattleRepository
RewardRepository
EventPublisher
OutboxRepository
LeaderboardCache
DistributedLock
TransactionManager
AnalyticsWriter
AnalyticsReader
Clock
IDGenerator
PasswordHasher
TokenManager
```

Ports define what the application needs, not how infrastructure works.

Good port example:

```go
type HeroRepository interface {
	Save(ctx context.Context, hero hero.Hero) error
	FindByID(ctx context.Context, id hero.ID) (hero.Hero, error)
}
```

Good analytics port example:

```go
type BattleAnalyticsWriter interface {
	RecordBattleFinished(ctx context.Context, event BattleFinishedAnalyticsEvent) error
}
```

Bad port example:

```go
type PostgresHeroRepository interface {
	ExecSQL(query string, args ...any) error
}
```

Bad analytics port example:

```go
type ClickHouseBattleAnalytics interface {
	InsertIntoClickHouse(ctx context.Context, query string, args ...any) error
}
```

Ports should be technology-neutral.

---

### `internal/adapters/`

Contains concrete implementations of ports and transport layers.

Examples:

```text
HTTP handlers
gRPC servers
PostgreSQL repositories
Redis cache implementation
MongoDB log repository
ClickHouse analytics repository
RabbitMQ publisher and consumer
Kafka producer and consumer
```

Adapters may depend on external libraries.

Adapters translate between external systems and internal application/domain types.

---

### `internal/platform/`

Contains process-level technical utilities.

Examples:

```text
configuration loading
structured logging setup
metrics setup
tracing setup
validation helpers
graceful shutdown
runtime diagnostics
```

Platform code should not contain business rules.

---

## Domain Boundaries

### User

Represents a registered platform user.

Owns:

```text
id
email
password hash
role
created at
updated at
```

Related capabilities:

```text
registration
login
authentication
authorization
refresh tokens
audit logging
```

---

### Auth

Represents authentication and authorization rules.

Owns:

```text
password hashing policy
token creation
token validation
roles
permissions
session invalidation
refresh token lifecycle
```

Auth should not be mixed into HTTP handlers directly.

HTTP handlers call application use cases.

---

### Hero

Represents a playable character.

Owns:

```text
id
user id
name
class
level
experience
combat stats
inventory references
equipped items
status
```

Hero domain rules include:

```text
name validation
class restrictions
level progression
damage calculation
stat management
equipment effects
inventory constraints
```

---

### Enemy

Represents an arena opponent.

Owns:

```text
id
name
level
combat stats
reward profile
behavior profile
```

Enemy domain rules include:

```text
damage receiving
enemy difficulty
reward eligibility
battle participation
```

---

### Battle

Represents a fight between a hero and an enemy.

Owns:

```text
battle id
hero participant
enemy participant
battle status
rounds
winner
result
started at
finished at
```

Battle domain rules include:

```text
round progression
damage application
winner detection
battle result calculation
battle event creation
```

Battle logic belongs in the domain layer.

Persistence, HTTP and messaging do not belong in battle domain code.

---

### Reward

Represents experience, items and other gains after battle or tournament activity.

Owns:

```text
experience amount
reward item
reward rules
reward conditions
reward event data
```

Reward domain rules include:

```text
base reward calculation
level-based reward calculation
enemy-specific reward rules
reward item selection
reward granting eligibility
```

---

### Inventory

Represents hero-owned items.

Owns:

```text
item list
item quantity
equipped state
capacity constraints
item snapshots
```

Inventory domain rules include:

```text
adding items
removing items
equipping items
copying inventory snapshots
capacity handling
```

---

### Rating

Represents leaderboard and competitive score.

Owns:

```text
hero id
wins
losses
score
rank projection
```

Rating domain rules include:

```text
win/loss updates
score calculation
leaderboard projection events
ranking consistency
```

---

### Event

Represents domain and integration events.

Owns:

```text
event id
event type
aggregate id
aggregate type
occurred at
payload
metadata
version
```

Events are used for:

```text
outbox pattern
Kafka streaming
RabbitMQ job publishing
audit logging
analytics
event replay and debugging
```

---

### Tournament

Represents structured competition between heroes.

Owns:

```text
tournament id
participants
state
rounds
winner
rewards
schedule
```

Tournament can be added after core battle and rating logic is stable.

---

### Audit

Represents security and operational audit records.

Owns:

```text
actor id
action
resource
timestamp
metadata
request id
correlation id
```

Audit data should be append-only.

---

## Application Use Cases

Application layer should be organized by use cases, not by generic service names.

Examples:

```text
CreateHero
GetHero
UpdateHero
StartBattle
GetBattleResult
GrantReward
GetInventory
EquipItem
GetLeaderboard
RecordBattleAnalytics
PublishOutboxEvents
RegisterUser
LoginUser
RefreshToken
```

A use case may coordinate:

```text
domain rules
repository ports
transaction ports
cache ports
event publisher ports
analytics ports
lock ports
clock and ID generator ports
```

Use cases should be testable with fake ports.

---

## HTTP API Architecture

HTTP is an adapter, not the core application.

HTTP handlers are responsible for:

```text
routing
request decoding
path and query parameters
headers
input validation
calling application use cases
mapping application errors to HTTP status codes
returning JSON responses
```

HTTP handlers must not contain battle logic, reward logic, analytics logic or persistence logic.

Target REST API groups:

```text
/health
/api/v1/auth
/api/v1/users
/api/v1/heroes
/api/v1/battles
/api/v1/inventory
/api/v1/leaderboard
/api/v1/tournaments
/api/v1/analytics
/api/v1/admin
```

---

## REST API Principles

The REST API should use:

```text
clear resource names
consistent status codes
JSON request and response bodies
unified error format
pagination
filtering
sorting
request ID
authentication middleware
authorization rules
rate limiting
graceful shutdown
```

Unified error response:

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

---

## gRPC Architecture

gRPC is used for internal service communication and streaming scenarios.

Target services:

```text
HeroService
BattleService
RatingService
InventoryService
AnalyticsService
```

gRPC responsibilities:

```text
internal service APIs
streaming battle rounds
service-to-service calls
typed contracts through Protocol Buffers
deadlines
metadata
status codes
interceptors
health checks
```

gRPC code belongs in adapters.

Generated code belongs in `gen/`.

Protocol Buffer definitions belong in `proto/`.

---

## Persistence Architecture

### PostgreSQL

PostgreSQL stores core relational state.

Planned tables:

```text
users
refresh_tokens
heroes
enemies
battles
battle_rounds
items
hero_inventory
ratings
tournaments
tournament_participants
outbox_events
```

PostgreSQL is responsible for:

```text
source-of-truth relational data
transactions
constraints
indexes
foreign keys
query consistency
outbox storage
```

PostgreSQL access must be implemented through repository adapters.

Application code depends on repository ports, not PostgreSQL directly.

---

### Redis

Redis is used for fast ephemeral and derived data.

Planned usage:

```text
hero cache
leaderboard cache
battle cooldowns
rate limiting
distributed locks
session or token metadata
pub/sub notifications where useful
```

Example keys:

```text
hero:{id}
leaderboard:global
battle:cooldown:hero:{id}
rate_limit:ip:{ip}
refresh_token:{token_id}
lock:battle:hero:{id}
```

Redis must not be the source of truth for critical relational state.

---

### MongoDB

MongoDB stores flexible document-oriented data.

Planned collections:

```text
battle_logs
audit_logs
event_debug_logs
analytics_snapshots
```

MongoDB is used for:

```text
battle replay documents
audit records
debug event payloads
analytics-oriented documents
flexible schema data
embedded documents
aggregation pipelines
```

MongoDB must not replace PostgreSQL for core transactional state.

---

### ClickHouse

ClickHouse stores analytical and aggregated data.

Planned tables:

```text
battle_analytics
hero_progression_analytics
reward_analytics
rating_snapshots
event_analytics
user_activity_analytics
```

ClickHouse is used for:

```text
OLAP queries
large analytical scans
battle statistics
reward distribution analytics
event analytics
leaderboard history
aggregated reporting
performance-oriented analytical reads
```

ClickHouse must not be used as the source of truth for transactional state.

PostgreSQL remains the source of truth for relational business data.

ClickHouse data should be populated through event streams, workers or explicit analytics pipelines.

Application code must depend on analytics ports, not ClickHouse directly.

---

## Transaction and Consistency Model

Application services coordinate transactions through ports.

Transaction-sensitive use cases include:

```text
CreateHero
StartBattle
GrantReward
EquipItem
UpdateRating
PublishOutboxEvents
```

Rules:

```text
Business decisions are made in application/domain code.
Database transaction boundaries are controlled by application use cases.
Repository adapters execute concrete queries.
Outbox events are written in the same transaction as state changes.
External event publishing happens after transaction commit.
Analytics storage is populated asynchronously when possible.
```

The preferred consistency pattern:

```text
PostgreSQL transaction
→ state change
→ outbox event inserted
→ transaction committed
→ outbox publisher reads event
→ Kafka/RabbitMQ publish
→ analytics worker consumes event
→ ClickHouse analytics row inserted
→ event marked as processed where needed
```

---

## Event-Driven Architecture

DevArena uses events for asynchronous processing and integration.

Event categories:

```text
domain events
integration events
outbox events
audit events
analytics events
background jobs
```

Example event types:

```text
hero.created
hero.updated
battle.started
battle.finished
reward.granted
item.equipped
rating.updated
tournament.finished
analytics.recorded
```

Event rules:

```text
Events must have stable type names.
Events must have explicit versions.
Events must be serializable.
Events must include aggregate information.
Consumers must be idempotent.
Event publishing should use the outbox pattern for reliability.
Analytics consumers should tolerate duplicates.
```

---

## RabbitMQ Architecture

RabbitMQ is used for background jobs and task processing.

Planned queues:

```text
reward.queue
notification.queue
levelup.queue
dead_letter.queue
```

Planned routing keys:

```text
battle.finished
reward.grant
hero.level_up
notification.send
```

RabbitMQ handles:

```text
job dispatching
retry
dead letter queue
manual ack/nack
background processing
worker scaling
```

RabbitMQ is best suited for commands/jobs that should be processed by workers.

---

## Kafka Architecture

Kafka is used for event streaming and analytics.

Planned topics:

```text
devarena.hero.events
devarena.battle.events
devarena.inventory.events
devarena.rating.events
devarena.audit.events
devarena.analytics.events
```

Kafka consumers:

```text
analytics-worker
audit-worker
rating-projector
event-debug-consumer
clickhouse-sink-worker
```

Kafka principles:

```text
partition by aggregate id when ordering matters
preserve ordering inside partition
at-least-once delivery
idempotent consumers
consumer groups
offset management
dead letter topic
event versioning
schema compatibility
```

Kafka is best suited for event streams and projections.

---

## Analytics Architecture

Analytics is separated from transactional business state.

Analytics data may come from:

```text
domain events
integration events
outbox events
worker-produced events
periodic projections
batch backfills
```

Analytics storage may include:

```text
ClickHouse for OLAP queries
MongoDB for flexible debug and replay documents
Redis for fast derived leaderboard views
```

Analytics use cases should depend on analytics ports.

Analytics adapters implement those ports using concrete systems.

Rules:

```text
Do not use ClickHouse for transactional writes.
Do not make domain logic depend on analytics storage.
Do not block critical user flows on non-critical analytics writes unless required.
Prefer asynchronous analytics ingestion.
Design analytics inserts to be idempotent or duplicate-tolerant.
```

---

## Worker Architecture

Workers are separate entrypoints under `cmd/`.

Examples:

```text
cmd/reward-worker
cmd/notification-worker
cmd/analytics-worker
cmd/outbox-publisher
```

Worker responsibilities:

```text
consume jobs or events
call application use cases
handle retries
handle idempotency
write analytics data
emit logs and metrics
shutdown gracefully
```

Workers should use the same application layer as the API.

Workers should not duplicate business logic.

---

## Configuration Architecture

Configuration belongs in `internal/platform/config`.

Configuration sources:

```text
environment variables
.env files for local development
Kubernetes secrets and config maps
CI/CD secrets
```

Configuration must be explicit.

Examples:

```text
HTTP_PORT
DATABASE_URL
REDIS_ADDR
MONGODB_URI
CLICKHOUSE_DSN
RABBITMQ_URL
KAFKA_BROKERS
JWT_SECRET
LOG_LEVEL
TRACING_ENABLED
```

Domain code must not read environment variables.

---

## Error Handling Architecture

Errors must be explicit and structured.

Error categories:

```text
validation errors
domain errors
application errors
repository errors
infrastructure errors
analytics errors
authentication errors
authorization errors
conflict errors
not found errors
retryable errors
non-retryable errors
```

Rules:

```text
Domain errors describe business rule violations.
Application errors describe use case failures.
Adapters map errors to protocol-specific responses.
HTTP maps errors to status codes.
gRPC maps errors to gRPC status codes.
Workers map retryable errors to retry behavior.
Analytics ingestion errors should be observable and retryable when possible.
```

Go error handling should use:

```text
errors.New
fmt.Errorf with %w
errors.Is
errors.As
sentinel errors where appropriate
custom error types where useful
```

---

## Observability Architecture

Observability includes:

```text
structured logging
request ID
correlation ID
trace ID
metrics
health checks
readiness probes
liveness probes
distributed tracing
alerts
dashboards
```

Planned tools and concepts:

```text
OpenTelemetry
Prometheus
Grafana
structured JSON logs
RED metrics
USE metrics
latency percentiles
error rate
throughput
analytics ingestion lag
worker processing lag
```

Observability should be added at adapter and platform boundaries.

Domain code should not depend on observability libraries.

---

## Security Architecture

Security areas:

```text
password hashing
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

Security rules:

```text
Passwords are never stored in plain text.
Tokens are validated in middleware and application services.
Authorization decisions are explicit.
Admin endpoints require role checks.
Sensitive configuration comes from environment or secret storage.
Audit logs are append-only.
Analytics data must not expose sensitive secrets.
```

---

## Testing Architecture

Testing levels:

```text
unit tests
table-driven tests
application use case tests
handler tests
repository integration tests
analytics adapter integration tests
worker tests
contract tests
benchmark tests
fuzz tests
end-to-end tests
```

Testing principles:

```text
Domain tests should be fast and isolated.
Application tests should use fake ports.
Adapter tests may use testcontainers or local infrastructure.
Repository tests should verify real database behavior.
Analytics tests should verify schema, inserts and query behavior.
Handler tests should use httptest.
Worker tests should verify retry and idempotency behavior.
Critical paths should have table-driven tests.
Performance-sensitive code should have benchmarks.
```

Commands:

```bash
go test ./...
go test -race ./...
go test -cover ./...
go test -bench=. ./...
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
ClickHouse query planning
ClickHouse table engines
ClickHouse partitioning
ClickHouse ordering keys
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

Performance rules:

```text
Optimize based on measurement.
Use benchmarks for hot paths.
Use indexes for query patterns.
Use Redis for derived fast-access data.
Use ClickHouse for analytical scans and aggregations.
Use context timeouts for external calls.
Avoid premature distributed complexity.
```

---

## Deployment Architecture

Local development uses Docker Compose.

Production-like deployment uses Kubernetes.

Planned Docker services:

```text
api
postgres
redis
mongodb
clickhouse
rabbitmq
kafka
reward-worker
notification-worker
analytics-worker
outbox-publisher
```

Planned Kubernetes resources:

```text
namespace
deployment
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

---

## CI/CD Architecture

GitHub Actions is used for continuous integration and deployment.

Planned workflows:

```text
ci.yml
docker.yml
deploy.yml
```

CI pipeline should include:

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

Deployment pipeline should include:

```text
build image
push image
apply Kubernetes manifests
wait for rollout
verify health checks
```

---

## Architecture Governance

Every new feature or learning topic must respect the architecture.

Before adding code, decide where it belongs:

```text
Is it business logic?
→ internal/domain

Is it a use case that coordinates multiple operations?
→ internal/app

Is it an interface required by the application?
→ internal/ports

Is it HTTP, gRPC, database, Redis, Kafka, RabbitMQ, MongoDB or ClickHouse code?
→ internal/adapters

Is it configuration, logging, tracing or shutdown logic?
→ internal/platform

Is it only process startup and dependency wiring?
→ cmd
```

Forbidden shortcuts:

```text
Do not put business logic in main.go.
Do not put business logic in HTTP handlers.
Do not put database-specific code in domain packages.
Do not put Kafka/RabbitMQ clients in application use cases directly.
Do not put ClickHouse clients in application use cases directly.
Do not create generic utility packages without clear ownership.
Do not create a central model package for all domain entities.
Do not duplicate domain rules in workers.
```

---

## Learning Integration Rule

DevArena is educational, but learning topics must be implemented as real project code.

Each learning topic should result in one of the following:

```text
A real domain capability
A useful application use case
A safer interface or port
A better adapter implementation
A meaningful test
A benchmark for a performance-sensitive path
A documented architectural decision
A production-readiness improvement
```

A topic should not be marked complete only because it was demonstrated in isolated terminal output.

Progress tracking belongs in:

```text
LEARNING_CHECKLIST.md
```

Architecture decisions belong in:

```text
ARCHITECTURE.md
```

Public project positioning belongs in:

```text
README.md
```

---

## Core Production Scenario

The final production-like scenario:

```text
1. User registers.
2. User logs in.
3. User creates a hero.
4. Hero starts a battle.
5. API validates JWT.
6. Middleware adds request ID and correlation ID.
7. Battle use case checks cooldown through Redis port.
8. Battle use case acquires distributed lock through lock port.
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
19. Analytics worker stores analytical data in ClickHouse.
20. Audit log is recorded.
21. API returns battle result as JSON.
22. Logs, metrics and traces are emitted.
```

---

## Final Architecture Goal

At the end of the project, DevArena should demonstrate:

```text
strong Go fundamentals
domain-first design
clear application use cases
ports and adapters
REST API design
gRPC internal communication
PostgreSQL persistence
Redis caching and locking
MongoDB document modeling
ClickHouse analytics and OLAP queries
RabbitMQ background jobs
Kafka event streaming
outbox pattern
Dockerized local development
Kubernetes deployment
CI/CD automation
structured testing
security basics
observability
performance awareness
production-readiness thinking
```

The final architecture should be explainable as:

```text
DevArena is a production-oriented educational backend platform written in Go.

It starts from a text-based arena game domain and evolves into a modular backend system.
The core business logic is kept in domain packages, application use cases coordinate behavior
through ports, and infrastructure integrations are implemented as adapters.

The system demonstrates REST API, PostgreSQL, Redis, MongoDB, ClickHouse, RabbitMQ, Kafka,
gRPC, Docker, Kubernetes, CI/CD, testing, observability, security, analytics and
production-readiness practices.
```