package command

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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
		version := strings.ReplaceAll(f.Name(), "terraform_", "")
		c.ui.Output(version)
	}
	return 0
}

func (c *listCommand) Synopsis() string {
	return "desplay available terraform versions in local."
}
