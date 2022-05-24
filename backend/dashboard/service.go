package dashboard

import (
	"time"
)

func LastMonthService() SummaryCountResponse {
	/**
	** @desc go time package usages
	** @see https://1minute-before6pm.tistory.com/11
	**/
	lastMonth := time.Now().AddDate(0, -1, 0)
	start := time.Date(lastMonth.Year(), lastMonth.Month(), 1, 0, 0, 0, 0, time.Local)
	end := time.Date(lastMonth.Year(), lastMonth.Month()+1, 1, 23, 59, 59, 0, time.Local).AddDate(0, 0, -1)
	orderCount := OrderCountByCreated(start, end)
	ordererCount := len(CountGroupByOrdererAndCreated(start, end))

	return SummaryCountResponse{
		OrdererCount: ordererCount,
		OrderCount:   orderCount,
	}
}

func TodayService() SummaryCountResponse {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	end := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.Local)
	orderCount := OrderCountByCreated(start, end)
	ordererCount := len(CountGroupByOrdererAndCreated(start, end))

	return SummaryCountResponse{
		OrdererCount: ordererCount,
		OrderCount:   orderCount,
	}
}

func ChartsByTypeService() ChartsResponse {
	return ChartsResponse{}
}
