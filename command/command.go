package command

import "github.com/mitchellh/cli"

// Factories generates each command
type Factories struct {
	UI       *cli.ColoredUi
	DataHome string
}

// Use creates useCommand
func (f *Factories) Use() (cli.Command, error) {
	return &useCommand{
		ui:       f.UI,
		dataHome: f.DataHome,
	}, nil
}

// List creates listCommand
func (f *Factories) List() (cli.Command, error) {
	return &listCommand{
		ui:       f.UI,
		dataHome: f.DataHome,
	}, nil
}
