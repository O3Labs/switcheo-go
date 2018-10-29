//Package deposit provides the binding for Switcheo Rest APIs deposit endpoints
//https://docs.switcheo.network/#create-deposit
package deposit

import (
	"fmt"

	"github.com/o3labs/switcheo-go/switcheo/client"
)

type Client struct {
	API client.API
}

//Available methods
type Interface interface {
	CreateDeposit(params DepositParams) (CreateDepositResponse, error)
	ExecuteDeposit(depositID string, params ExecuteDepositParams) (interface{}, error)
}

var _ Interface = (*Client)(nil)

//First endpoint required to deposit token into Switcheo smart contract.
//Create a deposit which can be executed through Execute Deposit.
func (c *Client) CreateDeposit(params DepositParams) (CreateDepositResponse, error) {
	out := CreateDepositResponse{}
	err := c.API.Request("POST", "/v2/deposits", params, &out)
	return out, err
}

//Second endpoint required to deposit token into Switcheo smart contract.
//Execute a deposit
func (c *Client) ExecuteDeposit(depositID string, params ExecuteDepositParams) (interface{}, error) {
	out := map[string]interface{}{}
	err := c.API.Request("POST", fmt.Sprintf("/v2/deposits/%v/broadcast", depositID), params, &out)
	return out, err
}
