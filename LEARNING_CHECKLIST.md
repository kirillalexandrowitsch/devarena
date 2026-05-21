# DevArena Learning Checklist

This file tracks all topics that must be learned during the development of DevArena.

The rule is simple:

```text
Every topic must be explained.
Every topic must be implemented in the project.
Every topic must be committed.
Every topic must be checked here.
```

The goal is not to finish quickly.

The goal is to build a serious backend project and understand every layer.

---

# Progress Levels

The project is designed to move through these levels:

```text
Beginner
→ Confident Intern
→ Junior Backend Developer
→ Strong Junior / Junior+
→ Middle-oriented Backend Developer
```

---

# Stage 0. Repository Initialization

- [x] Create GitHub repository
- [x] Clone repository locally
- [x] Initialize Go module
- [x] Create base project structure
- [x] Add README.md
- [x] Add ARCHITECTURE.md
- [x] Add LEARNING_CHECKLIST.md
- [x] Add .gitignore
- [x] Add first runnable Go program

---

# Stage 1. Go Core

## Basic Go Syntax

- [x] Variables
- [x] Basic data types
- [x] Functions
- [x] Conditions
- [x] Loops
- [x] Constants
- [x] Arrays
- [x] Slices
- [ ] Maps
- [ ] Structs
- [ ] Methods
- [ ] Pointers
- [ ] Interfaces
- [ ] Packages
- [ ] Modules

## Go Types

- [x] string
- [x] int
- [ ] int64
- [x] bool
- [x] float64
- [ ] rune
- [ ] byte
- [ ] zero values
- [ ] type conversion
- [ ] custom types
- [ ] type aliases

## Functions

- [x] Basic functions
- [x] Function parameters
- [x] Return values
- [ ] Multiple return values
- [ ] Named return values
- [ ] Variadic functions
- [ ] Anonymous functions
- [ ] Closures

## Control Flow

- [x] if / else
- [ ] switch
- [x] for loops
- [ ] break
- [ ] continue
- [ ] range

## Collections

- [x] Arrays
- [x] Slices
- [x] len
- [x] cap
- [x] append
- [x] copy
- [ ] Maps
- [ ] Map lookup
- [ ] Map delete
- [ ] Range over slice
- [ ] Range over map

## Structs and Methods

- [ ] Structs
- [ ] Embedded structs
- [ ] Methods
- [ ] Value receiver
- [ ] Pointer receiver
- [ ] Method sets

## Pointers

- [ ] Basic pointers
- [ ] Address operator
- [ ] Dereferencing
- [ ] nil
- [ ] Pointer receiver
- [ ] Mutating state through pointer

## Interfaces

- [ ] Interfaces
- [ ] Implicit implementation
- [ ] Empty interface / any
- [ ] Type assertion
- [ ] Type switch
- [ ] Interface composition

## Packages and Modules

- [ ] Packages
- [ ] Imports
- [ ] Exported names
- [ ] Unexported names
- [ ] Go modules
- [ ] go.mod
- [ ] go.sum
- [ ] go mod tidy
- [ ] init function

## Additional Core Go

- [ ] defer
- [ ] panic
- [ ] recover
- [ ] generics
- [ ] type parameters
- [ ] constraints
- [ ] functional options pattern
- [ ] dependency injection without framework

## Go Internals Basics

- [ ] Slice internals
- [ ] Map internals basics
- [ ] Interface internals basics
- [ ] Stack vs heap basics
- [ ] Escape analysis basics
- [ ] Garbage collector basics
- [ ] Struct alignment basics
- [ ] Memory allocation basics

---

# Stage 2. Go Error Handling

## Basic Error Handling

- [ ] error interface
- [ ] errors.New
- [ ] fmt.Errorf
- [ ] Returning errors
- [ ] Checking errors

## Error Wrapping

- [ ] Error wrapping with %w
- [ ] errors.Is
- [ ] errors.As
- [ ] Error chains

