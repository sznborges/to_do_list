.PHONY: help run docker/up tools swagger install test lint

GO = go
GOLANGCILINT = golangci-lint

# Load envs from .env file if it exists
-include: .env
export

DB_DSN=host=localhost port=5432 user=to-do-list-user password=to-do-list-password dbname=to-do-list sslmode=disable connect_timeout=4 statement_timeout=2s
_goose_ = goose -dir ./migration postgres "$(DB_DSN)"

# To add description to a target, just put a comment with two # after the target definition
# Ex:
# target_name: target_dep1 target_dep2  ## i'm a description
# do anything

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make [target]\033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_/-]+:.*?##/ { printf "\033[36m%-18s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

run: ## Run and serve an api
	@$(GO) run main.go api

dep/up:  ## Build and run docker apps
	@docker compose --env-file .docker-compose.env up --build -d
	until docker exec to-do-list-postgres-db-1 pg_isready ; do sleep 1 ; done
	make db/migrate-up

dep/down:  ## Stop docker apps
	@docker compose --env-file .docker-compose.env down

tools:  ## Install golangci-lint, swaggo and goose
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@$(GO) install github.com/swaggo/swag/cmd/swag@latest
	@$(GO) install github.com/pressly/goose/v3/cmd/goose@latest
	@$(GO) install github.com/marouni/adr@latest

docs/swagger/*: $(wildcard server/*.go)
	@swag init -generalInfo server/server.go -output ./docs/swagger

swagger: docs/swagger/*  ## Generate swagger docs

install: tools  ## Install go libs
	@$(GO) mod vendor && $(GO) mod tidy

upgrade/mod:
	go get -u -t ./...

upgrade: upgrade/mod install

test: dep/down dep/up  ## Run tests
	@$(GO) test -count=1 -v ./... -race

test/coverage: dep/down dep/up 
	go test ./... -coverprofile cover.out -short

test/coverage/html: test/coverage
	go tool cover -html cover.out	

lint:  ## Run lint tools
	@$(GOLANGCILINT) run

db/create-migration:
	$(_goose_) create $(MIGRATION_NAME) sql
		
db/migrate-up:
	$(_goose_) up

db/migrate-down:
	$(_goose_) down	

db/migration-status:
	$(_goose_) status

docker/build:
	docker build .	
