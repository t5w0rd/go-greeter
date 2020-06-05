
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
    
	protoc --proto_path=. --go_out=plugins=grpc:. proto/greeter/greeter.proto
    

.PHONY: build
build: proto

	go build -o greeter-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t greeter-service:latest
