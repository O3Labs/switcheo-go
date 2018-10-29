package deposit_test

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/o3labs/neo-utils/neoutils"
	"github.com/o3labs/switcheo-go/switcheo"
	"github.com/o3labs/switcheo-go/switcheo/client"
	"github.com/o3labs/switcheo-go/switcheo/deposit"
	"github.com/o3labs/switcheo-go/switcheo/signer"
)

var httpClient = &http.Client{Timeout: 80 * time.Second}
var testNetV2Contract = "a195c1549e7da61b8da315765a790ac7e7633b82"

func TestDeposit(t *testing.T) {
	api := client.New(switcheo.TestNetAPI, httpClient, switcheo.UserAgent)
	c := deposit.Client{API: *api}

	amount := uint64(100 * math.Pow10(8))
	params := deposit.DepositParams{
		Blockchain:   switcheo.NEOBlockchain,
		AssetID:      "SWTH",
		Amount:       strconv.FormatUint(amount, 10),
		Timestamp:    int64(time.Now().UnixNano() / int64(time.Millisecond)),
		ContractHash: testNetV2Contract,
	}
	wif := ""
	params.Sign(wif)

	response, err := c.CreateDeposit(params)
	if err != nil {
		log.Printf("error %v ", err)
		return
	}
	log.Printf("%+v", response)

	tx := response.Transaction

	b, _ := json.Marshal(tx)
	toSign := neoutils.NeonJSTransaction{}
	json.Unmarshal(b, &toSign)
	serialized := neoutils.NeonJSTXSerializer(toSign)

	signature, err := signer.SignBytesWithWIF(serialized, wif)
	if err != nil {
		log.Printf("error %v ", err)
		return
	}

	exeucuteDepositParams := deposit.ExecuteDepositParams{
		Signature: signature,
	}
	executeDepositResponse, err := c.ExecuteDeposit(response.ID, exeucuteDepositParams)
	if err != nil {
		log.Printf("error %v ", err)
		return
	}
	log.Printf("%+v", executeDepositResponse)

}
