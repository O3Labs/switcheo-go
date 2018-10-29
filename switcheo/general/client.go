//Package general provides the binding for Switcheo Rest APIs exchange information endpoints
//https://docs.switcheo.network/#exchange-information
package general

import "github.com/o3labs/switcheo-go/switcheo/client"

type Client struct {
	API client.API
}

//Available methods
type Interface interface {
	GetTimestamp() (TimestampResponse, error)
	GetContracts() (ContractsResponse, error)
	GetTokens() (map[string]TokenItemReponse, error)
	GetPairs() ([]string, error)
}

var _ Interface = (*Client)(nil)

//Get switcheo server timestamp
func (c *Client) GetTimestamp() (TimestampResponse, error) {
	out := TimestampResponse{}
	err := c.API.Request("GET", "/v2/exchange/timestamp", nil, &out)
	return out, err
}

//Get Switcheo contracts
func (c *Client) GetContracts() (ContractsResponse, error) {
	out := ContractsResponse{}
	err := c.API.Request("GET", "/v2/exchange/contracts", nil, &out)
	return out, err
}

//Get supported tokens
func (c *Client) GetTokens() (map[string]TokenItemReponse, error) {
	out := map[string]TokenItemReponse{}
	err := c.API.Request("GET", "/v2/exchange/tokens", nil, &out)
	return out, err
}

//Get supported pairs
func (c *Client) GetPairs() ([]string, error) {
	out := []string{}
	err := c.API.Request("GET", "/v2/exchange/pairs", nil, &out)
	return out, err
}
