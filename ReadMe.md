# 📘 Subscription Reminder Backend

## Overview

The **Subscription Reminder Backend** is a Go Fiber API that helps users track, manage, and receive reminders for their online subscriptions.  
It’s designed for personal use or as a foundation for a full SaaS subscription management system.

## The service supports adding subscriptions, scheduling reminders, and (in future versions) integrating AI-based analytics, payment tracking, and even direct subscription cancellation features

## 🧱 Tech Stack

- **Language:** Go (Golang)
- **Framework:** [Fiber](https://gofiber.io/)
- **Database:** PostgreSQL
- **ORM:** GORM
- **Hot Reload:** [Air](https://github.com/air-verse/air)
- **Environment Management:** `.env`
- **Documentation:** Swagger
- **Containerization:** Docker

---

## 📂 Project Structure

```bash
.
├──── main.go          # Entry point of the application
├── internal/
│   ├── routes/          # Route definitions
│   ├── handlers/        # HTTP handlers
│   ├── models/          # Database models
│   ├── services/        # Business logic
│   └── database/        # DB connection & setup
├── configs/
│   └── config.go        # Environment configs
├── .air.toml            # Air configuration file
├── go.mod
├── go.sum
├── .env
└── README.md
```
