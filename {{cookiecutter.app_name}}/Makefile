.PHONY: default help install build build-alpine package tag push helm helm-diff test clean

SHELL         = /bin/bash
APP_NAME      = {{ cookiecutter.app_name }}
VERSION      := $(shell git describe --always --tags)
GIT_COMMIT    = $(shell git rev-parse HEAD)
GIT_DIRTY     = $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE    = $(shell date '+%Y-%m-%d-%H:%M:%S')
SQUAD         = "{{ cookiecutter.squad }}"

default: help

help:
	@echo 'Management commands for ${APP_NAME}:'
	@echo
	@echo 'Usage:'
	@echo '    make build                 Compile the project.'
	@echo '    make build-alpine          Compile optimized for Alpine Linux.'
	@echo '    make package               Build final Docker image with just the Go binary inside.'
	@echo '    make tag                   Tag image created by package with latest, git commit and version.'
	@echo '    make push                  Push tagged images to registry.'
	@echo '    make helm                  Deploy to Kubernetes via Helm.'
	{% if cookiecutter.enable_auto_migration == "y" %}@echo '    make helm-migration        Run database migration via Helm.'{% endif %}
	@echo '    make run ARGS=             Run with supplied arguments.'
	@echo '    make test                  Run tests on a compiled project.'
	@echo '    make clean                 Clean the directory tree.'

	@echo

build:
	@echo "building ${APP_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X github.com/kitabisa/{{ cookiecutter.app_name }}/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/kitabisa/{{ cookiecutter.app_name }}/version.Version=${VERSION} -X github.com/kitabisa/{{ cookiecutter.app_name }}/version.Environment=${ENVIRONMENT} -X github.com/kitabisa/{{ cookiecutter.app_name }}/version.BuildDate=${BUILD_DATE}" -o bin/${APP_NAME}

package:
	@echo "building image ${APP_NAME} ${VERSION} ${GIT_COMMIT}"
	docker build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=${GIT_COMMIT}${GIT_DIRTY} -t ${REGISTRY_URL}/${APP_NAME}:local .

tag: package
	@echo "Tagging: latest ${VERSION} ${GIT_COMMIT}"
	docker tag ${REGISTRY_URL}/${APP_NAME}:local ${REGISTRY_URL}/${APP_NAME}:${GIT_COMMIT}
	docker tag ${REGISTRY_URL}/${APP_NAME}:local ${REGISTRY_URL}/${APP_NAME}:${VERSION}
	docker tag ${REGISTRY_URL}/${APP_NAME}:local ${REGISTRY_URL}/${APP_NAME}:latest

push: tag
	@echo "Pushing Docker image to registry: latest ${VERSION} ${GIT_COMMIT}"
	docker push ${REGISTRY_URL}/${APP_NAME}:${GIT_COMMIT}
	docker push ${REGISTRY_URL}/${APP_NAME}:${VERSION}
	docker push ${REGISTRY_URL}/${APP_NAME}:latest

helm:
	helm upgrade ${APP_NAME} ${ENVIRONMENT}/app --install \
		--namespace ${APP_NAME} \
		--values _infra/k8s/${ENVIRONMENT}.yaml \
		--set meta.env=${ENVIRONMENT},meta.squad=${SQUAD},meta.version=${VERSION},image.repository=${IMAGE_NAME},image.tag=${VERSION}
{% if cookiecutter.enable_auto_migration == "y" %}
helm-migration:
	helm upgrade ${APP_NAME}-migration ${ENVIRONMENT}/migration --install \
		--namespace ${APP_NAME} \
		--values _infra/k8s/${ENVIRONMENT}.yaml \
		--set meta.env=${ENVIRONMENT},meta.squad=${SQUAD},meta.version=${VERSION},image.repository=${REGISTRY_URL}/${APP_NAME},image.tag=${VERSION}
{% endif %}
run: build
	bin/${APP_NAME} ${ARGS}

test:
	go test ./...

clean:
	@test ! -e bin/${APP_NAME} || rm bin/${APP_NAME}
