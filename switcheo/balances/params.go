package balances

// Parameters for GetBalances
type GetBalancesParam struct {
	Addresses      []string `json:"addresses"`
	ContractHashes []string `json:"contract_hashes"`
	Group          bool     `json:"group"`
}
