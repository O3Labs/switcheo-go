package tickers_test

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/o3labs/switcheo-go/switcheo"
	"github.com/o3labs/switcheo-go/switcheo/client"
	"github.com/o3labs/switcheo-go/switcheo/tickers"
)

var httpClient = &http.Client{Timeout: 80 * time.Second}

func TestLast24Hours(t *testing.T) {
	api := client.New(switcheo.TestNetAPI, httpClient, switcheo.UserAgent)
	c := tickers.Client{API: *api}

	response, err := c.Last24Hours()
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", response)
}

func TestLastPrice(t *testing.T) {
	api := client.New(switcheo.TestNetAPI, httpClient, switcheo.UserAgent)
	c := tickers.Client{API: *api}

	response, err := c.LastPrice()
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", response)
}
