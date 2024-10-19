package entities

type FetchFilter struct {
	Cursor uint64   `json:"cursor,omitempty"`
	Limit  uint64   `json:"limit,omitempty"`
	Query  string   `json:"query,omitempty"`
	Sort   []string `json:"sort,omitempty"`
	ID     int64    `json:"-"`
	Alias  string   `json:"-"`
}

func SetDefaultFilter(filter *FetchFilter) {
	if filter.Cursor < 1 {
		filter.Cursor = 1
	}

	if filter.Limit < 1 {
		filter.Limit = 12
	}
}
