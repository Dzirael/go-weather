# 🌤 Weather Forecast API

A weather subscription service that allows users to receive weather updates for their city via email, with configurable frequency (hourly or daily).

---

## 📦 Features

- Get current weather by city
- Subscribe to email notifications
- Confirm or unsubscribe via tokenized email links
- Automatic email delivery every configured interval

---

## 🚀 Quick Start

### 🔧 Requirements

Make sure you have the following installed:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- A valid `.env` file (see example below)

---

### 📁 Example `.env` file

Create a `.env` file in the project root:

```env
LOGGER_DISABLE_STACKTRACE=true
ENVIRONMENT=dev

SERVER_ADDRESS=0.0.0.0:8080
SERVER_PROXY_HEADER=X-Forwarded-For

APPLICATION_URL=http://example.com/api
WEBSITE_URL=https://example.com

SMTP_HOST=smtp.gmail.com
SMTP_PORT=465
SMTP_USERNAME=email@gmail.com
SMTP_PASSWORD=

POSTGRES_DSN=postgres://test_user:test_pass@postgres:5432/postgres?sslmode=disable
WEATHER_SUBSCRIPTION_CHECK_INTERVAL=5m
