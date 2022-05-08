"""
Does the following:

1. Inits git if used
2. Deletes dockerfiles if not going to be used
3. Deletes config utils if not needed
"""
from __future__ import print_function
import os
import shutil
from subprocess import Popen

# Get the root project directory
PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)

def remove_file(filename):
    """
    generic remove file from project dir
    """
    fullpath = os.path.join(PROJECT_DIRECTORY, filename)
    if os.path.exists(fullpath):
        os.remove(fullpath)

def init_git():
    """
    Initialises git on the new project folder
    """
    GIT_COMMANDS = [
        ["git", "init"],
        ["git", "add", "."],
        ["git", "commit", "-a", "-m", "Initial Commit."]
    ]

    for command in GIT_COMMANDS:
        git = Popen(command, cwd=PROJECT_DIRECTORY)
        git.wait()


def remove_docker_files():
    """
    Removes files needed for docker if it isn't going to be used
    """
    for filename in ["Dockerfile",]:
        os.remove(os.path.join(
            PROJECT_DIRECTORY, filename
        ))

def run_go_generate():
    """
    Run go generate ./... on the new project folder to generate ent
    """
    GO_COMMANDS = [
      ["go", "get", "entgo.io/ent/cmd/ent@master"],
      ["go", "get", "entgo.io/contrib/entgql@master"],
      ["go", "get", "github.com/99designs/gqlgen@master"],
      ["go", "get", "github.com/google/addlicense"],
      ["go", "get", "github.com/rs/xid"],
      ["go", "get", "github.com/vmihailenco/msgpack/v5"], 
      ["go", "get", "github.com/hashicorp/go-multierror"],
      ["go", "get", "github.com/gin-contrib/cors"],
      ["go", "get", "github.com/gin-gonic/gin"],
      ["go", "get", "go.uber.org/zap"],
      ["go", "get", "github.com/mitchellh/go-homedir"],
      ["go", "get", "github.com/spf13/viper"],
      ["go", "generate", "./..."],
      ["go", "run", "cli/main.go", "gen-md-docs", "./docs"],

      ["go", "mod", "tidy"],
    ]
    
    for command in GO_COMMANDS:
      gorun = Popen(command, cwd=PROJECT_DIRECTORY)
      gorun.wait()

# 1. Remove Dockerfiles if docker is not going to be used
if '{{ cookiecutter.use_docker }}'.lower() != 'y':
    remove_docker_files()

# 2. Run go generate in the project
run_go_generate()

# 3. Initialize Git (should be run after all file have been modified or deleted)
if '{{ cookiecutter.use_git }}'.lower() == 'y':
    init_git()
else:
    remove_file(".gitignore")


