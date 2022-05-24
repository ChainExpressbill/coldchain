package dashboard

import "github.com/gofiber/fiber/v2"

// ordererCount    number // 주문 요청 업체 수
// orderCount    number    // 주문 건수
// 지난 달 주문 요약
func LastMonth(c *fiber.Ctx) error {
	LastMonthService()
	return c.SendString("LastMonth")
}

// ordererCount    number // 주문 요청 업체 수
// orderCount    number    // 주문 건수
// 금일 주문 현황
func Today(c *fiber.Ctx) error {
	TodayService()
	return c.SendString("Today")
}

// 최근 30일 간 데이터
// enum type {
// 	orderers    // 주문 업체 수
// 	orders    // 주문 건수
// 	successes    // 배송 완료 수
// 	failures    // 배송 실패 수
// }
// 200 {
// 	timestamp    string    // yyyy-mm-dd
// 	count    number
// }
func Charts(c *fiber.Ctx) error {
	ChartsByTypeService()
	return c.SendString("Charts")
}
