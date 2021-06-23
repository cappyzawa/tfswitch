package command

import (
	"net/http"

	"github.com/mitchellh/cli"
)

// Factories generates each command
type Factories struct {
	UI       *cli.ColoredUi
	DataHome string

	HttpClient *http.Client
}

// Use creates useCommand
func (f *Factories) Use() (cli.Command, error) {
	return &useCommand{
		ui:       f.UI,
		dataHome: f.DataHome,
	}, nil
}

// LocalList creates localListCommand
func (f *Factories) LocalList() (cli.Command, error) {
	return &localListCommand{
		ui:       f.UI,
		dataHome: f.DataHome,
	}, nil
}

// RemoteList creates remoteListCommand
func (f *Factories) RemoteList() (cli.Command, error) {
	return &remoteListCommand{
		ui:         f.UI,
		httpClient: f.HttpClient,
	}, nil
}