## Error Types

- [ ] Sentinel errors
- [ ] Custom error types
- [ ] Validation errors
- [ ] Domain errors
- [ ] Infrastructure errors
- [ ] Authentication errors
- [ ] Authorization errors
- [ ] Conflict errors
- [ ] Not found errors

## Backend Error Mapping

- [ ] Mapping errors to HTTP status codes
- [ ] Structured error responses
- [ ] Retryable errors
- [ ] Non-retryable errors
- [ ] Error boundaries
- [ ] Error classification

---

# Stage 3. Go Concurrency

## Goroutines

- [ ] Goroutines
- [ ] Starting goroutines
- [ ] Goroutine lifecycle
- [ ] Goroutine leaks

## Channels

- [ ] Channels
- [ ] Unbuffered channels
- [ ] Buffered channels
- [ ] Sending to channel
- [ ] Receiving from channel
- [ ] Closing channels
- [ ] Range over channel
- [ ] Channel ownership

## Select

- [ ] select
- [ ] default case in select
- [ ] time.After
- [ ] time.Ticker
- [ ] Timeout with select

## Sync Package

- [ ] sync.WaitGroup
- [ ] sync.Mutex
- [ ] sync.RWMutex
- [ ] sync.Once
- [ ] sync.Map basics
- [ ] atomic basics

## Race Conditions

- [ ] Race conditions
- [ ] go test -race
- [ ] Data race detection
- [ ] Fixing race conditions with Mutex
- [ ] Mutex contention basics

## Context

- [ ] context.Context
- [ ] context cancellation
- [ ] context timeout
- [ ] context deadline
- [ ] Cancellation propagation
- [ ] Context in HTTP handlers
- [ ] Context in repository methods

## Concurrency Patterns

- [ ] Worker pool
- [ ] Fan-out
- [ ] Fan-in
- [ ] Pipeline pattern
- [ ] Bounded concurrency
- [ ] Semaphore pattern
- [ ] errgroup
- [ ] Backpressure
- [ ] Graceful shutdown
- [ ] Rate-limited workers
- [ ] Retry workers

---

# Stage 4. Git and GitHub

## Git Basics

- [x] git clone
- [x] git status
- [x] git add
- [x] git commit
- [x] git push
- [ ] git pull
- [ ] git branch
- [ ] git checkout
- [ ] git switch
- [ ] git merge
- [ ] git rebase basics
- [x] git log
- [x] .gitignore

## GitHub Workflow

- [ ] Pull request
- [ ] Code review
- [ ] Merge conflicts
- [ ] Conventional commits
- [ ] Branch naming
- [ ] PR description
- [ ] Issues
- [ ] Milestones
- [ ] Project board

## Advanced Git

- [ ] Interactive rebase
- [ ] Squash commits
- [ ] cherry-pick
- [ ] revert
- [ ] reset
- [ ] reflog
- [ ] git tags
- [ ] release branches
- [ ] protected branches
- [ ] trunk-based development concept
- [ ] GitFlow concept

---

# Stage 5. Linux and Terminal

## Terminal Basics

- [x] pwd
- [x] cd
- [ ] ls
- [x] mkdir
- [x] touch
- [ ] rm
- [ ] cp
- [ ] mv
- [ ] cat
- [ ] less
- [ ] grep
- [x] find
- [ ] chmod

## Environment and Processes

- [ ] env
- [ ] export
- [ ] ps
- [ ] kill
- [ ] process signals
- [ ] background processes
- [ ] logs inspection

## Networking Tools

- [ ] curl
- [ ] netstat / ss
- [ ] lsof
- [ ] DNS basics with terminal tools
- [ ] /etc/hosts

## Advanced Terminal

- [ ] Pipes
- [ ] Redirects
- [ ] Bash scripts
- [ ] Cron basics
- [ ] systemd basics
- [ ] top / htop
- [ ] df / du
- [ ] tail
- [ ] head

---

# Stage 6. HTTP and REST API

