package tickers

type Last24HoursResponse []struct {
	Pair        string `json:"pair"`
	Open        string `json:"open"`
	Close       string `json:"close"`
	High        string `json:"high"`
	Low         string `json:"low"`
	Volume      string `json:"volume"`
	QuoteVolume string `json:"quote_volume"`
}
