package general

type TimestampResponse struct {
	Timestamp int64 `json:"timestamp"`
}

type ContractsResponse struct {
	NEO struct {
		V1  string `json:"V1"`
		V15 string `json:"V1_5"`
		V2  string `json:"V2"`
	} `json:"NEO"`
}

type TokenItemReponse struct {
	Hash          string `json:"hash"`
	Decimals      int    `json:"decimals"`
	TradingActive bool   `json:"trading_active"`
}
