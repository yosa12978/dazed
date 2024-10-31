.DEFAULT_GOAL = build

build:
	@go mod tidy
	@go build -o bin/dazed ./cmd/dazed/main.go
	@cp config.yaml bin
	@cp -r migrations bin

run: build
	@DAZED_ADDR=0.0.0.0:5000 ./bin/dazed