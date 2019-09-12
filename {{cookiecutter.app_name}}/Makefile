.PHONY: default help install build build-alpine package tag push helm helm-diff test clean

APP_NAME     = {{cookiecutter.app_name}}
VERSION     := $(shell git describe --always --tags)
GIT_COMMIT   = $(shell git rev-parse HEAD)
GIT_BRANCH   = $(shell git rev-parse --abbrev-ref HEAD)
GIT_DIRTY    = $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
ENVIRONMENT ?= $(shell ([[ $(GIT_BRANCH) == 'master' ]] && echo "prod") || ([[ $(GIT_BRANCH) == 'staging' ]] && echo "stg") || echo "dev")
BUILD_DATE   = $(shell date '+%Y-%m-%d-%H:%M:%S')
REGISTRY_URL  ?= "asia.gcr.io/kitabisa-dev"
SQUAD        = "{{cookiecutter.squad}}"

default: help

help:
	@echo 'Management commands for {{cookiecutter.app_name}}:'
	@echo
	@echo 'Usage:'
	@echo '    make install             Install the project`s dependencies.'
	@echo '    make build               Compile the project.'
	@echo '    make build-alpine        Compile optimized for Alpine Linux.'
	@echo '    make package             Build final Docker image with just the Go binary inside.'
	@echo '    make tag                 Tag image created by package with latest, git commit and version.'
	@echo '    make push                Push tagged images to registry.'
	@echo '    make helm                Deploy to Kubernetes via Helm.'
	@echo '    make helm-diff           Shows a diff explaing what a Helm upgrade would change.'
	@echo '    make run ARGS=           Run with supplied arguments.'
	@echo '    make test                Run tests on a compiled project.'
	@echo '    make clean               Clean the directory tree.'

	@echo

install:
	go get -v ./...

build: install
	@echo "building ${APP_NAME} $(VERSION)"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.GitCommit=$(GIT_COMMIT)$(GIT_DIRTY) -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.Version=$(VERSION) -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/Environment.Version=$(ENVIRONMENT) -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.BuildDate=$(BUILD_DATE)" -o bin/${APP_NAME}

build-alpine: install
	@echo "building ${APP_NAME} $(VERSION)"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags '-w -linkmode external -extldflags "-static" -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.GitCommit=$(GIT_COMMIT)$(GIT_DIRTY) -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.Version=$(VERSION) -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/Environment.Version=$(ENVIRONMENT) -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.BuildDate=$(BUILD_DATE)' -o bin/${APP_NAME}

package:
	@echo "building image ${APP_NAME} $(VERSION) $(GIT_COMMIT)"
	docker build --build-arg VERSION=$(VERSION) --build-arg GIT_COMMIT=$(GIT_COMMIT) --build-arg ENVIRONMENT=$(ENVIRONMENT) --build-arg GOPROXY=${GOPROXY} -t ${REGISTRY_URL}/${APP_NAME}:local .

tag: package
	@echo "Tagging: latest $(VERSION) $(GIT_COMMIT)"
	docker tag ${REGISTRY_URL}/${APP_NAME}:local ${REGISTRY_URL}/${APP_NAME}:$(GIT_COMMIT)
	docker tag ${REGISTRY_URL}/${APP_NAME}:local ${REGISTRY_URL}/${APP_NAME}:$(VERSION)
	docker tag ${REGISTRY_URL}/${APP_NAME}:local ${REGISTRY_URL}/${APP_NAME}:latest

push: tag
	@echo "Pushing Docker image to registry: latest $(VERSION) $(GIT_COMMIT)"
	docker push ${REGISTRY_URL}/${APP_NAME}:$(GIT_COMMIT)
	docker push ${REGISTRY_URL}/${APP_NAME}:$(VERSION)
	docker push ${REGISTRY_URL}/${APP_NAME}:latest

helm:
	helm upgrade ${APP_NAME} $(ENVIRONMENT)/app --install \
		--namespace ${APP_NAME} \
		--values _infra/k8s/$(ENVIRONMENT).yaml \
		--set meta.env=$(ENVIRONMENT),meta.squad=${SQUAD},meta.version=$(VERSION),image.repository=${REGISTRY_URL}/${APP_NAME},image.tag=$(VERSION)

helm-diff:
	helm diff upgrade ${APP_NAME} $(ENVIRONMENT)/app --allow-unreleased \
		--namespace ${APP_NAME} \
		--values _infra/k8s/$(ENVIRONMENT).yaml \
		--set meta.env=$(ENVIRONMENT),meta.squad=${SQUAD},meta.version=$(VERSION),image.repository=${REGISTRY_URL}/${APP_NAME},image.tag=$(VERSION)

run: build
        bin/${APP_NAME} ${ARGS}

test:
	go test ./...

clean:
	@test ! -e bin/${APP_NAME} || rm bin/${APP_NAME}
