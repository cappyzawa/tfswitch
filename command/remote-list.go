package command

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mitchellh/cli"
)

type remoteListCommand struct {
	ui         *cli.ColoredUi
	httpClient *http.Client
}

func (c *remoteListCommand) Help() string {
	return `This command displays available versions in remote. (https://releases.hashicorp.com/terraform/)

Usage:
  tfswitch remote-list
  `
}

func (c *remoteListCommand) Run(args []string) int {
	cc := &client{
		url: &url.URL{
			Scheme: "https",
			Host:   "releases.hashicorp.com",
			Path:   "terraform",
		},
		httpClient: c.httpClient,
	}
	versions, err := cc.List()
	if err != nil {
		c.ui.Error(err.Error())
		return 1
	}
	for _, v := range versions {
		c.ui.Output(v)
	}
	return 0
}

type client struct {
	url        *url.URL
	httpClient *http.Client
}

func (c *client) List() ([]string, error) {
	req, err := http.NewRequest(http.MethodGet, c.url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "tfswitch")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var versions []string
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		v := strings.TrimPrefix(strings.TrimSpace(s.Text()), "terraform_")
		if v != "../" {
			versions = append(versions, v)
		}
	})
	return versions, nil
}

func (c *remoteListCommand) Synopsis() string {
	return "display available terraform versions in remote. (https://releases.hashicorp.com/terraform/)"
}
