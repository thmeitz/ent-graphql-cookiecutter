{%- if cookiecutter.use_docker == "y" %}
.PHONY: build build-alpine clean test help default
{%- else %}
.PHONY: build clean test help default
{%- endif %}

BIN_NAME={{cookiecutter.app_name}}

VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE_NAME := "{{cookiecutter.docker_hub_username}}/{{cookiecutter.app_name}}"

default: test

help:
	@echo 'Management commands for {{cookiecutter.app_name}}:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	{% if cookiecutter.use_docker == "y" %}@echo '    make build-alpine    Compile optimized for alpine linux.'
	@echo '    make package         Build final docker image with just the go binary inside'
	@echo '    make tag             Tag image created by package with latest, git commit and version'
	@echo '    make test            Run tests on a compiled project.'
	@echo '    make push            Push tagged images to registry'{% endif %}
	@echo '    make clean           Clean the directory tree.'
	@echo

build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME}

get-deps:
	@echo "runs go get"
	go get entgo.io/contrib@master
	go get entgo.io/ent@master
	go get ariga.io/entimport@master
	go get ariga.io/atlas@master
	go get github.com/99designs/gqlgen
	go get github.com/gin-contrib/cors
	go get github.com/gin-contrib/zap
	go get github.com/gin-gonic/gin
	go get github.com/go-sql-driver/mysql
	go get github.com/golang-migrate/migrate/v4
	go get github.com/google/addlicense
	go get github.com/hashicorp/go-multierror
	go get github.com/lib/pq
	go get github.com/mattn/go-sqlite3
	go get github.com/mitchellh/go-homedir
	go get github.com/rs/xid
	go get github.com/spf13/cobra
	go get github.com/spf13/viper
	go get github.com/stretchr/testify
	go get github.com/vektah/gqlparser
	go get github.com/vektah/gqlparser/v2
	go get github.com/vmihailenco/msgpack/v5
	go get github.com/weppos/publicsuffix-go
	go get go.uber.org/zap
	go get golang.org/x/net
	go get gopkg.in/yaml.v2

{%- if cookiecutter.use_docker == "y" %}
build-alpine:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags '-w -linkmode external -extldflags "-static" -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.BuildDate=${BUILD_DATE}' -o bin/${BIN_NAME} cli/main.go

base-image:
	@echo "building base image ${BIN_NAME} ${VERSION} $(GIT_COMMIT)"
	docker build -t $(IMAGE_NAME):${VERSION} -f Dockerfile.base .

package:
	@echo "building image ${BIN_NAME} ${VERSION} $(GIT_COMMIT)"
	docker build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=$(GIT_COMMIT) -t $(IMAGE_NAME):local cli/main.go

tag: 
	@echo "Tagging: latest ${VERSION} $(GIT_COMMIT)"
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):$(GIT_COMMIT)
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):${VERSION}
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):latest

push: tag
	@echo "Pushing docker image to registry: latest ${VERSION} $(GIT_COMMIT)"
	docker push $(IMAGE_NAME):$(GIT_COMMIT)
	docker push $(IMAGE_NAME):${VERSION}
	docker push $(IMAGE_NAME):latest
{% endif %}
clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

test:
	go test ./...
