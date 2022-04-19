GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get go-micro.dev/v4/cmd/protoc-gen-micro


.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/starter.proto
	
.PHONY: build
build:
	go build -o starter *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t starter:latest
