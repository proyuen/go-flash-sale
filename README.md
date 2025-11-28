# Go Flash Sale (High-Concurrency Seckill System)

![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?style=flat&logo=docker)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-336791?style=flat&logo=postgresql)

An enterprise-grade high-concurrency flash sale system built with modern **Go** ecosystem. 
This project demonstrates the implementation of a distributed system capable of handling high traffic loads, ensuring data consistency (no overselling), and high availability.

> **Note:** This project is currently under active development.

## 🏗 Architecture & Tech Stack

This project follows the **Standard Go Project Layout** and **Domain-Driven Design (DDD)** principles.

### Backend
* **Language:** Go (Golang) 1.25+
* **Web Framework:** Gin
* **Database:** PostgreSQL 15
* **ORM / Data Access:** SQLC (Type-safe SQL generation, Database-First approach)
* **Cache (Planned):** Redis (Lua scripts for atomic inventory deduction)
* **Messaging (Planned):** RabbitMQ (Async processing & traffic peak clipping)
* **Config Management:** Viper
* **Migration:** Golang-Migrate

### Infrastructure
* **Containerization:** Docker & Docker Compose
* **Automation:** GNU Make (Makefile)

---

## 🚀 Getting Started

Follow these steps to set up the project locally.

### Prerequisites
* [Docker](https://www.docker.com/) & Docker Compose
* [Go](https://go.dev/) 1.25+
* Make (Optional, for running makefile commands)

### 1. Clone the Repository
```bash
git clone [https://github.com/proyuen/go-flash-sale.git](https://github.com/proyuen/go-flash-sale.git)
cd go-flash-sale