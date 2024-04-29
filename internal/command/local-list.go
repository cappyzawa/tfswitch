package command

import (
	"flag"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/cli"
)

type LocalListCommand struct {
	UI       *cli.ColoredUi
	DataHome string
}

func (c *LocalListCommand) Help() string {
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

func (c *LocalListCommand) Run(args []string) int {
	var filter string
	flags := flag.NewFlagSet("", flag.ExitOnError)
	flags.StringVar(&filter, "filter", "", "Filter by the specified version (Prefix Match)")
	if err := flags.Parse(args); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	tfDataPATH := filepath.Join(c.DataHome, "tfswitch")
	if d, err := os.Stat(tfDataPATH); os.IsNotExist(err) || !d.IsDir() {
		c.UI.Error(err.Error())
		return 1
	}
	files, _ := os.ReadDir(tfDataPATH)
	for _, f := range files {
		if f.IsDir() && strings.HasPrefix(f.Name(), filter) {
			c.UI.Output(f.Name())
		}
	}
	return 0
}

func (c *LocalListCommand) Synopsis() string {
	return "display available terraform versions in local."
}
