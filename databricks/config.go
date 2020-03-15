package databricks

import (
	"fmt"
	"net/url"

	"github.com/go-openapi/runtime"
	openapiClient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	cleanhttp "github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform/helper/logging"
	"github.com/innovationnorway/go-databricks/plumbing"
)

type Config struct {
	Token string
	Host  string
}

type Meta struct {
	Databricks *plumbing.Databricks
	AuthInfo   runtime.ClientAuthInfoWriter
}

func (c *Config) Client() (interface{}, error) {
	u, err := url.Parse(c.Host)
	if err != nil {
		return nil, fmt.Errorf("Error parsing host: %s", err)
	}

	if u.Scheme == "" {
		u.Scheme = "https"
	}

	if u.Path == "" {
		u.Path = plumbing.DefaultBasePath
	}

	client := openapiClient.NewWithClient(
		u.Host, u.Path, []string{u.Scheme},
		cleanhttp.DefaultClient())
	client.Transport = logging.NewTransport("Databricks", client.Transport)

	authInfo := runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		r.SetHeaderParam("User-Agent", "Terraform")
		r.SetHeaderParam("Authorization", "Bearer "+c.Token)
		return nil
	})

	return &Meta{
		Databricks: plumbing.New(client, strfmt.Default),
		AuthInfo:   authInfo,
	}, nil
}
