package messageTemplate

import "github.com/ArtisanCloud/power-wechat/src/kernel"

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}
