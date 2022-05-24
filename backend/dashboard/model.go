package dashboard

import "time"

type SummaryCountResponse struct {
	OrdererCount int `json:"ordererCount"`
	OrderCount   int `json:"orderCount"`
}

type ChartsResponse []struct {
	Timestamp time.Time `json:"timestamp"`
	Orderer   string    `json:"orderer,omitempty"`
	Count     int       `json:"count"`
}
