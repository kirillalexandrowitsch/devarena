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
- [x] Maps
- [x] Structs
- [x] Methods
- [x] Pointers
- [x] Interfaces
- [x] Packages
- [x] Modules

## Go Types

- [x] string
- [x] int
- [x] int64
- [x] bool
- [x] float64
- [x] rune
- [x] byte
- [x] zero values
- [x] type conversion
- [x] custom types
- [x] type aliases

## Functions

- [x] Basic functions
- [x] Function parameters
- [x] Return values
- [x] Multiple return values
- [x] Named return values
- [x] Variadic functions
- [x] Anonymous functions
- [x] Closures

## Control Flow

- [x] if / else
- [x] switch
- [x] for loops
- [x] break
- [x] continue
- [x] range

## Collections

- [x] Arrays
- [x] Slices
- [x] len
- [x] cap
- [x] append
- [x] copy
- [x] Maps
- [x] Map lookup
- [x] Map delete
- [x] Range over slice
- [x] Range over map

## Structs and Methods

- [x] Structs
- [x] Embedded structs
- [x] Methods
- [x] Value receiver
- [x] Pointer receiver
- [x] Method sets

## Pointers

- [x] Basic pointers
- [x] Address operator
- [x] Dereferencing
- [x] nil
- [x] Pointer receiver
- [x] Mutating state through pointer

## Interfaces

- [x] Interfaces
- [x] Implicit implementation
- [x] Empty interface / any
- [x] Type assertion
- [x] Type switch
- [x] Interface composition

## Packages and Modules

- [x] Packages
- [x] Imports
- [x] Exported names
- [x] Unexported names
- [x] Go modules
- [x] go.mod
- [ ] go.sum
- [x] go mod tidy
- [x] init function

## Additional Core Go

- [x] defer
- [x] panic
- [x] recover
- [x] generics
- [x] type parameters
- [x] constraints
- [x] functional options pattern
- [x] dependency injection without framework

## Go Internals Basics

- [x] Slice internals
- [x] Map internals basics
- [x] Interface internals basics
- [x] Stack vs heap basics
- [x] Escape analysis basics
- [x] Garbage collector basics
- [x] Struct alignment basics
- [x] Memory allocation basics

---

# Stage 2. Go Error Handling

## Basic Error Handling

- [x] error interface
- [x] errors.New
- [x] fmt.Errorf
- [x] Returning errors
- [x] Checking errors

## Error Wrapping

- [x] Error wrapping with %w
- [x] errors.Is
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
- [x] internal
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

## Domain-First Target Architecture

- [ ] Domain-first design
- [ ] Application use cases
- [ ] Ports and adapters
- [ ] Infrastructure behind interfaces
- [ ] Domain package boundaries
- [ ] Application package boundaries
- [ ] Adapter package boundaries
- [ ] Platform package boundaries
- [ ] Dependency direction enforcement
- [ ] Avoiding business logic in cmd
- [ ] Avoiding business logic in handlers
- [ ] Avoiding infrastructure imports in domain
- [ ] Avoiding concrete database clients in app layer
- [ ] Avoiding concrete analytics storage clients in app layer

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

## Analytical Modeling

- [ ] OLTP vs OLAP
- [ ] Analytical events
- [ ] Fact tables
- [ ] Dimension tables basics
- [ ] Aggregated tables
- [ ] Time-series event modeling
- [ ] Append-only analytics data
- [ ] Duplicate-tolerant analytics design
- [ ] Analytics data retention
- [ ] Analytics backfill strategy

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
- [ ] MongoDB test container
- [ ] ClickHouse test container
- [ ] RabbitMQ test container
- [ ] Kafka test container
- [ ] Database cleanup strategies
- [ ] Transactional tests
- [ ] Analytics test data cleanup
- [ ] Idempotent analytics test design

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
- [ ] Testing ClickHouse inserts
- [ ] Testing ClickHouse analytical queries

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

## Docker for Local Infrastructure

- [ ] Local PostgreSQL container
- [ ] Local Redis container
- [ ] Local MongoDB container
- [ ] Local ClickHouse container
- [ ] Local RabbitMQ container
- [ ] Local Kafka container
- [ ] Docker Compose service dependencies
- [ ] Docker Compose healthchecks
- [ ] Docker Compose volumes
- [ ] Docker Compose networks

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

