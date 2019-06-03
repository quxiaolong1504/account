.PHONY: help

export GO111MODULE=on

help:
	@printf "Usage: make [target]\n\n"
	@echo Targets:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'

test: ## Test npm dependencies for the api, admin, and frontend apps
	@echo test


gen:
	go generate ./pkg/...

dev: ## Run the web server
	air -c .air.toml