SHELL := /bin/bash

GIT_SHA := $(shell git rev-parse HEAD | cut -c 1-12)
VERSION := $(shell git describe --tags --dirty --always --abbrev=12)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
PWD := $(shell pwd)

PRODUCER_IMAGE_NAME = producer
CONSUMER_IMAGE_NAME = consumer
FETCHER_IMAGE_NAME = fetcher

default: build

all:

build:
	go mod vendor
	GOOS=linux go build -o bin/consumer github.com/Ezetowers/tennis-statistics/consumer
	GOOS=linux go build -o bin/producer github.com/Ezetowers/tennis-statistics/producer
	GOOS=linux go build -o bin/fetcher github.com/Ezetowers/tennis-statistics/fetcher
.PHONY: build

build-darwin:
	go mod vendor
	GOOS=darwin go build -o bin/consumer github.com/Ezetowers/tennis-statistics/consumer
	GOOS=darwin go build -o bin/producer github.com/Ezetowers/tennis-statistics/producer
	GOOS=darwin go build -o bin/fetcher github.com/Ezetowers/tennis-statistics/fetcher
.PHONY: build-darwin

docker-image:
	docker build -f ./producer/Dockerfile -t "$(PRODUCER_IMAGE_NAME):$(GIT_SHA)" .
	docker build -f ./consumer/Dockerfile -t "$(CONSUMER_IMAGE_NAME):$(GIT_SHA)" .
	docker build -f ./fetcher/Dockerfile -t "$(FETCHER_IMAGE_NAME):$(GIT_SHA)" .
.PHONY: docker-image

docker-compose-up:
	docker-compose -f docker-compose-dev.yaml up -d --build
.PHONY: docker-compose-up

docker-compose-down:
	docker-compose -f docker-compose-dev.yaml stop -t 1
	docker-compose -f docker-compose-dev.yaml down
.PHONY: docker-compose-down

docker-compose-logs:
	docker-compose -f docker-compose-dev.yaml logs -f
.PHONY: docker-compose-logs
