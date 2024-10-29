package entities

// export const AnalyticItemSchema = z.object({
//   id: z.coerce.number(),
//   buyer: z.string(),
//   style: z.string(),
//   amount: z.coerce.number(),
// })

type AnalyticItem struct {
	ID     int    `json:"id"`
	Buyer  string `json:"buyer"`
	Style  string `json:"style"`
	Amount int    `json:"amount"`
}

type AnalyticItemPagination struct {
	Items []AnalyticItem `json:"items"`
	Total int            `json:"total"`
}

type AggregateByFactory_Row struct {
	Code  string `json:"code"`
	Size  string `json:"size"`
	Color string `json:"color"`
	Count int    `json:"count"`
}

type AggregateByFactory struct {
	Factory string                   `json:"factory"`
	Rows    []AggregateByFactory_Row `json:"rows"`
	Total   int                      `json:"total"`
}