# Stage 18. ClickHouse

## ClickHouse Basics

- [ ] What ClickHouse is
- [ ] OLTP vs OLAP
- [ ] Column-oriented storage
- [ ] Analytical workloads
- [ ] Append-oriented data model
- [ ] ClickHouse server basics
- [ ] ClickHouse client basics
- [ ] ClickHouse SQL basics
- [ ] Databases
- [ ] Tables
- [ ] Columns
- [ ] Data types
- [ ] Nullable columns
- [ ] LowCardinality
- [ ] DateTime
- [ ] DateTime64
- [ ] UUID
- [ ] Array basics
- [ ] Map basics

## ClickHouse Table Engines

- [ ] MergeTree
- [ ] ReplacingMergeTree
- [ ] SummingMergeTree
- [ ] AggregatingMergeTree
- [ ] Distributed table concept
- [ ] Memory engine basics
- [ ] Choosing table engines
- [ ] Engine tradeoffs

## ClickHouse Schema Design

- [ ] ORDER BY key
- [ ] PARTITION BY
- [ ] PRIMARY KEY in ClickHouse
- [ ] Sorting key design
- [ ] Partitioning strategy
- [ ] Granules
- [ ] Sparse indexes
- [ ] Data skipping indexes
- [ ] Time-based partitioning
- [ ] Event table design
- [ ] Battle analytics table design
- [ ] Reward analytics table design
- [ ] Rating snapshot table design
- [ ] User activity analytics table design
- [ ] Denormalized analytics schema
- [ ] Fact table concept
- [ ] Dimension data handling
- [ ] Duplicate-tolerant schema design

## ClickHouse Queries

- [ ] SELECT
- [ ] INSERT
- [ ] WHERE
- [ ] GROUP BY
- [ ] ORDER BY
- [ ] LIMIT
- [ ] COUNT
- [ ] SUM
- [ ] AVG
- [ ] MIN
- [ ] MAX
- [ ] uniq
- [ ] uniqExact
- [ ] quantile
- [ ] quantiles
- [ ] argMax
- [ ] argMin
- [ ] toDate
- [ ] toStartOfHour
- [ ] toStartOfDay
- [ ] dateDiff
- [ ] Window functions basics
- [ ] WITH clause basics
- [ ] Subqueries basics

## ClickHouse Aggregations

- [ ] Aggregation functions
- [ ] AggregateFunction type
- [ ] SimpleAggregateFunction type
- [ ] AggregatingMergeTree basics
- [ ] SummingMergeTree basics
- [ ] Materialized views
- [ ] Incremental aggregation
- [ ] Pre-aggregated tables
- [ ] Rollup tables
- [ ] Daily aggregation tables
- [ ] Leaderboard analytics aggregation
- [ ] Battle result aggregation
- [ ] Reward distribution aggregation

## ClickHouse Performance

- [ ] Query performance basics
- [ ] Avoiding SELECT *
- [ ] Filtering by sorting key
- [ ] Partition pruning
- [ ] Compression basics
- [ ] Batch inserts
- [ ] Insert block size
- [ ] Avoiding many small inserts
- [ ] Query profiling basics
- [ ] system.query_log
- [ ] system.parts
- [ ] system.mutations
- [ ] EXPLAIN basics
- [ ] Memory usage basics
- [ ] Cardinality awareness

## ClickHouse with Go

- [ ] ClickHouse Go driver
- [ ] Opening ClickHouse connection
- [ ] Connection settings
- [ ] Context with ClickHouse queries
- [ ] Batch insert from Go
- [ ] Querying analytical data from Go
- [ ] Scanning query results
- [ ] Mapping ClickHouse rows to Go structs
- [ ] ClickHouse adapter package
- [ ] Analytics repository pattern
- [ ] Analytics ports
- [ ] Fake analytics ports for tests
- [ ] ClickHouse integration tests
- [ ] ClickHouse migrations from Go
- [ ] Error handling for ClickHouse queries
- [ ] Retry strategy for analytics writes
- [ ] Idempotent analytics writes

## ClickHouse Architecture in DevArena

