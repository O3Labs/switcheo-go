package deposit_test

import (
	"log"
	"math"
	"strconv"
	"testing"
	"time"

	"github.com/o3labs/switcheo-go/switcheo"
	"github.com/o3labs/switcheo-go/switcheo/deposit"
)

func TestSignDeposit(t *testing.T) {
	amount := uint64(100 * math.Pow10(8))
	params := deposit.DepositParams{
		Blockchain:   switcheo.NEOBlockchain,
		AssetID:      "SWTH",
		Amount:       strconv.FormatUint(amount, 10),
		Timestamp:    int64(time.Now().UnixNano() / int64(time.Millisecond)),
		ContractHash: testNetV2Contract,
	}
	//test
	wif := ""

	params.Sign(wif)
	log.Printf("%+v", params)
}
