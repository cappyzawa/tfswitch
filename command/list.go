package command

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/cli"
)

type listCommand struct {
	ui       *cli.ColoredUi
	dataHome string
}

func (c *listCommand) Help() string {
	return `This command desplays available versions in local.

Usage:
  tfswitch list
  `
}

func (c *listCommand) Run(args []string) int {
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

func (c *listCommand) Synopsis() string {
	return "desplay available terraform versions in local."
}
