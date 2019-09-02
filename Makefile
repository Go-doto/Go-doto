.PHONY: start
start: erase build up ## clean current environment, recreate dependencies and spin up again

.PHONY: stop
stop: ## stop environment
		docker-compose stop

.PHONY: rebuild
rebuild: start ## same as start

.PHONY: erase
erase: ## stop and delete containers, clean volumes.
		docker-compose stop
		docker-compose rm -v -f

.PHONY: build
build: ## build environment and initialize composer and project dependencies
		docker-compose build

.PHONY: up
up: ## spin up environment
		docker-compose up -d

.PHONY: queue-new
queue-new: ## check queue New topic
		bash ./scripts/queue-new.sh New

.PHONY: queue-solved
queue-solved: ## check queue Solved topic
		bash ./scripts/queue-new.sh Solved

.PHONY: queue-unsolved
queue-unsolved: ## check queue Unsolved topic
		bash ./scripts/queue-new.sh Unsolved

.PHONY: run
run: ## run app commands. make run command=hello
		GO111MODULE=on ${GOROOT}/bin/go run cmd/main.go $(command)

.PHONY: help
help: ## Display this help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
