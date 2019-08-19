.PHONY: default help install build build-alpine package tag push test clean

BIN_NAME={{cookiecutter.app_name}}

VERSION := $(shell git describe --tags)
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
ENVIRONMENT=$(shell ([[ $(GIT_BRANCH) == 'master' ]] && echo "Production") || ([[ $(GIT_BRANCH) == 'staging' ]] && echo "Staging") || echo "Development")
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')

IMAGE_NAME := "{{cookiecutter.docker_hub_username}}/{{cookiecutter.app_name}}"

default: test

help:
	@echo 'Management commands for {{cookiecutter.app_name}}:'
	@echo
	@echo 'Usage:'
	@echo '    make install             Install the projects dependencies.'
	@echo '    make build               Compile the project.'
	@echo '    make build-alpine        Compile optimized for alpine linux.'
	@echo '    make package             Build final docker image with just the go binary inside.'
	@echo '    make tag                 Tag image created by package with latest, git commit and version.'
	@echo '    make push                Push tagged images to registry.'
	@echo '    make test                Run tests on a compiled project.'
	@echo '    make clean               Clean the directory tree.'
	@echo

install:
	go get -v ./...

build: install
	@echo "building ${BIN_NAME} $(VERSION)"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.GitCommit=$(GIT_COMMIT)$(GIT_DIRTY) -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.Version=$(VERSION) -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/Environment.Version=$(ENVIRONMENT) -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.BuildDate=$(BUILD_DATE)" -o bin/${BIN_NAME}

build-alpine: install
	@echo "building ${BIN_NAME} $(VERSION)"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags '-w -linkmode external -extldflags "-static" -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.GitCommit=$(GIT_COMMIT)$(GIT_DIRTY) -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.Version=$(VERSION) -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/Environment.Version=$(ENVIRONMENT) -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.BuildDate=$(BUILD_DATE)' -o bin/${BIN_NAME}

package:
	@echo "building image ${BIN_NAME} $(VERSION) $(GIT_COMMIT)"
	docker build --build-arg VERSION=$(VERSION) --build-arg GIT_COMMIT=$(GIT_COMMIT) -t ${IMAGE_NAME}:local .

tag: package
	@echo "Tagging: latest $(VERSION) $(GIT_COMMIT)"
	docker tag ${IMAGE_NAME}:local ${IMAGE_NAME}:$(GIT_COMMIT)
	docker tag ${IMAGE_NAME}:local ${IMAGE_NAME}:$(VERSION)
	docker tag ${IMAGE_NAME}:local ${IMAGE_NAME}:latest

push: tag
	@echo "Pushing docker image to registry: latest $(VERSION) $(GIT_COMMIT)"
	docker push ${IMAGE_NAME}:$(GIT_COMMIT)
	docker push ${IMAGE_NAME}:$(VERSION)
	docker push ${IMAGE_NAME}:latest

test:
	go test ./...

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}
