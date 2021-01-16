package main

import (
	"io"
	"os"
	"runtime/debug"

	"github.com/cappyzawa/tfswitch/command"
	"github.com/mitchellh/cli"
)

const (
	tfPATH = "/usr/local/bin/terraform"
)

var (
	version string
)

type runnner struct {
	out io.Writer
	in  io.Reader
	err io.Writer

	dataHome string
}

func (r *runnner) Run(args []string) int {

	ui := &cli.ColoredUi{
		ErrorColor: cli.UiColorRed,
		WarnColor:  cli.UiColorYellow,
		Ui: &cli.BasicUi{
			Reader:      r.in,
			Writer:      r.out,
			ErrorWriter: r.err,
		},
	}

	if version == "" {
		version = getVersion()
	}
	c := cli.NewCLI(args[0], version)
	c.Args = args[1:]
	factories := command.Factories{
		UI:       ui,
		DataHome: r.dataHome,
	}
	c.Commands = map[string]cli.CommandFactory{
		"use":  factories.Use,
		"list": factories.List,
	}
	exitStatus, err := c.Run()
	if err != nil {
		ui.Error(err.Error())
	}
	return exitStatus
}

func getVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "(devel)"
	}
	return info.Main.Version
}

func main() {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		dataHome = "~/.local/share"
	}

	r := &runnner{
		out:      os.Stdout,
		err:      os.Stderr,
		in:       os.Stdin,
		dataHome: dataHome,
	}

	os.Exit(r.Run(os.Args))
}
