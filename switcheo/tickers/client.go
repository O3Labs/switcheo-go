//Package tickers provides the binding for Switcheo Rest APIs tickers endpoints
//https://docs.switcheo.network/#tickers
package tickers

import "github.com/o3labs/switcheo-go/switcheo/client"

type Client struct {
	API client.API
}

//Available methods
type Interface interface {
	Last24Hours() (Last24HoursResponse, error)
	LastPrice() (map[string]map[string]string, error)
}

var _ Interface = (*Client)(nil)

//Get last 24 hour data for all pairs and markets
func (c *Client) Last24Hours() (Last24HoursResponse, error) {
	out := Last24HoursResponse{}
	err := c.API.Request("GET", "/v2/tickers/last_24_hours", nil, &out)
	return out, err
}

//Get last price of all symbols
func (c *Client) LastPrice() (map[string]map[string]string, error) {
	out := map[string]map[string]string{}
	err := c.API.Request("GET", "/v2/tickers/last_price", nil, &out)
	return out, err
}
