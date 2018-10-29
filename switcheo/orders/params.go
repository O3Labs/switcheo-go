package orders

type ListOrdersParams struct {
	Address      string `json:"address"`
	ContractHash string `json:"contract_hash"`
}
