//Package signer provides methods to sign data and parameters with NEO keypair
//The only dependency is github.com/o3labs/neo-utils/neoutils
package signer

import (
	"encoding/json"
	"fmt"

	"github.com/o3labs/neo-utils/neoutils"
)

func SignBytesWithWIF(b []byte, wif string) (string, error) {
	wallet, err := neoutils.GenerateFromWIF(wif)
	if err != nil {
		return "", err
	}
	signatureB, err := neoutils.Sign(b, neoutils.BytesToHex(wallet.PrivateKey))
	if err != nil {
		return "", err
	}
	return neoutils.BytesToHex(signatureB), nil
}

func SignParams(params interface{}, wif string) (address string, signature string, e error) {
	wallet, err := neoutils.GenerateFromWIF(wif)
	if err != nil {
		return "", "", err
	}

	b, _ := json.Marshal(params)
	//we must sort json key first
	sortedKey := map[string]interface{}{}
	json.Unmarshal(b, &sortedKey)
	sortedB, _ := json.Marshal(sortedKey)
	hex := neoutils.BytesToHex(sortedB)
	tosignHex := fmt.Sprintf("010001f0%02x%v0000", len(string(sortedB)), hex)

	address = neoutils.BytesToHex(neoutils.ReverseBytes(wallet.HashedSignature))
	signatureB, err := neoutils.Sign(neoutils.HexTobytes(tosignHex), neoutils.BytesToHex(wallet.PrivateKey))
	e = err

	signature = neoutils.BytesToHex(signatureB)
	return
}
