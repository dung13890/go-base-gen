# go-base-gen
![workflow status](https://github.com/dung13890/go-base-ge/actions/workflows/go-ci.yml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


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
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print only the version (default: false)

```

## Features
- [x] Generate project base on clean architecture
- [x] Generate module


## Structure
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

## Installation
