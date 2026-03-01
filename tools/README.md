Tools
=====

This folder contains small helper programs for local development and CI.

Available:

- `smoke.go` — runs a small set of HTTP checks against the running API (health, product listing, sample product, user/admin login, and admin create product).
- `promote.go` — one-off helper to promote `admin@local.com` to role `admin` in the local `data.db`.

Usage examples:

Start the backend (from repo root):

```bash
cd backend
export JWT_SECRET=devsecret
go run cmd/server/main.go
```

Run smoke tests (from repo root):

```bash
go run ./tools/smoke.go
```

Promote admin (from repo root):

```bash
go run ./tools/promote.go
```

CI should run the backend and then `go run ./tools/smoke.go` to validate the service is healthy.
