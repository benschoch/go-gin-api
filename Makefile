# Misc
.DEFAULT_GOAL = help
.PHONY        = help build run update db-dump db-restore

## —— Golang & MongoDB Makefile ———————————————————————————————————————————————————————————————————————————————————————
help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## —— Golang ——————————————————————————————————————————————————————————————————————————————————————————————————————————
build: ## Build compiles the packages named by the import paths.
	$(MAKE) CACHE=--no-cache build-cache

run: ## Compiles and runs the named main Go package.
	go run cmd/main.go

update: ## Install and update all modules.
	go get -u ./...

## —— MongoDB —————————————————————————————————————————————————————————————————————————————————————————————————————————
db-dump: ## Creates a binary export of a database's contents.
	 mongodump --host=127.0.0.1 -u root -p pass --out=data

db-restore: ## Loads data from either a binary database dump.
	mongorestore --authenticationDatabase=admin --uri "mongodb://root:pass@localhost:27017/api" ./data/dev/api

## —— Tests ———————————————————————————————————————————————————————————————————————————————————————————————————————————
