package entities

type Style struct {
	ID          int64  `json:"id"`
	BuyerID     int    `json:"buyer_id"`
	Name        string `json:"name"`
	Destination string `json:"destination"`
	Amount      int64  `json:"amount"`
}
