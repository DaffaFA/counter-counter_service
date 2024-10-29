package entities

import "time"

type FetchFilter struct {
	Cursor uint64   `json:"cursor,omitempty"`
	Limit  uint64   `json:"limit,omitempty"`
	Query  string   `json:"query,omitempty"`
	Sort   []string `json:"sort,omitempty"`
	ID     int      `json:"-"`
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

type DashboardAnalyticFilter struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

func (f *DashboardAnalyticFilter) SetDefault() {
	if f.From == "" {
		t := time.Now()
		t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

		f.From = t.Format(time.RFC3339)
	}

	if f.To == "" {
		t := time.Now()
		t = time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999, t.Location())

		f.From = t.Format(time.RFC3339)
	}
}
