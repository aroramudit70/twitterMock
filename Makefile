MODULE_NAME=twittermock

GOCMD=go
GORUN=$(GOCMD) run
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

PROJECT_DIR=$(PWD)
GEN_CODE_DIR=$(PROJECT_DIR)/api/generated/

all: 
	clean generate build docker docker-run
build:
	@echo ">building $(MODULE_NAME) application..."
	@mkdir -p bin
	@go build -mod=vendor -o bin/$(MODULE_NAME)
	@echo "Build successful....."

clean:
	rm -rf $(GEN_CODE_DIR)
	rm -rf bin/*
	rm -rf logs/*

generate:
	@cd api && mkdir -p generated && GO111MODULE=off go generate

run:
	$(GOINSTALL)
	@test $(profile) || (echo "specify profile (profile=dev/TBD)"; exit 1)
	bin/$(MODULE_NAME) -config=configs/$(profile).env

Docker:
	@test $(profile) || (echo "specify profile (profile=dev/TBD)"; exit 1)
	docker build --build-arg MODULE=$(MODULE_NAME) -f docker/$(profile)/Dockerfile -t $(MODULE_NAME) .

docker-run:
	docker run -p 8090:8090 $(MODULE_NAME)


