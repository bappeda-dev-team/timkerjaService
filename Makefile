APP_NAME=timkerja-service

.PHONY: all build run env clean docs help

# DEFAULT TARGET
all: build

build: $(APP_NAME)

$(APP_NAME): *.go
	@echo ">>> Building $(APP_NAME)..."
	@go build -o $(APP_NAME) .
	@echo ">>> SUCCESS..."

run: clean build
	@echo ">>> SET REQUIRED ENV..."
	@set -a; \
	source .env; \
	set +a; \
	echo ">>> REQUIRED ENV"; \
	$(MAKE) --no-print-directory env; \
	echo ">>> Running $(APP_NAME)..."; \
	./$(APP_NAME)

env:
	@echo ">>> REQUIRED ENV"
	@printf "DB_HOST:     "; [ -n "$$DB_HOST" ] && echo "$$DB_HOST" || echo "NOT SET"
	@printf "DB_PORT:     "; [ -n "$$DB_PORT" ] && echo "$$DB_PORT" || echo "NOT SET"
	@printf "DB_USER:     "; [ -n "$$DB_USER" ] && echo "$$DB_USER" || echo "NOT SET"
	@printf "DB_PASSWORD: "; [ -n "$$DB_PASSWORD" ] && echo "SET" || echo "NOT SET"
	@printf "DB_NAME:     "; [ -n "$$DB_NAME" ] && echo "$$DB_NAME" || echo "NOT SET"

clean:
	@echo "CLEANING UP"
	rm -f $(APP_NAME)

docs:
	@echo "GENERATING DOCS"
	swag init
	@echo "DOCS GENERATED"

help:
	@echo "Available commands:"
	@echo "  make          Build application"
	@echo "  make build    Build application"
	@echo "  make run      Build, load .env, check env, and run"
	@echo "  make env      Check required environment"
	@echo "  make clean    Remove binary"
	@echo "  make docs     Generate Swagger documentation"
