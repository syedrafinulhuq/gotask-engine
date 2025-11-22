# **GoTask Engine — Master Guide**

## Overview

GoTask Engine is a **distributed task runner + job scheduler** built entirely in **Go**.

It teaches you:

* concurrency
* worker pools
* queues
* schedulers
* distributed worker nodes
* HTTP API
* config system
* logging
* storage
* CLI tools

This project is your roadmap to becoming a **real Go backend engineer**.

---

## Core Purpose

The goal of GoTask Engine is:

* create tasks
* schedule tasks
* enqueue tasks
* distribute tasks across workers
* run tasks on multiple machines
* track results

While building it, you will learn **every major concept** in Go.

---

## What You Must Know (List + Purpose + Where It’s Used)

---

# 1. Go Basics (Syntax, Functions, Packages)

**Purpose:**
Understand Go structure, functions, imports, modules.

**Used In Project:**
Every part — API handlers, services, workers, scheduler.

---

# 2. Structs & Methods

**Purpose:**
Represent data and logic cleanly.

**Used In Project:**

* `Task` model
* `TaskService` logic
* `WorkerService`
* Config structs
* Queue & Job objects

---

# 3. Error Handling

**Purpose:**
Safely return informative errors.

**Used In Project:**

* config loading
* API validation
* scheduler execution
* job runner failures

---

# 4. Pointers

**Purpose:**
Avoid copying data, share state, modify objects.

**Used In Project:**

* services (`*TaskService`)
* handlers (`*TaskHandler`)
* shared queues
* worker pools

---

# 5. Goroutines

**Purpose:**
Run tasks concurrently.

**Used In Project:**

* worker loop
* scheduler loop
* async job execution
* worker heartbeats

---

# 6. Channels

**Purpose:**
Communicate safely between goroutines.

**Used In Project:**

* job queue
* worker pool
* scheduler → queue
* queue → worker

---

# 7. Select Statement

**Purpose:**
Wait on multiple channels, timeouts, cancellation.

**Used In Project:**

* scheduler ticking
* worker listening
* graceful shutdown

---

# 8. Mutex / RWMutex

**Purpose:**
Synchronize shared data access.

**Used In Project:**

* storing tasks in memory
* worker registry
* scheduled tasks map

---

# 9. net/http

**Purpose:**
Build REST API.

**Used In Project:**

* /health
* /tasks (POST, GET)
* /workers/register
* /workers/heartbeat
* CLI → server communication

---

# 10. JSON Encoding/Decoding

**Purpose:**
Parse API requests and responses.

**Used In Project:**

* adding tasks
* listing tasks
* worker registration
* config file loading

---

# 11. Project Structure (cmd, internal, pkg)

**Purpose:**
Maintain clean, scalable architecture.

**Used In Project:**

* `cmd/server/main.go` → entry
* `cmd/worker/main.go` → worker node
* `internal/tasks` → business logic
* `pkg/logger` → shared utilities

---

# 12. Config Files (JSON/YAML)

**Purpose:**
Make server configurable without editing code.

**Used In Project:**

* port number
* worker concurrency
* queue size
* log settings

---

# 13. Task Model (Entity)

**Purpose:**
Define what a task is.

**Used In Project:**

A Task contains:

* id
* name
* command
* schedule
* created_at

---

# 14. Task Service (Business Logic)

**Purpose:**
Manage in-memory or DB tasks.

**Used In Project:**

* add task
* list tasks
* update task
* delete task

---

# 15. HTTP Handlers (API Layer)

**Purpose:**
Expose logic to the outside world.

**Used In Project:**

* POST /tasks
* GET /tasks
* error validation

---

# 16. Job Queue (Core Concurrency Component)

**Purpose:**
Store jobs for workers to process.

**Used In Project:**

* buffered channels for jobs
* workers consume jobs
* scheduler pushes jobs

---

# 17. Worker Pool (Goroutines)

**Purpose:**
Run multiple tasks concurrently and safely.

**Used In Project:**

* dynamic worker pool
* configurable concurrency
* workers pulling from queue

---

# 18. Scheduler (Ticker + Cron)

**Purpose:**
Automatically enqueue tasks based on schedule.

**Used In Project:**

* ticker loop
* calculate next run
* push job into queue

---

# 19. Distributed Workers (Networking)

**Purpose:**
Scale horizontally across machines.

**Used In Project:**

* worker registration
* server assigns jobs
* workers heartbeat
* workers fetch jobs

---

# 20. CLI Tool (Control Panel)

**Purpose:**
Manage tasks and workers from terminal.

**Used In Project:**

* `gotaskctl add`
* `gotaskctl list`
* `gotaskctl run`

---

# 21. Database (BoltDB or PostgreSQL)

**Purpose:**
Persist tasks + logs across restarts.

**Used In Project:**

* storing tasks
* storing runs
* storing worker info

---

# 22. Logging

**Purpose:**
Debug, monitor, and understand flow.

**Used In Project:**

* server logs
* worker logs
* scheduler logs

---

# 23. Testing (Unit + Concurrency)

**Purpose:**
Ensure system works reliably.

**Used In Project:**

* task service tests
* queue tests
* scheduler tests
* API handler tests

---

# 24. Profiling (Optional)

**Purpose:**
Optimize performance using pprof.

**Used In Project:**

* detect deadlocks
* track goroutine leaks
* memory usage

---

## **High-Level Architecture**

```
                 ┌──────────────────┐
                 │   HTTP Server    │
                 │  /tasks /workers │
                 └─────────┬────────┘
                           │
                           ▼
                 ┌──────────────────┐
                 │   Task Service    │
                 └─────────┬────────┘
                           │
                           ▼
                 ┌──────────────────┐
                 │    Scheduler      │
                 │ (Ticker + Cron)   │
                 └─────────┬────────┘
                           │
                           ▼
                 ┌──────────────────┐
                 │     Queue         │
                 │  (Jobs Channel)   │
                 └─────────┬────────┘
                           │
                           ▼
           ┌───────────────────────────────┐
           │          Worker Pool           │
           │  goroutines processing jobs    │
           └───────────────────────────────┘
```

---

## **What To Build First (Ordered Roadmap)**

```markdown
```

```
1. Basic server (health route)
2. Config loader
3. Task model
4. Task service (with mutex)
5. Task handlers (POST/GET)
6. Queue (Jobs chan)
7. Worker pool (goroutines)
8. Scheduler (time.Ticker)
9. Scheduler -> Queue connection
10. Worker registration API
11. Worker heartbeats
12. Worker fetch tasks
13. CLI tool
14. Database storage
15. Logging
16. Tests
```

---

## **Final Notes**

* This project is designed to **teach** you Go deeply.
* You must build each component slowly, testing each part.
* Do **NOT** jump ahead — finish one layer before moving to the next.
* You can re-read this README every day to track progress.

---
