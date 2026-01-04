run:
	go run cmd/hexlet-path-size/main.go

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

build-and-run:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
	bin/hexlet-path-size
