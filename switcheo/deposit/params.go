package deposit

import "github.com/o3labs/switcheo-go/switcheo/signer"

// Parameters for CreateDeposit
type DepositParams struct {
	Blockchain   string `json:"blockchain"`
	AssetID      string `json:"asset_id"`
	Amount       string `json:"amount"`
	Timestamp    int64  `json:"timestamp"`
	ContractHash string `json:"contract_hash"`
	Address      string `json:"address,omitempty"`
	Signature    string `json:"signature,omitempty"`
}

// A method to sign a deposit payload with WIF
func (d *DepositParams) Sign(wif string) error {
	address, signature, err := signer.SignParams(d, wif)
	if err != nil {
		return err
	}
	d.Address = address
	d.Signature = signature

	return nil
}

// Parameters for ExecuteDeposit
type ExecuteDepositParams struct {
	Signature string `json:"signature"`
}
