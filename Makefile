build:
	@go build -o deploy/app main.go

fmt:
	@goimports -w -local github.com/mpedrozoduran/go-orchestrator .
	@go fmt ./...

lint: fmt
	@golangci-lint run

test:
	@go test ./...

run:
	@go run main.go

run-local:
	GOOS=linux GOARCH=amd64 go build -o deploy/app main.go
	@docker build -t app:latest .
	@docker-compose -f ./scripts/docker-compose.yml up -d
	./scripts/deps.sh


.PHONY: build fmt lint test run