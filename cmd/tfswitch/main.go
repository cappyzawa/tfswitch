package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-exec/tfinstall"
)

const (
	tfPATH = "/usr/local/bin/terraform"
)

type cli struct {
	out io.Writer
	err io.Writer

	dataHome string
}

func (c *cli) Run(args []string) int {
	if len(args) != 2 {
		return c.exitErr(fmt.Errorf("usage: %s [version]", args[0]))
	}
	version := args[1]
	execPATH, err := c.installOrExtractTerraform(version)
	if err != nil {
		return c.exitErr(err)
	}
	if _, err := os.Lstat(tfPATH); err == nil {
		if err := os.Remove(tfPATH); err != nil {
			return c.exitErr(err)
		}
	}
	if err := c.updateSymlink(execPATH, tfPATH); err != nil {
		return c.exitErr(err)
	}
	fmt.Fprintf(c.out, "Switched terraform version to %s\n", version)
	return 0
}

func (c *cli) installOrExtractTerraform(version string) (string, error) {
	tfDataPATH := filepath.Join(c.dataHome, "tfswitch")
	if err := os.MkdirAll(tfDataPATH, 0755); err != nil {
		return "", err
	}
	execPATH, err := tfinstall.Find(context.Background(), tfinstall.ExactVersion(version, tfDataPATH))
	if err != nil {
		return "", err
	}
	renamedExecPATH := fmt.Sprintf("%s_%s", execPATH, version)
	if err := os.Rename(execPATH, renamedExecPATH); err != nil {
		return "", err
	}
	return renamedExecPATH, nil
}

func (c *cli) updateSymlink(oldname, newname string) error {
	if _, err := os.Lstat(newname); err == nil {
		if err := os.Remove(newname); err != nil {
			return err
		}
	}
	if err := os.Symlink(oldname, newname); err != nil {
		return err
	}
	return nil
}

func (c *cli) exitErr(err error) int {
	fmt.Fprintf(c.err, "%v\n", err)
	return 1
}

func main() {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		dataHome = "~/.local/share"
	}

	c := &cli{
		out:      os.Stdout,
		err:      os.Stderr,
		dataHome: dataHome,
	}

	os.Exit(c.Run(os.Args))
}