## HTTP Basics

- [ ] HTTP
- [ ] Request
- [ ] Response
- [ ] Headers
- [ ] Body
- [ ] Status codes
- [ ] Content-Type
- [ ] Authorization header

## HTTP Methods

- [ ] GET
- [ ] POST
- [ ] PUT
- [ ] PATCH
- [ ] DELETE

## Request Data

- [ ] Query parameters
- [ ] Path parameters
- [ ] JSON body
- [ ] Request validation
- [ ] Request size limits

## Routing and Middleware

- [ ] Routing
- [ ] net/http
- [ ] chi router
- [ ] Middleware
- [ ] Logger middleware
- [ ] Recovery middleware
- [ ] Auth middleware
- [ ] Request ID middleware
- [ ] Rate limiter middleware
- [ ] Timeout middleware
- [ ] CORS basics

## API Responses

- [ ] JSON responses
- [ ] Unified response format
- [ ] Error response format
- [ ] Mapping errors to status codes

## API Features

- [ ] Pagination
- [ ] Filtering
- [ ] Sorting
- [ ] Idempotency basics
- [ ] Idempotency keys
- [ ] API versioning
- [ ] Backward compatibility
- [ ] OpenAPI / Swagger
- [ ] Graceful shutdown HTTP server

---

# Stage 7. API Design

- [ ] Resource naming
- [ ] Plural nouns
- [ ] HTTP method semantics
- [ ] Consistent status codes
- [ ] Consistent error format
- [ ] Consistent validation format
- [ ] Pagination design
- [ ] Filtering design
- [ ] Sorting design
- [ ] Partial updates
- [ ] Public API vs internal API
- [ ] API contracts
- [ ] Breaking changes
- [ ] Deprecation strategy
- [ ] Backward compatibility

---

# Stage 8. Backend Application Architecture

## Project Structure

- [x] cmd
- [ ] internal
- [ ] config
- [ ] handler
- [ ] service
- [ ] repository
- [ ] model
- [ ] middleware
- [ ] app
- [ ] response
- [ ] validator
- [ ] infrastructure
- [ ] migrations
- [ ] docker-compose.yml
- [ ] .env.example
- [x] README.md
- [x] ARCHITECTURE.md

## Architecture Concepts

- [ ] Handler layer
- [ ] Service layer
- [ ] Repository layer
- [ ] Model/domain layer
- [ ] DTO
- [ ] Request models
- [ ] Response models
- [ ] Dependency injection
- [ ] Interface-based design
- [ ] Clean architecture basics
- [ ] Hexagonal architecture basics
- [ ] Modular monolith
- [ ] Monolith vs microservices
- [ ] Service boundaries
- [ ] Bounded context
- [ ] Domain events

## Distributed Architecture Concepts

- [ ] Event-driven architecture
- [ ] Eventual consistency
- [ ] Outbox pattern
- [ ] Transactional outbox
- [ ] Saga pattern concept
- [ ] CQRS concept
- [ ] Idempotency
- [ ] Retries
- [ ] Circuit breaker concept
- [ ] Bulkhead pattern concept

---

# Stage 9. Data Modeling

## Relational Modeling

- [ ] Entities
- [ ] Relationships
- [ ] One-to-one
- [ ] One-to-many
- [ ] Many-to-many
- [ ] Primary keys
- [ ] Foreign keys
- [ ] Normalization basics
- [ ] Denormalization tradeoffs
- [ ] Soft delete
- [ ] Hard delete
- [ ] Audit tables

## Schema Evolution

- [ ] Migrations
- [ ] Data migrations
- [ ] Safe migrations
- [ ] Schema evolution
- [ ] Backward-compatible schema changes
- [ ] Data retention

## Consistency

- [ ] Optimistic locking
- [ ] Pessimistic locking
- [ ] Event history
- [ ] Read models
- [ ] Write models

---

# Stage 10. PostgreSQL and SQL

