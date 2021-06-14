# cookiecutter-golang

Powered by [Cookiecutter](https://github.com/audreyr/cookiecutter), Cookiecutter Golang is a framework for jumpstarting production-ready go projects quickly.

## Features

- Generous `Makefile` with management commands.
- Uses `go modules`.
- injects build time and git hash at build time.

## Usage

Let's pretend you want to create a project called "golang-services". Rather than starting from scratch maybe copying
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
$ cookiecutter https://github.com/kitabisa/cookiecutter-golang.git
```

You will be asked about your basic info (name, project name, app name, etc.). This info will be used to customize your new project.

Warning: After this point, change 'Kitabisa', 'golang-services', etc to your own information.

Answer the prompts with your own desired options. For example:
```console
app_name [golangproject]: golang-services
Select squad:
1 - backend
2 - infra
3 - frontend
Choose from 1, 2, 3 (1, 2, 3) [1]: 1
business_unit [platform]: platform
use_migrate_migration [n]: n
use_rabbitmq [n]: n
is_worker [n]: n
is_server [y]: y
Select connectivity:
1 - public
2 - private
Choose from 1, 2 (1, 2) [1]: 1
enable_uat [n]: n
enable_dev [n]: n
```

Enter the project and take a look around:
```console
$ cd golang-services/
$ ls
```

Run `make help` to see the available management commands, or just run `make build` to build your project.
```console
$ make help
$ make build
$ ./bin/golang-services
```
