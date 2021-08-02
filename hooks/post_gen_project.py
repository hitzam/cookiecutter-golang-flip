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


def remove_dir(dirname):
    """
    generic remove directory from project dir
    """
    shutil.rmtree(os.path.join(
        PROJECT_DIRECTORY, dirname
    ))

def rename_file(filename_before, filename_after):
    """
    renaming file
    """
    fullpath_before = os.path.join(PROJECT_DIRECTORY, filename_before)
    if os.path.exists(fullpath_before):
        fullpath_after = os.path.join(PROJECT_DIRECTORY, filename_after)
        os.rename(fullpath_before, fullpath_after)


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

# 1. Remove database migration if not selected
if '{{ cookiecutter.use_migrate_migration }}'.lower() != 'y':
    remove_dir("migrations")
    remove_file("cmd/migration.go")
    remove_file(".infra/helm/dev/migration.yaml")
    remove_file(".infra/helm/stg/migration.yaml")
    remove_file(".infra/helm/uat/migration.yaml")
    remove_file(".infra/helm/prod/migration.yaml")

# 2. Remove worker if not selected
if '{{ cookiecutter.is_worker }}'.lower() != 'y':
    remove_file("cmd/worker.go")
    remove_file(".infra/helm/dev/worker.yaml")
    remove_file(".infra/helm/stg/worker.yaml")
    remove_file(".infra/helm/uat/worker.yaml")
    remove_file(".infra/helm/prod/worker.yaml")

# 3. Remove server if not selected
if '{{ cookiecutter.is_server }}'.lower() != 'y':
    remove_file(".infra/helm/dev/server.yaml")
    remove_file(".infra/helm/stg/server.yaml")
    remove_file(".infra/helm/uat/server.yaml")
    remove_file(".infra/helm/prod/server.yaml")

# 4. Initialize Git (should be run after all file have been modified or deleted)
init_git()
