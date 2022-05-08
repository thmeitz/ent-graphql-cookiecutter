# ent-graphql-cookiecutter

Powered by [Cookiecutter](https://github.com/audreyr/cookiecutter). This is a fork of [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang).

# WARNING: really early development stage

This is not production ready and many things do not yet work as they should. As an example, the whole Docker configuration and make commands!

## Features

- `Makefile` with management commands
- Go 1.18 version
- uses `go mod`
- uses [gin-gonic](https://github.com/gin-gonic/gin)
- uses [zap](https://github.com/uber-go/zap) logger
- injects build time and git hash at build time

### Datbase Migrations
- uses go-migrate Ariga Atlas integration

## Optional Integrations

- Can (currently not) create dockerfile for building go binary and dockerfile for final go binary (no code in final container)
- If docker is used adds docker management commands to makefile

## Constraints

- Uses `mod` for dependency management

## Usage

Let's pretend you want to create a project called "echoserver". Rather than starting from scratch maybe copying 
some files and then editing the results to include your name, email, and various configuration issues that always 
get forgotten until the worst possible moment, get cookiecutter to do all the work.

First, get Cookiecutter. Trust me, it's awesome:
```console
$ pip install cookiecutter
```

Alternatively, you can install `cookiecutter` with homebrew:
```console
$ brew install cookiecutter
```

Finally, to run it based on this template, type:
```console
$ cookiecutter https://github.com/thmeitz/ent-graphql-cookiecutter.git
```

You will be asked about your basic info (name, project name, app name, etc.). This info will be used to customize your new project.

Warning: After this point, change 'Thomas Meitz', 'thmeitz', etc to your own information.

