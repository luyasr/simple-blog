.PHONY: all build

APP_NAME = simple-blog
REGISTRY = registry.cn-hangzhou.aliyuncs.com
NAMESPACE = hubcn
COMMIT_SHA=$(shell git rev-parse --short HEAD)

all: build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}

build-docker:
	docker build -t ${REGISTRY}/${NAMESPACE}/${APP_NAME}:${COMMIT_SHA}

push:
	docker push ${REGISTRY}/${NAMESPACE}/${APP_NAME}:${COMMIT_SHA}