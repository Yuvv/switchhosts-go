# switchhosts-go

A simple command line tool to switch your hosts

## Usage

```
$ shs

NAME:
   SwitchHosts! - Switch your hosts!

USAGE:
   shs.exe [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
   Yuvv <yuvv_th@outlook.com>

COMMANDS:
     add, a         Add new hosts
     edit, e        Edit existed hosts
     del, d         Delete new hosts
     list, l, ls    List existed hosts
     switch, s, sw  Switch to hosts
     view, v        View current hosts
     help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --global, -g   Global hosts configuration
   --help, -h     show help
   --version, -v  print the version
```

## Build

Clone this project:

```bash
git clone https://github.com/Yuvv/switchhosts-go.git
```

`cd` to project directory, and then compile it:

```bash
cd switchhosts-go

go build -o bin/shs
# go build -o bin/shs.exe   # on widows
```