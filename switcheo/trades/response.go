package trades

import "time"

//Response for ListTrades
type ListTradesResponse []struct {
	ID         string    `json:"id"`
	FillAmount int64     `json:"fill_amount"`
	TakeAmount int       `json:"take_amount"`
	EventTime  time.Time `json:"event_time"`
	IsBuy      bool      `json:"is_buy"`
}
