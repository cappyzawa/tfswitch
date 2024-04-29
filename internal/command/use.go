package command

import (
	"fmt"
	"os"

	"github.com/cappyzawa/tfswitch/v2/internal/repository"
	"github.com/mitchellh/cli"
)

const (
	tfPATH = "/usr/local/bin/terraform"
)

// UseCommand describes "version" command
type UseCommand struct {
	UI       *cli.ColoredUi
	DataHome string

	Client repository.Client
}

func (c *UseCommand) Help() string {
	return `This command switches the terraform version.

In case of missing specified version is missing in local, this command install the terraform binary from https://releases.hashicorp.com/terraform/ before switching.

Usage:
  tfswitch use VERSION

Examples:
  tfswitch use 0.14.4
  `
}

func (c *UseCommand) Run(args []string) int {
	version := args[0]
	execPATH, err := c.Client.Install(c.DataHome, version)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	if err := updateSymlink(execPATH, tfPATH); err != nil {
		c.UI.Error(err.Error())
		return 2
	}
	c.UI.Output(fmt.Sprintf("Switched terraform version to %s", version))
	return 0
}

func updateSymlink(oldname, newname string) error {
	if _, err := os.Lstat(newname); err == nil {
		if err := os.Remove(newname); err != nil {
			return err
		}
	}
	if err := os.Symlink(oldname, newname); err != nil {
		return err
	}
	return nil
}

func (c *UseCommand) Synopsis() string {
	return "use specified terraform version."
}
