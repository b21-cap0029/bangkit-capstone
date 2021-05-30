# aaida-backend

## Required Environment Variables

- `TWITTER_CLIENT_ID`
- `TWITTER_CLIENT_SECRET`
- `TENSORFLOW_SERVING_HOST`

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

## Run using docker

```bash
docker build -t aaida .
docker run -p 8000:8000 aaida
```
