package di

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/cappyzawa/tfswitch/v2/internal/command"
	"github.com/cappyzawa/tfswitch/v2/internal/opentofu"
	"github.com/cappyzawa/tfswitch/v2/internal/repository"
	"github.com/cappyzawa/tfswitch/v2/internal/terraform"
	"github.com/mitchellh/cli"
)

const (
	TargetTerraform string = "terraform"
	TargetOpenTofu  string = "opentofu"
)

type Container struct {
	UI       *cli.ColoredUi
	DataHome string

	client repository.Client
}

func NewContainer(ui *cli.ColoredUi, dataHome string, target string) (*Container, error) {
	c := &Container{
		UI:       ui,
		DataHome: dataHome,
	}
	switch target {
	case TargetTerraform:
		c.client = c.terraformClient()
	case TargetOpenTofu:
		c.client = c.opentofuClient()
	default:
		return nil, errors.New("invalid target. must be either terraform or opentofu")
	}
	return c, nil
}

func (c *Container) UseCommand() (cli.Command, error) {
	return &command.UseCommand{
		UI:       c.UI,
		DataHome: c.DataHome,
		Client:   c.client,
	}, nil
}

func (c *Container) RemoteListCommand() (cli.Command, error) {
	return &command.RemoteListCommand{
		UI:     c.UI,
		Client: c.client,
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

func (c *Container) opentofuClient() repository.Client {
	return opentofu.NewClient()
}

func (c *Container) httpClient() *http.Client {
	return http.DefaultClient
}
