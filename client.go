package scalingo

import (
	"crypto/tls"
	"net/http"
	"time"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type backendConfiguration struct {
	TokenGenerator TokenGenerator
	Endpoint       string
	TLSConfig      *tls.Config
	APIVersion     string
	httpClient     HTTPClient
}

type Client struct {
	*AddonsClient
	*AddonProvidersClient
	*AppsClient
	*CollaboratorsClient
	*DeploymentsClient
	*DomainsClient
	*EventsClient
	*KeysClient
	*LoginClient
	*LogsArchivesClient
	*LogsClient
	*NotificationPlatformsClient
	*NotificationsClient
	*NotifiersClient
	*OperationsClient
	*RunsClient
	*SignUpClient
	*SourcesClient
	*TokensClient
	*UsersClient
	*VariablesClient

	*backendConfiguration
}

type ClientConfig struct {
	Timeout        time.Duration
	Endpoint       string
	TLSConfig      *tls.Config
	TokenGenerator TokenGenerator
}

func NewClient(cfg ClientConfig) *Client {
	if cfg.Timeout == 0 {
		cfg.Timeout = 30 * time.Second
	}
	if cfg.Endpoint == "" {
		cfg.Endpoint = defaultEndpoint
	}
	if cfg.TLSConfig == nil {
		cfg.TLSConfig = &tls.Config{}
	}
	c := Client{
		backendConfiguration: &backendConfiguration{
			TokenGenerator: cfg.TokenGenerator,
			Endpoint:       cfg.Endpoint,
			APIVersion:     defaultAPIVersion,
			TLSConfig:      cfg.TLSConfig,
			httpClient: &http.Client{
				Timeout: cfg.Timeout,
				Transport: &http.Transport{
					Proxy:           http.ProxyFromEnvironment,
					TLSClientConfig: cfg.TLSConfig,
				},
			},
		},
	}

	c.init()

	return &c
}

func (c *Client) HTTPClient() HTTPClient {
	return c.httpClient
}

func (c *backendConfiguration) HTTPClient() HTTPClient {
	return c.httpClient
}

func (c *Client) init() {
	c.AddonsClient = &AddonsClient{subresourceClient{c.backendConfiguration}}
	c.AddonProvidersClient = &AddonProvidersClient{c.backendConfiguration}
	c.AppsClient = &AppsClient{c.backendConfiguration}
	c.CollaboratorsClient = &CollaboratorsClient{subresourceClient{c.backendConfiguration}}
	c.DeploymentsClient = &DeploymentsClient{c.backendConfiguration}
	c.DomainsClient = &DomainsClient{subresourceClient{c.backendConfiguration}}
	c.EventsClient = &EventsClient{subresourceClient{c.backendConfiguration}}
	c.KeysClient = &KeysClient{c.backendConfiguration}
	c.LoginClient = &LoginClient{c.backendConfiguration}
	c.LogsArchivesClient = &LogsArchivesClient{c.backendConfiguration}
	c.LogsClient = &LogsClient{c.backendConfiguration}
	c.NotificationPlatformsClient = &NotificationPlatformsClient{c.backendConfiguration}
	c.NotificationsClient = &NotificationsClient{subresourceClient{c.backendConfiguration}}
	c.NotifiersClient = &NotifiersClient{subresourceClient{c.backendConfiguration}}
	c.OperationsClient = &OperationsClient{subresourceClient{c.backendConfiguration}}
	c.RunsClient = &RunsClient{c.backendConfiguration}
	c.SignUpClient = &SignUpClient{c.backendConfiguration}
	c.TokensClient = &TokensClient{c.backendConfiguration}
	c.UsersClient = &UsersClient{c.backendConfiguration}
	c.VariablesClient = &VariablesClient{subresourceClient{c.backendConfiguration}}
	c.SourcesClient = &SourcesClient{c.backendConfiguration}
}
