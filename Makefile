APP_NAME=timkerja-service

.PHONY: all build run env clean

# DEFAULT TARGET
all: build

build: $(APP_NAME)

$(APP_NAME): *.go
	@echo ">>> Building $(APP_NAME)..."
	@go build -o $(APP_NAME) .
	@echo ">>> SUCCESS..."

run: build env
	@echo ">>> Running $(APP_NAME)..."
	./$(APP_NAME)

env:
	@echo "REQUIRED ENV"
	@echo "DB_HOST: $(DB_HOST)"
	@echo "DB_PORT: $(DB_PORT)"
	@echo "DB_USER: $(DB_USER)"
	@echo "DB_PASSWORD: $(DB_PASSWORD)"
	@echo "DB_NAME: $(DB_NAME)"

clean:
	@echo "CLEANING UP"
	rm -f $(APP_NAME)
