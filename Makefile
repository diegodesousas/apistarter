.PHONY: test coverage

APPDIR=/go/src/github.com/diegodesousas/apistarter
APPNAME=apistarter
IMAGE=diegodesousas/apistarter
PWD=$(shell pwd)
HTTP_PORT=8080
LAST_VERSION=$(shell git tag | tail -1)

build-image:
ifeq ($(shell docker images -q ${IMAGE} | wc -l),0)
	@echo "Building image ${IMAGE}"
	@docker \
		build \
		--target developer \
		-t ${IMAGE} \
		.
else
	@echo "Image ${IMAGE} already exists"
endif

build-release:
ifeq ($(shell docker images -q ${IMAGE}:${LAST_VERSION} | wc -l),0)
	@echo "Building ${IMAGE}:${LAST_VERSION}"
	@docker \
		build \
		-t ${IMAGE}:${LAST_VERSION} \
		--target release \
		.
else
	@echo "Image ${IMAGE}:${LAST_VERSION} already exists"
endif

test:
	@echo "Running unit tests"
	@docker \
		run \
		-it \
		-v ${PWD}:${APPDIR} \
		-p ${HTTP_PORT}:${HTTP_PORT} \
		--name ${APPNAME} \
		--rm  \
		-w ${APPDIR} \
		${IMAGE} \
		go test -coverprofile=coverage.out -race ./...

coverage:
	@docker \
		run \
		-it \
		-v ${PWD}:${APPDIR} \
		-p ${HTTP_PORT}:${HTTP_PORT} \
		--name ${APPNAME} \
		--rm  \
		-w ${APPDIR} \
		${IMAGE} \
		go tool cover -html=coverage.out -o coverage.html

shell:
	@docker \
		run \
		-it \
		-v ${PWD}:${APPDIR} \
		-p ${HTTP_PORT}:${HTTP_PORT} \
		--name ${APPNAME} \
		--rm  \
		-w ${APPDIR} \
		${IMAGE} \
		bash

build-app:
	go mod vendor
	go build cmd/server.go

watch:
	@docker run \
		-it \
		-v ${PWD}:${APPDIR} \
		-p ${HTTP_PORT}:${HTTP_PORT} \
		--name ${APPNAME} \
		--rm  \
		-w ${APPDIR} \
		${IMAGE} \
		dogo
