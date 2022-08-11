SHELL := /bin/bash

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## NEW REPOSITORY
## Note: All services here use the development environment located in ./.env
#
## run/cogs: Running CoGS service (REST API)
.PHONY: run/pos
run/cogs:
	go run ./app/main.go

# run/swagger: Running swagger (API DOC) service
.PHONY: run/swagger
run/swagger:
	go run ./app/swagger

## setup: setup GOPRIVATE for private dependency
.PHONY: setup
setup:
	export GO111MODULE=on GOPRIVATE="github.com/upgradeskill/fp2022-beta-team/*" GOSUMDB=off


## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database '${MINIPOS_DB_DSN}' up

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #
## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	@echo 'Running tests...'
	go test -vet=off ./...