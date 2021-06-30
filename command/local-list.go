package command

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/cli"
)

type localListCommand struct {
	ui       *cli.ColoredUi
	dataHome string
}

func (c *localListCommand) Help() string {
	return `This command displays available versions in local.

Usage:
  tfswitch local-list [--filter=VERSION]

Options:
  --filter    Filter by the specified version (Prefix Match)

Examples:
  tfswitch local-list
  tfswitch local-list --filter 1.0.0
  tfswitch local-list --filter 1.0
  tfswitch local-list --filter 1
  `
}

func (c *localListCommand) Run(args []string) int {
	var filter string
	flags := flag.NewFlagSet("", flag.ExitOnError)
	flags.StringVar(&filter, "filter", "", "Filter by the specified version (Prefix Match)")
	if err := flags.Parse(args); err != nil {
		c.ui.Error(err.Error())
		return 1
	}

	tfDataPATH := filepath.Join(c.dataHome, "tfswitch")
	if d, err := os.Stat(tfDataPATH); os.IsNotExist(err) || !d.IsDir() {
		c.ui.Error(err.Error())
		return 1
	}
	files, _ := ioutil.ReadDir(tfDataPATH)
	for _, f := range files {
		if f.IsDir() && strings.HasPrefix(f.Name(), filter) {
			c.ui.Output(f.Name())
		}
	}
	return 0
}

func (c *localListCommand) Synopsis() string {
	return "display available terraform versions in local."
}
