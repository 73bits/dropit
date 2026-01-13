# dropit

**share text that expires**

dropit is a minimal API service that lets you share plain text via a unique link that automatically expires after a defined time (TTL). Think of it as a lightweight, self-hosted Pastebin with built-in expiration.

---

## ✨ Features

* ⏳ Time-based expiration (TTL)
* 🔗 Share via simple link
* 🧠 Clean API design
* 🔐 Cryptographically secure IDs
* 🧹 Automatic cleanup of expired data
* 🚀 Easy to extend (Redis / PostgreSQL ready)

---

## 🧩 Use cases

* Sharing secrets or tokens temporarily
* One-off messages
* Debug logs or stack traces
* Secure handoff between systems

---

## 📡 API Overview

### Create a temporary text

`POST /text`

**Request**

```json
{
  "content": "this message will expire",
  "ttl_seconds": 60
}
```

**Response**

```json
{
  "url": "http://localhost:8080/text/abc123"
}
```

---

### Retrieve the text

`GET /text/{id}`

**Response**

```json
{
  "content": "this message will expire"
}
```

If the text is expired or does not exist, the API returns:

```
404 Not Found
```

---

## 🛠️ Getting Started

### Prerequisites

* Go 1.21+

### Run locally

```bash
go run cmd/server/main.go
```

Server will start on:

```
http://localhost:8080
```

---

## 🧪 Testing with curl

### Create text

```bash
curl -X POST http://localhost:8080/text \
  -H "Content-Type: application/json" \
  -d '{"content":"hello","ttl_seconds":30}'
```

### Fetch text

```bash
curl http://localhost:8080/text/<id>
```

---

## 📦 Why dropit?

> "Some data should disappear by default."

dropit is built to encourage safer data sharing by making expiration the norm, not an afterthought.

---

## 📄 License

MIT License
