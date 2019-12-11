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


# remove buroq related file/folders if not selected
if '{{ cookiecutter.use_buroq }}'.lower() != 'y':
    remove_dir("api")
    remove_file("cmd/root_buroq.go")
    remove_file("cmd/migration.go")
    remove_file("cmd/config_buroq.go")
    remove_file("go_buroq.mod")
    remove_dir("internal")
    remove_dir("migrations/seeds")
    remove_dir("migrations/sql")
    remove_dir("params")
    remove_dir("pkg")

if  '{{ cookiecutter.use_buroq }}'.lower() == 'y':
    remove_file("cmd/root.go")
    remove_file("cmd/config.go")
    remove_file("go.mod")
    remove_dir("log")

    rename_file("cmd/root_buroq.go", "cmd/root.go")
    rename_file("cmd/config_buroq.go", "cmd/config.go")
    rename_file("go_buroq.mod", "go.mod")

# 1. Remove viper config if not seleted
if '{{ cookiecutter.use_viper_config }}'.lower() != 'y' and '{{ cookiecutter.use_buroq }}'.lower() != 'y':
    remove_dir("config")

# 2. Remove logrus utils if not seleted
if '{{ cookiecutter.use_logrus_logging }}'.lower() != 'y' and '{{ cookiecutter.use_buroq }}'.lower() != 'y':
    remove_dir("log")

# 3. Remove cobra utils if not seleted
if '{{ cookiecutter.use_cobra_cmd }}'.lower() != 'y' and '{{ cookiecutter.use_buroq }}'.lower() != 'y':
    remove_dir("cmd")

# 4. Remove migrate utils if not selected
if '{{ cookiecutter.use_migrate_migration }}'.lower() != 'y' and '{{ cookiecutter.use_buroq }}'.lower() != 'y':
    remove_dir("migrations")

# 5. Remove uat values if not enabled
if '{{ cookiecutter.enable_uat }}'.lower() != 'y':
    remove_file("_infra/k8s/uat.yaml")

# 6. Remove dev values if not enabled
if '{{ cookiecutter.enable_dev }}'.lower() != 'y':
    remove_file("_infra/k8s/dev.yaml")

# 7. Initialize Git (should be run after all file have been modified or deleted)
init_git()
