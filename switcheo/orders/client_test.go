package orders_test

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/o3labs/switcheo-go/switcheo"
	"github.com/o3labs/switcheo-go/switcheo/client"
	"github.com/o3labs/switcheo-go/switcheo/orders"
)

var httpClient = &http.Client{Timeout: 80 * time.Second}

func TestListOrders(t *testing.T) {
	api := client.New(switcheo.TestNetAPI, httpClient, switcheo.UserAgent)
	c := orders.Client{API: *api}

	params := orders.ListOrdersParams{
		Address:      "ANoW2zD8HmhbWJAjL4yKJWCZcF2WFb1ire",
		ContractHash: switcheo.TestNetV2,
	}
	response, err := c.ListOrders(params)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", response)
}
