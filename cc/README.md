# AAIDA Backend Service

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
