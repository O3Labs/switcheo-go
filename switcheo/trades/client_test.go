package trades_test

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/o3labs/switcheo-go/switcheo"
	"github.com/o3labs/switcheo-go/switcheo/client"
	"github.com/o3labs/switcheo-go/switcheo/trades"
)

var httpClient = &http.Client{Timeout: 80 * time.Second}

func TestListTrades(t *testing.T) {
	api := client.New(switcheo.TestNetAPI, httpClient, switcheo.UserAgent)
	c := trades.Client{API: *api}
	params := trades.ListTradesParams{
		Blockchain:   switcheo.NEOBlockchain,
		Pair:         "SWTH_NEO",
		Limit:        10,
		ContractHash: "a195c1549e7da61b8da315765a790ac7e7633b82",
	}
	response, err := c.ListTrades(params)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", response)
}
