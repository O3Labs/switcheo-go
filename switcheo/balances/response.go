package balances

import "time"

// Response for GetBalances
type GetBalancesResponse struct {
	Neo struct {
		Confirming map[string][]struct {
			EventType       string    `json:"event_type"`
			AssetID         string    `json:"asset_id"`
			Amount          int64     `json:"amount"`
			TransactionHash string    `json:"transaction_hash"`
			CreatedAt       time.Time `json:"created_at"`
		} `json:"confirming"`
		Confirmed map[string]string `json:"confirmed"`
		Locked    map[string]string `json:"locked"`
	} `json:"neo"`
}
