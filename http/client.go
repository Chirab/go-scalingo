package http

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"gopkg.in/errgo.v1"
)

const (
	AuthAPI     = "AUTHENTICATION_API"
	ScalingoAPI = "SCALINGO_API"
	DBAPI       = "DATABASES_API"
)

var apisConfig = map[string]apiConfig{
	AuthAPI: {
		DefaultEndpoint: "https://auth.scalingo.com",
		EnvironmentKey:  "SCALINGO_AUTH_URL",
	},
	ScalingoAPI: {
		DefaultEndpoint: "https://api.scalingo.com",
		EnvironmentKey:  "SCALINGO_API_URL",
	},
	DBAPI: {
		DefaultEndpoint: "https://db-api.scalingo.com",
		EnvironmentKey:  "SCALINGO_DB_API",
	},
}

type apiConfig struct {
	DefaultEndpoint string
	EnvironmentKey  string
}

type Client interface {
	ResourceList(resource string, payload, data interface{}) error
	ResourceAdd(resource string, payload, data interface{}) error
	ResourceGet(resource, resourceID string, payload, data interface{}) error
	ResourceUpdate(resource, resourceID string, payload, data interface{}) error
	ResourceDelete(resource, resourceID string) error

	SubresourceList(resource, resourceID, subresource string, payload, data interface{}) error
	SubresourceAdd(resource, resourceID, subresource string, payload, data interface{}) error
	SubresourceGet(resource, resourceID, subresource, id string, payload, data interface{}) error
	SubresourceUpdate(resource, resourceID, subresource, id string, payload, data interface{}) error
	SubresourceDelete(resource, resourceID, subresource, id string) error
	DoRequest(req *APIRequest, data interface{}) error
	Do(req *APIRequest) (*http.Response, error)

	TokenGenerator() TokenGenerator
	BaseURL() string
	HTTPClient() *http.Client
}

type ClientConfig struct {
	Timeout        time.Duration
	TLSConfig      *tls.Config
	APIVersion     string
	TokenGenerator TokenGenerator
}

type client struct {
	tokenGenerator TokenGenerator
	endpoint       string
	apiConfig      apiConfig
	apiVersion     string
	httpClient     *http.Client
}

func NewClient(api string, cfg ClientConfig) Client {
	if cfg.Timeout == 0 {
		cfg.Timeout = 30 * time.Second
	}
	if cfg.TLSConfig == nil {
		cfg.TLSConfig = &tls.Config{}
	}

	if cfg.APIVersion == "" {
		cfg.APIVersion = defaultAPIVersion
	}

	config := apisConfig[api]

	c := client{
		tokenGenerator: cfg.TokenGenerator,
		apiVersion:     cfg.APIVersion,
		apiConfig:      config,
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
			Transport: &http.Transport{
				Proxy:           http.ProxyFromEnvironment,
				TLSClientConfig: cfg.TLSConfig,
			},
		},
	}

	return &c
}

func (c *client) ResourceGet(resource, resourceID string, payload, data interface{}) error {
	return c.DoRequest(&APIRequest{
		Method:   "GET",
		Endpoint: "/" + resource + "/" + resourceID,
		Params:   payload,
	}, data)
}

func (c *client) ResourceList(resource string, payload, data interface{}) error {
	return c.DoRequest(&APIRequest{
		Method:   "GET",
		Endpoint: "/" + resource,
		Params:   payload,
	}, data)
}

func (c *client) ResourceAdd(resource string, payload, data interface{}) error {
	return c.DoRequest(&APIRequest{
		Method:   "POST",
		Endpoint: "/" + resource,
		Expected: Statuses{201},
		Params:   payload,
	}, data)
}

func (c client) ResourceUpdate(resource, resourceID string, payload, data interface{}) error {
	return c.DoRequest(&APIRequest{
		Method:   "PATCH",
		Endpoint: "/" + resource + "/" + resourceID,
		Params:   payload,
	}, data)
}

func (c *client) ResourceDelete(resource, resourceID string) error {
	return c.DoRequest(&APIRequest{
		Method:   "DELETE",
		Endpoint: "/" + resource + "/" + resourceID,
		Expected: Statuses{204},
	}, nil)
}

func (c *client) SubresourceGet(resource, resourceID, subresource, id string, payload, data interface{}) error {
	return c.DoRequest(&APIRequest{
		Method:   "GET",
		Endpoint: "/" + resource + "/" + resourceID + "/" + subresource + "/" + id,
		Params:   payload,
	}, data)
}

func (c *client) SubresourceList(resource, resourceID, subresource string, payload, data interface{}) error {
	return c.DoRequest(&APIRequest{
		Method:   "GET",
		Endpoint: "/" + resource + "/" + resourceID + "/" + subresource,
		Params:   payload,
	}, data)
}

func (c *client) SubresourceAdd(resource, resourceID, subresource string, payload, data interface{}) error {
	return c.DoRequest(&APIRequest{
		Method:   "POST",
		Endpoint: "/" + resource + "/" + resourceID + "/" + subresource,
		Expected: Statuses{201},
		Params:   payload,
	}, data)
}

func (c *client) SubresourceDelete(resource, resourceID, subresource, id string) error {
	return c.DoRequest(&APIRequest{
		Method:   "DELETE",
		Endpoint: "/" + resource + "/" + resourceID + "/" + subresource + "/" + id,
		Expected: Statuses{204},
	}, nil)
}

func (c *client) SubresourceUpdate(resource, resourceID, subresource, id string, payload, data interface{}) error {
	return c.DoRequest(&APIRequest{
		Method:   "PATCH",
		Endpoint: "/" + resource + "/" + resourceID + "/" + subresource + "/" + id,
		Params:   payload,
	}, data)
}

func (c *client) DoRequest(req *APIRequest, data interface{}) error {
	res, err := c.Do(req)
	if err != nil {
		return errgo.Mask(err, errgo.Any)
	}
	defer res.Body.Close()

	if data == nil {
		return nil
	}

	err = ParseJSON(res, data)
	if err != nil {
		return errgo.NoteMask(err, "fail to parse json of subresource response")
	}
	return nil
}

func (c *client) TokenGenerator() TokenGenerator {
	return c.tokenGenerator
}

func (c *client) BaseURL() string {
	endpoint := c.apiConfig.DefaultEndpoint
	if os.Getenv(c.apiConfig.EnvironmentKey) != "" {
		endpoint = os.Getenv(c.apiConfig.EnvironmentKey)
	}

	if c.apiVersion != "" {
		return fmt.Sprintf("%s/v%s", endpoint, c.apiVersion)
	}
	return endpoint
}

func (c *client) HTTPClient() *http.Client {
	return c.httpClient
}
