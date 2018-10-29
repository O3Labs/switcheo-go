//Package withdrawal provides the binding for Switcheo Rest APIs withdrawal endpoints
//https://docs.switcheo.network/#tickers
package withdrawal

import (
	"fmt"

	"github.com/o3labs/switcheo-go/switcheo/client"
)

type Client struct {
	API client.API
}

//Available methods
type Interface interface {
	CreateWithdrawal(params WithdrawalParams) (CreateWithdrawalResponse, error)
	ExecuteWithdrawal(params ExecuteWithdrawalParams) (interface{}, error)
}

var _ Interface = (*Client)(nil)

//First step required to withdraw token from Switcheo smart contract
func (c *Client) CreateWithdrawal(params WithdrawalParams) (CreateWithdrawalResponse, error) {
	out := CreateWithdrawalResponse{}
	err := c.API.Request("POST", "/v2/withdrawals", params, &out)
	return out, err
}

//Second step required to withdraw token from Switcheo smart contract
func (c *Client) ExecuteWithdrawal(params ExecuteWithdrawalParams) (interface{}, error) {
	out := map[string]interface{}{}
	err := c.API.Request("POST", fmt.Sprintf("/v2/withdrawals/%v/broadcast", params.ID), params, &out)
	return out, err
}
