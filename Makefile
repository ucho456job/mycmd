.PHONY: clip edit exec read view make_help

clip: ## Execute "go run main.go clip"
	go run main.go clip -g $(g) -t $(t)

edit: ## Execute "go run main.go edit"
	go run main.go edit -e $(e)

exec: ## Execute "go run main.go exec"
	go run main.go exec -g $(g) -t $(t)

read: ## Execute "go run main.go read"
	go run main.go read -g $(g) -t $(t)

view: ## Execute "go run main.go view"
	go run main.go view

make_help: ## Show Makefile options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
