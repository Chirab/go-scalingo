package scalingo

import (
	"net/http"

	httpclient "github.com/Scalingo/go-scalingo/http"
	"gopkg.in/errgo.v1"
)

type LogDrainsService interface {
	LogDrainsList(app string) ([]LogDrain, error)
	LogDrainsAddonList(app string, addonID string) ([]LogDrain, error)
	LogDrainAdd(app string, params LogDrainAddParams) (*LogDrainRes, error)
	LogDrainRemove(app, URL string) error
}

var _ LogDrainsService = (*Client)(nil)

type LogDrain struct {
	AppID       string `json:"app_id"`
	URL         string `json:"url"`
	Type        string `json:"type"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Token       string `json:"token"`
	DrainRegion string `json:"drain_region"`
}

type LogDrainRes struct {
	Drain LogDrain `json:"drain"`
}

type AddonRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type LogDrainsRes struct {
	Addon  AddonRes   `json:"addon"`
	Drains []LogDrain `json:"drains"`
}

func (c *Client) LogDrainsList(app string) ([]LogDrain, error) {
	var logDrainsRes LogDrainsRes
	err := c.ScalingoAPI().SubresourceList("apps", app, "log_drains", nil, &logDrainsRes)
	if err != nil {
		return nil, errgo.Notef(err, "fail to list the log drains")
	}
	return logDrainsRes.Drains, nil
}

func (c *Client) LogDrainsAddonList(app string, addonID string) ([]LogDrain, error) {
	var logDrainsRes LogDrainsRes

	req := &httpclient.APIRequest{
		Method:   "GET",
		Endpoint: "apps/" + app + "/addon/" + addonID + "/log_drains",
	}

	data := logDrainsRes
	err := c.ScalingoAPI().DoRequest(req, &data)
	if err != nil {
		return nil, errgo.Notef(err, "fail to list the log drains")
	}

	return logDrainsRes.Drains, nil
}

type LogDrainAddParams struct {
	Type        string `json:"type"`
	URL         string `json:"url"`
	Port        string `json:"port"`
	Host        string `json:"host"`
	Token       string `json:"token"`
	DrainRegion string `json:"drain_region"`
}

func (c *Client) LogDrainAdd(app string, params LogDrainAddParams) (*LogDrainRes, error) {
	var logDrainRes LogDrainRes
	payload := LogDrainRes{
		Drain: LogDrain{
			Type:        params.Type,
			URL:         params.URL,
			Host:        params.Host,
			Port:        params.Port,
			Token:       params.Token,
			DrainRegion: params.DrainRegion,
		},
	}

	err := c.ScalingoAPI().SubresourceAdd("apps", app, "log_drains", payload, &logDrainRes)
	if err != nil {
		return nil, errgo.Notef(err, "fail to add drain")
	}

	return &logDrainRes, nil
}

func (c *Client) LogDrainRemove(app, URL string) error {
	payload := map[string]string{
		"url": URL,
	}

	req := &httpclient.APIRequest{
		Method:   "DELETE",
		Endpoint: "/apps/" + app + "/log_drains",
		Expected: httpclient.Statuses{http.StatusNoContent},
		Params:   payload,
	}

	err := c.ScalingoAPI().DoRequest(req, nil)
	if err != nil {
		return errgo.Notef(err, "fail to delete log drain")
	}

	return nil
}