- [ ] ClickHouse as analytics storage
- [ ] ClickHouse is not source of truth
- [ ] PostgreSQL as transactional source of truth
- [ ] Event-driven analytics ingestion
- [ ] Analytics worker
- [ ] Kafka to ClickHouse pipeline
- [ ] Outbox to analytics pipeline
- [ ] Battle analytics events
- [ ] Reward analytics events
- [ ] Rating analytics events
- [ ] Audit analytics events
- [ ] User activity analytics events
- [ ] Analytics read use cases
- [ ] Analytics API endpoints
- [ ] ClickHouse migrations directory
- [ ] ClickHouse Docker Compose service
- [ ] ClickHouse configuration through environment variables
- [ ] Observability for analytics ingestion
- [ ] Backfill analytics data
- [ ] Handling duplicate analytics events

---

# Stage 19. gRPC

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

# Stage 20. Observability

## Logging

- [ ] Structured logging
- [ ] Log levels
- [ ] Request ID
- [ ] Correlation ID
- [ ] Error logging
- [ ] Access logs
- [ ] Worker logs
- [ ] Analytics ingestion logs

## Metrics

- [ ] Metrics basics
- [ ] Prometheus basics
- [ ] Grafana basics
- [ ] Counters
- [ ] Gauges
- [ ] Histograms
- [ ] RED metrics
- [ ] USE metrics
- [ ] Business metrics
- [ ] Analytics ingestion metrics
- [ ] ClickHouse query metrics
- [ ] Worker lag metrics

## Tracing

- [ ] Tracing basics
- [ ] OpenTelemetry basics
- [ ] Spans
- [ ] Trace ID
- [ ] Context propagation
- [ ] HTTP tracing
- [ ] gRPC tracing
- [ ] Worker tracing
- [ ] Database tracing
- [ ] Analytics pipeline tracing

## Health and Alerts

- [ ] Health checks
- [ ] Readiness
- [ ] Liveness
- [ ] Alerting basics
- [ ] Dashboards
- [ ] SLO / SLA basics
- [ ] Log aggregation basics
- [ ] Analytics ingestion lag alerts
- [ ] Database availability alerts
- [ ] Worker failure alerts

---

# Stage 21. Networking Basics

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

# Stage 22. Kubernetes

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
- [ ] StatefulSet basics
- [ ] Stateful services in Kubernetes
- [ ] Database deployment tradeoffs

---

# Stage 23. CI/CD

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

# Stage 24. Background Jobs

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

## Analytics Background Jobs

- [ ] Analytics worker
- [ ] Event consumer worker
- [ ] ClickHouse sink worker
- [ ] Analytics batch insert worker
- [ ] Analytics retry queue
- [ ] Analytics dead letter handling
- [ ] Analytics backfill job
- [ ] Analytics idempotency
- [ ] Analytics ingestion lag handling

---

# Stage 25. Performance Basics

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

## Analytics Performance

- [ ] OLAP query performance
- [ ] ClickHouse query optimization
- [ ] ClickHouse sorting key performance
- [ ] ClickHouse partition pruning
- [ ] Batch insert performance
- [ ] Materialized view performance
- [ ] Aggregated table performance
- [ ] Analytics query latency
- [ ] Analytics ingestion throughput
- [ ] Analytics storage growth

---

# Stage 26. Production Readiness

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

## Data and Analytics Readiness

- [ ] PostgreSQL backup basics
- [ ] MongoDB backup basics
- [ ] ClickHouse backup basics
- [ ] Data retention policy
- [ ] Analytics retention policy
- [ ] Schema migration rollback strategy
- [ ] Analytics backfill strategy
- [ ] Event replay strategy
- [ ] Duplicate event handling
- [ ] Data consistency checks

## Interview Readiness

- [ ] Explain Go project structure
- [ ] Explain domain-first architecture
- [ ] Explain ports and adapters
- [ ] Explain dependency direction
- [ ] Explain transaction boundaries
- [ ] Explain outbox pattern
- [ ] Explain PostgreSQL role
- [ ] Explain Redis role
- [ ] Explain MongoDB role
- [ ] Explain ClickHouse role
- [ ] Explain RabbitMQ role
- [ ] Explain Kafka role
- [ ] Explain gRPC role
- [ ] Explain Docker setup
- [ ] Explain Kubernetes basics
- [ ] Explain testing strategy
- [ ] Explain observability strategy
- [ ] Explain deployment strategy
- [ ] Explain production-readiness decisions