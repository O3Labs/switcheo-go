package withdrawal

import "github.com/o3labs/switcheo-go/switcheo/signer"

//Parameters for Withdrawal
type WithdrawalParams struct {
	Blockchain   string `json:"blockchain"`
	AssetID      string `json:"asset_id"`
	Amount       string `json:"amount"`
	Timestamp    int64  `json:"timestamp"`
	ContractHash string `json:"contract_hash"`
	Address      string `json:"address,omitempty"`
	Signature    string `json:"signature,omitempty"`
}

//A method to sign withdrawal parameters
func (d *WithdrawalParams) Sign(wif string) error {
	address, signature, err := signer.SignParams(d, wif)
	if err != nil {
		return err
	}
	d.Address = address
	d.Signature = signature

	return nil
}

//Parameters for ExecuteWithdrawal
type ExecuteWithdrawalParams struct {
	ID        string `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature,omitempty"`
}

//A method to sign execute withdrawal parameters
func (d *ExecuteWithdrawalParams) Sign(wif string) error {
	_, signature, err := signer.SignParams(d, wif)
	if err != nil {
		return err
	}

	d.Signature = signature

	return nil
}
