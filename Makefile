.PHONY: clip edit exec help read view make_help

clip: ## Execute "go run main.go clip"
	go run main.go clip -g $(g) -t $(t)

edit: ## Execute "go run main.go edit"
	go run main.go edit -e $(e)

exec: ## Execute "go run main.go exec"
	go run main.go exec -g $(g) -t $(t)

help: ## Execute "go run main.go help"
	go run main.go help

read: ## Execute "go run main.go read"
	go run main.go read -g $(g) -t $(t)

view: ## Execute "go run main.go view"
	go run main.go view

build_test: ## Build mycmd binary file and move it to /usr/local/bin/
	go build -o mycmd
	sudo mv mycmd /usr/local/bin/

build_release: ## Build binary files for release
	GOOS=linux GOARCH=amd64 go build -o mycmd_linux
	GOOS=darwin GOARCH=amd64 go build -o mycmd_mac
	GOOS=windows GOARCH=amd64 go build -o mycmd_windows_amd64.exe

make_help: ## Show Makefile options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
