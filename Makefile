.PHONY: test coverage

APPDIR=/go/src/github.com/diegodesousas/apistarter
APPNAME=go-apistarter
IMAGE=go-apistarter
GOPATH=/go
PWD=$(shell pwd)
PORT=9000

docker: 
	docker run \
		-it \
		-v ${PWD}:${APPDIR} \
		-p ${PORT}:${PORT} \
		--name ${APPNAME} \
		--hostname ${APPNAME} \
		--rm  \
		-w ${APPDIR} \
		${IMAGE} \
		${CMD} \

test:
	@$(MAKE) CMD="go test -coverprofile=coverage.out -race ./..." docker

coverage:
	@$(MAKE) CMD="go tool cover -html=coverage.out -o coverage.html"

sh:
	@$(MAKE) docker

build-app:
	go mod vendor
	go build cmd/server.go

watch:
	@$(MAKE) CMD="dogo" docker

build:
	 docker build -t ${IMAGE} .