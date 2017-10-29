DOCKER_IMAGE := 175914186171.dkr.ecr.ap-southeast-2.amazonaws.com/audience/api-content
DOCKER_BUILDER := ${DOCKER_IMAGE}:builder
BINARY_NAME := go-naive-chain

# Default target (since it's the first without '.' prefix)
build-local: clean fmt build

build-docker: clean fmt build-linux

clean:
	rm -f ./go-naive-chain

fmt:
	gofmt -w -s $$(find . -type f -name '*.go' -not -path "./vendor/*")

build:
	go build ./cmd/go-naive-chain

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/go-naive-chain

run: build
	./$(BINARY_NAME) 2>&1

docker:
	docker-compose build

run-docker: docker
	docker-compose up -d

stop-docker:
	docker-compose stop

# None of the Make tasks generate files with the name of the task, so all must be declared as 'PHONY'
.PHONY: build-all clean fmt build run docker run-docker stop-docker
