//Package balances provides the binding for Switcheo Rest APIs balances endpoints
//https://docs.switcheo.network/#balances
package balances

import (
	"encoding/binary"
	"log"

	"github.com/o3labs/neo-utils/neoutils"
	"github.com/o3labs/switcheo-go/switcheo/client"
)

type Client struct {
	API client.API
}

//Available methods
type Interface interface {
	GetBalances(params GetBalancesParam) (GetBalancesResponse, error)
}

var _ Interface = (*Client)(nil)

// Get contract balances of the given address
func (c *Client) GetBalances(params GetBalancesParam) (GetBalancesResponse, error) {
	out := GetBalancesResponse{}
	for i, v := range params.Addresses {
		params.Addresses[i] = neoutils.NEOAddressToScriptHashWithEndian(v, binary.BigEndian)
	}

	log.Printf("%+v", params)
	err := c.API.Request("GET", "/v2/balances", params, &out)
	return out, err
}
