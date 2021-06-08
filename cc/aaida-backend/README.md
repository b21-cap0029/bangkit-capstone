# aaida-backend

## Environment Variables

### Required

- `TWITTER_CLIENT_ID`
- `TWITTER_CLIENT_SECRET`
- `TENSORFLOW_BASE_URL`

### Optional

- `POSTGRES_DSN`

Will default to use sqlite3 if not provided

## Run server locally

```bash
go run main.go
```

# Run tests

```bash
make test
```

# Build binary

```bash
make
```

## Run using docker-compose

```bash
docker-compose up -d
```