## Basic SQL

- [ ] SELECT
- [ ] INSERT
- [ ] UPDATE
- [ ] DELETE
- [ ] WHERE
- [ ] JOIN
- [ ] LEFT JOIN
- [ ] GROUP BY
- [ ] HAVING
- [ ] ORDER BY
- [ ] LIMIT
- [ ] OFFSET

## Aggregate Functions

- [ ] COUNT
- [ ] SUM
- [ ] AVG
- [ ] MIN
- [ ] MAX

## Constraints

- [ ] Primary key
- [ ] Foreign key
- [ ] Unique constraint
- [ ] Not null constraint
- [ ] Check constraint

## Indexes

- [ ] Basic indexes
- [ ] Unique indexes
- [ ] Composite indexes
- [ ] Partial indexes
- [ ] Covering indexes

## Transactions

- [ ] Transactions
- [ ] BEGIN
- [ ] COMMIT
- [ ] ROLLBACK
- [ ] Transaction isolation levels
- [ ] Read committed
- [ ] Repeatable read
- [ ] Serializable
- [ ] Deadlocks
- [ ] Locks
- [ ] Row-level locking
- [ ] SELECT FOR UPDATE
- [ ] Optimistic locking

## PostgreSQL with Go

- [ ] pgx
- [ ] Connection pool
- [ ] Context with DB queries
- [ ] Repository pattern
- [ ] Migrations
- [ ] Repository tests

## Advanced PostgreSQL

- [ ] EXPLAIN
- [ ] EXPLAIN ANALYZE
- [ ] Query plans
- [ ] Index scan
- [ ] Sequential scan
- [ ] Upsert / ON CONFLICT
- [ ] CTE
- [ ] Window functions
- [ ] JSONB
- [ ] Full-text search basics
- [ ] Partitioning basics
- [ ] N+1 query problem
- [ ] Slow query analysis

---

# Stage 11. Authentication and Security

## Authentication

- [ ] Register
- [ ] Login
- [ ] Password hashing
- [ ] bcrypt
- [ ] JWT
- [ ] Access token
- [ ] Refresh token
- [ ] Refresh token storage
- [ ] Refresh token revocation
- [ ] Logout

## Authorization

- [ ] Authentication vs authorization
- [ ] Roles
- [ ] Permissions
- [ ] Admin routes
- [ ] User-owned resources
- [ ] 401 Unauthorized
- [ ] 403 Forbidden

## Security Basics

- [ ] Input validation
- [ ] SQL injection prevention
- [ ] Secrets in environment variables
- [ ] HTTPS basics
- [ ] Secure headers
- [ ] Rate limiting for security
- [ ] Brute force protection
- [ ] Audit logs
- [ ] Least privilege

## OWASP Basics

- [ ] OWASP Top 10 overview
- [ ] XSS concept
- [ ] CSRF concept
- [ ] SSRF concept
- [ ] Dependency vulnerability scanning

---

# Stage 12. Redis

## Redis Basics

- [ ] GET
- [ ] SET
- [ ] DEL
- [ ] EXPIRE
- [ ] TTL
- [ ] INCR
- [ ] HGET
- [ ] HSET
- [ ] Lists basics
- [ ] Sets basics
- [ ] Sorted sets basics

## Redis Use Cases

- [ ] Caching
- [ ] Cache TTL
- [ ] Cache invalidation
- [ ] Hero profile cache
- [ ] Leaderboard cache
- [ ] Rate limiting
- [ ] Session storage
- [ ] Distributed lock
- [ ] Pub/Sub

## Advanced Redis

- [ ] Cache-aside pattern
- [ ] Write-through cache
- [ ] Write-behind basics
- [ ] Cache stampede
- [ ] Cache penetration
- [ ] Cache avalanche
- [ ] Distributed locks with expiration
- [ ] Redlock concept
- [ ] Lua scripts basics
- [ ] Sorted set leaderboard
- [ ] Redis streams basics
- [ ] Eviction policies
- [ ] RDB/AOF persistence basics
- [ ] Redis Cluster basics
- [ ] Redis Sentinel basics

