package general_test

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/o3labs/switcheo-go/switcheo"
	"github.com/o3labs/switcheo-go/switcheo/client"
	"github.com/o3labs/switcheo-go/switcheo/general"
)

var httpClient = &http.Client{Timeout: 80 * time.Second}

func TestGetTimestamp(t *testing.T) {
	api := client.New(switcheo.TestNetAPI, httpClient, switcheo.UserAgent)
	c := general.Client{API: *api}
	response, err := c.GetTimestamp()
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", response)
}

func TestGetContract(t *testing.T) {
	api := client.New(switcheo.TestNetAPI, httpClient, switcheo.UserAgent)
	c := general.Client{API: *api}
	response, err := c.GetContracts()
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", response)
}

func TestGetTokens(t *testing.T) {
	api := client.New(switcheo.TestNetAPI, httpClient, switcheo.UserAgent)
	c := general.Client{API: *api}
	response, err := c.GetTokens()
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", response)
}

func TestGetPairs(t *testing.T) {
	api := client.New(switcheo.TestNetAPI, httpClient, switcheo.UserAgent)
	c := general.Client{API: *api}
	response, err := c.GetPairs()
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", response)
}
