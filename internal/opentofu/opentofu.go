package opentofu

import (
	"errors"

	"github.com/cappyzawa/tfswitch/v2/internal/repository"
)

var _ repository.Client = (*opentofu)(nil)

type opentofu struct{}

func (o *opentofu) Name() string {
	return "opentofu"
}

func (o *opentofu) Versions() ([]string, error) {
	return nil, errors.New("opentofu: not implemented")
}

func (o *opentofu) Install(dataHome string, version string) (string, error) {
	return "", errors.New("opentofu: not implemented")
}
