package terraform

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/cappyzawa/tfswitch/v2/internal/repository"
	goversion "github.com/hashicorp/go-version"
	install "github.com/hashicorp/hc-install"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/hc-install/src"
)

var _ repository.Client = (*terraform)(nil)

type terraform struct {
	url        *url.URL
	httpClient *http.Client
}

func NewClient(url *url.URL, httpClient *http.Client) repository.Client {
	return &terraform{
		url:        url,
		httpClient: httpClient,
	}
}

func (t *terraform) Name() string {
	return "terraform"
}

func (t *terraform) Versions() ([]string, error) {
	req, err := http.NewRequest(http.MethodGet, t.url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "tfswitch")
	res, err := t.httpClient.Do(req)
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

func (t *terraform) Install(dataHome string, version string) (string, error) {
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
