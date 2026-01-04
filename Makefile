run:
	go run cmd/hexlet-path-size/main.go

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

build-and-run:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
	bin/hexlet-path-size

lint:
	golangci-lint run ./... --config=./.github/config/.golangci.yml -v

lint-fix:
	golangci-lint run ./... --config=./.github/config/.golangci.yml --fix -v

tidy-vendor:
	go mod tidy
	go mod vendor

test:
	go clean -testcache && go test ./...