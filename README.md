# Sport Management System API

This is my backend project for Sport Management System using Go + Gin + PostgreSQL.

---

## Work Completion Summary

I have completed the requested tasks:

- Implemented pagination on list endpoints:
  - `GET /api/v1/teams/`
  - `GET /api/v1/players/`
  - `GET /api/v1/matches/`
- Added dummy/seeding migration for more data example.
- Updated Postman collection with pagination examples (`?page=1`).

Response includes metadata:

```json
"meta": {
  "total": 123,
  "page": 1,
  "page_size": 10
}
```

## Tech Stack

- Go
- Gin
- PostgreSQL
- sqlx
- golang-migrate
- JWT auth

## How to Run

1. Set env in `.env`:

```env
DB_USERNAME=postgres
DB_PASSWORD=
DB_HOST=localhost
DB_PORT=5432
DB_NAME=sport_management_db
DB_SSLMODE=disable
JWT_SECRET=evan-secret-key
APP_PORT=8080
```

2. Run:

```bash
go run main.go
```

> The app runs migration automatically on startup.

## Postman

Use file:

- `postman_collection.json`

Login first to get token, then call protected endpoints using:

- `Authorization: Bearer {{token}}`

## Project Structure

```text
config/
internal/
  dto/
  handlers/
  middleware/
  models/
  repository/
  routes/
  services/
  utils/
migrations/
main.go
```

## Notes

- This project uses layered architecture: `handler -> service -> repository`.
- Soft delete is used (`deleted_at`).
- List endpoints now support pagination and return metadata.
- Dockerization and CI/CD are not implemented in this version.

