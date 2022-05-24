package dashboard

type SummaryCountResponse struct {
	OrdererCount int `json:"ordererCount"`
	OrderCount   int `json:"orderCount"`
}
