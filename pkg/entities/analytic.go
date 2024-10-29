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