---

# Stage 13. Testing

## Go Testing Basics

- [ ] testing package
- [ ] Unit tests
- [ ] Table-driven tests
- [ ] Subtests
- [ ] t.Run
- [ ] Test helpers
- [ ] Test fixtures

## Backend Tests

- [ ] Service tests
- [ ] Handler tests
- [ ] HTTP handler tests with httptest
- [ ] Repository tests
- [ ] Integration tests
- [ ] API integration tests

## Test Doubles

- [ ] Mocks
- [ ] Fakes
- [ ] Stubs
- [ ] Mocking repository interfaces
- [ ] Mocking external services

## Test Infrastructure

- [ ] Testcontainers
- [ ] PostgreSQL test container
- [ ] Redis test container
- [ ] RabbitMQ test container
- [ ] Kafka test container
- [ ] Database cleanup strategies
- [ ] Transactional tests

## Advanced Testing

- [ ] go test ./...
- [ ] go test -race
- [ ] Coverage
- [ ] Golden files
- [ ] Benchmark tests
- [ ] Fuzz testing
- [ ] Load testing basics
- [ ] Deterministic tests
- [ ] Avoiding flaky tests
- [ ] Testing async code
- [ ] Testing Kafka consumers
- [ ] Testing RabbitMQ consumers

---

# Stage 14. Docker

## Docker Basics

- [ ] Image
- [ ] Container
- [ ] Dockerfile
- [ ] docker build
- [ ] docker run
- [ ] docker exec
- [ ] docker logs
- [ ] docker ps
- [ ] Docker Compose
- [ ] Ports
- [ ] Volumes
- [ ] Networks
- [ ] Environment variables
- [ ] Multi-stage build

## Docker for Production-like Setup

- [ ] Image layers
- [ ] Build cache
- [ ] .dockerignore
- [ ] Minimal images
- [ ] Distroless concept
- [ ] Non-root user
- [ ] Healthcheck
- [ ] Container resource limits
- [ ] Docker networking
- [ ] Service dependencies
- [ ] Reproducible builds
- [ ] Security scanning basics

---

# Stage 15. RabbitMQ

## RabbitMQ Basics

- [ ] Queue
- [ ] Exchange
- [ ] Routing key
- [ ] Producer
- [ ] Consumer
- [ ] Ack
- [ ] Nack
- [ ] Prefetch
- [ ] Dead letter queue
- [ ] Retry
- [ ] Direct exchange
- [ ] Fanout exchange
- [ ] Topic exchange

## RabbitMQ Reliability

- [ ] Message durability
- [ ] Persistent messages
- [ ] Manual ack
- [ ] Requeue behavior
- [ ] Poison messages
- [ ] Delayed retry pattern
- [ ] Exponential backoff retry
- [ ] Idempotent consumers
- [ ] Competing consumers
- [ ] Publisher confirms
- [ ] Mandatory messages
- [ ] Alternate exchange
- [ ] Connection recovery
- [ ] Channel lifecycle
- [ ] Backpressure basics

---

# Stage 16. Kafka

## Kafka Basics

- [ ] Broker
- [ ] Topic
- [ ] Partition
- [ ] Offset
- [ ] Producer
- [ ] Consumer
- [ ] Consumer group
- [ ] Event log
- [ ] Key-based partitioning
- [ ] Ordering inside partition
- [ ] At-least-once delivery
- [ ] Basic retries

## Kafka Reliability

- [ ] Replication factor
- [ ] Leader / follower replicas
- [ ] ISR
- [ ] Acknowledgements
- [ ] Producer acks
- [ ] Consumer offset commit
- [ ] Manual commit
- [ ] Auto commit
- [ ] Idempotent producer
- [ ] Exactly-once semantics concept
- [ ] Idempotent consumer
- [ ] Consumer lag
- [ ] Rebalancing
- [ ] Dead letter topic
- [ ] Compacted topics

