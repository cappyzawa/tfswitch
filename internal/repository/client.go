package repository

type Client interface {
	Name() string
	Versions() ([]string, error)
	Install(dataHome string, version string) (string, error)
}
