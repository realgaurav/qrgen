## Courtesy: https://www.padok.fr/en/blog/beautiful-makefile-awk
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

default: help

bin: qrgen ## Generate binary.

qrgen: cmd/main.go
	CGO_ENABLED=0 go build -o $@ $?

