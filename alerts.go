package scalingo

import (
	"gopkg.in/errgo.v1"
)

type AlertsService interface {
	AlertsList(app string) ([]Alert, error)
	AlertAdd(app string, params AlertAddParams) (Alert, error)
	AlertShow(app, id string) (*Alert, error)
	AlertUpdate(app, id string, params AlertUpdateParams) (*Alert, error)
	AlertRemove(app, id string) error
}

type AlertsClient struct {
	subresourceClient
}

type Alert struct {
	ID            string  `json:"id"`
	ContainerType string  `json:"container_type"`
	Metric        string  `json:"metric"`
	Limit         float64 `json:"limit"`
	Disabled      bool    `json:"disabled"`
	SendWhenBelow bool    `json:"send_when_below"`
}

type AlertsRes struct {
	Alerts []*Alert `json:"alerts"`
}

type AlertRes struct {
	Alert Alert `json:"alert"`
}

func (c *AlertsClient) AlertsList(app string) ([]*Alert, error) {
	var alertsRes AlertsRes
	err := c.subresourceList(app, "alerts", nil, &alertsRes)
	if err != nil {
		return nil, errgo.Mask(err)
	}
	return alertsRes.Alerts, nil
}

type AlertAddParams struct {
	ContainerType string
	Metric        string
	Limit         float64
	Disabled      bool
}

func (c *AlertsClient) AlertAdd(app string, params AlertAddParams) (*Alert, error) {
	var alertRes AlertRes
	err := c.subresourceAdd(app, "alerts", AlertRes{
		Alert: Alert{
			ContainerType: params.ContainerType,
			Metric:        params.Metric,
			Limit:         params.Limit,
			Disabled:      params.Disabled,
		},
	}, &alertRes)
	if err != nil {
		return nil, errgo.Mask(err)
	}
	return &alertRes.Alert, nil
}

func (c *AlertsClient) AlertShow(app, id string) (*Alert, error) {
	var alertRes AlertRes
	err := c.subresourceGet(app, "alerts", id, nil, &alertRes)
	if err != nil {
		return nil, errgo.Mask(err)
	}
	return &alertRes.Alert, nil
}

type AlertUpdateParams struct {
	ContainerType *string  `json:"container_type,omitempty"`
	Metric        *string  `json:"metric,omitempty"`
	Limit         *float64 `json:"limit,omitempty"`
	Disabled      *bool    `json:"disabled,omitempty"`
}

func (c *AlertsClient) AlertUpdate(app, id string, params AlertUpdateParams) (*Alert, error) {
	var alertRes AlertRes
	err := c.subresourceUpdate(app, "alerts", id, params, &alertRes)
	if err != nil {
		return nil, errgo.Mask(err)
	}
	return &alertRes.Alert, nil
}

func (c *AlertsClient) AlertRemove(app, id string) error {
	err := c.subresourceDelete(app, "alerts", id)
	if err != nil {
		return errgo.Mask(err, errgo.Any)
	}
	return nil
}
