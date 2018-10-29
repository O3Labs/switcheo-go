package offers_test

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/o3labs/switcheo-go/switcheo"
	"github.com/o3labs/switcheo-go/switcheo/client"
	"github.com/o3labs/switcheo-go/switcheo/offers"
)

var httpClient = &http.Client{Timeout: 80 * time.Second}

func TestListOffers(t *testing.T) {
	api := client.New(switcheo.TestNetAPI, httpClient, switcheo.UserAgent)
	c := offers.Client{API: *api}

	params := offers.ListOffersParams{
		Blockchain:   switcheo.NEOBlockchain,
		Pair:         "SWTH_NEO",
		ContractHash: switcheo.V2TestNet,
	}
	response, err := c.ListOffers(params)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", response)
}
