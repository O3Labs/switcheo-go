package offers

//Parameters for ListOffers
type ListOffersParams struct {
	Blockchain   string `json:"blockchain"`
	Pair         string `json:"pair"`
	ContractHash string `json:"contract_hash"`
}
