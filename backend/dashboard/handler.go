package dashboard

import (
	"github.com/ChainExpressbill/coldchain/utils"
	"github.com/gofiber/fiber/v2"
)

// @Description 지난 달 주문 요약
// @Summary 		지난 달 주문 요약
// @Tags 				Dashboard
// @Accept 			application/json;charset=UTF-8
// @Produce 		application/json;charset=UTF-8
// @Success 		200 {object} SummaryCountResponse
// @Failure 		400 {object} map[string]string
// @Failure 		500 {object} map[string]string
// @Router 			/dashboard/summary/last-month [get]
func LastMonth(c *fiber.Ctx) error {
	lastMonthSummaryCountResponse := LastMonthService()
	return c.Status(fiber.StatusOK).JSON(lastMonthSummaryCountResponse)
}

// @Description 금일 주문 현황
// @Summary 		금일 주문 현황
// @Tags 				Dashboard
// @Accept 			application/json;charset=UTF-8
// @Produce 		application/json;charset=UTF-8
// @Success 		200 {object} SummaryCountResponse
// @Failure 		400 {object} map[string]string
// @Failure 		500 {object} map[string]string
// @Router 			/dashboard/summary/today [get]
func Today(c *fiber.Ctx) error {
	todaySummaryCountResponse := TodayService()
	return c.Status(fiber.StatusOK).JSON(todaySummaryCountResponse)
}

// param name,param type,data type,is mandatory?,comment attribute(optional)
// @Description 최근 30일 간 데이터
// @Summary 		최근 30일 간 데이터
// @Tags 				Dashboard
// @Accept 			application/json;charset=UTF-8
// @Produce 		application/json;charset=UTF-8
// @Param 			type path string true "Chart Type"
// @Success 		200 {object} ChartsResponse
// @Failure 		400 {object} map[string]string
// @Failure 		500 {object} map[string]string
// @Router 			/dashboard/summary/charts/{type} [get]
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
