.PHONY: swaggo

# ==============================================================================
# Main

run:
	go run ./cmd/minitwitter/minitwitter.go

build:
	go build ./cmd/minitwitter/minitwitter.go

# ==============================================================================
# Tools commands

run-linter:
	echo "Starting linters"
	golangci-lint run ./...

swaggo:
	echo "Starting swagger generating"
	swag init -g **/**/*.go

# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps-clean-cache:
	go clean -modcache