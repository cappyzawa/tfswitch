# tfswitch

[![BuildStatus](https://github.com/cappyzawa/tfswitch/actions/workflows/ci.yml/badge.svg)](https://github.com/cappyzawa/tfswitch/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/cappyzawa/tfswitch)](https://goreportcard.com/report/github.com/cappyzawa/tfswitch)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/cappyzawa/tfswitch)](https://pkg.go.dev/github.com/cappyzawa/tfswitch)

## Motivation

[tfutils/tfenv: Terraform version manager](https://github.com/tfutils/tfenv) is great tool for switching the terraform version used in local.

However, There is an issue that it takes extra time to execute the command. (This is referred by [terraform executions delayed by ~ 1 second · Issue \#196 · tfutils/tfenv](https://github.com/tfutils/tfenv/issues/196)). If you want to display the terraform version in the prompt, you will have to wait almost a second every time the it is updated.

This tool can also switch the terraform version. This installs binary from https://releases.hashicorp.com/terraform/ if specified version does not found in your machine. 
`tfenv` runs own script instead of binary, but this tool runs binary directly.

## How to use

### Install

```bash
GO111MODULE=on go get github.com/cappyzawa/tfswitch/cmd/tfswitch
```

or download from [Releases · cappyzawa/tfswitch](https://github.com/cappyzawa/tfswitch/releases).

### Use
```bash
$ tfswitch -h
Usage: tfswitch [--version] [--help] <command> [<args>]

Available commands are:
    local-list     display available terraform versions in local.
    remote-list    display available terraform versions in remote. (https://releases.hashicorp.com/terraform/)
    use            use specified terraform version.
```

```bash
$ tfswitch use [version]
# e.g., If you want to use 0.14.4
$ tfswitch use 0.14.4
Switched terraform version to 0.14.4

# This tool creates or replaces symbolic links.
$ ls -l `which terraform`
lrwxr-xr-x  1 cappyzawa  admin  54 Jan 14 20:14 /usr/local/bin/terraform@ -> /Users/cappyzawa/.local/share/tfswitch/terraform_0.14.4
$ tfswitch 0.14.3
$ ls -l `which terraform`
lrwxr-xr-x  1 cappyzawa  admin  54 Jan 14 20:18 /usr/local/bin/terraform@ -> /Users/cappyzawa/.local/share/tfswitch/terraform_0.14.3
```

#### Available Environments

* `XDG_DATA_HOME` (Default: `~/.local/share`): The specified version of `terraform` will be saved as `$XDG_DATA_HOME/tfswitch/terraform_[version]`
