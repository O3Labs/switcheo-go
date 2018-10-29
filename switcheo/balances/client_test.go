package balances_test

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/o3labs/switcheo-go/switcheo"
	"github.com/o3labs/switcheo-go/switcheo/balances"
	"github.com/o3labs/switcheo-go/switcheo/client"
)

var httpClient = &http.Client{Timeout: 80 * time.Second}
var mainNetV2Contract = "91b83e96f2a7c4fdf0c1688441ec61986c7cae26"

func TestGetBalances(t *testing.T) {
	api := client.New(switcheo.MainNetAPI, httpClient, switcheo.UserAgent)
	c := balances.Client{API: *api}
	params := balances.GetBalancesParam{
		Addresses:      []string{"ASi48wqdF9avm91pWwdphcAmaDJQkPNdNt"},
		ContractHashes: []string{mainNetV2Contract},
		Group:          true,
	}
	response, err := c.GetBalances(params)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", response)
}
