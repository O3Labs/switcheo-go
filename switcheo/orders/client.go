//Package orders provides the binding for Switcheo Rest APIs orders endpoints
//https://docs.switcheo.network/#orders
package orders

import (
	"encoding/binary"

	"github.com/o3labs/neo-utils/neoutils"
	"github.com/o3labs/switcheo-go/switcheo/client"
)

type Client struct {
	API client.API
}

//Available methods
type Interface interface {
	ListOrders(params ListOrdersParams) (ListOrdersResponse, error)
}

var _ Interface = (*Client)(nil)

//List available orders
func (c *Client) ListOrders(params ListOrdersParams) (ListOrdersResponse, error) {
	out := ListOrdersResponse{}
	//auto convert to address hash
	params.Address = neoutils.NEOAddressToScriptHashWithEndian(params.Address, binary.BigEndian)

	err := c.API.Request("GET", "/v2/orders", params, &out)
	return out, err
}
