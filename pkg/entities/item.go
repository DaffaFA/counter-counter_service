package entities

type Item struct {
	Code    string   `json:"code,omitempty" db:"code"`
	BuyerID int      `json:"buyer_id,omitempty" db:"buyer_id"`
	Buyer   *Setting `json:"buyer,omitempty"`
	StyleID int      `json:"style_id,omitempty" db:"style_id"`
	Style   *Setting `json:"style,omitempty"`
	ColorID int      `json:"color_id,omitempty" db:"color_id"`
	Color   *Setting `json:"color,omitempty"`
	SizeID  int      `json:"size_id,omitempty" db:"size_id"`
	Size    *Setting `json:"size,omitempty"`
}

type ItemPagination struct {
	Items []Item `json:"items"`
	Total int    `json:"total"`
}
