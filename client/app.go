package client

import (
	"github.com/replicatedhq/replicated/pkg/types"
)

func (c *Client) ListApps() ([]types.AppAndChannels, error) {
	platformApps, err := c.PlatformClient.ListApps()
	if err != nil {
		return nil, err
	}

	shipApps, err := c.ShipClient.ListApps()
	if err != nil {
		return nil, err
	}

	apps := make([]types.AppAndChannels, 0, 0)
	for _, platformApp := range platformApps {
		channels := make([]types.Channel, 0, 0)
		for _, platformChannel := range platformApp.Channels {
			channel := types.Channel{
				ID:          platformChannel.Id,
				Name:        platformChannel.Name,
				Description: platformChannel.Description,
			}

			channels = append(channels, channel)
		}

		app := types.AppAndChannels{
			App: &types.App{
				ID:        platformApp.App.Id,
				Name:      platformApp.App.Name,
				Scheduler: platformApp.App.Scheduler,
				Slug:      platformApp.App.Scheduler,
			},
			Channels: channels,
		}

		apps = append(apps, app)
	}

	apps = append(apps, shipApps...)

	return apps, nil
}

func (c *Client) GetApp(appID string) (interface{}, error) {
	return nil, nil
}

func (c *Client) CreateApp(appOptions interface{}) (interface{}, error) {
	return nil, nil
}

func (c *Client) DeleteApp(appID string) error {
	return nil
}