## Kafka Architecture

- [ ] Schema registry concept
- [ ] Avro concept
- [ ] Protobuf with Kafka concept
- [ ] Event versioning
- [ ] Outbox pattern
- [ ] Transactional outbox
- [ ] Event-driven architecture

---

# Stage 17. MongoDB

## MongoDB Basics

- [ ] Document
- [ ] Collection
- [ ] BSON
- [ ] ObjectID
- [ ] _id
- [ ] CRUD
- [ ] Insert
- [ ] Find
- [ ] Update
- [ ] Delete
- [ ] Indexes
- [ ] Embedded documents
- [ ] References
- [ ] Schema design
- [ ] Aggregation pipeline basics

## MongoDB Advanced

- [ ] Compound indexes
- [ ] Text indexes
- [ ] TTL indexes
- [ ] $match
- [ ] $group
- [ ] $sort
- [ ] $project
- [ ] $lookup
- [ ] Schema validation
- [ ] Write concern basics
- [ ] Read concern basics
- [ ] Replica set basics
- [ ] Sharding basics
- [ ] Document growth problem
- [ ] Embedding vs referencing tradeoffs

---

# Stage 18. gRPC

## gRPC Basics

- [ ] Protocol Buffers
- [ ] .proto files
- [ ] message
- [ ] service
- [ ] rpc
- [ ] Unary RPC
- [ ] Generated Go code
- [ ] gRPC server
- [ ] gRPC client
- [ ] Status codes
- [ ] Metadata
- [ ] Deadlines
- [ ] Interceptors

## gRPC Advanced

- [ ] Server streaming
- [ ] Client streaming
- [ ] Bidirectional streaming
- [ ] Protobuf versioning
- [ ] Backward compatibility in proto
- [ ] Field numbering rules
- [ ] Optional fields
- [ ] Well-known types
- [ ] Error details
- [ ] Retry policies
- [ ] Load balancing basics
- [ ] Health checking
- [ ] Reflection
- [ ] grpcurl
- [ ] mTLS basics
- [ ] REST to gRPC gateway concept

---

# Stage 19. Observability

## Logging

- [ ] Structured logging
- [ ] Log levels
- [ ] Request ID
- [ ] Correlation ID
- [ ] Error logging
- [ ] Access logs

## Metrics

- [ ] Metrics basics
- [ ] Prometheus basics
- [ ] Grafana basics
- [ ] Counters
- [ ] Gauges
- [ ] Histograms
- [ ] RED metrics
- [ ] USE metrics

## Tracing

- [ ] Tracing basics
- [ ] OpenTelemetry basics
- [ ] Spans
- [ ] Trace ID
- [ ] Context propagation

## Health and Alerts

- [ ] Health checks
- [ ] Readiness
- [ ] Liveness
- [ ] Alerting basics
- [ ] Dashboards
- [ ] SLO / SLA basics
- [ ] Log aggregation basics

---

# Stage 20. Networking Basics

- [ ] TCP/IP basics
- [ ] DNS
- [ ] HTTP
- [ ] HTTPS
- [ ] TLS basics
- [ ] Ports
- [ ] localhost
- [ ] IP address
- [ ] Request/response lifecycle
- [ ] Latency
- [ ] Timeout
- [ ] Retry
- [ ] TCP connection lifecycle basics
- [ ] Keep-alive
- [ ] Connection pooling
- [ ] Load balancing basics
- [ ] Reverse proxy
- [ ] Gateway
- [ ] Service discovery
- [ ] DNS caching
- [ ] TLS certificates
- [ ] mTLS concept
- [ ] WebSocket basics
- [ ] HTTP/2 basics

---

# Stage 21. Kubernetes

## Kubernetes Basics

