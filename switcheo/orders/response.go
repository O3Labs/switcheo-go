package orders

import "time"

type FillOrder struct {
	ID              string      `json:"id"`
	OfferHash       string      `json:"offer_hash"`
	OfferAssetID    string      `json:"offer_asset_id"`
	WantAssetID     string      `json:"want_asset_id"`
	FillAmount      string      `json:"fill_amount"`
	WantAmount      string      `json:"want_amount"`
	FilledAmount    string      `json:"filled_amount"`
	FeeAssetID      string      `json:"fee_asset_id"`
	FeeAmount       string      `json:"fee_amount"`
	Price           string      `json:"price"`
	Txn             interface{} `json:"txn"`
	Status          string      `json:"status"`
	CreatedAt       time.Time   `json:"created_at"`
	TransactionHash string      `json:"transaction_hash"`
}

type MakeOrder struct {
	ID              string      `json:"id"`
	OfferHash       string      `json:"offer_hash"`
	AvailableAmount string      `json:"available_amount"`
	OfferAssetID    string      `json:"offer_asset_id"`
	OfferAmount     string      `json:"offer_amount"`
	WantAssetID     string      `json:"want_asset_id"`
	WantAmount      string      `json:"want_amount"`
	FilledAmount    string      `json:"filled_amount"`
	Txn             interface{} `json:"txn"`
	CancelTxn       interface{} `json:"cancel_txn"`
	Price           string      `json:"price"`
	Status          string      `json:"status"`
	CreatedAt       time.Time   `json:"created_at"`
	TransactionHash string      `json:"transaction_hash"`
	Trades          []struct {
		ID           string    `json:"id"`
		Status       string    `json:"status"`
		WantAmount   string    `json:"want_amount"`
		FilledAmount string    `json:"filled_amount"`
		FeeAmount    string    `json:"fee_amount"`
		FeeAssetID   string    `json:"fee_asset_id"`
		Price        string    `json:"price"`
		CreatedAt    time.Time `json:"created_at"`
	} `json:"trades"`
}

type ListOrdersResponse []struct {
	ID                      string      `json:"id"`
	Blockchain              string      `json:"blockchain"`
	ContractHash            string      `json:"contract_hash"`
	Address                 string      `json:"address"`
	Side                    string      `json:"side"`
	OfferAssetID            string      `json:"offer_asset_id"`
	WantAssetID             string      `json:"want_asset_id"`
	OfferAmount             string      `json:"offer_amount"`
	WantAmount              string      `json:"want_amount"`
	TransferAmount          string      `json:"transfer_amount"`
	PriorityGasAmount       string      `json:"priority_gas_amount"`
	UseNativeToken          bool        `json:"use_native_token"`
	NativeFeeTransferAmount int         `json:"native_fee_transfer_amount"`
	DepositTxn              interface{} `json:"deposit_txn"`
	CreatedAt               time.Time   `json:"created_at"`
	Status                  string      `json:"status"`
	Fills                   []FillOrder `json:"fills"`
	Makes                   []MakeOrder `json:"makes"`
}
