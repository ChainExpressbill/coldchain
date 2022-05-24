package dashboard

import (
	"github.com/ChainExpressbill/coldchain/utils"
	"github.com/gofiber/fiber/v2"
)

// 지난 달 주문 요약
func LastMonth(c *fiber.Ctx) error {
	lastMonthSummaryCountResponse := LastMonthService()
	return c.Status(fiber.StatusOK).JSON(lastMonthSummaryCountResponse)
}

// 금일 주문 현황
func Today(c *fiber.Ctx) error {
	todaySummaryCountResponse := TodayService()
	return c.Status(fiber.StatusOK).JSON(todaySummaryCountResponse)
}

// 최근 30일 간 데이터
func Charts(c *fiber.Ctx) error {
	var chartTypes = []string{"orderers", "orders", "successes", "failures"}

	param := c.Params("type")
	if param == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	isValidChartType := utils.ContainsInSlice(chartTypes, param)
	if !isValidChartType {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	chartsResponse := ChartsByTypeService(param)
	return c.Status(fiber.StatusOK).JSON(chartsResponse)
}
