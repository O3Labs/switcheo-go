package offers

type Offers []struct {
	ID              string `json:"id"`
	OfferAsset      string `json:"offer_asset"`
	WantAsset       string `json:"want_asset"`
	AvailableAmount int64  `json:"available_amount"`
	OfferAmount     int64  `json:"offer_amount"`
	WantAmount      int    `json:"want_amount"`
}
