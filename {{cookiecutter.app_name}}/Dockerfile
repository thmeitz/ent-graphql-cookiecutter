# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

LABEL app="build-{{cookiecutter.app_name}}"
LABEL REPO="https://github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}"

WORKDIR /app

ADD . /app
RUN make build-alpine
# RUN go mod download
# RUN go mod download github.com/apparentlymart/go-textseg
# RUN go mod download github.com/ugorji/go
# RUN go mod download github.com/vmihailenco/msgpack
# RUN go mod download github.com/vmihailenco/tagparser

# RUN go build  -o /{{cookiecutter.app_name}} /app/cli/main.go



##
## Deploy
##
FROM gcr.io/distroless/base-debian10

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

WORKDIR /

COPY --from=build /{{cookiecutter.app_name}} /{{cookiecutter.app_name}}

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/{{cookiecutter.app_name}}"]