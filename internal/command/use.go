package command

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	goversion "github.com/hashicorp/go-version"
	install "github.com/hashicorp/hc-install"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/hc-install/src"
	"github.com/mitchellh/cli"
)

const (
	tfPATH = "/usr/local/bin/terraform"
)

// useCommand describes "version" command
type useCommand struct {
	ui       *cli.ColoredUi
	dataHome string
}

func (c *useCommand) Help() string {
	return `This command switches the terraform version.

In case of missing specified version is missing in local, this command install the terraform binary from https://releases.hashicorp.com/terraform/ before switching.

Usage:
  tfswitch use VERSION

Examples:
  tfswitch use 0.14.4
  `
}

func (c *useCommand) Run(args []string) int {
	version := args[0]
	execPATH, err := installOrExtractTerraform(c.dataHome, version)
	if err != nil {
		c.ui.Error(err.Error())
		return 1
	}
	if err := updateSymlink(execPATH, tfPATH); err != nil {
		c.ui.Error(err.Error())
		return 2
	}
	c.ui.Output(fmt.Sprintf("Switched terraform version to %s", version))
	return 0
}

func installOrExtractTerraform(dataHome string, version string) (string, error) {
	// e.g., tfDataPATH = $HOME/.local/share/tfswitch/0.14.4
	tfDataPATH := filepath.Join(dataHome, "tfswitch", version)
	if err := os.MkdirAll(tfDataPATH, 0755); err != nil {
		return "", err
	}
	installer := install.NewInstaller()
	tfVer := goversion.Must(goversion.NewVersion(version))
	execPATH, err := installer.Ensure(context.Background(), []src.Source{
		&releases.ExactVersion{
			Product:    product.Terraform,
			Version:    tfVer,
			InstallDir: tfDataPATH,
		},
	})
	if err != nil {
		return "", err
	}
	return execPATH, nil
}

func updateSymlink(oldname, newname string) error {
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

func (c *useCommand) Synopsis() string {
	return "use specified terraform version."
}
