package withdrawal_test

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/o3labs/switcheo-go/switcheo"
	"github.com/o3labs/switcheo-go/switcheo/client"
	"github.com/o3labs/switcheo-go/switcheo/withdrawal"
)

var httpClient = &http.Client{Timeout: 80 * time.Second}
var testNetV2Contract = "a195c1549e7da61b8da315765a790ac7e7633b82"

func TestCreateWithdrawal(t *testing.T) {
	api := client.New(switcheo.TestNetAPI, httpClient, switcheo.UserAgent)
	c := withdrawal.Client{API: *api}

	amount := uint64(100 * math.Pow10(8))
	params := withdrawal.WithdrawalParams{
		Blockchain:   switcheo.NEOBlockchain,
		AssetID:      "SWTH",
		Amount:       strconv.FormatUint(amount, 10),
		Timestamp:    int64(time.Now().UnixNano() / int64(time.Millisecond)),
		ContractHash: testNetV2Contract,
	}
	wif := ""
	params.Sign(wif)

	response, err := c.CreateWithdrawal(params)
	if err != nil {
		log.Printf("error %v ", err)
		return
	}
	log.Printf("%+v", response)

	executeWithdrawalParams := withdrawal.ExecuteWithdrawalParams{
		ID:        response.ID,
		Timestamp: int64(time.Now().UnixNano() / int64(time.Millisecond)),
	}
	executeWithdrawalParams.Sign(wif)

	withdrawalResponse, err := c.ExecuteWithdrawal(executeWithdrawalParams)
	if err != nil {
		log.Printf("error %v ", err)
		return
	}
	log.Printf("%+v", withdrawalResponse)
}
