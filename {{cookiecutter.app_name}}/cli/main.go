package main

import (
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/cli/cmd"
	_ "github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/ent/runtime"
)

func main() {
	cmd.Execute()
}
