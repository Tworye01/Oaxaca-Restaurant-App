GOFILES := $(shell find . -name "*.go")
GOMODULES := $(shell go list ./...)

.phony: all
all:
	go run .

.phony: build
build:
	go build -o build/program/app .

.phony:
clean:
	rm -rf build
	rm -rf golds

.phony: test
test:
	go clean -testcache
	go mod tidy
	go test -cover $(GOMODULES)

.phony: swag
swag:
	swag init

.phony: fmt
fmt:
	gofmt "-s" -w $(GOFILES)

.phony: golds
golds:
	golds -gen -dir=golds -only-list-exporteds -nouses -compact -theme=dark ./...