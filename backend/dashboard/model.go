package dashboard

import "time"

type SummaryCountResponse struct {
	OrdererCount int `json:"ordererCount"`
	OrderCount   int `json:"orderCount"`
}

type ChartsResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Count     int       `json:"count"`
}
