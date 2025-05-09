.PHONY: help dev dev-shell test lint fmt build run test-local clean init tidy

# Environment detection
IS_CONTAINER := $(shell test -f /.dockerenv && echo true || echo false)

# Reusable command executor for Go commands, allowing commands to be run from
# inside the dev container, or on the host machine
define run_go
	@if [ "$(IS_CONTAINER)" = "true" ]; then \
		echo "🔧 Running: $(1)"; \
		$(1); \
	else \
		if ! docker compose ps --services --filter "status=running" | grep -q "^app$$"; then \
			echo "❌ Dev container 'app' is not running. Run 'make up' first."; \
			exit 1; \
		fi; \
		echo "🐳 Running in container: $(1)"; \
		docker compose exec app $(1); \
	fi
endef

# =====================
# Help
# =====================
.DEFAULT_GOAL := help
help:
	@echo ""
	@echo "🔧 Dev Environment Commands"
	@printf "  %-15s %s\n" "up" "Start development container"
	@printf "  %-15s %s\n" "up-rebuild" "Rebuild and start the dev container"
	@printf "  %-15s %s\n" "dev-shell" "Enter a shell in the running dev container"
	@echo ""
	@echo "🚀 In-Container Commands (auto-detect context)"
	@printf "  %-15s %s\n" "dev" "Run development server"
	@printf "  %-15s %s\n" "test" "Run tests"
	@printf "  %-15s %s\n" "lint" "Run linter"
	@printf "  %-15s %s\n" "fmt" "Format code"
	@printf "  %-15s %s\n" "build" "Build the Go binary"
	@printf "  %-15s %s\n" "run" "Run the Go application"
	@printf "  %-15s %s\n" "init" "Initialize a new go module"
	@printf "  %-15s %s\n" "tidy" "Tidy and verify go.mod/go.sum"
	@echo ""
	@printf "  %-15s %s\n" "clean" "Remove build artifacts"
	@echo ""

# =====================
# Dev Container Commands
# =====================

up:
	docker compose up

up-rebuild:
	docker compose up --build

dev-shell:
	docker compose exec app /bin/bash

# =====================
# In-Container Commands (run automatically inside or via docker exec)
# =====================

dev: ## Run the app in dev mode with hot reload
	$(call run_go,air -c .air.toml)

test:
	$(call run_go,go test ./...)

lint:
	$(call run_go,golangci-lint run)

fmt:
	$(call run_go,sh -c 'goimports -w . && go fmt ./...')

build: ## Build the application inside the container
	$(call run_go,go build -o bin/app ./cmd/app)

run: ## Run the application inside the container
	$(call run_go,go run ./cmd/app)


init: ## Initialize a new go module (inside container)
	$(call run_go,test -f go.mod || go mod init github.com/AGiantSquid/golang-template)

tidy: ## Tidy and verify dependencies (inside container)
	$(call run_go,sh -c 'go mod tidy && go mod verify')


clean: ## Remove build artifacts
	rm -rf bin/
	rm -rf tmp/
	rm -f coverage.out
	rm -f coverage.html
