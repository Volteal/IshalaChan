build:
	@go build -o ./bin/IshalaChan

tidy:
	@go mod tidy

run: build
	@./bin/IshalaChan

test:
	@go test -v ./...