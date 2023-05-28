# go-base-gen
![workflow status](https://github.com/dung13890/go-base-gen/actions/workflows/go-ci.yml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)


## Overview
Command line tool to generate a project from a template. It is a tool that helps you to init code base and start a new project quickly.

```bash
NAME:
   go-base-gen - Use this tool to generate base code

USAGE:
   go-base-gen [global options] command [command options] [arguments...]

VERSION:
   v1.0.0

COMMANDS:
   project  Generate base code for go project use clean architecture
   domain   Create new domain in project
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print only the version (default: false)
```

## Features
- [x] Generate project base on clean architecture
- [x] Generate module


## Installation
```bash
go install github.com/dung13890/go-base-gen@latest
```

## Using
- Init Project base on clean architecture
```bash
## Short
go-base-gen project -n <project-name>

## Long
go-base-gen project --name <project-name> --path <project-path>
```

- Create new domain
```bash
## Short
go-base-gen domain -n <domain-name> -pj <project-name> -m <module-name>

## Long
go-base-gen domain --name <domain-name> --project <project-name> --module <module-name> --path <project-path>
```

## Structure project after generate
```
.
├── cmd
│   ├── app
│   ├── migrate
│   └── seed
├── config
├── db
│   ├── migrations
│   └── seeds
├── internal
│   ├── app
│   ├── constants
│   ├── domain
│   │   └── auth
│   │       ├── delivery
│   │       ├── repository
│   │       └── usecase
│   ├── impl
│   │   ├── pubsub
│   │   └── service
│   ├── modules
│   └── registry
└── pkg
```

You can see more detail in [go-clean-architecture](https://github.com/dung13890/go-clean-architecture)

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/dung13890)

