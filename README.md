# Protego MCP Protocol v0.2

This branch introduces a **minimal Go backend** to prove the API round-trip works.

## What’s here
- `GET /api/health` → health check
- `POST /api/grade` → stub grader that fails if payload contains `SELECT * FROM`

> No database, no LLM, no frontend yet—just a tiny, testable API.

---

## Prerequisites
- Go 1.23+

## Run the server
```bash
cd server
go mod tidy
go run ./...
Server starts on: http://localhost:8080
```

---

## Test the endpoints
### Health
```bash
harish $ curl http://localhost:8080/api/health
ok
```

---
### Grade(stub)
- Fails when code contains 'SELECT * FROM'
```bash
harish $ curl -X POST http://localhost:8080/api/grade -d '{"code":"SELECT * FROM"}'
{"pass":false,"findings":[{"rule":"SQLI-1","detail":"Found raw 'SELECT * FROM' — avoid unsafe patterns. Use explicit columns and parameters.","severity":"high","passed":false}]}
```

- Passes with other code 
```bash
harish $ curl -X POST http://localhost:8080/api/grade -d '{"code":"print(\"hello\")"}'
{"pass":true,"findings":[]}
```
