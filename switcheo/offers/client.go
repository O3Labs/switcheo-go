//Package offers provides the binding for Switcheo Rest APIs offers endpoints
//https://docs.switcheo.network/#offers
package offers

import "github.com/o3labs/switcheo-go/switcheo/client"

type Client struct {
	API client.API
}

//Available methods
type Interface interface {
	ListOffers(params ListOffersParams) (Offers, error)
}

var _ Interface = (*Client)(nil)

//List available offers
func (c *Client) ListOffers(params ListOffersParams) (Offers, error) {
	out := Offers{}
	err := c.API.Request("GET", "/v2/offers", params, &out)
	return out, err
}
