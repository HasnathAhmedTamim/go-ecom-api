Backend (Go + Gin + SQLite)

Quick start

- Set dev secret and run:

  Windows PowerShell:

  $env:JWT_SECRET='devsecret'
  cd backend
  go run cmd/server/main.go

- Build:

  cd backend
  make build

Developer tasks

- Format: `make fmt`
- Vet: `make vet`
- Tidy modules: `make tidy`
- Run smoke tests: `make smoke` (requires server running on 127.0.0.1:8080)

Notes

- Database file is `backend/data.db` (ignored by .gitignore). Seeds run on first start to populate demo users and products.
