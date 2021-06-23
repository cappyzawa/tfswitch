package command

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/cli"
)

type localListCommand struct {
	ui       *cli.ColoredUi
	dataHome string
}

func (c *localListCommand) Help() string {
	return `This command displays available versions in local.

Usage:
  tfswitch local-list
  `
}

func (c *localListCommand) Run(args []string) int {
	tfDataPATH := filepath.Join(c.dataHome, "tfswitch")
	if d, err := os.Stat(tfDataPATH); os.IsNotExist(err) || !d.IsDir() {
		c.ui.Error(err.Error())
		return 1
	}
	files, _ := ioutil.ReadDir(tfDataPATH)
	for _, f := range files {
		if f.IsDir() {
			c.ui.Output(f.Name())
		}
	}
	return 0
}

func (c *localListCommand) Synopsis() string {
	return "display available terraform versions in local."
}
