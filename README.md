# bbpack-info

## Description
Reviewing packages within all recipes in Yocto is not the easiest task. The `bbpak` tool helps you to extract all the information from the packages within a Yocto build, and query all the information about packages of the currently compiled image, including patches tracking.

## Usage
This is straightforward: install it somewhere in your `/usr/local/bin` or `$HOME/bin` and simply invoke it.

```
NAME:
   bbpak - Query installed packages in Yocto's BitBake

USAGE:
   bbpak [global options] command [command options] [arguments...]

VERSION:
   0.1 Alpha

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --path value, -p value      Path to the build
   --manifest value, -m value  Name of the manifest to reference package index
   --format value, -f value    Output in: csv, md, json, txt (default: "txt")
   --list, -l                  List available manifests (default: false)
   --help, -h                  show help (default: false)
   --version, -v               print the version (default: false)
```

Todo:

- View single package thorough details
- Patches review

## Installation

You should compile it on your own or package it for your distribution of the day. :wink:

1. `apt-get install golang-go`
2. `git clone https://github.com/isbm/bbpack-info.git`
3. `cd bbpack-info/cmd`
4. `make`

Enjoy.

## Contribution

Sure. Your PR is always welcome.
