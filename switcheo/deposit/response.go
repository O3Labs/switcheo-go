package deposit

type DepositTransaction struct {
	Sha256 string `json:"sha256"`
	Hash   string `json:"hash"`
	Inputs []struct {
		PrevIndex int    `json:"prevIndex"`
		PrevHash  string `json:"prevHash"`
	} `json:"inputs"`
	Outputs []struct {
		AssetID    string  `json:"assetId"`
		ScriptHash string  `json:"scriptHash"`
		Value      float64 `json:"value"`
	} `json:"outputs"`
	Script     string `json:"script"`
	Version    int    `json:"version"`
	Type       int    `json:"type"`
	Attributes []struct {
		Usage int    `json:"usage"`
		Data  string `json:"data"`
	} `json:"attributes"`
	Scripts []interface{} `json:"scripts"`
	Gas     int           `json:"gas"`
}

// Response for CreateDeposit
type CreateDepositResponse struct {
	ID           string             `json:"id"`
	Transaction  DepositTransaction `json:"transaction"`
	ScriptParams struct {
		ScriptHash string        `json:"scriptHash"`
		Operation  string        `json:"operation"`
		Args       []interface{} `json:"args"`
	} `json:"script_params"`
}
