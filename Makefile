.PHONY: build container test

EXECUTABLE ?= optikon-api
IMAGE ?= bin/$(EXECUTABLE)
REPO = optikon/api
TAG = 0.1.1

all: build

build:
	CGO_ENABLED=0 go build --ldflags '${EXTLDFLAGS}' -o ${IMAGE} github.com/optikon/api

test:
	CGO_ENABLED=1 go test --cover --race github.com/optikon/api

container:
	docker run -t -w /go/src/github.com/optikon/api -v `pwd`:/go/src/github.com/optikon/api golang:1.10.1 make
	docker build -t $(REPO):$(TAG) .
