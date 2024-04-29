package command

import (
	"flag"
	"strings"

	"github.com/cappyzawa/tfswitch/v2/internal/flags"
	"github.com/cappyzawa/tfswitch/v2/internal/repository"
	"github.com/mitchellh/cli"
)

type RemoteListCommand struct {
	UI     *cli.ColoredUi
	Client repository.Client
}

func (c *RemoteListCommand) Help() string {
	return `This command displays available versions in remote. (https://releases.hashicorp.com/terraform/)

Usage:
  tfswitch remote-list [--filter=VERSION]

Options:
  --filter    Filter by the specified version (Prefix Match)

Examples:
  tfswitch remote-list
  tfswitch remote-list --filter 1.0.0
  tfswitch remote-list --filter 1.0
  tfswitch remote-list --filter 1
  `
}

func (c *RemoteListCommand) Run(args []string) int {
	var f flags.RemoteList
	fs := flag.NewFlagSet("remote-list", flag.ContinueOnError)
	fs.StringVar(&f.Filter, "filter", "", "Filter by the specified version (Prefix Match)")
	if err := fs.Parse(args); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	versions, err := c.Client.Versions()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	for _, v := range versions {
		if strings.HasPrefix(v, f.Filter) {
			c.UI.Output(v)
		}
	}
	return 0
}

func (c *RemoteListCommand) Synopsis() string {
	return "display available terraform versions in remote. (https://releases.hashicorp.com/terraform/)"
}
