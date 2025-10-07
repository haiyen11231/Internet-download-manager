# 🌐 Internet Download Manager

This project implements a backend microservice system for managing and tracking internet downloads.
It demonstrates how to build a Go application with modular design, gRPC communication, HTTP APIs, caching (Redis/in-memory), message queue (MQ), scheduled cron jobs and dependency injection using Google Wire.

## 📌 Key Features

### 🧩 Hybrid API Gateway (gRPC + HTTP)

- Exposes both **gRPC** and **RESTful HTTP/JSON** endpoints via `grpc-gateway`.
- Automatically converts JSON requests into gRPC calls for unified backend logic.

### ⚡ Message Queue for Download Tasks

- Download requests are **produced as messages** and **consumed asynchronously** by worker services.
- Ensures scalability, retry logic, and decoupled task execution.

### 🧠 Caching Layer

- Abstracted caching interface with **Redis** and **in-memory** support.
- Used to optimize operations such as:
  - Checking taken account names.
  - Storing token public keys.
- Easily switchable via YAML configuration.

### 🔐 Token-Based Authentication

- Implements **JWT authentication** with RSA key pairs.
- Public keys are cached for performance and can be rotated safely.
- Token lifetimes and regeneration durations are configurable.

### 🧵 Dependency Injection with Wire

- Uses **Google Wire** to automate dependency wiring.
- Reduces boilerplate in initializing complex structs and services.

### 🕒 Background Jobs / Cron Tasks

- Includes **scheduled jobs** to periodically process or clean up download tasks.
- Ensures failed or expired tasks are retried or removed automatically.

### 📦 Protocol Buffers with Buf

- Uses **Buf** for Protobuf schema management and validation.
- Integrates plugins:
  - `grpc-go`, `grpc-gateway`, `validate-go`, and `openapiv2`.
- Generates client/server stubs and OpenAPI documentation automatically.

### 💾 Configurable Download Backend

- Supports multiple storage modes:
  - **S3-compatible** (e.g., MinIO)
  - Future extensibility for local or other providers
- Configurable via YAML under `download:` section.

---

## 📂 Repository Structure

```

```

---

## 🚀 Getting Started

### Prerequisites

-

### Setup Instructions
