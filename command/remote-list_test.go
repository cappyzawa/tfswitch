package command

import (
	"net/http"
	"net/url"
	"testing"
)

func TestClientList(t *testing.T) {
	t.Parallel()

	cases := map[string]struct{}{
		"success": {},
	}

	for name := range cases {
		c := &client{
			url: &url.URL{
				Scheme: "https",
				Host:   "releases.hashicorp.com",
				Path:   "terraform",
			},
			httpClient: http.DefaultClient,
		}
		t.Run(name, func(t *testing.T) {
			versions, err := c.List()
			if err != nil {
				t.Errorf(err.Error())
			}
			if len(versions) <= 0 {
				t.Errorf("At least one or more versions should be present")
			}
		})
	}
}
