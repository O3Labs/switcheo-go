package switcheo

import (
	"net/http"
	"time"
)

const (
	TestNetAPI         = "https://test-api.switcheo.network"
	MainNetAPI         = "https://api.switcheo.network"
	clientVersion      = "0.1"
	defaultHTTPTimeout = 80 * time.Second
	UserAgent          = "github.com/apisit/switcheo-go version/" + clientVersion
	NEOBlockchain      = "neo"
	MainNetV2          = "91b83e96f2a7c4fdf0c1688441ec61986c7cae26"
	TestNetV2          = "a195c1549e7da61b8da315765a790ac7e7633b82"
)

var httpClient = &http.Client{Timeout: defaultHTTPTimeout}
