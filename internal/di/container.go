package di

import (
	"net/http"
	"net/url"

	"github.com/cappyzawa/tfswitch/v2/internal/command"
	"github.com/cappyzawa/tfswitch/v2/internal/repository"
	"github.com/cappyzawa/tfswitch/v2/internal/terraform"
	"github.com/mitchellh/cli"
)

type Container struct {
	UI       *cli.ColoredUi
	DataHome string
}

func NewContainer(ui *cli.ColoredUi, dataHome string) *Container {
	return &Container{
		UI:       ui,
		DataHome: dataHome,
	}
}

func (c *Container) UseCommand() (cli.Command, error) {
	return &command.UseCommand{
		UI:       c.UI,
		DataHome: c.DataHome,
		Client:   c.terraformClient(),
	}, nil
}

func (c *Container) RemoteListCommand() (cli.Command, error) {
	return &command.RemoteListCommand{
		UI:     c.UI,
		Client: c.terraformClient(),
	}, nil
}

func (c *Container) LocalListCommand() (cli.Command, error) {
	return &command.LocalListCommand{
		UI:       c.UI,
		DataHome: c.DataHome,
	}, nil
}

func (c *Container) terraformClient() repository.Client {
	u := &url.URL{
		Scheme: "https",
		Host:   "releases.hashicorp.com",
		Path:   "terraform",
	}
	hc := c.httpClient()
	return terraform.NewClient(u, hc)
}

func (c *Container) httpClient() *http.Client {
	return http.DefaultClient
}
