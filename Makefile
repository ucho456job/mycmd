.PHONY: edit exec make_help

edit: ## Execute "go run main.go edit"
	go run main.go edit -e $(e)

exec: ## Execute "go run main.go exec"
	go run main.go exec -g $(g) -t $(t)

make_help: ## Show Makefile options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
