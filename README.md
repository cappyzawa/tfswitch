# tfswitch

[![BuildStatus](https://github.com/cappyzawa/tfswitch/workflows/CI/badge.svg)](https://github.com/cappyzawa/tfswitch/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/cappyzawa/tfswitch)](https://goreportcard.com/report/github.com/cappyzawa/tfswitch)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/cappyzawa/tfswitch)](https://pkg.go.dev/github.com/cappyzawa/tfswitch)

## Motivation

[tfutils/tfenv: Terraform version manager](https://github.com/tfutils/tfenv) is great tool for switching terraform version used in local.

However, There is an issue that it takes extra time to execute the command. (This is reffered by [terraform executions delayed by ~ 1 second · Issue \#196 · tfutils/tfenv](https://github.com/tfutils/tfenv/issues/196)).

This tool can also switch terraform version. This tool runs terraform binary from https://releases.hashicorp.com/terraform/ directly. (`tfenv` runs own script instead of binary.)

## How to use

```bash
tfswitch [version]
# e.g., If you want to use 0.14.4
tfswitch 0.14.4
```

### Available Environments

* `XDG_DATA_HOME` (Default: `~/.local/share`): The specified version of `terraform` will be saved as `$XDG_DATA_HOME/tfswitch/terraform_[version]`
