package main

import (
	"flag"
	"io"
	"os"
	"runtime/debug"

	"github.com/cappyzawa/tfswitch/v2/internal/di"
	"github.com/cappyzawa/tfswitch/v2/internal/flags"
	"github.com/mitchellh/cli"
)

var version string

type runnner struct {
	out io.Writer
	in  io.Reader
	err io.Writer

	dataHome string
}

func (r *runnner) Run(args []string) int {
	var globalFlags flags.Global
	flag.StringVar(&globalFlags.Target, "target", "terraform", "Target version")
	flag.Parse()

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
	c.Args = flag.Args()
	dc, err := di.NewContainer(ui, r.dataHome, globalFlags.Target)
	if err != nil {
		ui.Error(err.Error())
		return 1
	}
	c.Commands = map[string]cli.CommandFactory{
		"use":         dc.UseCommand,
		"local-list":  dc.LocalListCommand,
		"remote-list": dc.RemoteListCommand,
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
	home := os.Getenv("HOME")
	if dataHome == "" {
		dataHome = home + "/.local/share"
	}

	r := &runnner{
		out:      os.Stdout,
		err:      os.Stderr,
		in:       os.Stdin,
		dataHome: dataHome,
	}

	os.Exit(r.Run(os.Args))
}
