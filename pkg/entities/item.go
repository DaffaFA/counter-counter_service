package entities

import "time"

type Item struct {
	Code      string     `json:"code,omitempty" db:"code"`
	BuyerID   int        `json:"buyer_id,omitempty" db:"buyer_id"`
	Buyer     *Setting   `json:"buyer,omitempty"`
	StyleID   int        `json:"style_id,omitempty" db:"style_id"`
	Style     *Style     `json:"style,omitempty"`
	ColorID   int        `json:"color_id,omitempty" db:"color_id"`
	Color     *Setting   `json:"color,omitempty"`
	SizeID    int        `json:"size_id,omitempty" db:"size_id"`
	Size      *Setting   `json:"size,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type ItemPagination struct {
	Items []Item `json:"items"`
	Total int    `json:"total"`
}

type ItemCreateParam struct {
	Code  string `json:"code,omitempty" db:"code"`
	Buyer string `json:"buyer,omitempty"`
	Style string `json:"style,omitempty"`
	Color string `json:"color,omitempty"`
	Size  string `json:"size,omitempty"`
}

type ItemCountChart struct {
	Bucket time.Time `json:"bucket"`
	Count  *int      `json:"count"`
}