- [ ] Cluster
- [ ] Node
- [ ] Pod
- [ ] Deployment
- [ ] ReplicaSet
- [ ] Service
- [ ] Ingress
- [ ] Namespace
- [ ] ConfigMap
- [ ] Secret
- [ ] Readiness probe
- [ ] Liveness probe
- [ ] Resource requests
- [ ] Resource limits
- [ ] kubectl basics
- [ ] kubectl logs
- [ ] kubectl describe
- [ ] kubectl exec

## Kubernetes Advanced

- [ ] Rolling updates
- [ ] Rollback
- [ ] Horizontal Pod Autoscaler
- [ ] PersistentVolume
- [ ] PersistentVolumeClaim
- [ ] StorageClass basics
- [ ] Jobs
- [ ] CronJobs
- [ ] ServiceAccount
- [ ] RBAC basics
- [ ] NetworkPolicy basics
- [ ] Init containers
- [ ] Sidecar containers
- [ ] Helm basics
- [ ] Kustomize basics
- [ ] PodDisruptionBudget
- [ ] Affinity / anti-affinity basics
- [ ] Node selectors
- [ ] Config rollout strategy
- [ ] Secrets management basics

---

# Stage 22. CI/CD

## CI/CD Basics

- [ ] Pipeline
- [ ] Build
- [ ] Test
- [ ] Lint
- [ ] Docker build
- [ ] Docker push
- [ ] Deploy
- [ ] GitHub Actions
- [ ] Workflow
- [ ] Job
- [ ] Step
- [ ] Secrets
- [ ] Artifacts
- [ ] Environments

## CI/CD Advanced

- [ ] Branch protection
- [ ] Pull request checks
- [ ] Required status checks
- [ ] Semantic versioning
- [ ] Release tags
- [ ] Changelog basics
- [ ] Deployment strategies
- [ ] Rolling deployment
- [ ] Blue-green deployment concept
- [ ] Canary deployment concept
- [ ] Rollback
- [ ] Caching dependencies
- [ ] Build matrix
- [ ] Reusable workflows
- [ ] Environment approvals
- [ ] Supply chain security basics
- [ ] SBOM concept

---

# Stage 23. Background Jobs

## Basic Background Jobs

- [ ] Worker
- [ ] Job
- [ ] Queue
- [ ] Retry
- [ ] Dead letter queue
- [ ] Scheduled job
- [ ] Idempotent processing

## Advanced Background Jobs

- [ ] Delayed jobs
- [ ] Exponential backoff
- [ ] Poison message handling
- [ ] Job deduplication
- [ ] Job status tracking
- [ ] Distributed workers
- [ ] Concurrency limits
- [ ] Graceful worker shutdown
- [ ] Backpressure

---

# Stage 24. Performance Basics

## Basic Performance

- [ ] Avoid unnecessary DB queries
- [ ] Pagination
- [ ] Indexes
- [ ] Caching
- [ ] Connection pool
- [ ] Timeouts
- [ ] Basic profiling awareness

## Advanced Performance

- [ ] pprof
- [ ] CPU profiling
- [ ] Memory profiling
- [ ] Allocation analysis
- [ ] Benchmarking
- [ ] N+1 problem
- [ ] Query optimization
- [ ] Cache hit ratio
- [ ] Latency percentiles
- [ ] Throughput
- [ ] Backpressure
- [ ] Load testing
- [ ] Horizontal scaling

---

# Stage 25. Production Readiness

## Basic Production Readiness

- [ ] Config via env
- [ ] Health checks
- [ ] Graceful shutdown
- [ ] Logging
- [ ] Error handling
- [ ] Migrations
- [ ] Docker Compose
- [ ] CI checks

## Advanced Production Readiness

- [ ] Zero-downtime deployment basics
- [ ] Rollback strategy
- [ ] Database migration safety
- [ ] Feature flags basics
- [ ] Readiness before traffic
- [ ] Alerting
- [ ] Dashboards
- [ ] Runbooks
- [ ] Incident analysis
- [ ] Postmortem basics