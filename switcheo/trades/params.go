package trades

//Parameters for ListTrades
type ListTradesParams struct {
	Blockchain   string `json:"blockchain"`
	Pair         string `json:"pair"`
	Limit        int    `json:"limit"`
	ContractHash string `json:"contract_hash"`
}
