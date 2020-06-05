
GOPATH:=$(shell go env GOPATH)
    

.PHONY: build
build:

	go build -o greeter-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t greeter-service:latest
