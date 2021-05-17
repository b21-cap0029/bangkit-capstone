MODULE=github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend

.PHONY: all
all: dist/aaida

dist/aaida:
	@go build -o $@

.PHONY: clean
clean:
	@-rm -rf dist
	
test:
	@go test $(MODULE)/internal/handler