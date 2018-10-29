//Package trades provides the binding for Switcheo Rest APIs trades endpoints
//https://docs.switcheo.network/#tickers
package trades

import "github.com/o3labs/switcheo-go/switcheo/client"

type Client struct {
	API client.API
}

//Available methods
type Interface interface {
	ListTrades(params ListTradesParams) (ListTradesResponse, error)
}

var _ Interface = (*Client)(nil)

//Retrieves trades that have already occurred on Switcheo Exchange filtered by the request parameters
func (c *Client) ListTrades(params ListTradesParams) (ListTradesResponse, error) {
	out := ListTradesResponse{}
	err := c.API.Request("GET", "/v2/trades", params, &out)
	return out, err
}
